package domain

// Room model
type Room struct {
	// ID string `key:"id" flag:"cell" client:"true" storedb:"true"`
	// CsvPos 是指配置表里面的pos，表示房间在整个建筑结构里面的方位
	CsvPos int32 `key:"csv_pos" flag:"cell" client:"true" storedb:"true"`
	// 所属建筑的id
	BuildID string `key:"build_id" flag:"cell" client:"true" storedb:"true"`
	// CsvID   int32  `key:"csv_id" flag:"cell" client:"true" storedb:"true"`
	// Left    string `key:"left" flag:"cell" client:"true" storedb:"true"`
	// Right   string `key:"right" flag:"cell" client:"true" storedb:"true"`
	// Top     string `key:"top" flag:"cell" client:"true" storedb:"true"`
	// Bottom  string `key:"bottom" flag:"cell" client:"true" storedb:"true"`
	// Front   string `key:"front" flag:"cell" client:"true" storedb:"true"`
	// Behind  string `key:"behind" flag:"cell" client:"true" storedb:"true"`

	// Extends map[int32]int32 `key:"extends" flag:"cell" client:"true" storedb:"true"`

	// Extends1 map[int32]string `key:"extends1" flag:"cell" client:"true" storedb:"true"`

	// Extends2 map[string]int32 `key:"extends2" flag:"cell" client:"true" storedb:"true"`

	// Extends3 map[string]string `key:"extends3" flag:"cell" client:"true" storedb:"true"`

	// Desk *Desk `key:"desk" flag:"cell" client:"true" storedb:"true"`

	// Desks map[int32]*Desk `key:"desks" flag:"cell" client:"true" storedb:"true"`
}

// 桌子
type Desk struct {
	Width  int32  `key:"width" flag:"cell" client:"true" storedb:"true"`
	Height int32  `key:"height" flag:"cell" client:"true" storedb:"true"`
	Name   string `key:"name" flag:"cell" client:"true" storedb:"true"`
}
