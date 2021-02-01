// Code generated by generator, DO NOT EDIT.
package entitydef

import (
	"encoding/json"
	attr "gitlab.gamesword.com/nut/entitygen/attr"
	bson "go.mongodb.org/mongo-driver/bson"
)

type Int8Slice attr.Slice

func EmptyInt8Slice() *Int8Slice {
	return NewInt8Slice(nil)
}
func NewInt8Slice(items []int8) *Int8Slice {
	var convertData []interface{} = []interface{}{}
	for k, v := range items {
		convertData[k] = v
	}
	return (*Int8Slice)(attr.NewSlice(convertData))
}
func (a *Int8Slice) Set(idx int, item int8) {
	(*attr.Slice)(a).Set(idx, item)
}
func (a *Int8Slice) Add(item int8) {
	(*attr.Slice)(a).Add(item)
}
func (a *Int8Slice) At(idx int) int8 {
	val := (*attr.Slice)(a).Int8(idx)
	return val
}
func (a *Int8Slice) DelAt(idx int) bool {
	return (*attr.Slice)(a).DeleteAt(idx)
}
func (a *Int8Slice) Count() int {
	return (*attr.Slice)(a).Len()
}
func (a *Int8Slice) setParent(k string, parent attr.Field) {
	(*attr.Slice)(a).SetParent(k, parent)
}
func (a *Int8Slice) ForEach(fn func(k int, v int8) bool) {
	(*attr.Slice)(a).ForEach(func(k int, v interface{}) bool {
		return fn(k, v.(int8))
	})
}
func (a *Int8Slice) Equal(other *Int8Slice) bool {
	return (*attr.Slice)(a).Equal((*attr.Slice)(other))
}
func (a *Int8Slice) data() []int8 {
	dd := []int8{}
	a.ForEach(func(idx int, v int8) bool {
		dd = append(dd, v)
		return true
	})
	return dd
}
func (a *Int8Slice) MarshalJSON() ([]byte, error) {
	return json.Marshal(a.data())
}
func (a *Int8Slice) UnmarshalJSON(b []byte) error {
	dd := []int8{}
	err := json.Unmarshal(b, &dd)
	if err != nil {
		return err
	}
	convertData := []interface{}{}
	for k, v := range dd {
		convertData[k] = v
	}
	(*attr.Slice)(a).SetData(convertData)
	return nil
}
func (a *Int8Slice) MarshalBSON() ([]byte, error) {
	return bson.Marshal(a.data())
}
func (a *Int8Slice) UnmarshalBSON(b []byte) error {
	dd := []int8{}
	err := bson.Unmarshal(b, &dd)
	if err != nil {
		return err
	}
	convertData := []interface{}{}
	for k, v := range dd {
		convertData[k] = v
	}
	(*attr.Slice)(a).SetData(convertData)
	return nil
}
