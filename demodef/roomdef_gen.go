// Code generated by generator, DO NOT EDIT.
package demodef

import (
	"encoding/json"
	attr "gitlab.gamesword.com/nut/entitygen/attr"
	bson "go.mongodb.org/mongo-driver/bson"
)

var RoomDefMeta *attr.Meta

func init() {
	RoomDefMeta = attr.NewMeta(func() interface{} {
		return EmptyRoomDef()
	}, func() interface{} {
		return &[]*RoomDef{}
	})

	RoomDefMeta.DefAttr("csvPos", attr.Int32, attr.AfOtherClients, true)
	RoomDefMeta.DefAttr("buildID", attr.String, attr.AfOtherClients, true)
	RoomDefMeta.DefAttr("extends", &KVInt32Int32{}, attr.AfOtherClients, true)
	RoomDefMeta.DefAttr("extends1", &KVInt32Str{}, attr.AfOtherClients, true)
	RoomDefMeta.DefAttr("extends2", &KVStrInt32{}, attr.AfOtherClients, true)
	RoomDefMeta.DefAttr("extends3", &KVStrStr{}, attr.AfOtherClients, true)
	RoomDefMeta.DefAttr("desk111", &Desk{}, attr.AfOtherClients, true)
	RoomDefMeta.DefAttr("desks222", &KVInt32Desk{}, attr.AfOtherClients, true)
	RoomDefMeta.DefAttr("desks333", &KVStrDesk{}, attr.AfOtherClients, true)
	RoomDefMeta.DefAttr("desksArr", &DeskSlice{}, attr.AfOtherClients, true)
	RoomDefMeta.DefAttr("int8ss", &Int8Slice{}, attr.AfOtherClients, true)
	RoomDefMeta.DefAttr("vec3", &Vector3{}, attr.AfOtherClients, true)
	RoomDefMeta.DefAttr("vec2Arr", &Vector2Slice{}, attr.AfOtherClients, true)
	// 实体内置的属性
	// 实体内置的 ID
	RoomDefMeta.DefAttr("id", attr.String, attr.AfOtherClients, true)
	// 实体内置的 位置
	RoomDefMeta.DefAttr("pos", attr.Vector3, attr.AfCell, true)
	// 实体内置的 朝向
	RoomDefMeta.DefAttr("rot", attr.Vector3, attr.AfCell, true)
}

type RoomDef attr.StrMap

func EmptyRoomDef() *RoomDef {
	return NewRoomDef(0, "", nil, nil, nil, nil, nil, nil, nil, nil, nil, nil, nil)
}
func NewRoomDef(csvPos int32, buildID string, extends *KVInt32Int32, extends1 *KVInt32Str, extends2 *KVStrInt32, extends3 *KVStrStr, desk111 *Desk, desks222 *KVInt32Desk, desks333 *KVStrDesk, desksArr *DeskSlice, int8ss *Int8Slice, vec3 *Vector3, vec2Arr *Vector2Slice) *RoomDef {
	m := (*RoomDef)(attr.NewStrMap(nil))
	m.SetCsvPos(csvPos)
	m.SetBuildID(buildID)
	m.SetExtends(extends)
	m.SetExtends1(extends1)
	m.SetExtends2(extends2)
	m.SetExtends3(extends3)
	m.SetDesk111(desk111)
	m.SetDesks222(desks222)
	m.SetDesks333(desks333)
	m.SetDesksArr(desksArr)
	m.SetInt8ss(int8ss)
	m.SetVec3(vec3)
	m.SetVec2Arr(vec2Arr)
	m.ClearChangeKey()
	// 实体内置的属性
	// 实体内置的 ID
	m.SetId("")
	// 实体内置的 位置
	m.SetPos(attr.EmptyVec3())
	// 实体内置的 朝向
	m.SetRot(attr.EmptyVec3())
	return m
}
func (a *RoomDef) GetCsvPos() int32 {
	val := (*attr.StrMap)(a).Int32("csvPos")
	return val
}
func (a *RoomDef) SetCsvPos(csvPos int32) {
	(*attr.StrMap)(a).Set("csvPos", csvPos)
}

func (a *RoomDef) GetBuildID() string {
	val := (*attr.StrMap)(a).Str("buildID")
	return val
}
func (a *RoomDef) SetBuildID(buildID string) {
	(*attr.StrMap)(a).Set("buildID", buildID)
}

func (a *RoomDef) GetExtends() *KVInt32Int32 {
	val := (*attr.StrMap)(a).Value("extends")
	if val == nil {
		return nil
	}
	return val.(*KVInt32Int32)
}
func (a *RoomDef) SetExtends(extends *KVInt32Int32) {
	extends.SetParent("extends", (*attr.StrMap)(a))
	(*attr.StrMap)(a).Set("extends", extends)
}

func (a *RoomDef) GetExtends1() *KVInt32Str {
	val := (*attr.StrMap)(a).Value("extends1")
	if val == nil {
		return nil
	}
	return val.(*KVInt32Str)
}
func (a *RoomDef) SetExtends1(extends1 *KVInt32Str) {
	extends1.SetParent("extends1", (*attr.StrMap)(a))
	(*attr.StrMap)(a).Set("extends1", extends1)
}

func (a *RoomDef) GetExtends2() *KVStrInt32 {
	val := (*attr.StrMap)(a).Value("extends2")
	if val == nil {
		return nil
	}
	return val.(*KVStrInt32)
}
func (a *RoomDef) SetExtends2(extends2 *KVStrInt32) {
	extends2.SetParent("extends2", (*attr.StrMap)(a))
	(*attr.StrMap)(a).Set("extends2", extends2)
}

func (a *RoomDef) GetExtends3() *KVStrStr {
	val := (*attr.StrMap)(a).Value("extends3")
	if val == nil {
		return nil
	}
	return val.(*KVStrStr)
}
func (a *RoomDef) SetExtends3(extends3 *KVStrStr) {
	extends3.SetParent("extends3", (*attr.StrMap)(a))
	(*attr.StrMap)(a).Set("extends3", extends3)
}

func (a *RoomDef) GetDesk111() *Desk {
	val := (*attr.StrMap)(a).Value("desk111")
	if val == nil {
		return nil
	}
	return val.(*Desk)
}
func (a *RoomDef) SetDesk111(desk111 *Desk) {
	desk111.SetParent("desk111", (*attr.StrMap)(a))
	(*attr.StrMap)(a).Set("desk111", desk111)
}

func (a *RoomDef) GetDesks222() *KVInt32Desk {
	val := (*attr.StrMap)(a).Value("desks222")
	if val == nil {
		return nil
	}
	return val.(*KVInt32Desk)
}
func (a *RoomDef) SetDesks222(desks222 *KVInt32Desk) {
	desks222.SetParent("desks222", (*attr.StrMap)(a))
	(*attr.StrMap)(a).Set("desks222", desks222)
}

func (a *RoomDef) GetDesks333() *KVStrDesk {
	val := (*attr.StrMap)(a).Value("desks333")
	if val == nil {
		return nil
	}
	return val.(*KVStrDesk)
}
func (a *RoomDef) SetDesks333(desks333 *KVStrDesk) {
	desks333.SetParent("desks333", (*attr.StrMap)(a))
	(*attr.StrMap)(a).Set("desks333", desks333)
}

func (a *RoomDef) GetDesksArr() *DeskSlice {
	val := (*attr.StrMap)(a).Value("desksArr")
	if val == nil {
		return nil
	}
	return val.(*DeskSlice)
}
func (a *RoomDef) SetDesksArr(desksArr *DeskSlice) {
	desksArr.SetParent("desksArr", (*attr.StrMap)(a))
	(*attr.StrMap)(a).Set("desksArr", desksArr)
}

func (a *RoomDef) GetInt8ss() *Int8Slice {
	val := (*attr.StrMap)(a).Value("int8ss")
	if val == nil {
		return nil
	}
	return val.(*Int8Slice)
}
func (a *RoomDef) SetInt8ss(int8ss *Int8Slice) {
	int8ss.SetParent("int8ss", (*attr.StrMap)(a))
	(*attr.StrMap)(a).Set("int8ss", int8ss)
}

func (a *RoomDef) GetVec3() *Vector3 {
	val := (*attr.StrMap)(a).Value("vec3")
	if val == nil {
		return nil
	}
	return val.(*Vector3)
}
func (a *RoomDef) SetVec3(vec3 *Vector3) {
	vec3.SetParent("vec3", (*attr.StrMap)(a))
	(*attr.StrMap)(a).Set("vec3", vec3)
}

func (a *RoomDef) GetVec2Arr() *Vector2Slice {
	val := (*attr.StrMap)(a).Value("vec2Arr")
	if val == nil {
		return nil
	}
	return val.(*Vector2Slice)
}
func (a *RoomDef) SetVec2Arr(vec2Arr *Vector2Slice) {
	vec2Arr.SetParent("vec2Arr", (*attr.StrMap)(a))
	(*attr.StrMap)(a).Set("vec2Arr", vec2Arr)
}

func (a *RoomDef) GetRot() *attr.Vec3 {
	val := (*attr.StrMap)(a).Value("rot")
	if val == nil {
		return nil
	}
	return val.(*attr.Vec3)
}
func (a *RoomDef) SetRot(rot *attr.Vec3) {
	rot.SetParent("rot", (*attr.StrMap)(a))
	(*attr.StrMap)(a).Set("rot", rot)
}
func (a *RoomDef) GetPos() *attr.Vec3 {
	val := (*attr.StrMap)(a).Value("pos")
	if val == nil {
		return nil
	}
	return val.(*attr.Vec3)
}
func (a *RoomDef) SetPos(pos *attr.Vec3) {
	pos.SetParent("pos", (*attr.StrMap)(a))
	(*attr.StrMap)(a).Set("pos", pos)
}
func (a *RoomDef) GetId() string {
	return (*attr.StrMap)(a).Str("id")
}
func (a *RoomDef) SetId(id string) {
	(*attr.StrMap)(a).Set("id", id)
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
func (a *RoomDef) SetParent(k string, parent attr.Field) {
	(*attr.StrMap)(a).SetParent(k, parent)
}
func (a *RoomDef) ForEach(fn func(k string, v interface{}) bool) {
	(*attr.StrMap)(a).ForEach(fn)
}
func (a *RoomDef) Equal(other *RoomDef) bool {
	return (*attr.StrMap)(a).Equal((*attr.StrMap)(other))
}
func (a *RoomDef) Undertype() interface{} {
	return (*attr.StrMap)(a)
}
func (a *RoomDef) Data() map[string]interface{} {
	dd := map[string]interface{}{}
	a.ForEach(func(k string, v interface{}) bool {
		dd[k] = v
		return true
	})
	return dd
}
func (a *RoomDef) MarshalJSON() ([]byte, error) {
	return json.Marshal((*attr.StrMap)(a).ToMap())
}
func (a *RoomDef) UnmarshalJSON(b []byte) error {
	_, err := RoomDefMeta.UnmarshalJson(b, (*attr.StrMap)(a))
	if err != nil {
		return err
	}
	return nil
}
func (a *RoomDef) MarshalBSON() ([]byte, error) {
	return bson.Marshal((*attr.StrMap)(a).ToMap())
}
func (a *RoomDef) UnmarshalBSON(b []byte) error {
	_, err := RoomDefMeta.UnmarshalBson(b, (*attr.StrMap)(a))
	if err != nil {
		return err
	}
	return nil
}
