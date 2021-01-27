package demo

import (
	"encoding/json"
	"entitygen/attr"

	"go.mongodb.org/mongo-driver/bson"
)

var deskAttrDef *attr.Def

func init() {
	deskAttrDef = &attr.Def{}

	deskAttrDef.DefAttr("width", attr.Int32, attr.AfBase, true)
	deskAttrDef.DefAttr("height", attr.Int32, attr.AfBase, true)
	deskAttrDef.DefAttr("name", attr.String, attr.AfBase, true)
}

type DeskDef attr.StrMap

func EmptyDesk() *DeskDef {
	return NewDesk(0, 0, "")
}

func NewDesk(
	width int32,
	height int32,
	name string,
) *DeskDef {
	m := (*DeskDef)(attr.NewStrMap(nil))

	m.SetWidth(width)
	m.SetHeight(height)
	m.SetName(name)

	(*attr.StrMap)(m).ClearChangeKey()
	return m
}

func (m *DeskDef) MarshalJSON() ([]byte, error) {
	return json.Marshal((*attr.StrMap)(m).ToMap())
}
func (m *DeskDef) UnmarshalJSON(b []byte) error {
	mm, err := deskAttrDef.UnmarshalJson(b)
	if err != nil {
		return err
	}
	(*attr.StrMap)(m).SetData(mm)
	(*attr.StrMap)(m).ForEach(func(k string, v interface{}) bool {
		if k != "id" && !deskAttrDef.GetDef(k).IsPrimary() {
			v.(IField).setParent(k, (*attr.StrMap)(m))
		}
		return true
	})
	return nil
}

func (m *DeskDef) MarshalBSON() ([]byte, error) {
	return bson.Marshal((*attr.StrMap)(m).ToMap())
}

func (m *DeskDef) UnmarshalBSON(b []byte) error {
	mm, err := deskAttrDef.UnmarshalBson(b)
	if err != nil {
		return err
	}
	(*attr.StrMap)(m).SetData(mm)
	(*attr.StrMap)(m).ForEach(func(k string, v interface{}) bool {
		if k != "id" && !deskAttrDef.GetDef(k).IsPrimary() {
			v.(IField).setParent(k, (*attr.StrMap)(m))
		}
		return true
	})
	return nil
}

func (m *DeskDef) ForEach(fn func(s string, v interface{}) bool) {
	(*attr.StrMap)(m).ForEach(fn)
}

func (m *DeskDef) GetWidth() int32 {
	return (*attr.StrMap)(m).Int32("width")
}

func (m *DeskDef) SetWidth(v int32) {
	(*attr.StrMap)(m).Set("width", v)
}

func (m *DeskDef) GetHeight() int32 {
	return (*attr.StrMap)(m).Int32("height")
}

func (m *DeskDef) SetHeight(v int32) {
	(*attr.StrMap)(m).Set("height", v)
}

func (m *DeskDef) GetName() string {
	return (*attr.StrMap)(m).Str("name")
}

func (m *DeskDef) SetName(v string) {
	(*attr.StrMap)(m).Set("name", v)
}

func (m *DeskDef) setParent(k string, parent attr.Field) {
	(*attr.StrMap)(m).SetParent(k, parent)
}

func (m *DeskDef) Equal(other *DeskDef) bool {
	return (*attr.StrMap)(m).Equal((*attr.StrMap)(other))
}

func (m *DeskDef) HasChange() bool {
	return (*attr.StrMap)(m).HasChange()
}

func (m *DeskDef) ChangeKey() map[string]struct{} {
	return (*attr.StrMap)(m).ChangeKey()
}

func (m *DeskDef) ClearChangeKey() {
	(*attr.StrMap)(m).ClearChangeKey()
}
