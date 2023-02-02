# 1. Gorm

Gorm是一个已经迭代了10年+的功能强大的ORM框架，在字节内部被广泛使用并且拥有非常丰富的开源扩展。

如果查询或者更新时结构体中有条件为零值，需要使用Map传。

Gorm删除数据包括物理删除和软删除，拥有软删除能力的Model调用Delete时，记录不会被从数据库中真正删除。但GORM会将DeletedAt置为当前时间，并且你不能再通过正常的查询方法找到该记录。使用Unscoped可以查询到被软删的数据。

事务上，GORM提供了Begin、Commit、Rollback方法用于使用事务。

GORM在提供了CURD的Hook能力。Hook是在创建、查询、更新、删除等操作之前、之后自动调用的函数。
如果任何Hook返回错误，GORM将停止后续的操作并回滚事务。

对于写操作(创建、更新、删除)， 为了确保数据的完整性，GORM会将它们封装在事务内运行。
但这会降低性能，你可以使用SkipDefaultTransaction关闭默认事务。使用PrepareStmt缓存预编译语句可以提高后续调用的速度，本机测试提高大约35 %左右。

```go
package main 
import ( 
	"gorm.io/gorm" 
	"gorm.io/driver/sqlite" 
) 
type Product struct { 
	gorm.Model 
	Code string 
	Price uint 
} 
func main() { 
	db, err := gorm.Open(sqlite.Open("test.db"), &gorm.Config{}) 
	if err != nil { 
	panic("failed to connect database") 
	} 
	
	// 迁移 schema 
	db.AutoMigrate(&Product{}) 
	
	// Create 
	db.Create(&Product{Code: "D42", Price: 100}) 
	
	// Read 
	var product Product db.First(&product, 1) // 根据整型主键查找 
	db.First(&product, "code = ?", "D42") // 查找 code 字段值为 D42 的记录 
	
	// Update - 将 product 的 price 更新为 200 
	db.Model(&product).Update("Price", 200) 
	// Update - 更新多个字段 
	db.Model(&product).Updates(Product{Price: 200, Code: "F42"}) // 仅更新非零值字段 
	db.Model(&product).Updates(map[string]interface{}{"Price": 200, "Code": "F42"}) 
	
	// Delete - 删除 product 
	db.Delete(&product, 1) 
}
```

生态
GORM代码生成工具 https://github.com/go-gorm/gen
GORM分片库方案 https://github.com/go-gorm/sharding
GORM手动索引 https://github.com/go-gorm/hints
GORM乐观锁  https://github.com/go-gorm/optimisticlock
GORM读写分离  https://github.com/go-gorm/dbresolver
GORM OpenTelemetry扩展  https://github.com/go-gorm/opentelemetry




# 2. Kitex

Kitex是字节内部的Golang微服务RPC框架，具有高性能、强可扩展的主要特点，支持多协议并且拥有丰富的开源扩展。

## 2.1 什么是RPC？

RPC（Remote procedure call，远程过程调用）是指计算机程序导致过程（子例程）在不同的地址空间（通常在共享网络上的另一台计算机上）中执行，其编码就像普通（本地）过程调用一样，程序员没有显式编码远程交互的详细信息。也就是说，程序员编写的代码基本相同，无论子例程是执行程序的本地还是远程的。简单来说，通过使用 RPC，我们可以像调用方法一样快捷的与远程服务进行交互。

## 2.2 Kitex相关内容

如果本地开发环境是Windows建议是使用虚拟机或者WSL2。

```go
// 安装代码生成工具
go install github.com/cloudwego/kitex/tool/cmd/kitex@latest 
go install github.com/cloudwego/thriftgo@latest
```

## 2.3 定义IDL

IDL是Interface description language的缩写，指接口描述语言，是规范的一部分，是跨平台开发的基础。我们可以使用 IDL 来支持 RPC 的信息传输定义。Kitex 默认支持 `thrift` 和 `proto3` 两种 IDL，而在底层传输上，Kitex 使用扩展的 `thrift` 作为底层的传输协议。

## 2.4 生成代码

```arduino
kitex -module example -service example echo.thrift
```

生成后的项目结构如下：

```go
. 
|-- build.sh 
|-- echo.thrift 
|-- handler.go 
|-- kitex_gen 
|   `-- api 
|       |-- echo 
|       |   |-- client.go 
|       |   |-- echo.go 
|       |   |-- invoker.go 
|       |   `-- server.go 
|       |-- echo.go 
|       `-- k-echo.go 
|-- main.go 
`-- script   
	|-- bootstrap.sh   
	`-- settings.py
```

- build.sh：构建脚本
- kitex_gen：IDL内容相关的生成代码，主要是基础的Server/Client代码
- main.go：程序入口
- handler.go：用户在该文件里实现IDLservice定义的方法

服务默认监听8888端口，

```go
package main ​ 
import (        
	"context"        
	api "exmaple/kitex_gen/api" 
) ​ 

// EchoImpl implements the last service interface defined in the IDL. 
type EchoImpl struct{} ​ 

// Echo implements the EchoImpl interface. 
func (s *EchoImpl) Echo(ctx context.Context, req *api.Request) (resp *api.Response, err error) {        
// TODO: Your code here...        
	return 
}
```

以上是服务端代码，下面可以创建一个客户端。

```go
import "example/kitex_gen/api/echo" 
import "github.com/cloudwego/kitex/client" 
... 
c, err := echo.NewClient("example", client.WithHostPorts("0.0.0.0:8888")) 
if err != nil {  
	log.Fatal(err) 
}
```

此外，Kitex还对接了主流的服务注册与发现中心，如ETCD，Nacos等。

Kitex生态：

详见：`https://www.cloudwego.io/zh/docs/kitex/`

# 3.Hertz

Hertz是字节内部的HTTP框架，参考了其他开源框架的优势，结合字节跳动内部的需求，具有高易用性、高性能、高扩展性特点。

**Hertz路由**：
- Hertz提供了GET、POST、PUT、DELETE、ANY等方法用于注册路由；
- 提供了路由组Group的能力，用于支持路由分组功能；
- 提供了参数路由和通配路由，路由的优先级为：静态路由 > 命名路由 > 通配路由

**Hertz参数绑定**：
- Hertz提供了Bind、Validate、BindAndValidate函数用于进行参数绑定和校验；

**Hertz中间件**：
- Hertz的中间件主要分为客户端中间件与服务端中间件；

**Hertz Client**：
- Hertz提供了HTTP Client用于帮助用户发送HTTP请求；

**Hertz代码生成工具**：
- Hertz提供了代码生成工具Hz，通过定义IDL文件即可生成对应的基础服务代码；




