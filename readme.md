>`jen`的用法：[参考代码](https://dev.to/hlubek/metaprogramming-with-go-or-how-to-build-code-generators-that-parse-go-code-2k3j)


#### 修改/维护
1. 主要是修改 `entitygen/` 和 `domain/` 目录下的代码，`attr/` 目录下的改动需要讨论才可以改动
2. 修改完成后，执行  `go generate ./...`，生成成功后， `go test -coverprofile=./profile/cover.txt ./...` 执行测试，测试通过再提交


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



#### TODO
1. [x] 类型 ID, UID 这样命名的字段如果改为合适的 key，应该改为 id,uid 比较合适。目前的策略是 实体定义的 field 名字为大写开头驼峰，
   然后代码转换时，是转为小写开头的驼峰写法。不过这种对`ID`,`UID` 就很不友好了，输出的是 `iD` 和 `uID` 这个要想一下办法
   解决办法：需要自定义 key 的在定义文件里面写好 key 的 tag 就好了
2. 针对通用的 pos(`*math.Vector3`), rot(`*math.Vector3`) 看能否抽出来，不放到实体定义里面去