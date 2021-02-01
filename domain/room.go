package domain

// Room model
type Room struct {
	// ID string `flag:"cell" client:"true" storedb:"true"`
	// CsvPos 是指配置表里面的pos，表示房间在整个建筑结构里面的方位
	CsvPos int32 `flag:"cell" client:"true" storedb:"true"`
	// 所属建筑的id
	BuildID string `flag:"cell" client:"true" storedb:"true"`
	// CsvID   int32  `flag:"cell" client:"true" storedb:"true"`
	// Left    string `flag:"cell" client:"true" storedb:"true"`
	// Right   string `flag:"cell" client:"true" storedb:"true"`
	// Top     string `flag:"cell" client:"true" storedb:"true"`
	// Bottom  string `flag:"cell" client:"true" storedb:"true"`
	// Front   string `flag:"cell" client:"true" storedb:"true"`
	// Behind  string `flag:"cell" client:"true" storedb:"true"`

	Extends map[int32]int32 `flag:"cell" client:"true" storedb:"true"`

	Extends1 map[int32]string `flag:"cell" client:"true" storedb:"true"`

	Extends2 map[string]int32 `flag:"cell" client:"true" storedb:"true"`

	Extends3 map[string]string `flag:"cell" client:"true" storedb:"true"`

	Desk111 *Desk `flag:"cell" client:"true" storedb:"true"`

	Desks222 map[int32]*Desk `flag:"cell" client:"true" storedb:"true"`

	Desks333 map[string]*Desk `flag:"cell" client:"true" storedb:"true"`

	DesksArr []*Desk `flag:"cell" client:"true" storedb:"true"`
	// Strss []string `flag:"cell" client:"true" storedb:"true"`

	Int8ss []int8 `flag:"cell" client:"true" storedb:"true"`
}

// 桌子
type Desk struct {
	Width  int32  `flag:"cell" client:"true" storedb:"true"`
	Height int32  `flag:"cell" client:"true" storedb:"true"`
	Name   string `flag:"cell" client:"true" storedb:"true"`
	CsvID  int32  `flag:"cell" client:"true" storedb:"true"`
}
