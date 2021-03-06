package domain

import "github.com/g3n/engine/math32"

// RoomDef model
// 以 Def 结尾的表示是实体定义
type RoomDef struct {
	// CsvPos 是指配置表里面的pos，表示房间在整个建筑结构里面的方位
	CsvPos int32 `flag:"cell" client:"true" storedb:"true"`
	// 所属建筑的id
	BuildID string `flag:"cell" client:"true" storedb:"true"`

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

	Vec3 *math32.Vector3 `flag:"cell" client:"true" storedb:"true"`

	Vec2Arr []*math32.Vector2 `flag:"cell" client:"true" storedb:"true"`

	Pos *math32.Vector2 `flag:"cell" client:"true" storedb:"true"`
	Rot *math32.Vector2 `flag:"base" client:"true" storedb:"false"`
}

// 桌子
type Desk struct {
	Width  int32
	Height int32
	Name   string
	CsvID  int32
	Below  *Desk
}
