package demo

import (
	"encoding/json"
	"entitygen/attr"

	"go.mongodb.org/mongo-driver/bson"
)

type IField interface {
	setParent(k string, parent attr.AttrField)
}

type Desks attr.Int32Map

func EmptyDesks() *Desks {
	return NewDesks(nil)
}

func NewDesks(data map[int32]*Desk) *Desks {
	var convertData map[int32]interface{} = map[int32]interface{}{}
	for k, v := range data {
		convertData[k] = v
	}
	return (*Desks)(attr.NewInt32Map(convertData))
}

func (m *Desks) MarshalJSON() ([]byte, error) {
	return json.Marshal((*attr.Int32Map)(m).ToMap())
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

	(*attr.Int32Map)(m).SetData(convertData)
	return nil
}

func (m *Desks) MarshalBSON() ([]byte, error) {
	return bson.Marshal((*attr.Int32Map)(m).ToMap())
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
	(*attr.Int32Map)(m).SetData(convertData)
	return nil
}

func (m *Desks) Set(k int32, v *Desk) {
	v.setParent("", (*attr.Int32Map)(m))
	(*attr.Int32Map)(m).Set(k, v)
}

func (m *Desks) Get(key int32) *Desk {
	v := (*attr.Int32Map)(m).Value(key)
	if v != nil {
		return v.(*Desk)
	}
	return nil
}

func (m *Desks) Delete(key int32) bool {
	return (*attr.Int32Map)(m).Delete(key)
}

func (m *Desks) ForEach(fn func(k int32, v *Desk) bool) {
	(*attr.Int32Map)(m).ForEach(func(k int32, v interface{}) bool {
		return fn(k, v.(*Desk))
	})
}

func (m *Desks) setParent(k string, parent attr.AttrField) {
	(*attr.Int32Map)(m).SetParent(k, parent)
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
	(*attr.Int32Map)(m).ForEach(func(k int32, v interface{}) bool {
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

type Room attr.StrMap

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
	m := (*Room)(attr.NewStrMap(nil))

	m.SetCsvPos(csvPos)
	m.SetBuildID(buildID)
	m.SetExtends(extends)
	m.SetExtends1(extends1)
	m.SetExtends2(extends2)
	m.SetExtends3(extends3)
	m.SetDesk(desk)
	m.SetDesks(desks)

	(*attr.StrMap)(m).ClearChangeKey()
	return m
}

func (m *Room) HasChange() bool {
	return (*attr.StrMap)(m).HasChange()
}

func (m *Room) ChangeKey() map[string]struct{} {
	return (*attr.StrMap)(m).ChangeKey()
}

func (m *Room) ClearChangeKey() {
	(*attr.StrMap)(m).ClearChangeKey()
}

func (m *Room) MarshalJSON() ([]byte, error) {
	return json.Marshal((*attr.StrMap)(m).ToMap())
}
func (m *Room) UnmarshalJSON(b []byte) error {
	mm, err := room.UnmarshalJson(b)
	if err != nil {
		return err
	}
	(*attr.StrMap)(m).SetData(mm)
	(*attr.StrMap)(m).ForEach(func(k string, v interface{}) bool {
		if k != "id" && !room.GetDef(k).IsPrimary() {
			v.(IField).setParent(k, (*attr.StrMap)(m))
		}
		return true
	})
	return nil
}

func (m *Room) MarshalBSON() ([]byte, error) {
	return bson.Marshal((*attr.StrMap)(m).ToMap())
}

func (m *Room) UnmarshalBSON(b []byte) error {
	mm, err := room.UnmarshalBson(b)
	if err != nil {
		return err
	}
	(*attr.StrMap)(m).SetData(mm)
	(*attr.StrMap)(m).ForEach(func(k string, v interface{}) bool {
		if k != "id" && !room.GetDef(k).IsPrimary() {
			v.(IField).setParent(k, (*attr.StrMap)(m))
		}
		return true
	})
	return nil
}

func (m *Room) ForEach(fn func(s string, v interface{}) bool) {
	(*attr.StrMap)(m).ForEach(fn)
}

func (m *Room) GetBuildID() string {
	return (*attr.StrMap)(m).Str("build_id")
}

func (m *Room) SetBuildID(v string) {
	(*attr.StrMap)(m).Set("build_id", v)
}

func (m *Room) GetCsvPos() int32 {
	return (*attr.StrMap)(m).Int32("csv_pos")
}

func (m *Room) SetCsvPos(v int32) {
	(*attr.StrMap)(m).Set("csv_pos", v)
}

func (m *Room) GetExtends() *KVInt32Int32 {
	return (*attr.StrMap)(m).Value("extends").(*KVInt32Int32)
}

func (m *Room) SetExtends(extends *KVInt32Int32) {
	extends.setParent("extends", (*attr.StrMap)(m))
	(*attr.StrMap)(m).Set(
		"extends",
		extends,
	)
}

func (m *Room) GetExtends1() *KVInt32Str {
	return (*attr.StrMap)(m).Value("extends1").(*KVInt32Str)
}

func (m *Room) SetExtends1(extends *KVInt32Str) {
	extends.setParent("extends1", (*attr.StrMap)(m))
	(*attr.StrMap)(m).Set(
		"extends1",
		extends,
	)
}

func (m *Room) GetExtends2() *KVStrInt32 {
	return (*attr.StrMap)(m).Value("extends2").(*KVStrInt32)
}

func (m *Room) SetExtends2(extends *KVStrInt32) {
	extends.setParent("extends2", (*attr.StrMap)(m))
	(*attr.StrMap)(m).Set(
		"extends2",
		extends,
	)
}

func (m *Room) GetExtends3() *KVStrStr {
	return (*attr.StrMap)(m).Value("extends3").(*KVStrStr)
}

func (m *Room) SetExtends3(extends *KVStrStr) {
	extends.setParent("extends3", (*attr.StrMap)(m))
	(*attr.StrMap)(m).Set(
		"extends3",
		extends,
	)
}

func (m *Room) GetDesk() *Desk {
	return (*attr.StrMap)(m).Value("desk").(*Desk)
}

func (m *Room) SetDesk(desk *Desk) {
	desk.setParent("desk", (*attr.StrMap)(m))
	(*attr.StrMap)(m).Set(
		"desk",
		desk,
	)
}

func (m *Room) GetDesks() *Desks {
	return (*attr.StrMap)(m).Value("desks").(*Desks)
}

func (m *Room) SetDesks(desks *Desks) {
	desks.setParent("desks", (*attr.StrMap)(m))
	(*attr.StrMap)(m).Set(
		"desks",
		desks,
	)
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
