# Client

[![Go doc](https://img.shields.io/badge/go.dev-reference-brightgreen?logo=go&logoColor=white&style=flat)](https://pkg.go.dev/github.com/kercylan98/minotaur)
![](https://img.shields.io/badge/Email-kercylan@gmail.com-green.svg?style=flat)

暂无介绍...


## 目录导航
列出了该 `package` 下所有的函数及类型定义，可通过目录导航进行快捷跳转 ❤️
<details>
<summary>展开 / 折叠目录导航</summary>


> 包级函数定义

|函数名称|描述
|:--|:--
|[NewClient](#NewClient)|创建客户端
|[CloneClient](#CloneClient)|克隆客户端
|[NewTCP](#NewTCP)|暂无描述...
|[NewUnixDomainSocket](#NewUnixDomainSocket)|暂无描述...
|[NewWebsocket](#NewWebsocket)|创建 websocket 客户端


> 类型定义

|类型|名称|描述
|:--|:--|:--
|`STRUCT`|[Client](#client)|客户端
|`INTERFACE`|[Core](#core)|暂无描述...
|`STRUCT`|[ConnectionClosedEventHandle](#connectionclosedeventhandle)|暂无描述...
|`STRUCT`|[Packet](#packet)|暂无描述...
|`STRUCT`|[TCP](#tcp)|暂无描述...
|`STRUCT`|[UnixDomainSocket](#unixdomainsocket)|暂无描述...
|`STRUCT`|[Websocket](#websocket)|websocket 客户端

</details>


***
## 详情信息
#### func NewClient(core Core) *Client
<span id="NewClient"></span>
> 创建客户端

***
#### func CloneClient(client *Client) *Client
<span id="CloneClient"></span>
> 克隆客户端

***
#### func NewTCP(addr string) *Client
<span id="NewTCP"></span>

***
#### func NewUnixDomainSocket(addr string) *Client
<span id="NewUnixDomainSocket"></span>

***
#### func NewWebsocket(addr string) *Client
<span id="NewWebsocket"></span>
> 创建 websocket 客户端

***
### Client `STRUCT`
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
<details>
<summary>查看 / 收起单元测试</summary>


```go

func TestClient_WriteWS(t *testing.T) {
	var wait sync.WaitGroup
	wait.Add(1)
	srv := server.New(server.NetworkWebsocket)
	srv.RegConnectionReceivePacketEvent(func(srv *server.Server, conn *server.Conn, packet []byte) {
		srv.Shutdown()
	})
	srv.RegStopEvent(func(srv *server.Server) {
		wait.Done()
	})
	srv.RegMessageReadyEvent(func(srv *server.Server) {
		cli := client.NewWebsocket("ws://127.0.0.1:9999")
		cli.RegConnectionOpenedEvent(func(conn *client.Client) {
			conn.WriteWS(2, []byte("Hello"))
		})
		if err := cli.Run(); err != nil {
			panic(err)
		}
	})
	if err := srv.Run(":9999"); err != nil {
		panic(err)
	}
	wait.Wait()
}

```


</details>


***
#### func (*Client) Write(packet []byte, callback ...func (err error))
> 向连接中写入数据
***
#### func (*Client) GetServerAddr()  string
> 获取服务器地址
***
### Core `INTERFACE`

```go
type Core interface {
	Run(runState chan error, receive func(wst int, packet []byte))
	Write(packet *Packet) error
	Close()
	GetServerAddr() string
	Clone() Core
}
```
### ConnectionClosedEventHandle `STRUCT`

```go
type ConnectionClosedEventHandle func(conn *Client, err any)
```
### Packet `STRUCT`

```go
type Packet struct {
	wst      int
	data     []byte
	callback func(err error)
}
```
### TCP `STRUCT`

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
### UnixDomainSocket `STRUCT`

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
<details>
<summary>查看 / 收起单元测试</summary>


```go

func TestUnixDomainSocket_Write(t *testing.T) {
	var closed = make(chan struct{})
	srv := server.New(server.NetworkUnix)
	srv.RegConnectionReceivePacketEvent(func(srv *server.Server, conn *server.Conn, packet []byte) {
		t.Log(string(packet))
		conn.Write(packet)
	})
	srv.RegStartFinishEvent(func(srv *server.Server) {
		time.Sleep(time.Second)
		cli := client.NewUnixDomainSocket("./test.sock")
		cli.RegConnectionOpenedEvent(func(conn *client.Client) {
			conn.Write([]byte("Hello~"))
		})
		cli.RegConnectionReceivePacketEvent(func(conn *client.Client, wst int, packet []byte) {
			t.Log(packet)
			closed <- struct{}{}
		})
		if err := cli.Run(); err != nil {
			panic(err)
		}
	})
	go func() {
		if err := srv.Run("./test.sock"); err != nil {
			panic(err)
		}
	}()
	<-closed
	srv.Shutdown()
}

```


</details>


***
#### func (*UnixDomainSocket) Close()
***
#### func (*UnixDomainSocket) GetServerAddr()  string
***
#### func (*UnixDomainSocket) Clone()  Core
***
### Websocket `STRUCT`
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
