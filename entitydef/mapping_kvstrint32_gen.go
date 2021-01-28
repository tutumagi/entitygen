// Code generated by generator, DO NOT EDIT.
package entitydef

import attr "entitygen/attr"

type KVStrInt32 attr.StrMap

func EmptyKVStrInt32() *KVStrInt32 {
	return NewKVStrInt32(nil)
}
func NewKVStrInt32(data map[string]int32) *KVStrInt32 {
	var convertData map[string]interface{} = map[string]interface{}{}
	for k, v := range data {
		convertData[k] = v
	}
	return (*KVStrInt32)(attr.NewStrMap(convertData))
}
func (a *KVStrInt32) Set(k string, v int32) {
	(*attr.StrMap)(a).Set(k, v)
}
func (a *KVStrInt32) Get(k string) int32 {
	return (*attr.StrMap)(a).Int32(k)
}
func (a *KVStrInt32) Delete(k string) bool {
	return (*attr.StrMap)(a).Delete(k)
}
func (a *KVStrInt32) setParent(k string, parent attr.Field) {
	(*attr.StrMap)(a).SetParent(k, parent)
}
func (a *KVStrInt32) ForEach(fn func(k string, v int32) bool) {
	(*attr.StrMap)(a).ForEach(func(k string, v interface{}) bool {
		return fn(k, v.(int32))
	})
}
func (a *KVStrInt32) Equal(other *KVStrInt32) bool {
	return (*attr.StrMap)(a).Equal((*attr.StrMap)(other))
}
func (a *KVStrInt32) data() map[string]int32 {
	dd := map[string]int32{}
	a.ForEach(func(k string, v int32) bool {
		dd[k] = v
		return true
	})
	return dd
}
