# sono-server-test

###sono-server的本地测试注意事项：

1. 需要自己在本地的sono-server目录下新建logs目录
2. conf模块下的conf.toml中需要将logs的目录设置为绝对路径
3. src/conf/conf.go中将conf文件的路径设置为绝对路径,如下：（否则在每个***_test.go里面都读取不到配置文件，因为路径变了）

```
func init() {
    flag.StringVar(&confPath, "conf", "/Users/sonofelice/goframe/conf/conf.toml", "-conf path")
}

```
