package demo

import (
	"encoding/json"

	"gitlab.gamesword.com/nut/entitygen/attr"

	"go.mongodb.org/mongo-driver/bson"
)

var deskAttrDef *attr.Meta

func init() {
	deskAttrDef = &attr.Meta{}

	deskAttrDef.DefAttr("width", attr.Int32, attr.AfOtherClients, true)
	deskAttrDef.DefAttr("height", attr.Int32, attr.AfOtherClients, true)
	deskAttrDef.DefAttr("name", attr.String, attr.AfOtherClients, true)
}

type DeskDef attr.StrMap

func EmptyDesk() *DeskDef {
	return NewDesk(0, 0, "")
}
func NewDesk(width int32, height int32, name string) *DeskDef {
	m := (*DeskDef)(attr.NewStrMap(nil))
	m.SetWidth(width)
	m.SetHeight(height)
	m.SetName(name)
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
func (a *DeskDef) ForEach(fn func(s string, v interface{}) bool) {
	(*attr.StrMap)(a).ForEach(fn)
}
func (a *DeskDef) Equal(other *DeskDef) bool {
	return (*attr.StrMap)(a).Equal((*attr.StrMap)(other))
}
func (a *DeskDef) MarshalJSON() ([]byte, error) {
	return json.Marshal((*attr.StrMap)(a).ToMap())
}
func (a *DeskDef) UnmarshalJSON(b []byte) error {
	_, err := deskAttrDef.UnmarshalJson(b, (*attr.StrMap)(a))
	if err != nil {
		return err
	}

	return nil
}
func (a *DeskDef) MarshalBSON() ([]byte, error) {
	return bson.Marshal((*attr.StrMap)(a).ToMap())
}
func (a *DeskDef) UnmarshalBSON(b []byte) error {
	_, err := deskAttrDef.UnmarshalBson(b, (*attr.StrMap)(a))
	if err != nil {
		return err
	}

	return nil
}
