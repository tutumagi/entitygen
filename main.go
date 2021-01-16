package main

import (
	"fmt"
	"go/types"
	"os"
	"strings"

	"golang.org/x/tools/go/packages"
)

func main() {
	// 1. 获取参数
	if len(os.Args) != 2 {
		failErr(fmt.Errorf("expected exactly on argument"))
	}

	// 2. 拿到参数
	sourceType := os.Args[1] // 包名.类型名
	sourceTypePackage, sourceTypeName := splitSourceType(sourceType)

	// 3. 加载包
	pkg := loadPackage(sourceTypePackage)

	// 4. 搜索这个包有没有这个类型
	obj := pkg.Types.Scope().Lookup(sourceTypeName)
	if obj == nil {
		failErr(fmt.Errorf("%s not found in declared types of %s", sourceTypeName, pkg))
	}

	// 5. 判断 sourceTypeName 是否是个类型名字
	if _, ok := obj.(*types.TypeName); !ok {
		failErr(fmt.Errorf("%s is not a named type", obj))
	}

	// 6. 判断 类型名字 的含义是否是个 struct
	structType, ok := obj.Type().Underlying().(*types.Struct)
	if !ok {
		failErr(fmt.Errorf("type %v is not a struct", obj))
	}

	// 7. 遍历所有的 struct 的名字
	for i := 0; i < structType.NumFields(); i++ {
		field := structType.Field(i)
		tagValue := structType.Tag(i)
		fmt.Println(field.Name(), tagValue, field.Type())
	}

}

func loadPackage(path string) *packages.Package {
	cfg := &packages.Config{
		Mode: packages.NeedTypes | packages.NeedImports,
	}
	pkgs, err := packages.Load(cfg, path)
	if err != nil {
		failErr(fmt.Errorf("loading packages for inspection:%v", err))
	}
	if packages.PrintErrors(pkgs) > 0 {
		os.Exit(1)
	}
	return pkgs[0]
}

func splitSourceType(sourceType string) (string, string) {
	idx := strings.LastIndexByte(sourceType, '.')
	if idx == -1 {
		failErr(fmt.Errorf(`expected qualified type as "pkg/path.Mytype"`))
	}
	sourceTypePackage := sourceType[0:idx]
	sourceTypeName := sourceType[idx+1:]
	return sourceTypePackage, sourceTypeName
}

func failErr(err error) {
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error: %v\n", err)
		os.Exit(1)
	}
}
