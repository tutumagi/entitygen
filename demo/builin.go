package demo

import (
	"encoding/json"
	"entitygen/attr"

	"go.mongodb.org/mongo-driver/bson"
)

// RoomExtends 外观
type RoomExtends struct {
	_data *attr.Int32Map
}

func EmptyRoomExtends() *RoomExtends {
	return NewRoomExtends(nil)
}

func NewRoomExtends(data map[int32]int32) *RoomExtends {
	var convertData map[int32]interface{} = map[int32]interface{}{}
	for k, v := range data {
		convertData[k] = v
	}
	dd := attr.NewInt32InterfaceMap(convertData)

	return &RoomExtends{
		_data: dd,
	}
}

func (m *RoomExtends) MarshalJSON() ([]byte, error) {
	return json.Marshal(m._data.ToMap())
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

	m._data = attr.NewInt32InterfaceMap(convertData)
	return nil
}

func (m *RoomExtends) MarshalBSON() ([]byte, error) {
	return bson.Marshal(m._data.ToMap())
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

	m._data = attr.NewInt32InterfaceMap(convertData)
	return nil
}

func (m *RoomExtends) Set(k int32, v int32) {
	m._data.Set(k, v)
}

func (m *RoomExtends) Get(key int32) int32 {
	return m._data.Int32(key)
}

func (m *RoomExtends) Delete(key int32) bool {
	return m._data.Delete(key)
}

func (m *RoomExtends) ForEach(fn func(k int32, v int32) bool) {
	m._data.ForEach(func(k int32, v interface{}) bool {
		return fn(k, v.(int32))
	})
}

func (m *RoomExtends) setParent(k string, parent attr.AttrField) {
	m._data.SetParent(k, parent)
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

func (m *RoomExtends) data() map[int32]int32 {
	var dd map[int32]int32 = map[int32]int32{}
	m._data.ForEach(func(k int32, v interface{}) bool {
		dd[k] = v.(int32)
		return true
	})
	return dd
}

// RoomExtends1 外观1
type RoomExtends1 struct {
	_data *attr.Int32Map
}

func EmptyRoomExtends1() *RoomExtends1 {
	return NewRoomExtends1(nil)
}

func NewRoomExtends1(data map[int32]string) *RoomExtends1 {
	var convertData map[int32]interface{} = map[int32]interface{}{}
	for k, v := range data {
		convertData[k] = v
	}
	dd := attr.NewInt32InterfaceMap(convertData)

	return &RoomExtends1{
		_data: dd,
	}
}

func (m *RoomExtends1) MarshalJSON() ([]byte, error) {
	return json.Marshal(m._data.ToMap())
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

	m._data = attr.NewInt32InterfaceMap(convertData)
	return nil
}

func (m *RoomExtends1) MarshalBSON() ([]byte, error) {
	return bson.Marshal(m._data.ToMap())
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

	m._data = attr.NewInt32InterfaceMap(convertData)
	return nil
}

func (m *RoomExtends1) Set(k int32, v string) {
	m._data.Set(k, v)
}

func (m *RoomExtends1) Get(key int32) string {
	return m._data.Str(key)
}

func (m *RoomExtends1) Delete(key int32) bool {
	return m._data.Delete(key)
}

func (m *RoomExtends1) ForEach(fn func(k int32, v string) bool) {
	m._data.ForEach(func(k int32, v interface{}) bool {
		return fn(k, v.(string))
	})
}

func (m *RoomExtends1) setParent(k string, parent attr.AttrField) {
	m._data.SetParent(k, parent)
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

func (m *RoomExtends1) data() map[int32]string {
	var dd map[int32]string = map[int32]string{}
	m._data.ForEach(func(k int32, v interface{}) bool {
		dd[k] = v.(string)
		return true
	})
	return dd
}

type RoomExtends2 struct {
	_data *attr.AttrMap
}

func EmptyRoomExtends2() *RoomExtends2 {
	return NewRoomExtends2(nil)
}

func NewRoomExtends2(data map[string]int32) *RoomExtends2 {
	var convertData map[string]interface{} = map[string]interface{}{}
	for k, v := range data {
		convertData[k] = v
	}
	dd := attr.NewStringInterfaceMap(convertData)

	return &RoomExtends2{
		_data: dd,
	}
}

func (m *RoomExtends2) MarshalJSON() ([]byte, error) {
	return json.Marshal(m._data.ToMap())
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

	m._data = attr.NewStringInterfaceMap(convertData)
	return nil
}

func (m *RoomExtends2) MarshalBSON() ([]byte, error) {
	return bson.Marshal(m._data.ToMap())
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

	m._data = attr.NewStringInterfaceMap(convertData)
	return nil
}

func (m *RoomExtends2) Set(k string, v int32) {
	m._data.Set(k, v)
}

func (m *RoomExtends2) Get(key string) int32 {
	return m._data.Int32(key)
}

func (m *RoomExtends2) Delete(key string) bool {
	return m._data.Delete(key)
}

func (m *RoomExtends2) ForEach(fn func(k string, v int32) bool) {
	m._data.ForEach(func(k string, v interface{}) bool {
		return fn(k, v.(int32))
	})
}

func (m *RoomExtends2) setParent(k string, parent attr.AttrField) {
	m._data.SetParent(k, parent)
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

func (m *RoomExtends2) data() map[string]int32 {
	var dd map[string]int32 = map[string]int32{}
	m._data.ForEach(func(k string, v interface{}) bool {
		dd[k] = v.(int32)
		return true
	})
	return dd
}

type RoomExtends3 struct {
	_data *attr.AttrMap
}

func EmptyRoomExtends3() *RoomExtends3 {
	return NewRoomExtends3(nil)
}

func NewRoomExtends3(data map[string]string) *RoomExtends3 {
	var convertData map[string]interface{} = map[string]interface{}{}
	for k, v := range data {
		convertData[k] = v
	}
	dd := attr.NewStringInterfaceMap(convertData)

	return &RoomExtends3{
		_data: dd,
	}
}

func (m *RoomExtends3) MarshalJSON() ([]byte, error) {
	return json.Marshal(m._data.ToMap())
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

	m._data = attr.NewStringInterfaceMap(convertData)
	return nil
}

func (m *RoomExtends3) MarshalBSON() ([]byte, error) {
	return bson.Marshal(m._data.ToMap())
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

	m._data = attr.NewStringInterfaceMap(convertData)
	return nil
}

func (m *RoomExtends3) Set(k string, v string) {
	m._data.Set(k, v)
}

func (m *RoomExtends3) Get(key string) string {
	return m._data.Str(key)
}

func (m *RoomExtends3) Delete(key string) bool {
	return m._data.Delete(key)
}

func (m *RoomExtends3) ForEach(fn func(k string, v string) bool) {
	m._data.ForEach(func(k string, v interface{}) bool {
		return fn(k, v.(string))
	})
}

func (m *RoomExtends3) setParent(k string, parent attr.AttrField) {
	m._data.SetParent(k, parent)
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

func (m *RoomExtends3) data() map[string]string {
	var dd map[string]string = map[string]string{}
	m._data.ForEach(func(k string, v interface{}) bool {
		dd[k] = v.(string)
		return true
	})
	return dd
}
