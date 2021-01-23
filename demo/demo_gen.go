package demo

import (
	"encoding/json"
	"entitygen/attr"

	"go.mongodb.org/mongo-driver/bson"
)

type IField interface {
	setRootKey(k string)
	setAncestry(ancestry *attr.AttrMap)
}

// RoomExtends 外观
type RoomExtends struct {
	data *attr.Int32Map
}

func NewRoomExtends(rootKey string, ancestry *attr.AttrMap, data map[int32]int32) *RoomExtends {
	var convertData map[int32]interface{} = map[int32]interface{}{}
	for k, v := range data {
		convertData[k] = v
	}
	dd := attr.NewInt32InterfaceMap(rootKey, ancestry, convertData)

	return &RoomExtends{
		data: dd,
	}
}

func (m *RoomExtends) MarshalJSON() ([]byte, error) {
	return json.Marshal(m.data.ToMap())
}
func (m *RoomExtends) UnmarshalJSON(b []byte) error {
	var dd map[int32]int32 = map[int32]int32{}
	err := json.Unmarshal(b, &dd)
	if err != nil {
		return err
	}
	var convertData map[int32]interface{} = map[int32]interface{}{}
	for k, v := range dd {
		convertData[k] = v
	}

	m.data = attr.NewInt32InterfaceMap("", nil, convertData)
	return nil
}

func (m *RoomExtends) MarshalBSON() ([]byte, error) {
	return bson.Marshal(m.data.ToMap())
}

func (m *RoomExtends) UnmarshalBSON(b []byte) error {
	var dd map[int32]int32 = map[int32]int32{}
	err := bson.Unmarshal(b, &dd)
	if err != nil {
		return err
	}
	var convertData map[int32]interface{} = map[int32]interface{}{}
	for k, v := range dd {
		convertData[k] = v
	}

	m.data = attr.NewInt32InterfaceMap("", nil, convertData)
	return nil
}

func (m *RoomExtends) Set(k int32, v int32) {
	m.data.Set(k, v)
}

func (m *RoomExtends) Get(key int32) int32 {
	return m.data.Int32(key)
}

func (m *RoomExtends) Delete(key int32) bool {
	return m.data.Delete(key)
}

func (m *RoomExtends) ForEach(fn func(k int32, v int32) bool) {
	m.data.ForEach(func(k int32, v interface{}) bool {
		return fn(k, v.(int32))
	})
}

func (m *RoomExtends) setRootKey(k string) {
	m.data.SetRootKey(k)
}

func (m *RoomExtends) setAncestry(ancestry *attr.AttrMap) {
	m.data.SetAncestry(ancestry)
}

func (m *RoomExtends) Equal(other *RoomExtends) bool {
	equal := true
	m.ForEach(func(k, v int32) bool {
		if v != other.Get(k) {
			equal = false
			return false
		}
		return true
	})
	return equal
}

// RoomExtends1 外观1
type RoomExtends1 struct {
	data *attr.Int32Map
}

func NewRoomExtends1(rootKey string, ancestry *attr.AttrMap, data map[int32]string) *RoomExtends1 {
	var convertData map[int32]interface{} = map[int32]interface{}{}
	for k, v := range data {
		convertData[k] = v
	}
	dd := attr.NewInt32InterfaceMap(rootKey, ancestry, convertData)

	return &RoomExtends1{
		data: dd,
	}
}

func (m *RoomExtends1) MarshalJSON() ([]byte, error) {
	return json.Marshal(m.data.ToMap())
}
func (m *RoomExtends1) UnmarshalJSON(b []byte) error {
	var dd map[int32]string = map[int32]string{}
	err := json.Unmarshal(b, &dd)
	if err != nil {
		return err
	}
	var convertData map[int32]interface{} = map[int32]interface{}{}
	for k, v := range dd {
		convertData[k] = v
	}

	m.data = attr.NewInt32InterfaceMap("", nil, convertData)
	return nil
}

func (m *RoomExtends1) MarshalBSON() ([]byte, error) {
	return bson.Marshal(m.data.ToMap())
}

func (m *RoomExtends1) UnmarshalBSON(b []byte) error {
	var dd map[int32]string = map[int32]string{}
	err := bson.Unmarshal(b, &dd)
	if err != nil {
		return err
	}
	var convertData map[int32]interface{} = map[int32]interface{}{}
	for k, v := range dd {
		convertData[k] = v
	}

	m.data = attr.NewInt32InterfaceMap("", nil, convertData)
	return nil
}

func (m *RoomExtends1) Set(k int32, v string) {
	m.data.Set(k, v)
}

func (m *RoomExtends1) Get(key int32) string {
	return m.data.Str(key)
}

func (m *RoomExtends1) Delete(key int32) bool {
	return m.data.Delete(key)
}

func (m *RoomExtends1) ForEach(fn func(k int32, v string) bool) {
	m.data.ForEach(func(k int32, v interface{}) bool {
		return fn(k, v.(string))
	})
}

func (m *RoomExtends1) setRootKey(k string) {
	m.data.SetRootKey(k)
}

func (m *RoomExtends1) setAncestry(ancestry *attr.AttrMap) {
	m.data.SetAncestry(ancestry)
}

func (m *RoomExtends1) Equal(other *RoomExtends1) bool {
	equal := true
	m.ForEach(func(k int32, v string) bool {
		if v != other.Get(k) {
			equal = false
			return false
		}
		return true
	})
	return equal
}

type RoomExtends2 struct {
	data *attr.AttrMap
}

func NewRoomExtends2(rootKey string, ancestry *attr.AttrMap, data map[string]int32) *RoomExtends2 {
	var convertData map[string]interface{} = map[string]interface{}{}
	for k, v := range data {
		convertData[k] = v
	}
	dd := attr.NewStringInterfaceMap(rootKey, ancestry, convertData)

	return &RoomExtends2{
		data: dd,
	}
}

func (m *RoomExtends2) MarshalJSON() ([]byte, error) {
	return json.Marshal(m.data.ToMap())
}
func (m *RoomExtends2) UnmarshalJSON(b []byte) error {
	var dd map[string]int32 = map[string]int32{}
	err := json.Unmarshal(b, &dd)
	if err != nil {
		return err
	}
	var convertData map[string]interface{} = map[string]interface{}{}
	for k, v := range dd {
		convertData[k] = v
	}

	m.data = attr.NewStringInterfaceMap("", nil, convertData)
	return nil
}

func (m *RoomExtends2) MarshalBSON() ([]byte, error) {
	return bson.Marshal(m.data.ToMap())
}

func (m *RoomExtends2) UnmarshalBSON(b []byte) error {
	var dd map[string]int32 = map[string]int32{}
	err := bson.Unmarshal(b, &dd)
	if err != nil {
		return err
	}
	var convertData map[string]interface{} = map[string]interface{}{}
	for k, v := range dd {
		convertData[k] = v
	}

	m.data = attr.NewStringInterfaceMap("", nil, convertData)
	return nil
}

func (m *RoomExtends2) Set(k string, v int32) {
	m.data.Set(k, v)
}

func (m *RoomExtends2) Get(key string) int32 {
	return m.data.Int32(key)
}

func (m *RoomExtends2) Delete(key string) bool {
	return m.data.Delete(key)
}

func (m *RoomExtends2) ForEach(fn func(k string, v int32) bool) {
	m.data.ForEach(func(k string, v interface{}) bool {
		return fn(k, v.(int32))
	})
}

func (m *RoomExtends2) setRootKey(k string) {
	m.data.SetRootKey(k)
}

func (m *RoomExtends2) setAncestry(ancestry *attr.AttrMap) {
	m.data.SetAncestry(ancestry)
}

func (m *RoomExtends2) Equal(other *RoomExtends2) bool {
	equal := true
	m.ForEach(func(k string, v int32) bool {
		if v != other.Get(k) {
			equal = false
			return false
		}
		return true
	})
	return equal
}

type RoomExtends3 struct {
	data *attr.AttrMap
}

func NewRoomExtends3(rootKey string, ancestry *attr.AttrMap, data map[string]string) *RoomExtends3 {
	var convertData map[string]interface{} = map[string]interface{}{}
	for k, v := range data {
		convertData[k] = v
	}
	dd := attr.NewStringInterfaceMap(rootKey, ancestry, convertData)

	return &RoomExtends3{
		data: dd,
	}
}

func (m *RoomExtends3) MarshalJSON() ([]byte, error) {
	return json.Marshal(m.data.ToMap())
}
func (m *RoomExtends3) UnmarshalJSON(b []byte) error {
	var dd map[string]string = map[string]string{}
	err := json.Unmarshal(b, &dd)
	if err != nil {
		return err
	}
	var convertData map[string]interface{} = map[string]interface{}{}
	for k, v := range dd {
		convertData[k] = v
	}

	m.data = attr.NewStringInterfaceMap("", nil, convertData)
	return nil
}

func (m *RoomExtends3) MarshalBSON() ([]byte, error) {
	return bson.Marshal(m.data.ToMap())
}

func (m *RoomExtends3) UnmarshalBSON(b []byte) error {
	var dd map[string]string = map[string]string{}
	err := bson.Unmarshal(b, &dd)
	if err != nil {
		return err
	}
	var convertData map[string]interface{} = map[string]interface{}{}
	for k, v := range dd {
		convertData[k] = v
	}

	m.data = attr.NewStringInterfaceMap("", nil, convertData)
	return nil
}

func (m *RoomExtends3) Set(k string, v string) {
	m.data.Set(k, v)
}

func (m *RoomExtends3) Get(key string) string {
	return m.data.Str(key)
}

func (m *RoomExtends3) Delete(key string) bool {
	return m.data.Delete(key)
}

func (m *RoomExtends3) ForEach(fn func(k string, v string) bool) {
	m.data.ForEach(func(k string, v interface{}) bool {
		return fn(k, v.(string))
	})
}

func (m *RoomExtends3) setRootKey(k string) {
	m.data.SetRootKey(k)
}

func (m *RoomExtends3) setAncestry(ancestry *attr.AttrMap) {
	m.data.SetAncestry(ancestry)
}

func (m *RoomExtends3) Equal(other *RoomExtends3) bool {
	equal := true
	m.ForEach(func(k string, v string) bool {
		if v != other.Get(k) {
			equal = false
			return false
		}
		return true
	})
	return equal
}

var room *attr.DataDef

func init() {
	room = &attr.DataDef{}

	room.DefAttr("csv_pos", attr.Int32Attr, attr.AfBase, true)
	room.DefAttr("build_id", attr.StringAttr, attr.AfBase, true)
	room.DefAttr("extends", &RoomExtends{}, attr.AfBase, true)
	room.DefAttr("extends1", &RoomExtends1{}, attr.AfBase, true)
	room.DefAttr("extends2", &RoomExtends2{}, attr.AfBase, true)
	room.DefAttr("extends3", &RoomExtends3{}, attr.AfBase, true)

}

type Room struct {
	// parent    AttrMapImp
	// parentKey string
	data *attr.AttrMap
}

func NewRoom(
	csvPos int32,
	buildID string,
	extends map[int32]int32,
	extends1 map[int32]string,
	extends2 map[string]int32,
	extends3 map[string]string,
) *Room {
	m := &Room{}
	m.data = attr.NewAttrMap()

	m.SetCsvPos(csvPos)
	m.SetBuildID(buildID)
	m.SetExtends(extends)
	m.SetExtends1(extends1)
	m.SetExtends2(extends2)
	m.SetExtends3(extends3)

	m.data.ClearChangeKey()
	return m
}

func (m *Room) MarshalJSON() ([]byte, error) {
	return json.Marshal(m.data.ToMap())
}
func (m *Room) UnmarshalJSON(b []byte) error {
	mm, err := room.UnmarshalJson(b)
	if err != nil {
		return err
	}
	m.data.SetData(mm)
	m.data.ForEach(func(k string, v interface{}) bool {
		if k != "id" && !room.GetDef(k).IsPrimary() {
			v.(IField).setRootKey(k)
			v.(IField).setAncestry(m.data)
		}
		return true
	})
	return nil
}

func (m *Room) MarshalBSON() ([]byte, error) {
	return bson.Marshal(m.data.ToMap())
}

func (m *Room) UnmarshalBSON(b []byte) error {
	mm, err := room.UnmarshalBson(b)
	if err != nil {
		return err
	}
	m.data.SetData(mm)
	return nil
}

func (m *Room) InitAttrMap() {
	m.data = attr.NewAttrMap()
}

func (m *Room) ForEach(fn func(s string, v interface{}) bool) {
	m.data.ForEach(fn)
}

func (m *Room) GetBuildID() string {
	return m.data.Str("build_id")
}

func (m *Room) SetBuildID(v string) {
	m.data.Set("build_id", v)
}

func (m *Room) GetCsvPos() int32 {
	return m.data.Int32("csv_pos")
}

func (m *Room) SetCsvPos(v int32) {
	m.data.Set("csv_pos", v)
}

func (m *Room) GetExtends() *RoomExtends {
	return m.data.Value("extends").(*RoomExtends)
}

func (m *Room) SetExtends(extends map[int32]int32) {
	m.data.Set(
		"extends",
		NewRoomExtends("extends", m.data, extends),
	)
}

func (m *Room) GetExtends1() *RoomExtends1 {
	return m.data.Value("extends1").(*RoomExtends1)
}

func (m *Room) SetExtends1(extends map[int32]string) {
	m.data.Set(
		"extends1",
		NewRoomExtends1("extends1", m.data, extends),
	)
}

func (m *Room) GetExtends2() *RoomExtends2 {
	return m.data.Value("extends2").(*RoomExtends2)
}

func (m *Room) SetExtends2(extends map[string]int32) {
	m.data.Set(
		"extends2",
		NewRoomExtends2("extends2", m.data, extends),
	)
}

func (m *Room) GetExtends3() *RoomExtends3 {
	return m.data.Value("extends3").(*RoomExtends3)
}

func (m *Room) SetExtends3(extends map[string]string) {
	m.data.Set(
		"extends3",
		NewRoomExtends3("extends3", m.data, extends),
	)
}

func (m *Room) setRootKey(k string) {
	m.data.SetRootKey(k)
}

func (m *Room) setAncestry(ancestry *attr.AttrMap) {
	m.data.SetAncestry(ancestry)
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
