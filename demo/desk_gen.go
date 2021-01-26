package demo

import (
	"encoding/json"
	"entitygen/attr"

	"go.mongodb.org/mongo-driver/bson"
)

var desk *attr.Def

func init() {
	desk = &attr.Def{}

	desk.DefAttr("width", attr.Int32Attr, attr.AfBase, true)
	desk.DefAttr("height", attr.Int32Attr, attr.AfBase, true)
	desk.DefAttr("name", attr.StringAttr, attr.AfBase, true)
}

type Desk attr.StrMap

func EmptyDesk() *Desk {
	return NewDesk(0, 0, "")
}

func NewDesk(
	width int32,
	height int32,
	name string,
) *Desk {
	m := (*Desk)(attr.NewStrMap(nil))

	m.SetWidth(width)
	m.SetHeight(height)
	m.SetName(name)

	(*attr.StrMap)(m).ClearChangeKey()
	return m
}

func (m *Desk) MarshalJSON() ([]byte, error) {
	return json.Marshal((*attr.StrMap)(m).ToMap())
}
func (m *Desk) UnmarshalJSON(b []byte) error {
	mm, err := desk.UnmarshalJson(b)
	if err != nil {
		return err
	}
	(*attr.StrMap)(m).SetData(mm)
	(*attr.StrMap)(m).ForEach(func(k string, v interface{}) bool {
		if k != "id" && !desk.GetDef(k).IsPrimary() {
			v.(IField).setParent(k, (*attr.StrMap)(m))
		}
		return true
	})
	return nil
}

func (m *Desk) MarshalBSON() ([]byte, error) {
	return bson.Marshal((*attr.StrMap)(m).ToMap())
}

func (m *Desk) UnmarshalBSON(b []byte) error {
	mm, err := desk.UnmarshalBson(b)
	if err != nil {
		return err
	}
	(*attr.StrMap)(m).SetData(mm)
	(*attr.StrMap)(m).ForEach(func(k string, v interface{}) bool {
		if k != "id" && !desk.GetDef(k).IsPrimary() {
			v.(IField).setParent(k, (*attr.StrMap)(m))
		}
		return true
	})
	return nil
}

func (m *Desk) ForEach(fn func(s string, v interface{}) bool) {
	(*attr.StrMap)(m).ForEach(fn)
}

func (m *Desk) GetWidth() int32 {
	return (*attr.StrMap)(m).Int32("width")
}

func (m *Desk) SetWidth(v int32) {
	(*attr.StrMap)(m).Set("width", v)
}

func (m *Desk) GetHeight() int32 {
	return (*attr.StrMap)(m).Int32("height")
}

func (m *Desk) SetHeight(v int32) {
	(*attr.StrMap)(m).Set("height", v)
}

func (m *Desk) GetName() string {
	return (*attr.StrMap)(m).Str("name")
}

func (m *Desk) SetName(v string) {
	(*attr.StrMap)(m).Set("name", v)
}

func (m *Desk) setParent(k string, parent attr.Field) {
	(*attr.StrMap)(m).SetParent(k, parent)
}

func (m *Desk) Equal(other *Desk) bool {
	return m.GetHeight() == other.GetHeight() && m.GetWidth() == other.GetWidth() && m.GetName() == other.GetName()
}

func (m *Desk) HasChange() bool {
	return (*attr.StrMap)(m).HasChange()
}

func (m *Desk) ChangeKey() map[string]struct{} {
	return (*attr.StrMap)(m).ChangeKey()
}

func (m *Desk) ClearChangeKey() {
	(*attr.StrMap)(m).ClearChangeKey()
}
