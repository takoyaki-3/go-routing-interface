package goroutinginterface

// クエリ用構造体
type QueryStr struct {
	Origin      QueryNodeStr  `json:"origin"`
	Destination QueryNodeStr  `json:"destination"`
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
