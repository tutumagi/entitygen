package main

import (
	"fmt"
	"go/types"
	"os"
	"path/filepath"
	"reflect"
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

func generate(sourceTypeName string, structType *types.Struct) error {
	goPackage := os.Getenv("GOPACKAGE")

	f := NewFile(goPackage)

	f.PackageComment("Code generated by generator, DO NOT EDIT.")

	fmt.Printf("begin generate type:%s\n", sourceTypeName)

	// 读取 types.Struct 所有字段信息，计算出我们要的信息，并做合法性判断
	fields := getStructFields(structType)

	// 生成的结构体名字 XXXDef
	structName := sourceTypeName + "Def"
	// 生成的对应的数据结构描述的名字 XXXAttrDef
	attrDefName := strings.ToLower(sourceTypeName) + "AttrDef"

	// 一些预设的类型或者关键字
	// *attr.StrMap
	attrStrMap := func() *Statement { return Id("*").Qual("entitygen/attr", "StrMap") }
	// attr.Field
	attrField := func() *Statement { return Qual("entitygen/attr", "Field") }
	// 将 name 变量转为 *attr.StrMap类型: (*attr.StrMap)(name)
	convertAttrStrMap := func(name string) *Statement { return Parens(attrStrMap()).Parens(Id(name)) }
	// a *XXXDef
	thisFn := func() *Statement { return Id("a").Op("*").Id(structName) }
	// 将 "a" 转为 *attr.StrMap 类型：(*attr.StrMap)(a)
	convertThisFn := func() *Statement { return convertAttrStrMap("a") }

	// 1. 写 attrDef
	writeAttrDef(f, attrDefName, fields)

	// 2. 写定义  type XXXDef attr.StrMap
	f.Type().Id(structName).Qual(
		"entitygen/attr",
		"StrMap",
	)

	// 3. 写构造函数
	writeCtor(f, structName, sourceTypeName, fields)

	// 4. 写所有字段的 getter/setter
	err := writeGetterSetter(f, fields, thisFn, convertThisFn)
	if err != nil {
		failErr(err)
	}

	// 5. 写自定义方法
	writeCustomMethod(f, structName, attrField, thisFn, convertThisFn, convertAttrStrMap)

	// 6. 写 marshal & unmarshal
	writeEncodeDecode(f, thisFn, convertThisFn, attrDefName)

	// 7. dump 到 文件
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

type structField struct {
	name     string
	key      string
	typ      types.Type
	storeDB  bool
	base     bool
	cell     bool
	client   bool
	getter   Code
	setter   Code
	setParam Code
	// zero value 对应的 Code
	emptyValue Code
}

func getStructFields(structType *types.Struct) []*structField {
	result := make([]*structField, 0, structType.NumFields())
	for i := 0; i < structType.NumFields(); i++ {
		field := structType.Field(i)
		name := field.Name()

		if name == "id" {
			// NOTE: 目前生成的代码里面跳过 id 的处理
			continue
		}

		typ := field.Type()
		storeDB := false
		flagBase := true // 目前的实现里面属性肯定会存储在 base 里面
		flagCell := false
		client := false

		tagValue := reflect.StructTag(structType.Tag(i))
		key, ok := tagValue.Lookup("key")
		if !ok {
			failErr(fmt.Errorf("field:%s 必须有tag:key", name))
		}
		{
			storeDBStr, ok := tagValue.Lookup("storedb")
			if !ok {
				failErr(fmt.Errorf("field:%s 必须有tag:storedb", name))
			}
			if storeDBStr != "true" && storeDBStr != "false" {
				failErr(fmt.Errorf("field:%s storedb(%s) 必须是 true 或者 false", name, storeDBStr))
			}
			if storeDBStr == "true" {
				storeDB = true
			}
		}

		{
			clientStr, ok := tagValue.Lookup("client")
			if !ok {
				failErr(fmt.Errorf("field:%s 必须有tag:client", name))
			}
			if clientStr != "true" && clientStr != "false" {
				failErr(fmt.Errorf("field:%s client(%s) 必须是 true 或者 false", name, clientStr))
			}
			if clientStr == "true" {
				client = true
			}
		}

		{
			flagStr, ok := tagValue.Lookup("flag")
			if !ok {
				failErr(fmt.Errorf("field:%s 必须有tag:flag", name))
			}
			if flagStr != "base" && flagStr != "cell" {
				failErr(fmt.Errorf("field:%s flag(%s) 必须是 base 或者 cell", name, flagStr))
			}
			if flagStr == "cell" {
				flagCell = true
			}
		}

		result = append(result, &structField{
			name:       name,
			key:        key,
			typ:        typ,
			storeDB:    storeDB,
			base:       flagBase,
			cell:       flagCell,
			client:     client,
			emptyValue: getEmptyValue(typ),
			getter:     Id(fmt.Sprintf("Get%s", name)),
			setter:     Id(fmt.Sprintf("Set%s", name)),
			setParam:   Id(key).Id(typ.String()),
		})
	}
	return result
}

func getEmptyValue(typ types.Type) Code {
	switch v := typ.(type) {
	case *types.Basic:
		switch v.Kind() {
		case types.String, types.UntypedString:
			return Lit("")
		case types.Int, types.Uint, types.Int8, types.Uint8, types.Int16, types.Uint16, types.Int32, types.Uint32, types.Int64, types.Uint64, types.Float32, types.Float64:
			return Lit(0)
		case types.Bool:
			return Lit(false)
		default:
			return Lit(nil)
		}
	default:
		return Lit(nil)
	}
}

func writeEncodeDecode(f *File, thisFn func() *Statement, convertThisFn func() *Statement, attrDefName string) {
	writer := func(encodeFnName string, encodePackageFn *Statement, decodeFnName string, decodePackageFn *Statement) {
		// marshal
		f.Func().Params(thisFn()).Id(encodeFnName).Params().Params(Index().Byte(), Error()).
			Block(
				Return(encodePackageFn.Call(convertThisFn().Dot("ToMap").Call())),
			)
		// unmarshal
		f.Func().Params(thisFn()).Id(decodeFnName).Params(Id("b").Index().Byte()).Error().
			BlockFunc(func(g *Group) {
				g.Id("mm").Id(",").Id("err").Op(":=").Add(decodePackageFn.Call(Id("b")))
				g.If(Id("err").Op("!=").Nil()).Block(
					Return(Id("err")),
				)
				g.Add(convertThisFn().Dot("SetData").Params(Id("mm")))
				g.Add(convertThisFn().Dot("ForEach").Params(
					Func().Params(Id("k").String(), Id("v").Interface()).Bool().
						BlockFunc(func(g *Group) {
							g.If(Id("k").Op("!=").Lit("id").Op("&&").Op("!").Id(attrDefName).Dot("GetDef").Params(Id("k")).Dot("IsPrimary").Params().Block(
								Id("v").Dot("").Parens(Id("IField")).Dot("setParent").Params(Id("k"), convertThisFn()),
							))
							g.Return(True())
						}),
				))
				g.Return(Nil())
			},
			)
	}
	writer(
		"MarshalJSON",
		Qual("encoding/json", "Marshal"),
		"UnmarshalJSON",
		Id(attrDefName).Dot("UnmarshalJson"),
	)

	writer(
		"MarshalBSON",
		Qual("go.mongodb.org/mongo-driver/bson", "Marshal"),
		"UnmarshalBSON",
		Id(attrDefName).Dot("UnmarshalBson"),
	)
	// // marshal
	// f.Func().Params(thisFn()).Id("MarshalJSON").Params().Params(Index().Byte(), Error()).
	// 	Block(
	// 		Return(Qual("encoding/json", "Marshal").Call(convertThisFn().Dot("ToMap").Params())),
	// 	)
	// // unmarshal
	// f.Func().Params(thisFn()).Id("UnmarshalJSON").Params(Id("b").Index().Byte()).Error().
	// 	BlockFunc(func(g *Group) {
	// 		g.Id("mm").Id(",").Id("err").Op(":=").Id(attrDefName).Dot("UnmarshalJson").Params(Id("b"))
	// 		g.If(Id("err").Op("!=").Nil()).Block(
	// 			Return(Id("err")),
	// 		)
	// 		g.Add(convertThisFn().Dot("SetData").Params(Id("mm")))
	// 		g.Add(convertThisFn().Dot("ForEach").Params(
	// 			Func().Params(Id("k").String(), Id("v").Interface()).Bool().
	// 				BlockFunc(func(g *Group) {
	// 					g.If(Id("k").Op("!=").Lit("id").Op("&&").Op("!").Id(attrDefName).Dot("GetDef").Params(Id("k")).Dot("IsPrimary").Params().Block(
	// 						Id("v").Dot("").Parens(Id("IField")).Dot("setParent").Params(Id("k"), convertThisFn()),
	// 					))
	// 					g.Return(True())
	// 				}),
	// 		))
	// 		g.Return(Nil())
	// 	},
	// 	)
}

func writeCustomMethod(
	f *File,
	structName string,
	attrField func() *Statement,
	thisFn func() *Statement,
	convertThisFn func() *Statement,
	convertAttrStrMap func(string) *Statement,
) {
	// 3. 写 changekey 相关的
	f.Func().Params(thisFn()).Id("HasChange").Params().Bool().
		Block(
			Return(convertThisFn().Dot("HasChange").Call()),
		)

	f.Func().Params(thisFn()).Id("ChangeKey").Params().Map(String()).Struct().
		Block(
			Return(convertThisFn().Dot("ChangeKey").Call()),
		)

	f.Func().Params(thisFn()).Id("ClearChangeKey").Params().
		Block(
			convertThisFn().Dot("ClearChangeKey").Call(),
		)

	// 4. 写 setParent
	f.Func().Params(thisFn()).Id("setParent").Params(Id("k").String(), Id("parent").Add(attrField())).
		Block(
			convertThisFn().Dot("SetParent").Call(Id("k"), Id("parent")),
		)

	// 5. ForEach
	f.Func().Params(thisFn()).Id("ForEach").Params(Id("fn").Func().Params(Id("s").String(), Id("v").Interface()).Bool()).
		Block(
			convertThisFn().Dot("ForEach").Call(Id("fn")),
		)

	// 写 Equal
	f.Func().Params(thisFn()).Id("Equal").Params(Id("other").Op("*").Id(structName)).Bool().Block(
		Return(convertThisFn().Dot("Equal").Call(convertAttrStrMap("other"))),
	)
}

func writeGetterSetter(f *File, fields []*structField, thisFn func() *Statement, convertThisFn func() *Statement) error {
	for i := 0; i < len(fields); i++ {

		field := fields[i]

		switch v := field.typ.(type) {
		case *types.Basic:
			// 写 getter
			// attr.StrMap 的 get 方法
			// 如果是基础类型，则直接大写第一个字母的方法进行 getter 比如 int32 就是 .Int32("xxx")
			// 如果是 string 类型，则使用 Str 方法，比如 .Str("yyy")
			attrGetFuncName := strings.Title(v.Name())
			switch v.Kind() {
			case types.String, types.UntypedString:
				attrGetFuncName = "Str"
			}

			// func (a *XXXDef) GetField() FieldType
			f.Func().Params(thisFn()).Add(field.getter).Params().Id(v.Name()).
				Block(
					Return(
						convertThisFn().Dot(attrGetFuncName).Params(Lit(field.key)),
					),
				)

			//  写 setter
			f.Func().Params(thisFn()).Add(field.setter).Params(field.setParam).
				Block(
					convertThisFn().Dot("Set").Params(Lit(field.key), Id(field.key)),
				)

			// 换行符
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
	}
	return nil
}

func writeCtor(f *File, structName string, sourceTypeName string, fields []*structField) {
	// EmptyXXXX 和 NewXXX
	emptyCtorName := fmt.Sprintf("Empty%s", sourceTypeName)
	normalCtorName := fmt.Sprintf("New%s", sourceTypeName)
	// 写 EmptyXXX
	f.Func().Id(emptyCtorName).Params().Op("*").Id(structName).
		Block(
			Return(Id(normalCtorName).CallFunc(func(g *Group) {
				for _, field := range fields {
					g.Add(field.emptyValue)
				}
			})),
		)
	// 写 NewXXX
	f.Func().Id(normalCtorName).ParamsFunc(func(g *Group) {
		for _, field := range fields {
			g.Add(field.setParam)
		}
	}).Op("*").Id(structName).
		BlockFunc(func(g *Group) {

			g.Id("m").Op(":=").Parens(Op("*").Id(structName)).Parens(Qual("entitygen/attr", "NewStrMap").Call(Nil()))

			for _, field := range fields {
				g.Id("m").Dot("").Add(field.setter).Call(Id(field.key))
			}

			g.Id("m").Dot("ClearChangeKey").Call()

			g.Return(Id("m"))
		})
}

func writeAttrDef(f *File, attrDefName string, fields []*structField) {
	// *attr.Def
	attrDef := func() *Statement { return Id("*").Qual("entitygen/attr", "Def") }

	// var xxxAttrDef *attr.Def
	f.Var().Id(attrDefName).Add(attrDef())
	f.Func().Id("init").Params().
		BlockFunc(
			func(g *Group) {
				g.Id(attrDefName).Op("=").Op("&").Qual("entitygen/attr", "Def").Block()
				g.Line()

				for i := 0; i < len(fields); i++ {
					field := fields[i]

					switch v := field.typ.(type) {
					case *types.Basic:
						g.Id(attrDefName).Dot("DefAttr").CallFunc(func(ig *Group) {
							ig.Lit(field.key)
							ig.Qual("entitygen/attr", strings.Title(v.Name()))

							if field.cell {
								ig.Qual("entitygen/attr", "AfCell")
							} else {
								ig.Qual("entitygen/attr", "AfBase")
							}

							if field.storeDB {
								ig.True()
							} else {
								ig.False()
							}
						})
					}
				}
			},
		)
}
