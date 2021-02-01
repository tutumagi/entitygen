// Code generated by generator, DO NOT EDIT.
package entitydef

import (
	"encoding/json"
	attr "gitlab.gamesword.com/nut/entitygen/attr"
	bson "go.mongodb.org/mongo-driver/bson"
)

type KVInt32Int32 attr.Int32Map

func EmptyKVInt32Int32() *KVInt32Int32 {
	return NewKVInt32Int32(nil)
}
func NewKVInt32Int32(data map[int32]int32) *KVInt32Int32 {
	var convertData map[int32]interface{} = map[int32]interface{}{}
	for k, v := range data {
		convertData[k] = v
	}
	return (*KVInt32Int32)(attr.NewInt32Map(convertData))
}
func (a *KVInt32Int32) Set(k int32, v int32) {
	(*attr.Int32Map)(a).Set(k, v)
}
func (a *KVInt32Int32) Get(k int32) int32 {
	return (*attr.Int32Map)(a).Int32(k)
}
func (a *KVInt32Int32) Delete(k int32) bool {
	return (*attr.Int32Map)(a).Delete(k)
}
func (a *KVInt32Int32) setParent(k string, parent attr.Field) {
	(*attr.Int32Map)(a).SetParent(k, parent)
}
func (a *KVInt32Int32) ForEach(fn func(k int32, v int32) bool) {
	(*attr.Int32Map)(a).ForEach(func(k int32, v interface{}) bool {
		return fn(k, v.(int32))
	})
}
func (a *KVInt32Int32) Equal(other *KVInt32Int32) bool {
	return (*attr.Int32Map)(a).Equal((*attr.Int32Map)(other))
}
func (a *KVInt32Int32) Has(k int32) bool {
	return (*attr.Int32Map)(a).Has(k)
}
func (a *KVInt32Int32) data() map[int32]int32 {
	dd := map[int32]int32{}
	a.ForEach(func(k int32, v int32) bool {
		dd[k] = v
		return true
	})
	return dd
}
func (a *KVInt32Int32) MarshalJSON() ([]byte, error) {
	return json.Marshal((*attr.Int32Map)(a).ToMap())
}
func (a *KVInt32Int32) UnmarshalJSON(b []byte) error {
	dd := map[int32]int32{}
	err := json.Unmarshal(b, &dd)
	if err != nil {
		return err
	}
	convertData := map[int32]interface{}{}
	for k, v := range dd {
		convertData[k] = v
	}
	(*attr.Int32Map)(a).SetData(convertData)
	return nil
}
func (a *KVInt32Int32) MarshalBSON() ([]byte, error) {
	return bson.Marshal((*attr.Int32Map)(a).ToMap())
}
func (a *KVInt32Int32) UnmarshalBSON(b []byte) error {
	dd := map[int32]int32{}
	err := bson.Unmarshal(b, &dd)
	if err != nil {
		return err
	}
	convertData := map[int32]interface{}{}
	for k, v := range dd {
		convertData[k] = v
	}
	(*attr.Int32Map)(a).SetData(convertData)
	return nil
}
