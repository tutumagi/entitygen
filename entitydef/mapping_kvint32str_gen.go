// Code generated by generator, DO NOT EDIT.
package entitydef

import (
	"encoding/json"
	attr "gitlab.gamesword.com/nut/entitygen/attr"
	bson "go.mongodb.org/mongo-driver/bson"
)

type KVInt32Str attr.Int32Map

func EmptyKVInt32Str() *KVInt32Str {
	return NewKVInt32Str(nil)
}
func NewKVInt32Str(data map[int32]string) *KVInt32Str {
	var convertData map[int32]interface{} = map[int32]interface{}{}
	for k, v := range data {
		convertData[k] = v
	}
	return (*KVInt32Str)(attr.NewInt32Map(convertData))
}
func (a *KVInt32Str) Set(k int32, v string) {
	(*attr.Int32Map)(a).Set(k, v)
}
func (a *KVInt32Str) Get(k int32) string {
	return (*attr.Int32Map)(a).Str(k)
}
func (a *KVInt32Str) Delete(k int32) bool {
	return (*attr.Int32Map)(a).Delete(k)
}
func (a *KVInt32Str) setParent(k string, parent attr.Field) {
	(*attr.Int32Map)(a).SetParent(k, parent)
}
func (a *KVInt32Str) ForEach(fn func(k int32, v string) bool) {
	(*attr.Int32Map)(a).ForEach(func(k int32, v interface{}) bool {
		return fn(k, v.(string))
	})
}
func (a *KVInt32Str) Equal(other *KVInt32Str) bool {
	return (*attr.Int32Map)(a).Equal((*attr.Int32Map)(other))
}
func (a *KVInt32Str) data() map[int32]string {
	dd := map[int32]string{}
	a.ForEach(func(k int32, v string) bool {
		dd[k] = v
		return true
	})
	return dd
}
func (a *KVInt32Str) MarshalJSON() ([]byte, error) {
	return json.Marshal((*attr.Int32Map)(a).ToMap())
}
func (a *KVInt32Str) UnmarshalJSON(b []byte) error {
	dd := map[int32]string{}
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
func (a *KVInt32Str) MarshalBSON() ([]byte, error) {
	return bson.Marshal((*attr.Int32Map)(a).ToMap())
}
func (a *KVInt32Str) UnmarshalBSON(b []byte) error {
	dd := map[int32]string{}
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
