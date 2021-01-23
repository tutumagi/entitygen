package attr

import (
	"fmt"
	"strings"
	"sync"
)

// AttrMap map attr
type AttrMap struct {
	// 如果这个是 根AttrMap，则 rootKey 为空，ancestry 为 nil
	rootKey string
	// 存这个数据的祖宗
	ancestry *AttrMap

	data map[string]interface{}

	changedkey map[string]struct{}
}

var attrPool *sync.Pool = &sync.Pool{
	New: func() interface{} {
		return &AttrMap{
			data:       map[string]interface{}{},
			changedkey: map[string]struct{}{},
		}
	},
}

func NewAttrMap() *AttrMap {
	a := attrPool.Get().(*AttrMap)
	a.rootKey = ""
	a.ancestry = nil

	a.data = map[string]interface{}{}
	for k := range a.changedkey {
		delete(a.changedkey, k)
	}
	return a
}

func NewSubAttrMap(key string, ancestry *AttrMap) *AttrMap {
	a := attrPool.Get().(*AttrMap)
	a.rootKey = key
	a.ancestry = ancestry

	a.data = map[string]interface{}{}
	for k := range a.changedkey {
		delete(a.changedkey, k)
	}
	return a
}

func ReleaseAttrMap(mm *AttrMap) {
	attrPool.Put(mm)
}

func (a *AttrMap) SetRootKey(k string) {
	if a.rootKey != "" && a.rootKey != k {
		panic(fmt.Sprintf("key is already exit old:%s new:%s", a.rootKey, k))
	}
	a.rootKey = k
}

func (a *AttrMap) SetAncestry(ancestry *AttrMap) {
	if a.ancestry != nil && a.ancestry != ancestry {
		panic(fmt.Sprintf("ancestry is already exit old:%s new:%s", a.ancestry, ancestry))
	}
	a.ancestry = ancestry
}

// func (a *AttrMap) ToMap(filter ...func(k string) bool) map[string]interface{} {
func (a *AttrMap) ToMap() map[string]interface{} {
	// result := map[string]interface{}{}
	// var f func(k string) bool = nil
	// if len(filter) > 0 {
	// 	f = filter[0]
	// }
	// for k, v := range a.data {
	// 	if f != nil {
	// 		if f(k) {
	// 			result[k] = v
	// 		}
	// 	} else {
	// 		result[k] = v
	// 	}
	// }
	// return result
	return a.data
}

func (a *AttrMap) SetData(data map[string]interface{}) {
	a.data = data
}

func (a *AttrMap) ForEach(fn func(k string, v interface{}) bool) {
	for k, v := range a.data {
		if !fn(k, v) {
			break
		}
	}
}

func (a *AttrMap) Delete(key string) bool {
	if a.isRoot() {
		panic("can not delete key in root attr")
	}
	if _, ok := a.data[key]; ok {
		delete(a.data, key)
		a.change()
		return true
	}
	return false
}

func (a *AttrMap) FastDelete(key string) {
	if a.isRoot() {
		panic("can not delete key in root attr")
	}
	delete(a.data, key)
	a.change()
}

func (a *AttrMap) Set(key string, val interface{}) {

	a.data[key] = val
	// 这里缓存 修改的 key
	// 还有一种做法是 改变立马通知除去
	a.SetChangeKey(key)
}

func (a *AttrMap) change() {
	if a.isRoot() {

	} else {
		a.ancestry.SetChangeKey(a.rootKey)
	}
}

func (a *AttrMap) isRoot() bool {
	return a.ancestry == nil
}

func (a *AttrMap) HasChange() bool {
	return len(a.changedkey) > 0
}

func (a *AttrMap) ClearChangeKey() {
	for k := range a.changedkey {
		delete(a.changedkey, k)
	}
}

func (a *AttrMap) SetChangeKey(key string) {
	a.changedkey[key] = struct{}{}
}

// Bool returns value with Bool type
func (a *AttrMap) Bool(key string) bool {
	v, ok := a.data[key]
	if !ok {
		return false
	}
	r, ok := v.(bool)
	if !ok {
		return false
	}
	return r
}

// String returns value with String type
func (a *AttrMap) Str(key string) string {
	v, ok := a.data[key]
	if !ok {
		return ""
	}
	r, ok := v.(string)
	if !ok {
		return ""
	}
	return r
}

// Value returns value with interface{} type
func (a *AttrMap) Value(key string) interface{} {
	v, ok := a.data[key]
	if !ok {
		return nil
	}
	return v
}

// Int returns value with Int type
func (a *AttrMap) Int(key string) int {
	v, ok := a.data[key]
	if !ok {
		return 0
	}
	r, ok := v.(int)
	if !ok {
		return 0
	}
	return r
}

// Int8 returns value with Int8 type
func (a *AttrMap) Int8(key string) int8 {
	v, ok := a.data[key]
	if !ok {
		return 0
	}
	r, ok := v.(int8)
	if !ok {
		return 0
	}
	return r
}

// Int16 returns value with Int16 type
func (a *AttrMap) Int16(key string) int16 {
	v, ok := a.data[key]
	if !ok {
		return 0
	}
	r, ok := v.(int16)
	if !ok {
		return 0
	}
	return r
}

// Int32 returns value with Int32 type
func (a *AttrMap) Int32(key string) int32 {
	v, ok := a.data[key]
	if !ok {
		return 0
	}
	r, ok := v.(int32)
	if !ok {
		return 0
	}
	return r
}

// Int64 returns value with Int64 type
func (a *AttrMap) Int64(key string) int64 {
	v, ok := a.data[key]
	if !ok {
		return 0
	}
	r, ok := v.(int64)
	if !ok {
		return 0
	}
	return r
}

// UInt returns value with UInt type
func (a *AttrMap) UInt(key string) uint {
	v, ok := a.data[key]
	if !ok {
		return 0
	}
	r, ok := v.(uint)
	if !ok {
		return 0
	}
	return r
}

// UInt8 returns value with UInt8 type
func (a *AttrMap) UInt8(key string) uint8 {
	v, ok := a.data[key]
	if !ok {
		return 0
	}
	r, ok := v.(uint8)
	if !ok {
		return 0
	}
	return r
}

// UInt16 returns value with UInt16 type
func (a *AttrMap) UInt16(key string) uint16 {
	v, ok := a.data[key]
	if !ok {
		return 0
	}
	r, ok := v.(uint16)
	if !ok {
		return 0
	}
	return r
}

// UInt32 returns value with UInt32 type
func (a *AttrMap) UInt32(key string) uint32 {
	v, ok := a.data[key]
	if !ok {
		return 0
	}
	r, ok := v.(uint32)
	if !ok {
		return 0
	}
	return r
}

// UInt64 returns value with UInt64 type
func (a *AttrMap) UInt64(key string) uint64 {
	v, ok := a.data[key]
	if !ok {
		return 0
	}
	r, ok := v.(uint64)
	if !ok {
		return 0
	}

	return r
}

// Float32 returns value with float32 type
func (a *AttrMap) Float32(key string) float32 {
	v, ok := a.data[key]
	if !ok {
		return 0
	}
	r, ok := v.(float32)
	if !ok {
		return 0
	}
	return r
}

// Float64 returns value with float64 type
func (a *AttrMap) Float64(key string) float64 {
	v, ok := a.data[key]
	if !ok {
		return 0
	}
	r, ok := v.(float64)
	if !ok {
		return 0
	}

	return r
}

func (a *AttrMap) String() string {
	var sb strings.Builder
	sb.WriteString("MapAttr{")
	isFirstField := true
	for k, v := range a.data {
		if !isFirstField {
			sb.WriteString(", ")
		}

		fmt.Fprintf(&sb, "%#v", k)
		sb.WriteString(": ")
		switch a := v.(type) {
		case *AttrMap:
			sb.WriteString(a.String())
		// case *AttrList:
		// 	sb.WriteString(a.String())
		default:
			fmt.Fprintf(&sb, "%#v", v)
		}
		isFirstField = false
	}
	sb.WriteString("}")
	return sb.String()
}
