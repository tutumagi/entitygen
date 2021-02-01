// Code generated by generator, DO NOT EDIT.
package entitydef

import (
	"encoding/json"
	attr "gitlab.gamesword.com/nut/entitygen/attr"
	bson "go.mongodb.org/mongo-driver/bson"
)

var deskMeta *attr.Meta

func init() {
	deskMeta = &attr.Meta{}

	deskMeta.DefAttr("width", attr.Int32, attr.AfCell, true)
	deskMeta.DefAttr("height", attr.Int32, attr.AfCell, true)
	deskMeta.DefAttr("name", attr.String, attr.AfCell, true)
	deskMeta.DefAttr("csv_id", attr.Int32, attr.AfCell, true)
}

type DeskDef attr.StrMap

func EmptyDeskDef() *DeskDef {
	return NewDeskDef(0, 0, "", 0)
}
func NewDeskDef(width int32, height int32, name string, csv_id int32) *DeskDef {
	m := (*DeskDef)(attr.NewStrMap(nil))
	m.SetWidth(width)
	m.SetHeight(height)
	m.SetName(name)
	m.SetCsvID(csv_id)
	m.ClearChangeKey()
	return m
}
func (a *DeskDef) GetWidth() int32 {
	return (*attr.StrMap)(a).Int32("width")
}
func (a *DeskDef) SetWidth(width int32) {
	(*attr.StrMap)(a).Set("width", width)
}

func (a *DeskDef) GetHeight() int32 {
	return (*attr.StrMap)(a).Int32("height")
}
func (a *DeskDef) SetHeight(height int32) {
	(*attr.StrMap)(a).Set("height", height)
}

func (a *DeskDef) GetName() string {
	return (*attr.StrMap)(a).Str("name")
}
func (a *DeskDef) SetName(name string) {
	(*attr.StrMap)(a).Set("name", name)
}

func (a *DeskDef) GetCsvID() int32 {
	return (*attr.StrMap)(a).Int32("csv_id")
}
func (a *DeskDef) SetCsvID(csv_id int32) {
	(*attr.StrMap)(a).Set("csv_id", csv_id)
}

func (a *DeskDef) HasChange() bool {
	return (*attr.StrMap)(a).HasChange()
}
func (a *DeskDef) ChangeKey() map[string]struct{} {
	return (*attr.StrMap)(a).ChangeKey()
}
func (a *DeskDef) ClearChangeKey() {
	(*attr.StrMap)(a).ClearChangeKey()
}
func (a *DeskDef) setParent(k string, parent attr.Field) {
	(*attr.StrMap)(a).SetParent(k, parent)
}
func (a *DeskDef) ForEach(fn func(k string, v interface{}) bool) {
	(*attr.StrMap)(a).ForEach(fn)
}
func (a *DeskDef) Equal(other *DeskDef) bool {
	return (*attr.StrMap)(a).Equal((*attr.StrMap)(other))
}
func (a *DeskDef) data() map[string]interface{} {
	dd := map[string]interface{}{}
	a.ForEach(func(k string, v interface{}) bool {
		dd[k] = v
		return true
	})
	return dd
}
func (a *DeskDef) MarshalJSON() ([]byte, error) {
	return json.Marshal((*attr.StrMap)(a).ToMap())
}
func (a *DeskDef) UnmarshalJSON(b []byte) error {
	mm, err := deskMeta.UnmarshalJson(b)
	if err != nil {
		return err
	}
	(*attr.StrMap)(a).SetData(mm)
	(*attr.StrMap)(a).ForEach(func(k string, v interface{}) bool {
		if k != "id" && !deskMeta.GetDef(k).IsPrimary() {
			v.(IField).setParent(k, (*attr.StrMap)(a))
		}
		return true
	})
	return nil
}
func (a *DeskDef) MarshalBSON() ([]byte, error) {
	return bson.Marshal((*attr.StrMap)(a).ToMap())
}
func (a *DeskDef) UnmarshalBSON(b []byte) error {
	mm, err := deskMeta.UnmarshalBson(b)
	if err != nil {
		return err
	}
	(*attr.StrMap)(a).SetData(mm)
	(*attr.StrMap)(a).ForEach(func(k string, v interface{}) bool {
		if k != "id" && !deskMeta.GetDef(k).IsPrimary() {
			v.(IField).setParent(k, (*attr.StrMap)(a))
		}
		return true
	})
	return nil
}