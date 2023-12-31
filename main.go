package main

import (
	"fmt"
	"go/types"
	"os"
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
	sourceArg := os.Args[1] // 包名.类型名

	if isGenerateSingleType(sourceArg) {
		sourceTypePackage, sourceTypeName := splitSourceType(sourceArg)

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

		err := generate(sourceTypeName, structType)
		if err != nil {
			failErr(err)
		}
	} else {
		pkg := loadPackage(sourceArg)

		allStructs := make(map[string]*ssInfo, len(pkg.TypesInfo.Types))

		for _, tt := range pkg.TypesInfo.Defs {
			// 这里为 nil 情况就是类似 go 源文件第一行写着 `package domain`，则 ast 是有的，不过类型定义为 nil
			if tt != nil {
				collectTypes(tt.Name(), tt.Type(), allStructs)
				// if ss != nil {
				// allStructs[ss.name] = *ss
				// }
				// switch v := tt.Type().(type) {
				// case *types.Basic:
				// 	fmt.Printf("collect basic types name:%s type:%s. skip it\n", tt.Name(), tt.Type())
				// case *types.Map:
				// 	fmt.Printf("collect map types name:%s type:%s. \n", tt.Name(), tt.Type())
				// 	allStructs = append(allStructs, ssInfo{
				// 		name: MapTypeName(v),
				// 		typ:  v,
				// 	})
				// case *types.Struct:
				// 	fmt.Printf("collect struct types name:%s type:%s. \n", tt.Name(), tt.Type())
				// 	allStructs = append(allStructs, ssInfo{
				// 		name: tt.Name(),
				// 		typ:  v,
				// 	})
				// case *types.Named: // 某个字段是自定义类型 会跑到这里
				// 	// fmt.Printf("collect types is named: %s. \n", tt.Type.String())
				// 	fmt.Printf("collect named types name:%s type:%s. \n", tt.Name(), tt.Type())
				// 	allStructs = append(allStructs, ssInfo{
				// 		name: tt.Name(),
				// 		typ:  tt.Type().Underlying().(*types.Struct),
				// 	})
				// case *types.Slice:
				// 	fmt.Printf("collect slice types name:%s type:%s. \n", tt.Name(), tt.Type())
				// 	allStructs = append(allStructs, ssInfo{
				// 		name: SliceTypeName(v),
				// 		typ:  v,
				// 	})
				// case *types.Pointer:
				// 	v.Elem()
				// default:
				// 	fmt.Printf("collect other types name:%s type:%s. \n", tt.Name(), tt.Type())
				// }
			}
		}
		// for _, tt := range pkg.TypesInfo.Types {
		// 	// fmt.Printf("collect types is buildin:%v %s\n", tt.IsBuiltin(), tt.Type.String())

		// 	switch v := tt.Type.(type) {
		// 	case *types.Basic:
		// 		fmt.Printf("collect types is %s. skip it.\n", v.Name())
		// 	case *types.Map:
		// 		fmt.Printf("collect types is %s.\n", tt.Type.String())
		// 	case *types.Struct:
		// 		fmt.Printf("collect types is %s value:%v.\n", v.String(), tt.Value)
		// 		allStrucs = append(allStrucs, v)
		// 	case *types.Named: // 某个字段是自定义类型 会跑到这里
		// 		// fmt.Printf("collect types is named: %s. \n", tt.Type.String())
		// 		fmt.Printf("collect types is named: %s. \n", v.Obj().Name())

		// 	}
		// }
		for _, info := range allStructs {
			err := generate(info.name, info.typ)
			if err != nil {
				failErr(err)
			}
		}
		return
	}

}

type ssInfo struct {
	name string
	typ  types.Type
}

func collectTypes(objName string, tt types.Type, outTypes map[string]*ssInfo) {
	switch v := tt.(type) {
	case *types.Basic:
		fmt.Printf("collect basic types name:%s type:%s. skip it\n", objName, v)
	case *types.Map:
		fmt.Printf("collect map types name:%s type:%s. \n", objName, v)
		// 如果 map 的 val 类型不是基础类型，则也要生成对应的类型代码
		if _, ok := v.Elem().(*types.Basic); !ok {
			collectTypes("", v.Elem(), outTypes)
		}
		ss := &ssInfo{
			name: MapTypeName(v),
			typ:  v,
		}
		outTypes[ss.name] = ss
	case *types.Struct:
		fmt.Printf("collect struct types name:%s type:%s. \n", objName, v)
		ss := &ssInfo{
			name: objName,
			typ:  v,
		}
		outTypes[ss.name] = ss
	case *types.Named: // 某个字段是自定义类型 会跑到这里
		// fmt.Printf("collect types is named: %s. \n", tt.Type.String())
		fmt.Printf("collect named types name:%s type:%s. \n", objName, v)
		ss := &ssInfo{
			name: v.Obj().Name(),
			typ:  v.Underlying().(*types.Struct),
		}
		outTypes[ss.name] = ss
	case *types.Slice:
		fmt.Printf("collect slice types name:%s type:%s. \n", objName, v)
		// 如果 slice 的 item 类型不是基础类型，则也要生成对应的类型代码
		if _, ok := v.Elem().(*types.Basic); !ok {
			collectTypes("", v.Elem(), outTypes)
		}

		ss := &ssInfo{
			name: SliceTypeName(v),
			typ:  v,
		}
		outTypes[ss.name] = ss
	case *types.Pointer:
		fmt.Printf("collect pointer types name:%s type:%s. \n", objName, v)
		collectTypes(objName, v.Elem(), outTypes)
	default:
		fmt.Printf("collect other types name:%s type:%s. \n", objName, tt)
	}
}

func generate(sourceTypeName string, structType types.Type) error {
	fmt.Printf("begin generate type:%s\n", sourceTypeName)
	// 1. 获取环境变量，初始化要生成的定义的文件`句柄`
	goPackage := os.Getenv("GOPACKAGE")

	f := NewFile(goPackage)

	// 第一行写注释
	f.PackageComment("Code generated by generator, DO NOT EDIT.")

	desTypeName := ""
	var err error
	// 2. 写这个类型
	switch vv := structType.(type) {
	case *types.Struct:
		desTypeName = writeStruct(f, sourceTypeName, vv)
	case *types.Map:
		desTypeName, err = writeMap(f, vv)
		if err != nil {
			failErr(err)
		}
	case *types.Slice:
		desTypeName, err = writeSlice(f, vv)
		if err != nil {
			failErr(err)
		}
	default:
		fmt.Printf("cannot generate the type %s\n", vv)
	}

	// 3. dump 到 文件
	// goFile := os.Getenv("GOFILE")
	// ext := filepath.Ext(goFile)
	// baseFilename := goFile[0 : len(goFile)-len(ext)]
	// targetFilename := baseFilename + "_" + strings.ToLower(desTypeName) + "_gen.go"
	targetFilename := strings.ToLower(desTypeName) + "_gen.go"

	fmt.Printf("save generate type:%s to file %s\n", sourceTypeName, targetFilename)

	return f.Save(targetFilename)
}

func loadPackage(path string) *packages.Package {
	cfg := &packages.Config{
		Mode: packages.NeedTypes | packages.NeedImports | packages.NeedTypesInfo,
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

// 解析字符串，拿到哪个包里面的，哪个类型字符串
func splitSourceType(sourceType string) (string, string) {
	idx := strings.LastIndexByte(sourceType, '.')
	if idx == -1 {
		failErr(fmt.Errorf(`expected qualified type as "pkg/path.Mytype"`))
	}
	sourceTypePackage := sourceType[0:idx]
	sourceTypeName := sourceType[idx+1:]
	return sourceTypePackage, sourceTypeName
}

// 判断是否生成单个类型，还是生成整个包
func isGenerateSingleType(sourePath string) bool {
	// 如果要生成的是 包名/路径.类型格式的，则加载指定类型
	paths := strings.Split(sourePath, "/")
	lastPath := paths[len(paths)-1]

	return len(strings.Split(lastPath, ".")) >= 2
}

func failErr(err error) {
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error: %v\n", err)
		os.Exit(1)
	}
}
