
会报如下错误信息：
```go
# command-line-arguments[command-line-arguments.test]
./xxx_test.go:30:14: undefined: xxxx
```

需要的操作是：
1. 执行`go mod init xxx`
2. 执行`go mod tidy`

这里实际上我执行了`go mod tidy`，但是有答案说`go build`也可以，所以`go build`和`go mod tidy`有什么区别呢？

`go mod tidy`：添加需要用到但go.mod中查不到的模块、删除未使用的模块

`go build`：在当前目录下编译生成可执行文件

然后运行，返回测试结果：
```go
=== RUN   TestHelloTom
    unit_test_example_test.go:17: Expected Tom do not match actual Jerry
--- FAIL: TestHelloTom (0.00s)
```

