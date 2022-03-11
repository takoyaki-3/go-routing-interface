package goroutinginterface

import (
	geojson "github.com/takoyaki-3/go-geojson"
	gtfs "github.com/takoyaki-3/go-gtfs"
)

type StopTimeStr struct {
	StopID        string  `json:"stop_id"`
	ZoneID        string  `json:"zone_id"`
	StopLat       float64 `json:"stop_lat"`
	StopLon       float64 `json:"stop_lon"`
	StopName      string  `json:"stop_name"`
	ArrivalTime   string  `json:"arrival_time"`
	DepartureTime string  `json:"departure_time"`
}

// 経路探索時の重み
type CostStr struct {
	Time     *float64 `json:"time"`
	Walk     *float64 `json:"walk"`
	Transfer *float64 `json:"transfer"`
	Distance *float64 `json:"distance"`
	Fare     *float64 `json:"fare"`
}

// 経路
type TripStr struct {
	Legs       []LegStr `json:"legs"`
	Properties struct {
		TotalTime     int    `json:"total_time"`
		ArrivalTime   string `json:"arrival_time"`
		DepartureTime string `json:"departure_time"`
	} `json:"properties"`
	Costs CostStr `json:"costs"`
}

// １乗車
type LegStr struct {
	Type       string            `json:"type"`
	Trip       gtfs.Trip         `json:"trip"`
	Route      gtfs.Route        `json:"route"`
	StopTimes  []StopTimeStr     `json:"stop_times"`
	Geometry   *geojson.Geometry `json:"geometry"`
	Costs      CostStr           `json:"cost"`
	Properties PropertyStr       `json:"properties"`
}

// プロパティ
type PropertyStr struct {
	ArrivalTime   string `json:"arrival_time"`
	DepartureTime string `json:"departure_time"`
}

// 新たなコスト関数
func NewCostStr() (c CostStr) {
	c.Time = floatPointer(0.0)
	c.Walk = floatPointer(0.0)
	c.Transfer = floatPointer(0.0)
	c.Distance = floatPointer(0.0)
	c.Fare = floatPointer(0.0)
	return c
}

func floatPointer(i float64) *float64 {
	return &i
}

// コストの加算
func CostAdder(a, b CostStr) (c CostStr) {
	c.Time = floatPointer(*a.Time + *b.Time)
	c.Walk = floatPointer(*a.Walk + *b.Walk)
	c.Transfer = floatPointer(*a.Transfer + *b.Transfer)
	c.Distance = floatPointer(*a.Distance + *b.Distance)
	c.Fare = floatPointer(*a.Fare + *b.Fare)
	if *a.Fare < 0 || *b.Fare < 0 {
		c.Fare = floatPointer(-1)
	}
	return c
}

// レスポンス用構造体
type ResponsStr struct {
	Trips []TripStr `json:"trips,omitempty"`
	Meta  struct {
		EngineVersion string `json:"engine_version,omitempty"`
		LastUpdated   string `json:"last_updated,omitempty"`
	} `json:"meta,omitempty"`
	Status string `json:"status,omitempty"`
}
