// Code generated by generator, DO NOT EDIT.
package entitydef

import (
	"encoding/json"
	attr "gitlab.gamesword.com/nut/entitygen/attr"
	bson "go.mongodb.org/mongo-driver/bson"
)

type KVStrDeskDef attr.StrMap

func EmptyKVStrDeskDef() *KVStrDeskDef {
	return NewKVStrDeskDef(nil)
}
func NewKVStrDeskDef(data map[string]*DeskDef) *KVStrDeskDef {
	var convertData map[string]interface{} = map[string]interface{}{}
	for k, v := range data {
		convertData[k] = v
	}
	return (*KVStrDeskDef)(attr.NewStrMap(convertData))
}
func (a *KVStrDeskDef) Set(k string, v *DeskDef) {
	v.setParent(k, (*attr.StrMap)(a))
	(*attr.StrMap)(a).Set(k, v)
}
func (a *KVStrDeskDef) Get(k string) *DeskDef {
	return (*attr.StrMap)(a).Value(k).(*DeskDef)
}
func (a *KVStrDeskDef) Delete(k string) bool {
	return (*attr.StrMap)(a).Delete(k)
}
func (a *KVStrDeskDef) setParent(k string, parent attr.Field) {
	(*attr.StrMap)(a).SetParent(k, parent)
}
func (a *KVStrDeskDef) ForEach(fn func(k string, v *DeskDef) bool) {
	(*attr.StrMap)(a).ForEach(func(k string, v interface{}) bool {
		return fn(k, v.(*DeskDef))
	})
}
func (a *KVStrDeskDef) Equal(other *KVStrDeskDef) bool {
	return (*attr.StrMap)(a).Equal((*attr.StrMap)(other))
}
func (a *KVStrDeskDef) data() map[string]*DeskDef {
	dd := map[string]*DeskDef{}
	a.ForEach(func(k string, v *DeskDef) bool {
		dd[k] = v
		return true
	})
	return dd
}
func (a *KVStrDeskDef) MarshalJSON() ([]byte, error) {
	return json.Marshal((*attr.StrMap)(a).ToMap())
}
func (a *KVStrDeskDef) UnmarshalJSON(b []byte) error {
	dd := map[string]*DeskDef{}
	err := json.Unmarshal(b, &dd)
	if err != nil {
		return err
	}
	convertData := map[string]interface{}{}
	for k, v := range dd {
		v.setParent(k, (*attr.StrMap)(a))
		convertData[k] = v
	}
	(*attr.StrMap)(a).SetData(convertData)
	return nil
}
func (a *KVStrDeskDef) MarshalBSON() ([]byte, error) {
	return bson.Marshal((*attr.StrMap)(a).ToMap())
}
func (a *KVStrDeskDef) UnmarshalBSON(b []byte) error {
	dd := map[string]*DeskDef{}
	err := bson.Unmarshal(b, &dd)
	if err != nil {
		return err
	}
	convertData := map[string]interface{}{}
	for k, v := range dd {
		v.setParent(k, (*attr.StrMap)(a))
		convertData[k] = v
	}
	(*attr.StrMap)(a).SetData(convertData)
	return nil
}
