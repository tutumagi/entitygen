// Code generated by generator, DO NOT EDIT.
package demodef

import (
	"encoding/json"
	attr "gitlab.testkaka.com/usm/game/entitygen/attr"
	bson "go.mongodb.org/mongo-driver/bson"
)

type Int8Slice attr.Slice

func EmptyInt8Slice() *Int8Slice {
	return NewInt8Slice(nil)
}
func NewInt8Slice(items []int8) *Int8Slice {
	var convertData []interface{} = []interface{}{}
	for _, v := range items {
		convertData = append(convertData, v)
	}
	return (*Int8Slice)(attr.NewSlice(convertData))
}
func CopyInt8Slice(value *Int8Slice) *Int8Slice {
	if value == nil {
		return nil
	}
	a := EmptyInt8Slice()
	value.ForEach(func(_ int, v int8) bool {
		a.Add(v)
		return true
	})
	return a
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
func (a *Int8Slice) SetParent(k string, parent attr.Field) {
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
func (a *Int8Slice) Undertype() interface{} {
	return (*attr.Slice)(a)
}
func (a *Int8Slice) Data() []int8 {
	dd := []int8{}
	a.ForEach(func(idx int, v int8) bool {
		dd = append(dd, v)
		return true
	})
	return dd
}
func (a *Int8Slice) MarshalJSON() ([]byte, error) {
	return json.Marshal(map[string][]int8{
		"d": a.Data(),
	})
}
func (a *Int8Slice) UnmarshalJSON(b []byte) error {
	dd := map[string][]int8{}
	err := json.Unmarshal(b, &dd)
	if err != nil {
		return err
	}
	convertData := []interface{}{}
	for k, v := range dd["d"] {
		_ = k
		convertData = append(convertData, v)
	}
	(*attr.Slice)(a).SetData(convertData)
	return nil
}
func (a *Int8Slice) MarshalBSON() ([]byte, error) {
	return bson.Marshal(map[string][]int8{
		"d": a.Data(),
	})
}
func (a *Int8Slice) UnmarshalBSON(b []byte) error {
	dd := map[string][]int8{}
	err := bson.Unmarshal(b, &dd)
	if err != nil {
		return err
	}
	convertData := []interface{}{}
	for k, v := range dd["d"] {
		_ = k
		convertData = append(convertData, v)
	}
	(*attr.Slice)(a).SetData(convertData)
	return nil
}
