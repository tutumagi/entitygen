package demo

import (
	"encoding/json"
	"entitygen/attr"

	"go.mongodb.org/mongo-driver/bson"
)

var desk *attr.DataDef

func init() {
	desk = &attr.DataDef{}

	desk.DefAttr("width", attr.Int32Attr, attr.AfBase, true)
	desk.DefAttr("height", attr.Int32Attr, attr.AfBase, true)
	desk.DefAttr("name", attr.StringAttr, attr.AfBase, true)
}

type Desk struct {
	// parent    AttrMapImp
	// parentKey string
	_data *attr.StrMap
}

func EmptyDesk() *Desk {
	return NewDesk(0, 0, "")
}

func NewDesk(
	width int32,
	height int32,
	name string,
) *Desk {
	m := &Desk{}
	m._data = attr.NewStrMap(nil)

	m.SetWidth(width)
	m.SetHeight(height)
	m.SetName(name)

	m._data.ClearChangeKey()
	return m
}

func (m *Desk) MarshalJSON() ([]byte, error) {
	return json.Marshal(m._data.ToMap())
}
func (m *Desk) UnmarshalJSON(b []byte) error {
	mm, err := desk.UnmarshalJson(b)
	if err != nil {
		return err
	}
	m._data = attr.NewStrMap(mm)
	m._data.ForEach(func(k string, v interface{}) bool {
		if k != "id" && !desk.GetDef(k).IsPrimary() {
			v.(IField).setParent(k, m._data)
		}
		return true
	})
	return nil
}

func (m *Desk) MarshalBSON() ([]byte, error) {
	return bson.Marshal(m._data.ToMap())
}

func (m *Desk) UnmarshalBSON(b []byte) error {
	mm, err := desk.UnmarshalBson(b)
	if err != nil {
		return err
	}
	m._data = attr.NewStrMap(mm)
	m._data.ForEach(func(k string, v interface{}) bool {
		if k != "id" && !desk.GetDef(k).IsPrimary() {
			v.(IField).setParent(k, m._data)
		}
		return true
	})
	return nil
}

func (m *Desk) ForEach(fn func(s string, v interface{}) bool) {
	m._data.ForEach(fn)
}

func (m *Desk) GetWidth() int32 {
	return m._data.Int32("width")
}

func (m *Desk) SetWidth(v int32) {
	m._data.Set("width", v)
}

func (m *Desk) GetHeight() int32 {
	return m._data.Int32("height")
}

func (m *Desk) SetHeight(v int32) {
	m._data.Set("height", v)
}

func (m *Desk) GetName() string {
	return m._data.Str("name")
}

func (m *Desk) SetName(v string) {
	m._data.Set("name", v)
}

func (m *Desk) setParent(k string, parent attr.AttrField) {
	m._data.SetParent(k, parent)
}

func (m *Desk) Equal(other *Desk) bool {
	return m.GetHeight() == other.GetHeight() && m.GetWidth() == other.GetWidth() && m.GetName() == other.GetName()
}
