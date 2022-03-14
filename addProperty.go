package goroutinginterface

import (
	"errors"
	"sort"

	geojson "github.com/takoyaki-3/go-geojson"
	gtfs "github.com/takoyaki-3/go-gtfs/v2"
	// fare "github.com/takoyaki-3/go-gtfs-fare"
)

// Legに対して付加情報を追加する関数
// leg.Trip.ID,stopTimesのstop_id，arrival_time or departure_time が判明している必要あり
// walkの場合は上記に加え、type="walk",geometryが必要
func (leg *LegStr) AddProperty(g *gtfs.GTFS) error {

	leg.Trip = g.GetTrip(leg.Trip.ID)
	leg.Route = g.GetRoute(leg.Trip.RouteID)

	// StopTimes
	// はじめに経由地を埋める
	if len(leg.StopTimes) == 2 {
		tripStopTimes := []gtfs.StopTime{}
		for _, stopTime := range g.StopsTimes {
			if stopTime.TripID == leg.Trip.ID {
				tripStopTimes = append(tripStopTimes, stopTime)
			}
		}
		sort.Slice(tripStopTimes, func(i, j int) bool {
			return tripStopTimes[i].StopSeq < tripStopTimes[j].StopSeq
		})
		// 乗車停留所，下車停留所を特定する
		fromStopIndex, toStopIndex := -1, -1
		fromStopTime := leg.StopTimes[0]
		toStopTime := leg.StopTimes[len(leg.StopTimes)-1]
		for i, stopTime := range tripStopTimes {
			if stopTime.StopID == fromStopTime.StopID && (stopTime.Arrival == fromStopTime.ArrivalTime || stopTime.Departure == fromStopTime.DepartureTime) {
				fromStopIndex = i
			} else if stopTime.StopID == toStopTime.StopID && (stopTime.Arrival == toStopTime.ArrivalTime || stopTime.Departure == toStopTime.DepartureTime) {
				toStopIndex = i
			}
		}
		if fromStopIndex < 0 || toStopIndex < 0 {
			return errors.New("from_stop or to_stop is not found in stop_times")
		}
		leg.StopTimes = make([]StopTimeStr, 0)
		for i := fromStopIndex; i <= toStopIndex; i++ {
			stopTime := tripStopTimes[i]
			leg.StopTimes = append(leg.StopTimes, StopTimeStr{
				StopID:        stopTime.StopID,
				ArrivalTime:   stopTime.Arrival,
				DepartureTime: stopTime.Departure,
			})
		}
	}
	// 次にstop_id以外の要素を埋める
	for i, stopTime := range leg.StopTimes {
		if stopTime.ArrivalTime == "" {
			stopTime.ArrivalTime = stopTime.DepartureTime
		} else if stopTime.DepartureTime == "" {
			stopTime.DepartureTime = stopTime.ArrivalTime
		}
		stop := g.GetStop(stopTime.StopID)
		stopTime.StopLat = stop.Latitude
		stopTime.StopLon = stop.Longitude
		stopTime.StopName = stop.Name
		stopTime.ZoneID = stop.ZoneID
		leg.StopTimes[i] = stopTime
	}

	// Geometry
	if leg.Geometry == nil {
		coordinates := [][]float64{}
		for _, stopTime := range leg.StopTimes {
			coordinates = append(coordinates, []float64{stopTime.StopLon, stopTime.StopLat})
		}
		geom := &geojson.Geometry{
			Type:        "LineString",
			Coordinates: coordinates,
			Properties:  nil,
		}

		leg.Geometry = geom
	}

	// Property
	leg.Properties = PropertyStr{
		DepartureTime: leg.StopTimes[0].DepartureTime,
		ArrivalTime:   leg.StopTimes[len(leg.StopTimes)-1].ArrivalTime,
	}

	// Cost
	timeCost := float64(gtfs.HHMMSS2Sec(leg.Properties.ArrivalTime) - gtfs.HHMMSS2Sec(leg.Properties.DepartureTime))
	walkDistance := 0.0
	transfer := 0.0
	distance := leg.Geometry.Distance()
	if leg.Type == "walk" {
		walkDistance = distance
	}

	// fare cost
	fareCost := 0.0
	// f := fare.
	if leg.Type != "walk" && leg.Type != "wait" {
		p,err := g.GetFareAttributeFromOD(leg.StopTimes[0].StopID,leg.StopTimes[len(leg.StopTimes)-1].StopID,leg.Route.ID)
		if err != nil {
			p,err := g.GetFareAttributeFromOD(leg.StopTimes[0].ZoneID,leg.StopTimes[len(leg.StopTimes)-1].ZoneID,leg.Route.ID)
			if err != nil {
				fareCost = -1.0
			} else {
				fareCost = p.Price
			}
		} else {
			fareCost = p.Price
		}
	}

	leg.Costs = CostStr{
		Time:     floatPointer(timeCost),
		Walk:     floatPointer(walkDistance),
		Transfer: floatPointer(transfer),
		Distance: floatPointer(distance),
		Fare:     floatPointer(fareCost),
	}

	if leg.Type != "walk" && leg.Type != "wait" {
		// Type 設定
		switch leg.Route.Type {
		case 0:
			leg.Type = "tram"
		case 1:
			leg.Type = "subway"
		case 2:
			leg.Type = "train"
		case 3:
			leg.Type = "bus"
		case 4:
			leg.Type = "ferry"
		case 5:
			leg.Type = "cable_tram"
		case 6:
			leg.Type = "lift"
		case 7:
			leg.Type = "lift"
		case 11:
			leg.Type = "trolley_bus"
		case 12:
			leg.Type = "monorail"
		}
	}

	return nil
}
