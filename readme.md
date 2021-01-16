[参考代码](https://dev.to/hlubek/metaprogramming-with-go-or-how-to-build-code-generators-that-parse-go-code-2k3j)

1. 生成指定类型的实体定义
`go run main.go entitygen/domain.LandAttr`

2. 生成整个包的实体定义
`go run main.go entitygen/domain`

3. 使用 `go generate ./...` 需要在包里面写上对应的执行命令
`//go:generate go run entitygen entitygen/domain.LandAttr`