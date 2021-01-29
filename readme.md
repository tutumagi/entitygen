>`jen`的用法：[参考代码](https://dev.to/hlubek/metaprogramming-with-go-or-how-to-build-code-generators-that-parse-go-code-2k3j)

1. 使用 `go generate ./...` 需要在包里面写上对应的执行命令
`//go:generate go run entitygen entitygen/domain`