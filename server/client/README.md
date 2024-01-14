# Client



[![Go doc](https://img.shields.io/badge/go.dev-reference-brightgreen?logo=go&logoColor=white&style=flat)](https://pkg.go.dev/github.com/kercylan98/minotaur/client)
![](https://img.shields.io/badge/Email-kercylan@gmail.com-green.svg?style=flat)

## 目录
列出了该 `package` 下所有的函数，可通过目录进行快捷跳转 ❤️
<details>
<summary>展开 / 折叠目录</summary


> 包级函数定义

|函数|描述
|:--|:--
|[NewClient](#NewClient)|创建客户端
|[CloneClient](#CloneClient)|克隆客户端
|[NewTCP](#NewTCP)|暂无描述...
|[NewUnixDomainSocket](#NewUnixDomainSocket)|暂无描述...
|[NewWebsocket](#NewWebsocket)|创建 websocket 客户端


> 结构体定义

|结构体|描述
|:--|:--
|[Client](#client)|客户端
|[Core](#core)|暂无描述...
|[ConnectionClosedEventHandle](#connectionclosedeventhandle)|暂无描述...
|[Packet](#packet)|暂无描述...
|[TCP](#tcp)|暂无描述...
|[UnixDomainSocket](#unixdomainsocket)|暂无描述...
|[Websocket](#websocket)|websocket 客户端

</details>


#### func NewClient(core Core)  *Client
<span id="NewClient"></span>
> 创建客户端
***
#### func CloneClient(client *Client)  *Client
<span id="CloneClient"></span>
> 克隆客户端
***
#### func NewTCP(addr string)  *Client
<span id="NewTCP"></span>
***
#### func NewUnixDomainSocket(addr string)  *Client
<span id="NewUnixDomainSocket"></span>
***
#### func NewWebsocket(addr string)  *Client
<span id="NewWebsocket"></span>
> 创建 websocket 客户端
***
### Client
客户端
```go
type Client struct {
	*events
	core           Core
	mutex          sync.Mutex
	closed         bool
	pool           *hub.ObjectPool[*Packet]
	loop           *writeloop.Channel[*Packet]
	loopBufferSize int
	block          chan struct{}
}
```
#### func (*Client) Run(block ...bool)  error
> 运行客户端，当客户端已运行时，会先关闭客户端再重新运行
>   - block 以阻塞方式运行
***
#### func (*Client) RunByBufferSize(size int, block ...bool)  error
> 指定写入循环缓冲区大小运行客户端，当客户端已运行时，会先关闭客户端再重新运行
>   - block 以阻塞方式运行
***
#### func (*Client) IsConnected()  bool
> 是否已连接
***
#### func (*Client) Close(err ...error)
> 关闭
***
#### func (*Client) WriteWS(wst int, packet []byte, callback ...func (err error))
> 向连接中写入指定 websocket 数据类型
>   - wst: websocket模式中指定消息类型
***
#### func (*Client) Write(packet []byte, callback ...func (err error))
> 向连接中写入数据
***
#### func (*Client) GetServerAddr()  string
> 获取服务器地址
***
### Core

```go
type Core struct{}
```
### ConnectionClosedEventHandle

```go
type ConnectionClosedEventHandle struct{}
```
### Packet

```go
type Packet struct {
	wst      int
	data     []byte
	callback func(err error)
}
```
### TCP

```go
type TCP struct {
	conn   net.Conn
	addr   string
	closed bool
}
```
#### func (*TCP) Run(runState chan error, receive func (wst int, packet []byte))
***
#### func (*TCP) Write(packet *Packet)  error
***
#### func (*TCP) Close()
***
#### func (*TCP) GetServerAddr()  string
***
#### func (*TCP) Clone()  Core
***
### UnixDomainSocket

```go
type UnixDomainSocket struct {
	conn   net.Conn
	addr   string
	closed bool
}
```
#### func (*UnixDomainSocket) Run(runState chan error, receive func (wst int, packet []byte))
***
#### func (*UnixDomainSocket) Write(packet *Packet)  error
***
#### func (*UnixDomainSocket) Close()
***
#### func (*UnixDomainSocket) GetServerAddr()  string
***
#### func (*UnixDomainSocket) Clone()  Core
***
### Websocket
websocket 客户端
```go
type Websocket struct {
	addr   string
	conn   *websocket.Conn
	closed bool
	mu     sync.Mutex
}
```
#### func (*Websocket) Run(runState chan error, receive func (wst int, packet []byte))
***
#### func (*Websocket) Write(packet *Packet)  error
***
#### func (*Websocket) Close()
***
#### func (*Websocket) GetServerAddr()  string
***
#### func (*Websocket) Clone()  Core
***
