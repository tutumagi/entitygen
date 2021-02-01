>`jen`的用法：[参考代码](https://dev.to/hlubek/metaprogramming-with-go-or-how-to-build-code-generators-that-parse-go-code-2k3j)


#### 用法
##### 第一次生成
1. 确保 `go env` 中的环境变量如下（如果不是，执行下面这个命令 `go env -w GOPRIVATE=gitlab.gamesword.com`）
    ```
    GONOPROXY="gitlab.gamesword.com"
    GONOSUMDB="gitlab.gamesword.com"
    GOPRIVATE="gitlab.gamesword.com"
    ```

2. 新建一个实体定义的 package，在里面定义实体，tag 中各个key-value 的含义见 `domain/room.go` 中的描述，目前 map 类型的 key 只支持 `int32` 和 `string`

3. 在导出实体文件的目标 package 中新建一个导出定义文件，将 `entitydef/mappping.go` 里面的文件内容拷贝过去

4. 在一级目录下，执行 `go generate -v ./...`

##### 如果是修改实体，或者新增实体
1. 在实体定义 package 里面，新增实体定义
2. 在一级目录下，执行 `go generate -v ./...`