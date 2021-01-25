package demo

import (
	"encoding/json"
	"entitygen/attr"

	"go.mongodb.org/mongo-driver/bson"
)

// KVInt32Int32 外观
type KVInt32Int32 attr.Int32Map

func EmptyKVInt32Int32() *KVInt32Int32 {
	return NewKVInt32Int32(nil)
}

func NewKVInt32Int32(data map[int32]int32) *KVInt32Int32 {
	var convertData map[int32]interface{} = map[int32]interface{}{}
	for k, v := range data {
		convertData[k] = v
	}
	dd := attr.NewInt32Map(convertData)

	return (*KVInt32Int32)(dd)

}

func (m *KVInt32Int32) MarshalJSON() ([]byte, error) {
	return json.Marshal((*attr.Int32Map)(m).ToMap())
}
func (m *KVInt32Int32) UnmarshalJSON(b []byte) error {
	var dd map[int32]int32 = map[int32]int32{}
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

func (m *KVInt32Int32) MarshalBSON() ([]byte, error) {
	return bson.Marshal((*attr.Int32Map)(m).ToMap())
}

func (m *KVInt32Int32) UnmarshalBSON(b []byte) error {
	var dd map[int32]int32 = map[int32]int32{}
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

func (m *KVInt32Int32) Set(k int32, v int32) {
	(*attr.Int32Map)(m).Set(k, v)
}

func (m *KVInt32Int32) Get(key int32) int32 {
	return (*attr.Int32Map)(m).Int32(key)
}

func (m *KVInt32Int32) Delete(key int32) bool {
	return (*attr.Int32Map)(m).Delete(key)
}

func (m *KVInt32Int32) ForEach(fn func(k int32, v int32) bool) {
	(*attr.Int32Map)(m).ForEach(func(k int32, v interface{}) bool {
		return fn(k, v.(int32))
	})
}

func (m *KVInt32Int32) setParent(k string, parent attr.AttrField) {
	(*attr.Int32Map)(m).SetParent(k, parent)
}

func (m *KVInt32Int32) Equal(other *KVInt32Int32) bool {
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

func (m *KVInt32Int32) data() map[int32]int32 {
	var dd map[int32]int32 = map[int32]int32{}
	(*attr.Int32Map)(m).ForEach(func(k int32, v interface{}) bool {
		dd[k] = v.(int32)
		return true
	})
	return dd
}

type KVInt32Str struct {
	_data *attr.Int32Map
}

func EmptyKVInt32Str() *KVInt32Str {
	return NewKVInt32Str(nil)
}

func NewKVInt32Str(data map[int32]string) *KVInt32Str {
	var convertData map[int32]interface{} = map[int32]interface{}{}
	for k, v := range data {
		convertData[k] = v
	}
	dd := attr.NewInt32Map(convertData)

	return &KVInt32Str{
		_data: dd,
	}
}

func (m *KVInt32Str) MarshalJSON() ([]byte, error) {
	return json.Marshal(m._data.ToMap())
}
func (m *KVInt32Str) UnmarshalJSON(b []byte) error {
	var dd map[int32]string = map[int32]string{}
	err := json.Unmarshal(b, &dd)
	if err != nil {
		return err
	}
	var convertData map[int32]interface{} = map[int32]interface{}{}
	for k, v := range dd {
		convertData[k] = v
	}

	m._data = attr.NewInt32Map(convertData)
	return nil
}

func (m *KVInt32Str) MarshalBSON() ([]byte, error) {
	return bson.Marshal(m._data.ToMap())
}

func (m *KVInt32Str) UnmarshalBSON(b []byte) error {
	var dd map[int32]string = map[int32]string{}
	err := bson.Unmarshal(b, &dd)
	if err != nil {
		return err
	}
	var convertData map[int32]interface{} = map[int32]interface{}{}
	for k, v := range dd {
		convertData[k] = v
	}

	m._data = attr.NewInt32Map(convertData)
	return nil
}

func (m *KVInt32Str) Set(k int32, v string) {
	m._data.Set(k, v)
}

func (m *KVInt32Str) Get(key int32) string {
	return m._data.Str(key)
}

func (m *KVInt32Str) Delete(key int32) bool {
	return m._data.Delete(key)
}

func (m *KVInt32Str) ForEach(fn func(k int32, v string) bool) {
	m._data.ForEach(func(k int32, v interface{}) bool {
		return fn(k, v.(string))
	})
}

func (m *KVInt32Str) setParent(k string, parent attr.AttrField) {
	m._data.SetParent(k, parent)
}

func (m *KVInt32Str) Equal(other *KVInt32Str) bool {
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

func (m *KVInt32Str) data() map[int32]string {
	var dd map[int32]string = map[int32]string{}
	m._data.ForEach(func(k int32, v interface{}) bool {
		dd[k] = v.(string)
		return true
	})
	return dd
}

type KVStrInt32 struct {
	_data *attr.StrMap
}

func EmptyStrInt32() *KVStrInt32 {
	return NewKVStrInt32(nil)
}

func NewKVStrInt32(data map[string]int32) *KVStrInt32 {
	var convertData map[string]interface{} = map[string]interface{}{}
	for k, v := range data {
		convertData[k] = v
	}
	dd := attr.NewStrMap(convertData)

	return &KVStrInt32{
		_data: dd,
	}
}

func (m *KVStrInt32) MarshalJSON() ([]byte, error) {
	return json.Marshal(m._data.ToMap())
}
func (m *KVStrInt32) UnmarshalJSON(b []byte) error {
	var dd map[string]int32 = map[string]int32{}
	err := json.Unmarshal(b, &dd)
	if err != nil {
		return err
	}
	var convertData map[string]interface{} = map[string]interface{}{}
	for k, v := range dd {
		convertData[k] = v
	}

	m._data = attr.NewStrMap(convertData)
	return nil
}

func (m *KVStrInt32) MarshalBSON() ([]byte, error) {
	return bson.Marshal(m._data.ToMap())
}

func (m *KVStrInt32) UnmarshalBSON(b []byte) error {
	var dd map[string]int32 = map[string]int32{}
	err := bson.Unmarshal(b, &dd)
	if err != nil {
		return err
	}
	var convertData map[string]interface{} = map[string]interface{}{}
	for k, v := range dd {
		convertData[k] = v
	}

	m._data = attr.NewStrMap(convertData)
	return nil
}

func (m *KVStrInt32) Set(k string, v int32) {
	m._data.Set(k, v)
}

func (m *KVStrInt32) Get(key string) int32 {
	return m._data.Int32(key)
}

func (m *KVStrInt32) Delete(key string) bool {
	return m._data.Delete(key)
}

func (m *KVStrInt32) ForEach(fn func(k string, v int32) bool) {
	m._data.ForEach(func(k string, v interface{}) bool {
		return fn(k, v.(int32))
	})
}

func (m *KVStrInt32) setParent(k string, parent attr.AttrField) {
	m._data.SetParent(k, parent)
}

func (m *KVStrInt32) Equal(other *KVStrInt32) bool {
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

func (m *KVStrInt32) data() map[string]int32 {
	var dd map[string]int32 = map[string]int32{}
	m._data.ForEach(func(k string, v interface{}) bool {
		dd[k] = v.(int32)
		return true
	})
	return dd
}

type KVStrStr struct {
	_data *attr.StrMap
}

func EmptyKVStrStr() *KVStrStr {
	return NewKVStrStr(nil)
}

func NewKVStrStr(data map[string]string) *KVStrStr {
	var convertData map[string]interface{} = map[string]interface{}{}
	for k, v := range data {
		convertData[k] = v
	}
	dd := attr.NewStrMap(convertData)

	return &KVStrStr{
		_data: dd,
	}
}

func (m *KVStrStr) MarshalJSON() ([]byte, error) {
	return json.Marshal(m._data.ToMap())
}
func (m *KVStrStr) UnmarshalJSON(b []byte) error {
	var dd map[string]string = map[string]string{}
	err := json.Unmarshal(b, &dd)
	if err != nil {
		return err
	}
	var convertData map[string]interface{} = map[string]interface{}{}
	for k, v := range dd {
		convertData[k] = v
	}

	m._data = attr.NewStrMap(convertData)
	return nil
}

func (m *KVStrStr) MarshalBSON() ([]byte, error) {
	return bson.Marshal(m._data.ToMap())
}

func (m *KVStrStr) UnmarshalBSON(b []byte) error {
	var dd map[string]string = map[string]string{}
	err := bson.Unmarshal(b, &dd)
	if err != nil {
		return err
	}
	var convertData map[string]interface{} = map[string]interface{}{}
	for k, v := range dd {
		convertData[k] = v
	}

	m._data = attr.NewStrMap(convertData)
	return nil
}

func (m *KVStrStr) Set(k string, v string) {
	m._data.Set(k, v)
}

func (m *KVStrStr) Get(key string) string {
	return m._data.Str(key)
}

func (m *KVStrStr) Delete(key string) bool {
	return m._data.Delete(key)
}

func (m *KVStrStr) ForEach(fn func(k string, v string) bool) {
	m._data.ForEach(func(k string, v interface{}) bool {
		return fn(k, v.(string))
	})
}

func (m *KVStrStr) setParent(k string, parent attr.AttrField) {
	m._data.SetParent(k, parent)
}

func (m *KVStrStr) Equal(other *KVStrStr) bool {
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

func (m *KVStrStr) data() map[string]string {
	var dd map[string]string = map[string]string{}
	m._data.ForEach(func(k string, v interface{}) bool {
		dd[k] = v.(string)
		return true
	})
	return dd
}
