package demo

import (
	"encoding/json"
	"entitygen/attr"

	"go.mongodb.org/mongo-driver/bson"
)

type IField interface {
	setParent(k string, parent attr.AttrField)
}

// Desks 外观
type Desks struct {
	_data *attr.Int32Map
}

func EmptyDesks() *Desks {
	return NewDesks(nil)
}

func NewDesks(data map[int32]*Desk) *Desks {
	var convertData map[int32]interface{} = map[int32]interface{}{}
	for k, v := range data {
		convertData[k] = v
	}
	dd := attr.NewInt32InterfaceMap(convertData)

	return &Desks{
		_data: dd,
	}
}

func (m *Desks) MarshalJSON() ([]byte, error) {
	return json.Marshal(m._data.ToMap())
}
func (m *Desks) UnmarshalJSON(b []byte) error {
	var dd map[int32]*Desk = map[int32]*Desk{}
	err := json.Unmarshal(b, &dd)
	if err != nil {
		return err
	}
	var convertData map[int32]interface{} = map[int32]interface{}{}
	for k, v := range dd {
		convertData[k] = v
	}

	m._data = attr.NewInt32InterfaceMap(convertData)
	return nil
}

func (m *Desks) MarshalBSON() ([]byte, error) {
	return bson.Marshal(m._data.ToMap())
}

func (m *Desks) UnmarshalBSON(b []byte) error {
	var dd map[int32]*Desk = map[int32]*Desk{}
	err := bson.Unmarshal(b, &dd)
	if err != nil {
		return err
	}
	var convertData map[int32]interface{} = map[int32]interface{}{}
	for k, v := range dd {
		convertData[k] = v
	}

	m._data = attr.NewInt32InterfaceMap(convertData)
	return nil
}

func (m *Desks) Set(k int32, v *Desk) {
	v.setParent("", m._data)
	m._data.Set(k, v)
}

func (m *Desks) Get(key int32) *Desk {
	v := m._data.Value(key)
	if v != nil {
		return v.(*Desk)
	}
	return nil
}

func (m *Desks) Delete(key int32) bool {
	return m._data.Delete(key)
}

func (m *Desks) ForEach(fn func(k int32, v *Desk) bool) {
	m._data.ForEach(func(k int32, v interface{}) bool {
		return fn(k, v.(*Desk))
	})
}

func (m *Desks) setParent(k string, parent attr.AttrField) {
	m._data.SetParent(k, parent)
}

func (m *Desks) Equal(other *Desks) bool {
	equal := true
	m.ForEach(func(k int32, v *Desk) bool {
		if !v.Equal(other.Get(k)) {
			equal = false
			return false
		}
		return true
	})
	return equal
}

func (m *Desks) data() map[int32]*Desk {
	var dd map[int32]*Desk = map[int32]*Desk{}
	m._data.ForEach(func(k int32, v interface{}) bool {
		dd[k] = v.(*Desk)
		return true
	})
	return dd
}

var room *attr.DataDef

func init() {
	room = &attr.DataDef{}

	room.DefAttr("csv_pos", attr.Int32Attr, attr.AfBase, true)
	room.DefAttr("build_id", attr.StringAttr, attr.AfBase, true)
	room.DefAttr("extends", &KVInt32Int32{}, attr.AfBase, true)
	room.DefAttr("extends1", &KVInt32Str{}, attr.AfBase, true)
	room.DefAttr("extends2", &KVStrInt32{}, attr.AfBase, true)
	room.DefAttr("extends3", &KVStrStr{}, attr.AfBase, true)
	room.DefAttr("desk", &Desk{}, attr.AfBase, true)
	room.DefAttr("desks", &Desks{}, attr.AfBase, true)
}

type Room struct {
	_data *attr.AttrMap
}

func EmptyRoom() *Room {
	return NewRoom(
		0,
		"",
		EmptyKVInt32Int32(),
		EmptyKVInt32Str(),
		EmptyStrInt32(),
		EmptyKVStrStr(),
		EmptyDesk(),
		EmptyDesks(),
	)
}

func NewRoom(
	csvPos int32,
	buildID string,
	extends *KVInt32Int32,
	extends1 *KVInt32Str,
	extends2 *KVStrInt32,
	extends3 *KVStrStr,
	desk *Desk,
	desks *Desks,
) *Room {
	m := &Room{}
	m._data = attr.NewAttrMap()

	m.SetCsvPos(csvPos)
	m.SetBuildID(buildID)
	m.SetExtends(extends)
	m.SetExtends1(extends1)
	m.SetExtends2(extends2)
	m.SetExtends3(extends3)
	m.SetDesk(desk)
	m.SetDesks(desks)

	m._data.ClearChangeKey()
	return m
}

func (m *Room) MarshalJSON() ([]byte, error) {
	return json.Marshal(m._data.ToMap())
}
func (m *Room) UnmarshalJSON(b []byte) error {
	mm, err := room.UnmarshalJson(b)
	if err != nil {
		return err
	}
	m._data.SetData(mm)
	m._data.ForEach(func(k string, v interface{}) bool {
		if k != "id" && !room.GetDef(k).IsPrimary() {
			v.(IField).setParent(k, m._data)
		}
		return true
	})
	return nil
}

func (m *Room) MarshalBSON() ([]byte, error) {
	return bson.Marshal(m._data.ToMap())
}

func (m *Room) UnmarshalBSON(b []byte) error {
	mm, err := room.UnmarshalBson(b)
	if err != nil {
		return err
	}
	m._data.SetData(mm)
	return nil
}

func (m *Room) InitAttrMap() {
	m._data = attr.NewAttrMap()
}

func (m *Room) ForEach(fn func(s string, v interface{}) bool) {
	m._data.ForEach(fn)
}

func (m *Room) GetBuildID() string {
	return m._data.Str("build_id")
}

func (m *Room) SetBuildID(v string) {
	m._data.Set("build_id", v)
}

func (m *Room) GetCsvPos() int32 {
	return m._data.Int32("csv_pos")
}

func (m *Room) SetCsvPos(v int32) {
	m._data.Set("csv_pos", v)
}

func (m *Room) GetExtends() *KVInt32Int32 {
	return m._data.Value("extends").(*KVInt32Int32)
}

func (m *Room) SetExtends(extends *KVInt32Int32) {
	extends.setParent("extends", m._data)
	m._data.Set(
		"extends",
		extends,
	)
}

func (m *Room) GetExtends1() *KVInt32Str {
	return m._data.Value("extends1").(*KVInt32Str)
}

func (m *Room) SetExtends1(extends *KVInt32Str) {
	extends.setParent("extends1", m._data)
	m._data.Set(
		"extends1",
		extends,
	)
}

func (m *Room) GetExtends2() *KVStrInt32 {
	return m._data.Value("extends2").(*KVStrInt32)
}

func (m *Room) SetExtends2(extends *KVStrInt32) {
	extends.setParent("extends2", m._data)
	m._data.Set(
		"extends2",
		extends,
	)
}

func (m *Room) GetExtends3() *KVStrStr {
	return m._data.Value("extends3").(*KVStrStr)
}

func (m *Room) SetExtends3(extends *KVStrStr) {
	extends.setParent("extends3", m._data)
	m._data.Set(
		"extends3",
		extends,
	)
}

func (m *Room) GetDesk() *Desk {
	return m._data.Value("desk").(*Desk)
}

func (m *Room) SetDesk(desk *Desk) {
	desk.setParent("desk", m._data)
	m._data.Set(
		"desk",
		desk,
	)
}

func (m *Room) GetDesks() *Desks {
	return m._data.Value("desks").(*Desks)
}

func (m *Room) SetDesks(desks *Desks) {
	desks.setParent("desks", m._data)
	m._data.Set(
		"desks",
		desks,
	)
}

func (m *Room) setParent(k string, parent attr.AttrField) {
	m._data.SetParent(k, parent)
}

// func (m *Room) Equal(other *Room) bool {
// 	equal := true
// 	m.ForEach(func(k string, v interface{}) bool {
// 		def := room.GetDef(k)
// 		if def.IsPrimary() {
// 			if v != other {
// 				equal = false
// 				return false
// 			}
// 		} else {
// 			v.(IEqualable)
// 		}
// 	})
// 	return equal
// }
