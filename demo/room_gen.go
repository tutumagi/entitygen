package demo

import (
	"encoding/json"

	"gitlab.gamesword.com/nut/entitygen/attr"

	"go.mongodb.org/mongo-driver/bson"
)

type IField interface {
	setParent(k string, parent attr.Field)
}

type Desks attr.Int32Map

func EmptyDesks() *Desks {
	return NewDesks(nil)
}

func NewDesks(data map[int32]*DeskDef) *Desks {
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
	var dd map[int32]*DeskDef = map[int32]*DeskDef{}
	err := json.Unmarshal(b, &dd)
	if err != nil {
		return err
	}
	var convertData map[int32]interface{} = map[int32]interface{}{}
	for k, v := range dd {
		convertData[k] = v

		v.setParent("", (*attr.Int32Map)(m))
	}

	(*attr.Int32Map)(m).SetData(convertData)
	return nil
}

func (m *Desks) MarshalBSON() ([]byte, error) {
	return bson.Marshal((*attr.Int32Map)(m).ToMap())
}

func (m *Desks) UnmarshalBSON(b []byte) error {
	var dd map[int32]*DeskDef = map[int32]*DeskDef{}
	err := bson.Unmarshal(b, &dd)
	if err != nil {
		return err
	}
	var convertData map[int32]interface{} = map[int32]interface{}{}
	for k, v := range dd {
		convertData[k] = v

		v.setParent("", (*attr.Int32Map)(m))
	}
	(*attr.Int32Map)(m).SetData(convertData)
	return nil
}

func (m *Desks) Set(k int32, v *DeskDef) {
	v.setParent("", (*attr.Int32Map)(m))
	(*attr.Int32Map)(m).Set(k, v)
}

func (m *Desks) Get(key int32) *DeskDef {
	v := (*attr.Int32Map)(m).Value(key)
	if v != nil {
		return v.(*DeskDef)
	}
	return nil
}

func (m *Desks) Delete(key int32) bool {
	return (*attr.Int32Map)(m).Delete(key)
}

func (m *Desks) ForEach(fn func(k int32, v *DeskDef) bool) {
	(*attr.Int32Map)(m).ForEach(func(k int32, v interface{}) bool {
		return fn(k, v.(*DeskDef))
	})
}

func (m *Desks) setParent(k string, parent attr.Field) {
	(*attr.Int32Map)(m).SetParent(k, parent)
}

func (m *Desks) Equal(other *Desks) bool {
	equal := true
	m.ForEach(func(k int32, v *DeskDef) bool {
		if !v.Equal(other.Get(k)) {
			equal = false
			return false
		}
		return true
	})
	return equal
}

func (m *Desks) data() map[int32]*DeskDef {
	var dd map[int32]*DeskDef = map[int32]*DeskDef{}
	(*attr.Int32Map)(m).ForEach(func(k int32, v interface{}) bool {
		dd[k] = v.(*DeskDef)
		return true
	})
	return dd
}

var roomAttrDef *attr.Meta

func init() {
	roomAttrDef = &attr.Meta{}

	roomAttrDef.DefAttr("csv_pos", attr.Int32, attr.AfBase, true)
	roomAttrDef.DefAttr("build_id", attr.String, attr.AfBase, true)
	roomAttrDef.DefAttr("extends", (*KVInt32Int32)(nil), attr.AfBase, true)
	roomAttrDef.DefAttr("extends1", (*KVInt32Str)(nil), attr.AfBase, true)
	roomAttrDef.DefAttr("extends2", (*KVStrInt32)(nil), attr.AfBase, true)
	roomAttrDef.DefAttr("extends3", (*KVStrStr)(nil), attr.AfBase, true)
	roomAttrDef.DefAttr("desk", (*DeskDef)(nil), attr.AfBase, true)
	roomAttrDef.DefAttr("desks", (*Desks)(nil), attr.AfBase, true)
}

type RoomDef attr.StrMap

func EmptyRoom() *RoomDef {
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
	desk *DeskDef,
	desks *Desks,
) *RoomDef {
	m := (*RoomDef)(attr.NewStrMap(nil))

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

func (m *RoomDef) String() string {
	return (*attr.StrMap)(m).String()
}

func (m *RoomDef) HasChange() bool {
	return (*attr.StrMap)(m).HasChange()
}

func (m *RoomDef) ChangeKey() map[string]struct{} {
	return (*attr.StrMap)(m).ChangeKey()
}

func (m *RoomDef) ClearChangeKey() {
	(*attr.StrMap)(m).ClearChangeKey()
}

func (m *RoomDef) MarshalJSON() ([]byte, error) {
	return json.Marshal((*attr.StrMap)(m).ToMap())
}
func (m *RoomDef) UnmarshalJSON(b []byte) error {
	mm, err := roomAttrDef.UnmarshalJson(b)
	if err != nil {
		return err
	}
	(*attr.StrMap)(m).SetData(mm)
	(*attr.StrMap)(m).ForEach(func(k string, v interface{}) bool {
		if k != "id" && !roomAttrDef.GetDef(k).IsPrimary() {
			v.(IField).setParent(k, (*attr.StrMap)(m))
		}
		return true
	})
	return nil
}

func (m *RoomDef) MarshalBSON() ([]byte, error) {
	return bson.Marshal((*attr.StrMap)(m).ToMap())
}

func (m *RoomDef) UnmarshalBSON(b []byte) error {
	mm, err := roomAttrDef.UnmarshalBson(b)
	if err != nil {
		return err
	}
	(*attr.StrMap)(m).SetData(mm)
	(*attr.StrMap)(m).ForEach(func(k string, v interface{}) bool {
		if k != "id" && !roomAttrDef.GetDef(k).IsPrimary() {
			v.(IField).setParent(k, (*attr.StrMap)(m))
		}
		return true
	})
	return nil
}

func (m *RoomDef) ForEach(fn func(s string, v interface{}) bool) {
	(*attr.StrMap)(m).ForEach(fn)
}

func (m *RoomDef) GetBuildID() string {
	return (*attr.StrMap)(m).Str("build_id")
}

func (m *RoomDef) SetBuildID(v string) {
	(*attr.StrMap)(m).Set("build_id", v)
}

func (m *RoomDef) GetCsvPos() int32 {
	return (*attr.StrMap)(m).Int32("csv_pos")
}

func (m *RoomDef) SetCsvPos(v int32) {
	(*attr.StrMap)(m).Set("csv_pos", v)
}

func (m *RoomDef) GetExtends() *KVInt32Int32 {
	return (*attr.StrMap)(m).Value("extends").(*KVInt32Int32)
}

func (m *RoomDef) SetExtends(extends *KVInt32Int32) {
	extends.setParent("extends", (*attr.StrMap)(m))
	(*attr.StrMap)(m).Set(
		"extends",
		extends,
	)
}

func (m *RoomDef) GetExtends1() *KVInt32Str {
	return (*attr.StrMap)(m).Value("extends1").(*KVInt32Str)
}

func (m *RoomDef) SetExtends1(extends *KVInt32Str) {
	extends.setParent("extends1", (*attr.StrMap)(m))
	(*attr.StrMap)(m).Set(
		"extends1",
		extends,
	)
}

func (m *RoomDef) GetExtends2() *KVStrInt32 {
	return (*attr.StrMap)(m).Value("extends2").(*KVStrInt32)
}

func (m *RoomDef) SetExtends2(extends *KVStrInt32) {
	extends.setParent("extends2", (*attr.StrMap)(m))
	(*attr.StrMap)(m).Set(
		"extends2",
		extends,
	)
}

func (m *RoomDef) GetExtends3() *KVStrStr {
	return (*attr.StrMap)(m).Value("extends3").(*KVStrStr)
}

func (m *RoomDef) SetExtends3(extends *KVStrStr) {
	extends.setParent("extends3", (*attr.StrMap)(m))
	(*attr.StrMap)(m).Set(
		"extends3",
		extends,
	)
}

func (m *RoomDef) GetDesk() *DeskDef {
	return (*attr.StrMap)(m).Value("desk").(*DeskDef)
}

func (m *RoomDef) SetDesk(desk *DeskDef) {
	desk.setParent("desk", (*attr.StrMap)(m))
	(*attr.StrMap)(m).Set(
		"desk",
		desk,
	)
}

func (m *RoomDef) GetDesks() *Desks {
	return (*attr.StrMap)(m).Value("desks").(*Desks)
}

func (m *RoomDef) SetDesks(desks *Desks) {
	desks.setParent("desks", (*attr.StrMap)(m))
	(*attr.StrMap)(m).Set(
		"desks",
		desks,
	)
}

func (m *RoomDef) Equal(other *RoomDef) bool {
	return (*attr.StrMap)(m).Equal((*attr.StrMap)(other))
}
