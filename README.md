# Minotaur

[![Go doc](https://img.shields.io/badge/go.dev-reference-brightgreen?logo=go&logoColor=white&style=flat)](https://pkg.go.dev/github.com/kercylan98/minotaur)
![](https://img.shields.io/badge/Email-kercylan@gmail.com-green.svg?style=flat)
![](https://komarev.com/ghpvc/?username=kercylan98)
<a target="_blank" href="https://goreportcard.com/report/github.com/kercylan98/minotaur"><img src="https://goreportcard.com/badge/github.com/kercylan98/minotaur?style=flat-square" /></a>

Minotaur 是一个基于Golang 1.20 编写的服务端开发支持库，其中采用了大量泛型设计，用于游戏服务器开发。

## 目录结构概况
```mermaid
mindmap
  root((Minotaur))
    /component 通用组件接口定义
      /components 通用组件内置实现
    /config 针对配置导表的配置加载
    /game 游戏通用功能接口定义
      /builtin 游戏通用功能内置实现
    /notify 通知功能接口定义
    /planner 策划相关工具目录
      /configexport 配置导表功能实现
    /report 数据埋点及上报功能
    /server 网络服务器支持
      /cross 内置跨服功能实现
      /router 内置路由器功能实现
    /tools 第三方工具
    /utils 常用辅助函数包
```

## Server 架构预览
![server-gdi.jpg](.github/images/server-gdi.jpg)

## 安装
注意：依赖于 **[Go](https://go.dev/) 1.20 +**

运行以下 Go 命令来安装软件包：`minotaur`
```sh
$ go get -u github.com/kercylan98/minotaur
```

## 用法
- 在`Minotaur`中大量使用了 **[泛型](https://go.dev/doc/tutorial/generics)** 、 **[观察者(事件)](https://www.runoob.com/design-pattern/observer-pattern.html)** 和 **[选项模式](https://juejin.cn/post/6844903729313873927)**，在使用前建议先进行相应了解；
- 项目文档可访问 **[pkg.go.dev](https://pkg.go.dev/github.com/kercylan98/minotaur)** 进行查阅；

### 本地文档
可使用 `godoc` 搭建本地文档服务器
#### 安装 godoc
```shell
git clone golang.org/x/tools
cd tools/cmd
go install ...
```
#### 使用 `godoc` 启动本地文档服务器
```shell
godoc -http=:9998 -play
```
#### Windows
```shell
.\local-doc.bat
```

#### Linux or MacOS
```shell
chmod 777 ./local-doc.sh
./local-doc.sh
```

#### 文档地址
**[http://localhost:9998/pkg/github.com/kercylan98/minotaur/](http://localhost:9998/pkg/github.com/kercylan98/minotaur/)**

### 简单示例
创建一个基于Websocket的回响服务器。
```go
package main

import (
	"github.com/kercylan98/minotaur/server"
)

func main() {
	srv := server.New(server.NetworkWebsocket)
	srv.RegConnectionReceiveWebsocketPacketEvent(func(srv *server.Server, conn *server.Conn, packet []byte, messageType int) {
		conn.Write(packet, messageType)
	})
	if err := srv.Run(":9999"); err != nil {
		panic(err)
	}
}
```
访问 **[WebSocket 在线测试](http://www.websocket-test.com/)** 进行验证。
> Websocket地址: ws://127.0.0.1:9999

### 事件
在 `Miontaur` 中，绝大多数场景都采用事件注册的方式进行回调，可通过`XXX.RegXXXEvent`的方式进行事件注册，例如：
```go
package main

import (
	"github.com/kercylan98/minotaur/server"
)

func main() {
	srv := server.New(server.NetworkWebsocket)
	srv.RegConsoleCommandEvent("command", nil)
	srv.RegStartBeforeEvent(nil)
	srv.RegStartFinishEvent(nil)
	srv.RegConnectionClosedEvent(nil)
	srv.RegConnectionOpenedEvent(nil)
	srv.RegConnectionReceivePacketEvent(nil)
	srv.RegConnectionReceiveWebsocketPacketEvent(nil)
	srv.RegReceiveCrossPacketEvent(nil)
	srv.RegMessageErrorEvent(nil)
	srv.RegMessageLowExecEvent(nil)
}
```
其他的一些支持事件的结构体（非所有）：
 - `game.Room` 游戏房间实现
 - `synchronization.Map` 并发安全的`Map`实现
 - ...
### 可选项
大部分的 `New` 函数均可使用可选项进行创建，具体函数前缀通常为 `With`。
```go
package main

import (
	"github.com/kercylan98/minotaur/server"
)

func main() {
	server.New(server.NetworkWebsocket, 
		server.WithTicker(50, false),
		server.WithProd(),
		server.WithMultiCore(3),
		// ...
	)
}
```
### 更多的网络模式
```go
package main

import (
	"github.com/kercylan98/minotaur/server"
)

func main() {
	server.New(server.NetworkTcp)
	server.New(server.NetworkTcp4)
	server.New(server.NetworkTcp6)
	server.New(server.NetworkUdp)
	server.New(server.NetworkUdp4)
	server.New(server.NetworkUdp6)
	server.New(server.NetworkWebsocket)
	server.New(server.NetworkKcp)
	server.New(server.NetworkUnix)
	server.New(server.NetworkHttp)
	server.New(server.NetworkGRPC)
}
```

# JetBrains OS licenses

`Minotaur` had been being developed with `GoLand` IDE under the **free JetBrains Open Source license(s)** granted by JetBrains s.r.o., hence I would like to express my thanks here.

<a href="https://www.jetbrains.com/?from=minotaur" target="_blank"><img src="https://resources.jetbrains.com/storage/products/company/brand/logos/jb_beam.png?_gl=1*1vt713y*_ga*MTEzMjEzODQxNC4xNjc5OTY3ODUw*_ga_9J976DJZ68*MTY4ODU0MDUyMy4yMC4xLjE2ODg1NDA5NDAuMjUuMC4w&_ga=2.261225293.1519421387.1688540524-1132138414.1679967850" width="250" align="middle"/></a>