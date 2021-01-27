package main

import (
	"fmt"
	"go/types"
	"strings"

	. "github.com/dave/jennifer/jen"
)

func genMapTypName(v *types.Map) string {
	key := strings.Title(v.Key().String())
	if key == "String" {
		key = "Str"
	}

	val := getTypString(v.Elem())
	if strings.HasPrefix(val, "*") {
		val = strings.TrimLeft(val, "*")
	}
	val = strings.Title(val)
	if val == "String" {
		val = "Str"
	}
	return fmt.Sprintf("KV%s%s", key, val)
}

func checkMapKey(v *types.Map) error {
	switch mapK := v.Key().(type) {
	case *types.Basic:
		if mapK.Kind() == types.Int32 || mapK.Kind() == types.String {
		} else {
			return fmt.Errorf("不支持的map key，目前 map key 只支持 int32 和 string. %T", mapK)
		}
	default:
		return fmt.Errorf("不支持的map key，目前 map key 只支持 zint32 和 string. %T", mapK)
	}
	return nil
}

func writeMap(f *File, v *types.Map) error {
	err := checkMapKey(v)
	if err != nil {
		return err
	}

	return nil
}
