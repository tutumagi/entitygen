// Code generated by generator, DO NOT EDIT.
package demodef

import (
	"encoding/json"
	"fmt"
	attr "gitlab.testkaka.com/usm/game/entitygen/attr"
	bson "go.mongodb.org/mongo-driver/bson"
)

type Vector2Slice attr.Slice

func EmptyVector2Slice() *Vector2Slice {
	return NewVector2Slice(nil)
}
func NewVector2Slice(items []*Vector2) *Vector2Slice {
	var convertData []interface{} = []interface{}{}
	for _, v := range items {
		convertData = append(convertData, v)
	}
	return (*Vector2Slice)(attr.NewSlice(convertData))
}
func CopyVector2Slice(value *Vector2Slice) *Vector2Slice {
	if value == nil {
		return nil
	}
	a := EmptyVector2Slice()
	value.ForEach(func(_ int, v *Vector2) bool {
		a.Add(CopyVector2(v))
		return true
	})
	return a
}
func (a *Vector2Slice) Set(idx int, item *Vector2) {
	item.SetParent(fmt.Sprintf("ik%d", idx), (*attr.Slice)(a))
	(*attr.Slice)(a).Set(idx, item)
}
func (a *Vector2Slice) Add(item *Vector2) {
	idx := a.Count()
	item.SetParent(fmt.Sprintf("ik%d", idx), (*attr.Slice)(a))
	(*attr.Slice)(a).Add(item)
}
func (a *Vector2Slice) At(idx int) *Vector2 {
	val := (*attr.Slice)(a).Value(idx)
	if val == nil {
		return nil
	}
	return val.(*Vector2)
}
func (a *Vector2Slice) DelAt(idx int) bool {
	return (*attr.Slice)(a).DeleteAt(idx)
}
func (a *Vector2Slice) Count() int {
	return (*attr.Slice)(a).Len()
}
func (a *Vector2Slice) SetParent(k string, parent attr.Field) {
	(*attr.Slice)(a).SetParent(k, parent)
}
func (a *Vector2Slice) ForEach(fn func(k int, v *Vector2) bool) {
	(*attr.Slice)(a).ForEach(func(k int, v interface{}) bool {
		return fn(k, v.(*Vector2))
	})
}
func (a *Vector2Slice) Equal(other *Vector2Slice) bool {
	return (*attr.Slice)(a).Equal((*attr.Slice)(other))
}
func (a *Vector2Slice) Undertype() interface{} {
	return (*attr.Slice)(a)
}
func (a *Vector2Slice) Data() []*Vector2 {
	dd := []*Vector2{}
	a.ForEach(func(idx int, v *Vector2) bool {
		dd = append(dd, v)
		return true
	})
	return dd
}
func (a *Vector2Slice) MarshalJSON() ([]byte, error) {
	return json.Marshal(map[string][]*Vector2{
		"d": a.Data(),
	})
}
func (a *Vector2Slice) UnmarshalJSON(b []byte) error {
	dd := map[string][]*Vector2{}
	err := json.Unmarshal(b, &dd)
	if err != nil {
		return err
	}
	convertData := []interface{}{}
	for k, v := range dd["d"] {
		v.SetParent(fmt.Sprintf("ik%d", k), (*attr.Slice)(a))
		convertData = append(convertData, v)
	}
	(*attr.Slice)(a).SetData(convertData)
	return nil
}
func (a *Vector2Slice) MarshalBSON() ([]byte, error) {
	return bson.Marshal(map[string][]*Vector2{
		"d": a.Data(),
	})
}
func (a *Vector2Slice) UnmarshalBSON(b []byte) error {
	dd := map[string][]*Vector2{}
	err := bson.Unmarshal(b, &dd)
	if err != nil {
		return err
	}
	convertData := []interface{}{}
	for k, v := range dd["d"] {
		v.SetParent(fmt.Sprintf("ik%d", k), (*attr.Slice)(a))
		convertData = append(convertData, v)
	}
	(*attr.Slice)(a).SetData(convertData)
	return nil
}
