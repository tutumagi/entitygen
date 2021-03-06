package attr

import (
	"reflect"
	"unicode"
)

func isNil(v interface{}) bool {
	if v == nil {
		return true
	}

	t := reflect.ValueOf(v)
	switch t.Kind() {
	case reflect.Bool,
		reflect.Int,
		reflect.Int8,
		reflect.Int16,
		reflect.Int32,
		reflect.Int64,
		reflect.Uint,
		reflect.Uint8,
		reflect.Uint16,
		reflect.Uint32,
		reflect.Uint64,
		reflect.Float32,
		reflect.Float64,
		reflect.Complex64,
		reflect.Complex128,
		reflect.String:
		return false
	case reflect.Interface:
		return isNil(t.Elem())
	default:
		return t.IsNil()
	}

	// if v == nil || (reflect.ValueOf(v).Kind() == reflect.Ptr && reflect.ValueOf(v).IsNil()) {
	// 	return true
	// }
	// return false
}

func LowerFirst(s string) string {
	for _, c := range s {
		return string(unicode.ToLower(c)) + s[1:]
	}
	return s
}
