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
|`STRUCT`|[Client](#struct_Client)|客户端
|`INTERFACE`|[Core](#struct_Core)|暂无描述...
|`STRUCT`|[ConnectionClosedEventHandle](#struct_ConnectionClosedEventHandle)|暂无描述...
|`STRUCT`|[Packet](#struct_Packet)|暂无描述...
|`STRUCT`|[TCP](#struct_TCP)|暂无描述...
|`STRUCT`|[UnixDomainSocket](#struct_UnixDomainSocket)|暂无描述...
|`STRUCT`|[Websocket](#struct_Websocket)|websocket 客户端

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
<span id="struct_Client"></span>
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
<span id="struct_Client_Run"></span>

#### func (*Client) Run(block ...bool)  error
> 运行客户端，当客户端已运行时，会先关闭客户端再重新运行
>   - block 以阻塞方式运行

***
<span id="struct_Client_RunByBufferSize"></span>

#### func (*Client) RunByBufferSize(size int, block ...bool)  error
> 指定写入循环缓冲区大小运行客户端，当客户端已运行时，会先关闭客户端再重新运行
>   - block 以阻塞方式运行

***
<span id="struct_Client_IsConnected"></span>

#### func (*Client) IsConnected()  bool
> 是否已连接

***
<span id="struct_Client_Close"></span>

#### func (*Client) Close(err ...error)
> 关闭

***
<span id="struct_Client_WriteWS"></span>

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
<span id="struct_Client_Write"></span>

#### func (*Client) Write(packet []byte, callback ...func (err error))
> 向连接中写入数据

***
<span id="struct_Client_GetServerAddr"></span>

#### func (*Client) GetServerAddr()  string
> 获取服务器地址

***
<span id="struct_Core"></span>
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
<span id="struct_ConnectionClosedEventHandle"></span>
### ConnectionClosedEventHandle `STRUCT`

```go
type ConnectionClosedEventHandle func(conn *Client, err any)
```
<span id="struct_Packet"></span>
### Packet `STRUCT`

```go
type Packet struct {
	wst      int
	data     []byte
	callback func(err error)
}
```
<span id="struct_TCP"></span>
### TCP `STRUCT`

```go
type TCP struct {
	conn   net.Conn
	addr   string
	closed bool
}
```
<span id="struct_TCP_Run"></span>

#### func (*TCP) Run(runState chan error, receive func (wst int, packet []byte))

***
<span id="struct_TCP_Write"></span>

#### func (*TCP) Write(packet *Packet)  error

***
<span id="struct_TCP_Close"></span>

#### func (*TCP) Close()

***
<span id="struct_TCP_GetServerAddr"></span>

#### func (*TCP) GetServerAddr()  string

***
<span id="struct_TCP_Clone"></span>

#### func (*TCP) Clone()  Core

***
<span id="struct_UnixDomainSocket"></span>
### UnixDomainSocket `STRUCT`

```go
type UnixDomainSocket struct {
	conn   net.Conn
	addr   string
	closed bool
}
```
<span id="struct_UnixDomainSocket_Run"></span>

#### func (*UnixDomainSocket) Run(runState chan error, receive func (wst int, packet []byte))

***
<span id="struct_UnixDomainSocket_Write"></span>

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
<span id="struct_UnixDomainSocket_Close"></span>

#### func (*UnixDomainSocket) Close()

***
<span id="struct_UnixDomainSocket_GetServerAddr"></span>

#### func (*UnixDomainSocket) GetServerAddr()  string

***
<span id="struct_UnixDomainSocket_Clone"></span>

#### func (*UnixDomainSocket) Clone()  Core

***
<span id="struct_Websocket"></span>
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
<span id="struct_Websocket_Run"></span>

#### func (*Websocket) Run(runState chan error, receive func (wst int, packet []byte))

***
<span id="struct_Websocket_Write"></span>

#### func (*Websocket) Write(packet *Packet)  error

***
<span id="struct_Websocket_Close"></span>

#### func (*Websocket) Close()

***
<span id="struct_Websocket_GetServerAddr"></span>

#### func (*Websocket) GetServerAddr()  string

***
<span id="struct_Websocket_Clone"></span>

#### func (*Websocket) Clone()  Core

***
