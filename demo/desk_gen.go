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
	data *attr.AttrMap
}

func defaultDesk() *Desk {
	return NewDesk(0, 0, "")
}

func NewDesk(
	width int32,
	height int32,
	name string,
) *Desk {
	m := &Desk{}
	m.data = attr.NewAttrMap()

	m.SetWidth(width)
	m.SetHeight(height)
	m.SetName(name)

	m.data.ClearChangeKey()
	return m
}

func (m *Desk) MarshalJSON() ([]byte, error) {
	return json.Marshal(m.data.ToMap())
}
func (m *Desk) UnmarshalJSON(b []byte) error {
	mm, err := desk.UnmarshalJson(b)
	if err != nil {
		return err
	}
	m.data.SetData(mm)
	m.data.ForEach(func(k string, v interface{}) bool {
		if k != "id" && !desk.GetDef(k).IsPrimary() {
			v.(IField).setRootKey(k)
			v.(IField).setAncestry(m.data)
		}
		return true
	})
	return nil
}

func (m *Desk) MarshalBSON() ([]byte, error) {
	return bson.Marshal(m.data.ToMap())
}

func (m *Desk) UnmarshalBSON(b []byte) error {
	mm, err := desk.UnmarshalBson(b)
	if err != nil {
		return err
	}
	m.data.SetData(mm)
	return nil
}

func (m *Desk) InitAttrMap() {
	m.data = attr.NewAttrMap()
}

func (m *Desk) ForEach(fn func(s string, v interface{}) bool) {
	m.data.ForEach(fn)
}

func (m *Desk) GetWidth() int32 {
	return m.data.Int32("width")
}

func (m *Desk) SetWidth(v int32) {
	m.data.Set("width", v)
}

func (m *Desk) GetHeight() int32 {
	return m.data.Int32("height")
}

func (m *Desk) SetHeight(v int32) {
	m.data.Set("height", v)
}

func (m *Desk) GetName() string {
	return m.data.Str("name")
}

func (m *Desk) SetName(v string) {
	m.data.Set("name", v)
}

func (m *Desk) setRootKey(k string) {
	m.data.SetRootKey(k)
}

func (m *Desk) setAncestry(ancestry *attr.AttrMap) {
	m.data.SetAncestry(ancestry)
}
