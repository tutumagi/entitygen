// Code generated by generator, DO NOT EDIT.
package entitydef

import (
	"encoding/json"
	attr "gitlab.gamesword.com/nut/entitygen/attr"
	bson "go.mongodb.org/mongo-driver/bson"
)

type KVStrStr attr.StrMap

func EmptyKVStrStr() *KVStrStr {
	return NewKVStrStr(nil)
}
func NewKVStrStr(data map[string]string) *KVStrStr {
	var convertData map[string]interface{} = map[string]interface{}{}
	for k, v := range data {
		convertData[k] = v
	}
	return (*KVStrStr)(attr.NewStrMap(convertData))
}
func (a *KVStrStr) Set(k string, v string) {
	(*attr.StrMap)(a).Set(k, v)
}
func (a *KVStrStr) Get(k string) string {
	return (*attr.StrMap)(a).Str(k)
}
func (a *KVStrStr) Delete(k string) bool {
	return (*attr.StrMap)(a).Delete(k)
}
func (a *KVStrStr) setParent(k string, parent attr.Field) {
	(*attr.StrMap)(a).SetParent(k, parent)
}
func (a *KVStrStr) ForEach(fn func(k string, v string) bool) {
	(*attr.StrMap)(a).ForEach(func(k string, v interface{}) bool {
		return fn(k, v.(string))
	})
}
func (a *KVStrStr) Equal(other *KVStrStr) bool {
	return (*attr.StrMap)(a).Equal((*attr.StrMap)(other))
}
func (a *KVStrStr) Has(k string) bool {
	return (*attr.StrMap)(a).Has(k)
}
func (a *KVStrStr) data() map[string]string {
	dd := map[string]string{}
	a.ForEach(func(k string, v string) bool {
		dd[k] = v
		return true
	})
	return dd
}
func (a *KVStrStr) MarshalJSON() ([]byte, error) {
	return json.Marshal((*attr.StrMap)(a).ToMap())
}
func (a *KVStrStr) UnmarshalJSON(b []byte) error {
	dd := map[string]string{}
	err := json.Unmarshal(b, &dd)
	if err != nil {
		return err
	}
	convertData := map[string]interface{}{}
	for k, v := range dd {
		convertData[k] = v
	}
	(*attr.StrMap)(a).SetData(convertData)
	return nil
}
func (a *KVStrStr) MarshalBSON() ([]byte, error) {
	return bson.Marshal((*attr.StrMap)(a).ToMap())
}
func (a *KVStrStr) UnmarshalBSON(b []byte) error {
	dd := map[string]string{}
	err := bson.Unmarshal(b, &dd)
	if err != nil {
		return err
	}
	convertData := map[string]interface{}{}
	for k, v := range dd {
		convertData[k] = v
	}
	(*attr.StrMap)(a).SetData(convertData)
	return nil
}
