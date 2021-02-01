package demo

import "gitlab.gamesword.com/nut/entitygen/attr"

type Int8Arr attr.Slice

func EmptyInt8Arr() *Int8Arr {
	return NewInt8Arr(nil)
}

func NewInt8Arr(items []int8) *Int8Arr {
	var convertData []interface{} = []interface{}{}
	for _, i := range items {
		convertData = append(convertData, i)
	}
	return (*Int8Arr)(attr.NewSlice(convertData))
}

func (a *Int8Arr) Add(item int8) {
	(*attr.Slice)(a).Add(item)
}

func (a *Int8Arr) At(idx int) (int8, bool) {
	val := (*attr.Slice)(a).Value(idx)
	if val == nil {
		return 0, false
	}
	return val.(int8), true
}

func (a *Int8Arr) Set(idx int, item int8) {
	(*attr.Slice)(a).Set(idx, item)
}

func (a *Int8Arr) DelIndex(idx int) bool {
	return (*attr.Slice)(a).DeleteAt(idx)
}

func (a *Int8Arr) Count() int {
	return (*attr.Slice)(a).Len()
}

func (a *Int8Arr) setParent(k string, parent attr.Field) {
	(*attr.Slice)(a).SetParent(k, parent)
}
func (a *Int8Arr) ForEach(fn func(idx int, v int8) bool) {
	(*attr.Slice)(a).ForEach(func(idx int, v interface{}) bool {
		return fn(idx, v.(int8))
	})
}
func (a *Int8Arr) Equal(other *Int8Arr) bool {
	return (*attr.Slice)(a).Equal((*attr.Slice)(other))
}

func (a *Int8Arr) data() []int8 {
	dd := []int8{}
	a.ForEach(func(idx int, v int8) bool {
		dd = append(dd, v)
		return true
	})
	return dd
}
