package goroutinginterface

import (
	"math"
	
	gm "github.com/takoyaki-3/go-map/v2"
	gtfs "github.com/takoyaki-3/go-gtfs/v2"
)

// クエリ用構造体
type QueryStr struct {
	Origin      QueryNodeStr  `json:"origin"`
	Destination QueryNodeStr  `json:"destination"`
	Limit       QueryLimit    `json:"limit"`
	Properties  PropertiesStr `json:"properties"`
}

// 始点発着点重み
type QueryNodeStr struct {
	StopId *string  `json:"stop_id"`
	Lat    *float64 `json:"lat"`
	Lon    *float64 `json:"lon"`
	Time   *int     `json:"time"`
}

// リクエストプロパティ
type PropertiesStr struct {
	WalkingSpeed float64 `json:"walking_speed"`
	Timetable    string  `json:"timetable"`
}

type QueryLimit struct {
	Time     int `json:"time"`
	Transfer int `json:"transfer"`
}

// QueryNodeStr から最も近い頂点を出力
func FindNearestNode(qns QueryNodeStr, g *gtfs.GTFS) string {
	stopId := ""
	minD := math.MaxFloat64
	if qns.StopId == nil {
		for _, stop := range g.Stops {
			d := gm.HubenyDistance(gm.Node{
				Lat: stop.Latitude,
				Lon: stop.Longitude},
				gm.Node{
					Lat: *qns.Lat,
					Lon: *qns.Lon})
			if d < minD {
				stopId = stop.ID
				minD = d
			}
		}
	} else {
		stopId = *qns.StopId
	}
	return stopId
}

