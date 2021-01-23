package demo

// Room model
type RoomXXX struct {
	// ID string `bson:"id" json:"id"`
	// CsvPos 是指配置表里面的pos，表示房间在整个建筑结构里面的方位
	CsvPos int32 `bson:"csv_pos" json:"csv_pos"`
	// 所属建筑的id
	BuildID string `bson:"build_id" json:"build_id"`
	// CsvID   int32  `bson:"csv_id" json:"csv_id"`
	// Left    string `bson:"left" json:"left"`
	// Right   string `bson:"right" json:"right"`
	// Top     string `bson:"top" json:"top"`
	// Bottom  string `bson:"bottom" json:"bottom"`
	// Front   string `bson:"front" json:"front"`
	// Behind  string `bson:"behind" json:"behind"`

	Extends map[int32]int32 `bson:"extends" json:"extends"`

	Extends1 map[int32]string `bson:"extends1" json:"extends1"`

	Extends2 map[string]int32 `bson:"extends2" json:"extends2"`

	Extends3 map[string]string `bson:"extends3" json:"extends3"`
}
