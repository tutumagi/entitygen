package main

import (
	"fmt"
	"go/types"
	"os"
	"path/filepath"
	"regexp"
	"strings"

	"golang.org/x/tools/go/packages"

	// 代码生成库
	. "github.com/dave/jennifer/jen"
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
	// for i := 0; i < structType.NumFields(); i++ {
	// 	field := structType.Field(i)
	// 	tagValue := structType.Tag(i)
	// 	fmt.Println(field.Name(), tagValue, field.Type())
	// }

	err := generate(sourceTypeName, structType)
	if err != nil {
		failErr(err)
	}
}

var structJsonPattern = regexp.MustCompile(`json:"([^"]+)"`)
var structBsonPattern = regexp.MustCompile(`bson:"([^"]+)"`)

const (
	ParentKeyName  = "parentKey"
	AttrsFieldName = "attrs"
)

func generate(sourceTypeName string, structType *types.Struct) error {
	goPackage := os.Getenv("GOPACKAGE")

	f := NewFile(goPackage)

	f.PackageComment("Code generated by generator, DO NOT EDIT.")

	// 1. 写定义
	attrName := sourceTypeName + "Def"
	f.Type().Id(attrName).Struct(
		Id(ParentKeyName).Id("string"),
		Id(AttrsFieldName).Op("*").Qual(
			"entitygen/attr",
			"AttrMap",
		),
	)

	// 2. 写字段的 getter/setter
	// var getterSetterFunc []Code

	for i := 0; i < structType.NumFields(); i++ {
		field := structType.Field(i)
		// tagValue := structType.Tag(i)

		fieldName := field.Name()
		setParamName := strings.ToLower(fieldName)

		switch v := field.Type().(type) {
		case *types.Basic:
			// getter
			f.Func().
				Params(Id("a").Op("*").Id(attrName)).
				Id(fmt.Sprintf("Get%s", fieldName)).
				Params().Id(v.Name()).
				Block(
					Return(
						Id("a").Dot(AttrsFieldName).Index(Lit(fieldName)),
					),
				)

			// setter
			f.Func().
				Params(Id("a").Op("*").Id(attrName)).
				Id(fmt.Sprintf("Set%s", fieldName)).
				Params(Id(setParamName).Id(v.Name())).
				Block()

			f.Line()
		case *types.Map:
			switch mapK := v.Key().(type) {
			case *types.Basic:
				if mapK.Kind() == types.Int32 || mapK.Kind() == types.String {

				} else {
					return fmt.Errorf("不支持的map key，目前 map key 只支持 int32 和 string. %T", mapK)
				}
			default:
				return fmt.Errorf("不支持的map key，目前 map key 只支持 zint32 和 string. %T", mapK)
			}
			// getter

		case *types.Named:

			// typName := v.Obj()
			// // Qual automatically imports packages
			// code.Op("*").Qual(
			// 	typName.Pkg().Path(),
			// 	typName.Name(),
			// )
		default:
			return fmt.Errorf("struct field type not handled: %T", v)
		}

		// 	changeSetFields = append(changeSetFields, code)
	}

	// changeSetName := sourceTypeName + "ChangeSet"
	// f.Type().Id(changeSetName).Struct(changeSetFields...)

	goFile := os.Getenv("GOFILE")
	ext := filepath.Ext(goFile)
	baseFilename := goFile[0 : len(goFile)-len(ext)]
	targetFilename := baseFilename + "_" + strings.ToLower(sourceTypeName) + "_gen.go"

	return f.Save(targetFilename)
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
