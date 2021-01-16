package attr

import (
	"fmt"
	"strings"
	"sync"
)

// AttrMap map attr
type AttrMap struct {
	attrs map[string]interface{}

	changedkey []string
}

var attrPool *sync.Pool = &sync.Pool{
	New: func() interface{} {
		return &AttrMap{
			attrs:      map[string]interface{}{},
			changedkey: []string{},
		}
	},
}

func NewAttrMap() *AttrMap {
	return attrPool.Get().(*AttrMap)
}

func ReleaseAttrMap(mm *AttrMap) {
	mm.attrs = map[string]interface{}{}
	mm.changedkey = mm.changedkey[0:0]
	attrPool.Put(mm)
}

func (a *AttrMap) String() string {
	var sb strings.Builder
	sb.WriteString("MapAttr{")
	isFirstField := true
	for k, v := range a.attrs {
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

func (a *AttrMap) ToMap(filter ...func(k string) bool) map[string]interface{} {
	result := map[string]interface{}{}
	var f func(k string) bool = nil
	if len(filter) > 0 {
		f = filter[0]
	}
	for k, v := range a.attrs {
		if f != nil {
			if f(k) {
				result[k] = v
			}
		} else {
			result[k] = v
		}
	}
	return result
}

func (a *AttrMap) Set(key string, val interface{}) {
	// var flag attrFlag
	// a.attrs[key] = val
	a.set(key, val)
	// switch sa := val.(type) {
	// case *AttrMap:
	// case *AttrList:
	// }
}

func (a *AttrMap) set(key string, val interface{}) {
	// a.owner.typeDesc.attrsDef[key].typ.(type)
	switch t := val.(type) {
	case int:
		a.attrs[key] = float64(t)
	case uint:
		a.attrs[key] = float64(t)
	case int8:
		a.attrs[key] = float64(t)
	case uint8:
		a.attrs[key] = float64(t)
	case int16:
		a.attrs[key] = float64(t)
	case uint16:
		a.attrs[key] = float64(t)
	case int32:
		a.attrs[key] = float64(t)
	case uint32:
		a.attrs[key] = float64(t)
	case int64:
		a.attrs[key] = float64(t)
	case uint64:
		a.attrs[key] = float64(t)
	case float32:
		a.attrs[key] = float64(t)
	case float64:
		a.attrs[key] = t
	default:
		a.attrs[key] = val
	}

	// 这里缓存 修改的 key
	// 还有一种做法是 改变立马通知除去
	a.changedkey = append(a.changedkey, key)
}

func (a *AttrMap) HasChange() bool {
	return len(a.changedkey) > 0
}

func (a *AttrMap) ClearChangeKey() {
	a.changedkey = a.changedkey[0:0]
}

// Map return AttrMap value from index
func (a *AttrMap) Map(key string) *AttrMap {
	v, ok := a.attrs[key]
	if !ok {
		return nil
	}
	r, ok := v.(*AttrMap)
	if !ok {
		return nil
	}
	return r
}

// Bool returns value with Bool type
func (a *AttrMap) Bool(key string) bool {
	v, ok := a.attrs[key]
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
	v, ok := a.attrs[key]
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
	v, ok := a.attrs[key]
	if !ok {
		return nil
	}
	return v
}

// Int returns value with Int type
func (a *AttrMap) Int(key string) int {
	v, ok := a.attrs[key]
	if !ok {
		return 0
	}
	r, ok := v.(float64)
	if !ok {
		return 0
	}
	return int(r)
}

// Int8 returns value with Int8 type
func (a *AttrMap) Int8(key string) int8 {
	v, ok := a.attrs[key]
	if !ok {
		return 0
	}
	r, ok := v.(float64)
	if !ok {
		return 0
	}
	return int8(r)
}

// Int16 returns value with Int16 type
func (a *AttrMap) Int16(key string) int16 {
	v, ok := a.attrs[key]
	if !ok {
		return 0
	}
	r, ok := v.(float64)
	if !ok {
		return 0
	}
	return int16(r)
}

// Int32 returns value with Int32 type
func (a *AttrMap) Int32(key string) int32 {
	v, ok := a.attrs[key]
	if !ok {
		return 0
	}
	r, ok := v.(float64)
	if !ok {
		return 0
	}
	return int32(r)
}

// Int64 returns value with Int64 type
func (a *AttrMap) Int64(key string) int64 {
	v, ok := a.attrs[key]
	if !ok {
		return 0
	}
	r, ok := v.(float64)
	if !ok {
		return 0
	}
	return int64(r)
}

// UInt returns value with UInt type
func (a *AttrMap) UInt(key string) uint {
	v, ok := a.attrs[key]
	if !ok {
		return 0
	}
	r, ok := v.(float64)
	if !ok {
		return 0
	}
	return uint(r)
}

// UInt8 returns value with UInt8 type
func (a *AttrMap) UInt8(key string) uint8 {
	v, ok := a.attrs[key]
	if !ok {
		return 0
	}
	r, ok := v.(float64)
	if !ok {
		return 0
	}
	return uint8(r)
}

// UInt16 returns value with UInt16 type
func (a *AttrMap) UInt16(key string) uint16 {
	v, ok := a.attrs[key]
	if !ok {
		return 0
	}
	r, ok := v.(float64)
	if !ok {
		return 0
	}
	return uint16(r)
}

// UInt32 returns value with UInt32 type
func (a *AttrMap) UInt32(key string) uint32 {
	v, ok := a.attrs[key]
	if !ok {
		return 0
	}
	r, ok := v.(uint32)
	if !ok {
		return 0
	}
	return uint32(r)
}

// UInt64 returns value with UInt64 type
func (a *AttrMap) UInt64(key string) uint64 {
	v, ok := a.attrs[key]
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
	v, ok := a.attrs[key]
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
	v, ok := a.attrs[key]
	if !ok {
		return 0
	}
	r, ok := v.(float64)
	if !ok {
		return 0
	}

	return r
}
