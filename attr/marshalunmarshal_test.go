package attr

import (
	"encoding/json"
	"reflect"
	"testing"

	dynamicstruct "github.com/ompluscator/dynamic-struct"
	"go.mongodb.org/mongo-driver/bson"

	. "github.com/go-playground/assert/v2"
)

type TT struct {
	Name string `bson:"name"`
	Age  int32  `bson:"age"`
	T    *TT    `bson:"t"`
}

type TTMap StrMap

func (a *TTMap) MarshalBSON() ([]byte, error) {
	return bson.Marshal((*StrMap)(a).ToMap())
}
func (a *TTMap) UnmarshalBSON(b []byte) error {
	mm, err := ttMeta.UnmarshalBson(b)
	if err != nil {
		return err
	}
	(*StrMap)(a).SetData(mm)

	return nil
}

func (a *TTMap) MarshalJSON() ([]byte, error) {
	return json.Marshal((*StrMap)(a).ToMap())
}
func (a *TTMap) UnmarshalJSON(b []byte) error {
	mm, err := ttMeta.UnmarshalJson(b)
	if err != nil {
		return err
	}
	(*StrMap)(a).SetData(mm)

	return nil
}

func (a *TTMap) IsZero() bool {
	return a == nil
}

var ttMeta *Meta

func init() {
	ttMeta = &Meta{}
	ttMeta.DefAttr("name", "", AfBase, true)
	ttMeta.DefAttr("age", int32(0), AfBase, true)
	ttMeta.DefAttr("t", &TTMap{}, AfBase, true)
}

func TestDynamicStruct(t *testing.T) {

	emptyTT := (*TTMap)(&StrMap{
		data: map[string]interface{}{
			"name": "tt",
			"age":  int32(33),
			"t":    (*TTMap)(nil),
		},
	})
	{
		bb, err := json.Marshal(emptyTT)
		Equal(t, err, nil)

		mm, err := ttMeta.UnmarshalJson(bb)
		Equal(t, err, nil)
		NotEqual(t, mm, nil)

		m := (*StrMap)(emptyTT).data
		Equal(t, err, nil)
		Equal(t, m["name"].(string), mm["name"].(string))
		Equal(t, m["age"].(int32), mm["age"].(int32))
		Equal(t, m["t"].(*TTMap), mm["t"].(*TTMap))
	}

	{
		bb, err := bson.Marshal(emptyTT)
		Equal(t, err, nil)

		mm, err := ttMeta.UnmarshalBson(bb)
		Equal(t, err, nil)
		NotEqual(t, mm, nil)

		m := (*StrMap)(emptyTT).data
		Equal(t, err, nil)
		Equal(t, m["name"].(string), mm["name"].(string))
		Equal(t, m["age"].(int32), mm["age"].(int32))
		Equal(t, m["t"].(*TTMap), mm["t"].(*TTMap))
	}

	// {
	// 	ttmap := &TTMap{}
	// 	err := bson.Unmarshal(bb, ttmap)
	// 	Equal(t, err, nil)
	// 	NotEqual(t, ttmap, nil)

	// 	mm := (*StrMap)(ttmap)
	// 	Equal(t, err, nil)
	// 	Equal(t, emptyTT.Name, mm.data["name"].(string))
	// 	Equal(t, emptyTT.Age, mm.data["age"].(int32))
	// 	Equal(t, nil, mm.data["t"].(*TTMap))
	// }

}

func TestSelf(t *testing.T) {
	emptyTT := &TT{}

	bb, err := bson.Marshal(emptyTT)
	Equal(t, err, nil)

	{
		newTT := &TT{}
		err = bson.Unmarshal(bb, newTT)
		Equal(t, err, nil)
		Equal(t, emptyTT.Name, newTT.Name)
		Equal(t, emptyTT.Age, newTT.Age)
		Equal(t, emptyTT.T, newTT.T)
	}

	{
		newTT2 := reflect.New(reflect.TypeOf(emptyTT).Elem()).Interface()
		err = bson.Unmarshal(bb, newTT2)
		Equal(t, err, nil)
		Equal(t, emptyTT.Name, newTT2.(*TT).Name)
		Equal(t, emptyTT.Age, newTT2.(*TT).Age)
		Equal(t, emptyTT.T, newTT2.(*TT).T)
	}

	{
		builder := dynamicstruct.ExtendStruct(&TT{})
		newTT3 := builder.Build().New()
		err = bson.Unmarshal(bb, newTT3)
		r := dynamicstruct.NewReader(newTT3)

		Equal(t, err, nil)
		Equal(t, emptyTT.Name, r.GetField("Name").String())
		Equal(t, emptyTT.Age, r.GetField("Age").Int32())
		Equal(t, emptyTT.T, r.GetField("T").Interface().(*TT))
	}

	{
		builder := dynamicstruct.NewStruct().
			AddField("Name", "", `bson:"bson"`).
			AddField("Age", int32(0), `bson:"age"`).
			AddField("T", &TT{}, `bson:"t"`)
		newTT3 := builder.Build().New()
		err = bson.Unmarshal(bb, newTT3)
		r := dynamicstruct.NewReader(newTT3)

		Equal(t, err, nil)
		Equal(t, emptyTT.Name, r.GetField("Name").String())
		Equal(t, emptyTT.Age, r.GetField("Age").Int32())
		Equal(t, emptyTT.T, r.GetField("T").Interface().(*TT))
	}
}
