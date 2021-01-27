// Code generated by generator, DO NOT EDIT.
package entitydef

import (
	"encoding/json"
	attr "entitygen/attr"
	bson "go.mongodb.org/mongo-driver/bson"
)

var roomAttrDef *attr.Def

func init() {
	roomAttrDef = &attr.Def{}

	roomAttrDef.DefAttr("csv_pos", attr.Int32, attr.AfCell, true)
	roomAttrDef.DefAttr("build_id", attr.String, attr.AfCell, true)
}

type RoomDef attr.StrMap

func EmptyRoom() *RoomDef {
	return NewRoom(0, "")
}
func NewRoom(csv_pos int32, build_id string) *RoomDef {
	m := (*RoomDef)(attr.NewStrMap(nil))
	m.SetCsvPos(csv_pos)
	m.SetBuildID(build_id)
	m.ClearChangeKey()
	return m
}
func (a *RoomDef) GetCsvPos() int32 {
	return (*attr.StrMap)(a).Int32("csv_pos")
}
func (a *RoomDef) SetCsvPos(csv_pos int32) {
	(*attr.StrMap)(a).Set("csv_pos", csv_pos)
}

func (a *RoomDef) GetBuildID() string {
	return (*attr.StrMap)(a).Str("build_id")
}
func (a *RoomDef) SetBuildID(build_id string) {
	(*attr.StrMap)(a).Set("build_id", build_id)
}

func (a *RoomDef) HasChange() bool {
	return (*attr.StrMap)(a).HasChange()
}
func (a *RoomDef) ChangeKey() map[string]struct{} {
	return (*attr.StrMap)(a).ChangeKey()
}
func (a *RoomDef) ClearChangeKey() {
	(*attr.StrMap)(a).ClearChangeKey()
}
func (a *RoomDef) setParent(k string, parent attr.Field) {
	(*attr.StrMap)(a).SetParent(k, parent)
}
func (a *RoomDef) ForEach(fn func(s string, v interface{}) bool) {
	(*attr.StrMap)(a).ForEach(fn)
}
func (a *RoomDef) Equal(other *RoomDef) bool {
	return (*attr.StrMap)(a).Equal((*attr.StrMap)(other))
}
func (a *RoomDef) MarshalJSON() ([]byte, error) {
	return json.Marshal((*attr.StrMap)(a).ToMap())
}
func (a *RoomDef) UnmarshalJSON(b []byte) error {
	mm, err := roomAttrDef.UnmarshalJson(b)
	if err != nil {
		return err
	}
	(*attr.StrMap)(a).SetData(mm)
	(*attr.StrMap)(a).ForEach(func(k string, v interface{}) bool {
		if k != "id" && !roomAttrDef.GetDef(k).IsPrimary() {
			v.(IField).setParent(k, (*attr.StrMap)(a))
		}
		return true
	})
	return nil
}
func (a *RoomDef) MarshalBSON() ([]byte, error) {
	return bson.Marshal((*attr.StrMap)(a).ToMap())
}
func (a *RoomDef) UnmarshalBSON(b []byte) error {
	mm, err := roomAttrDef.UnmarshalBson(b)
	if err != nil {
		return err
	}
	(*attr.StrMap)(a).SetData(mm)
	(*attr.StrMap)(a).ForEach(func(k string, v interface{}) bool {
		if k != "id" && !roomAttrDef.GetDef(k).IsPrimary() {
			v.(IField).setParent(k, (*attr.StrMap)(a))
		}
		return true
	})
	return nil
}
