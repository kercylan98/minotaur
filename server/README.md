# Server

[![Go doc](https://img.shields.io/badge/go.dev-reference-brightgreen?logo=go&logoColor=white&style=flat)](https://pkg.go.dev/github.com/kercylan98/minotaur)
![](https://img.shields.io/badge/Email-kercylan@gmail.com-green.svg?style=flat)

server 提供了包含多种网络类型的服务器实现


## 目录导航
列出了该 `package` 下所有的函数及类型定义，可通过目录导航进行快捷跳转 ❤️
<details>
<summary>展开 / 折叠目录导航</summary>


> 包级函数定义

|函数名称|描述
|:--|:--
|[NewBot](#NewBot)|创建一个机器人，目前仅支持 Socket 服务器
|[WithBotNetworkDelay](#WithBotNetworkDelay)|设置机器人网络延迟及波动范围
|[WithBotWriter](#WithBotWriter)|设置机器人写入器，默认为 os.Stdout
|[DefaultWebsocketUpgrader](#DefaultWebsocketUpgrader)|暂无描述...
|[NewHttpHandleWrapper](#NewHttpHandleWrapper)|创建一个新的 http 处理程序包装器
|[NewHttpContext](#NewHttpContext)|基于 gin.Context 创建一个新的 HttpContext
|[NewGinWrapper](#NewGinWrapper)|创建 gin 包装器，用于对 NewHttpWrapper 函数的替代
|[HasMessageType](#HasMessageType)|检查是否存在指定的消息类型
|[NewMultipleServer](#NewMultipleServer)|暂无描述...
|[GetNetworks](#GetNetworks)|获取所有支持的网络模式
|[WithLowMessageDuration](#WithLowMessageDuration)|通过指定慢消息时长的方式创建服务器，当消息处理时间超过指定时长时，将会输出 WARN 类型的日志
|[WithAsyncLowMessageDuration](#WithAsyncLowMessageDuration)|通过指定异步消息的慢消息时长的方式创建服务器，当消息处理时间超过指定时长时，将会输出 WARN 类型的日志
|[WithWebsocketConnInitializer](#WithWebsocketConnInitializer)|通过 websocket 连接初始化的方式创建服务器，当 initializer 返回错误时，服务器将不会处理该连接的后续逻辑
|[WithWebsocketUpgrade](#WithWebsocketUpgrade)|通过指定 websocket.Upgrader 的方式创建服务器
|[WithConnWriteBufferSize](#WithConnWriteBufferSize)|通过连接写入缓冲区大小的方式创建服务器
|[WithDispatcherBufferSize](#WithDispatcherBufferSize)|通过消息分发器缓冲区大小的方式创建服务器
|[WithMessageStatistics](#WithMessageStatistics)|通过消息统计的方式创建服务器
|[WithPacketWarnSize](#WithPacketWarnSize)|通过数据包大小警告的方式创建服务器，当数据包大小超过指定大小时，将会输出 WARN 类型的日志
|[WithLimitLife](#WithLimitLife)|通过限制最大生命周期的方式创建服务器
|[WithWebsocketWriteCompression](#WithWebsocketWriteCompression)|通过数据写入压缩的方式创建Websocket服务器
|[WithWebsocketCompression](#WithWebsocketCompression)|通过数据压缩的方式创建Websocket服务器
|[WithDeadlockDetect](#WithDeadlockDetect)|通过死锁、死循环、永久阻塞检测的方式创建服务器
|[WithDisableAsyncMessage](#WithDisableAsyncMessage)|通过禁用异步消息的方式创建服务器
|[WithAsyncPoolSize](#WithAsyncPoolSize)|通过指定异步消息池大小的方式创建服务器
|[WithWebsocketReadDeadline](#WithWebsocketReadDeadline)|设置 Websocket 读取超时时间
|[WithTicker](#WithTicker)|通过定时器创建服务器，为服务器添加定时器功能
|[WithTLS](#WithTLS)|通过安全传输层协议TLS创建服务器
|[WithGRPCServerOptions](#WithGRPCServerOptions)|通过GRPC的可选项创建GRPC服务器
|[WithWebsocketMessageType](#WithWebsocketMessageType)|设置仅支持特定类型的Websocket消息
|[WithPProf](#WithPProf)|通过性能分析工具PProf创建服务器
|[New](#New)|根据特定网络类型创建一个服务器
|[BindService](#BindService)|绑定服务到特定 Server，被绑定的服务将会在 Server 初始化时执行 Service.OnInit 方法


> 类型定义

|类型|名称|描述
|:--|:--|:--
|`STRUCT`|[Bot](#struct_Bot)|暂无描述...
|`STRUCT`|[BotOption](#struct_BotOption)|暂无描述...
|`STRUCT`|[Conn](#struct_Conn)|服务器连接单次消息的包装
|`STRUCT`|[ConsoleParams](#struct_ConsoleParams)|控制台参数
|`STRUCT`|[MessageReadyEventHandler](#struct_MessageReadyEventHandler)|暂无描述...
|`STRUCT`|[Http](#struct_Http)|基于 gin.Engine 包装的 http 服务器
|`STRUCT`|[HttpContext](#struct_HttpContext)|基于 gin.Context 的 http 请求上下文
|`STRUCT`|[HandlerFunc](#struct_HandlerFunc)|暂无描述...
|`STRUCT`|[ContextPacker](#struct_ContextPacker)|暂无描述...
|`STRUCT`|[HttpRouter](#struct_HttpRouter)|暂无描述...
|`STRUCT`|[HttpWrapperHandleFunc](#struct_HttpWrapperHandleFunc)|暂无描述...
|`STRUCT`|[HttpWrapper](#struct_HttpWrapper)|http 包装器
|`STRUCT`|[HttpWrapperGroup](#struct_HttpWrapperGroup)|http 包装器
|`STRUCT`|[MessageType](#struct_MessageType)|暂无描述...
|`STRUCT`|[Message](#struct_Message)|服务器消息
|`STRUCT`|[MultipleServer](#struct_MultipleServer)|暂无描述...
|`STRUCT`|[Network](#struct_Network)|暂无描述...
|`STRUCT`|[Option](#struct_Option)|暂无描述...
|`STRUCT`|[Server](#struct_Server)|网络服务器
|`INTERFACE`|[Service](#struct_Service)|兼容传统 service 设计模式的接口

</details>


***
## 详情信息
#### func NewBot(srv *Server, options ...BotOption) *Bot
<span id="NewBot"></span>
> 创建一个机器人，目前仅支持 Socket 服务器

<details>
<summary>查看 / 收起单元测试</summary>


```go

func TestNewBot(t *testing.T) {
	srv := server.New(server.NetworkWebsocket)
	srv.RegConnectionOpenedEvent(func(srv *server.Server, conn *server.Conn) {
		t.Logf("connection opened: %s", conn.GetID())
		conn.Close()
		conn.Write([]byte("hello"))
	})
	srv.RegConnectionClosedEvent(func(srv *server.Server, conn *server.Conn, err any) {
		t.Logf("connection closed: %s", conn.GetID())
	})
	srv.RegConnectionReceivePacketEvent(func(srv *server.Server, conn *server.Conn, packet []byte) {
		t.Logf("connection %s receive packet: %s", conn.GetID(), string(packet))
		conn.Write([]byte("world"))
	})
	srv.RegStartFinishEvent(func(srv *server.Server) {
		bot := server.NewBot(srv, server.WithBotNetworkDelay(100, 20), server.WithBotWriter(func(bot *server.Bot) io.Writer {
			return &Writer{t: t, bot: bot}
		}))
		bot.JoinServer()
		time.Sleep(time.Second)
		bot.SendPacket([]byte("hello"))
	})
	_ = srv.Run(":9600")
}

```


</details>


***
#### func WithBotNetworkDelay(delay time.Duration, fluctuation time.Duration) BotOption
<span id="WithBotNetworkDelay"></span>
> 设置机器人网络延迟及波动范围
>   - delay 延迟
>   - fluctuation 波动范围

***
#### func WithBotWriter(construction func (bot *Bot)  io.Writer) BotOption
<span id="WithBotWriter"></span>
> 设置机器人写入器，默认为 os.Stdout

***
#### func DefaultWebsocketUpgrader() *websocket.Upgrader
<span id="DefaultWebsocketUpgrader"></span>

***
#### func NewHttpHandleWrapper\[Context any\](srv *Server, packer ContextPacker[Context]) *Http[Context]
<span id="NewHttpHandleWrapper"></span>
> 创建一个新的 http 处理程序包装器
>   - 默认使用 server.HttpContext 作为上下文，如果需要依赖其作为新的上下文，可以通过 NewHttpContext 创建

***
#### func NewHttpContext(ctx *gin.Context) *HttpContext
<span id="NewHttpContext"></span>
> 基于 gin.Context 创建一个新的 HttpContext

***
#### func NewGinWrapper\[CTX any\](server *gin.Engine, pack func (ctx *gin.Context)  CTX) *HttpWrapper[CTX]
<span id="NewGinWrapper"></span>
> 创建 gin 包装器，用于对 NewHttpWrapper 函数的替代

***
#### func HasMessageType(mt MessageType) bool
<span id="HasMessageType"></span>
> 检查是否存在指定的消息类型

***
#### func NewMultipleServer(serverHandle ...func () ((addr string, srv *Server))) *MultipleServer
<span id="NewMultipleServer"></span>

***
#### func GetNetworks() []Network
<span id="GetNetworks"></span>
> 获取所有支持的网络模式

***
#### func WithLowMessageDuration(duration time.Duration) Option
<span id="WithLowMessageDuration"></span>
> 通过指定慢消息时长的方式创建服务器，当消息处理时间超过指定时长时，将会输出 WARN 类型的日志
>   - 默认值为 DefaultLowMessageDuration
>   - 当 duration <= 0 时，表示关闭慢消息检测

***
#### func WithAsyncLowMessageDuration(duration time.Duration) Option
<span id="WithAsyncLowMessageDuration"></span>
> 通过指定异步消息的慢消息时长的方式创建服务器，当消息处理时间超过指定时长时，将会输出 WARN 类型的日志
>   - 默认值为 DefaultAsyncLowMessageDuration
>   - 当 duration <= 0 时，表示关闭慢消息检测

***
#### func WithWebsocketConnInitializer(initializer func (writer http.ResponseWriter, request *http.Request, conn *websocket.Conn)  error) Option
<span id="WithWebsocketConnInitializer"></span>
> 通过 websocket 连接初始化的方式创建服务器，当 initializer 返回错误时，服务器将不会处理该连接的后续逻辑
>   - 该选项仅在创建 NetworkWebsocket 服务器时有效

***
#### func WithWebsocketUpgrade(upgrader *websocket.Upgrader) Option
<span id="WithWebsocketUpgrade"></span>
> 通过指定 websocket.Upgrader 的方式创建服务器
>   - 默认值为 DefaultWebsocketUpgrader
>   - 该选项仅在创建 NetworkWebsocket 服务器时有效

***
#### func WithConnWriteBufferSize(size int) Option
<span id="WithConnWriteBufferSize"></span>
> 通过连接写入缓冲区大小的方式创建服务器
>   - 默认值为 DefaultConnWriteBufferSize
>   - 设置合适的缓冲区大小可以提高服务器性能，但是会占用更多的内存

***
#### func WithDispatcherBufferSize(size int) Option
<span id="WithDispatcherBufferSize"></span>
> 通过消息分发器缓冲区大小的方式创建服务器
>   - 默认值为 DefaultDispatcherBufferSize
>   - 设置合适的缓冲区大小可以提高服务器性能，但是会占用更多的内存

***
#### func WithMessageStatistics(duration time.Duration, limit int) Option
<span id="WithMessageStatistics"></span>
> 通过消息统计的方式创建服务器
>   - 默认不开启，当 duration 和 limit 均大于 0 的时候，服务器将记录每 duration 期间的消息数量，并保留最多 limit 条

***
#### func WithPacketWarnSize(size int) Option
<span id="WithPacketWarnSize"></span>
> 通过数据包大小警告的方式创建服务器，当数据包大小超过指定大小时，将会输出 WARN 类型的日志
>   - 默认值为 DefaultPacketWarnSize
>   - 当 size <= 0 时，表示不设置警告

***
#### func WithLimitLife(t time.Duration) Option
<span id="WithLimitLife"></span>
> 通过限制最大生命周期的方式创建服务器
>   - 通常用于测试服务器，服务器将在到达最大生命周期时自动关闭

***
#### func WithWebsocketWriteCompression() Option
<span id="WithWebsocketWriteCompression"></span>
> 通过数据写入压缩的方式创建Websocket服务器
>   - 默认不开启数据压缩

***
#### func WithWebsocketCompression(level int) Option
<span id="WithWebsocketCompression"></span>
> 通过数据压缩的方式创建Websocket服务器
>   - 默认不开启数据压缩

***
#### func WithDeadlockDetect(t time.Duration) Option
<span id="WithDeadlockDetect"></span>
> 通过死锁、死循环、永久阻塞检测的方式创建服务器
>   - 当检测到死锁、死循环、永久阻塞时，服务器将会生成 WARN 类型的日志，关键字为 "SuspectedDeadlock"
>   - 默认不开启死锁检测

***
#### func WithDisableAsyncMessage() Option
<span id="WithDisableAsyncMessage"></span>
> 通过禁用异步消息的方式创建服务器

***
#### func WithAsyncPoolSize(size int) Option
<span id="WithAsyncPoolSize"></span>
> 通过指定异步消息池大小的方式创建服务器
>   - 当通过 WithDisableAsyncMessage 禁用异步消息时，此选项无效
>   - 默认值为 DefaultAsyncPoolSize

***
#### func WithWebsocketReadDeadline(t time.Duration) Option
<span id="WithWebsocketReadDeadline"></span>
> 设置 Websocket 读取超时时间
>   - 默认： DefaultWebsocketReadDeadline
>   - 当 t <= 0 时，表示不设置超时时间

***
#### func WithTicker(poolSize int, size int, connSize int, autonomy bool) Option
<span id="WithTicker"></span>
> 通过定时器创建服务器，为服务器添加定时器功能
>   - poolSize：指定服务器定时器池大小，当池子内的定时器数量超出该值后，多余的定时器在释放时将被回收，该值小于等于 0 时将使用 timer.DefaultTickerPoolSize
>   - size：服务器定时器时间轮大小
>   - connSize：服务器连接定时器时间轮大小，当该值小于等于 0 的时候，在新连接建立时将不再为其创建定时器
>   - autonomy：定时器是否独立运行（独立运行的情况下不会作为服务器消息运行，会导致并发问题）

***
#### func WithTLS(certFile string, keyFile string) Option
<span id="WithTLS"></span>
> 通过安全传输层协议TLS创建服务器
>   - 支持：Http、Websocket

***
#### func WithGRPCServerOptions(options ...grpc.ServerOption) Option
<span id="WithGRPCServerOptions"></span>
> 通过GRPC的可选项创建GRPC服务器

***
#### func WithWebsocketMessageType(messageTypes ...int) Option
<span id="WithWebsocketMessageType"></span>
> 设置仅支持特定类型的Websocket消息

***
#### func WithPProf(pattern ...string) Option
<span id="WithPProf"></span>
> 通过性能分析工具PProf创建服务器

***
#### func New(network Network, options ...Option) *Server
<span id="New"></span>
> 根据特定网络类型创建一个服务器

**示例代码：**

```go

func ExampleNew() {
	srv := server.New(server.NetworkWebsocket, server.WithLimitLife(time.Millisecond))
	srv.RegConnectionReceivePacketEvent(func(srv *server.Server, conn *server.Conn, packet []byte) {
		conn.Write(packet)
	})
	if err := srv.Run(":9999"); err != nil {
		panic(err)
	}
}

```

<details>
<summary>查看 / 收起单元测试</summary>


```go

func TestNew(t *testing.T) {
	srv := server.New(server.NetworkWebsocket, server.WithPProf())
	srv.RegStartBeforeEvent(func(srv *server.Server) {
		fmt.Println("启动前")
	})
	srv.RegStartFinishEvent(func(srv *server.Server) {
		fmt.Println("启动完成")
	})
	srv.RegConnectionClosedEvent(func(srv *server.Server, conn *server.Conn, err any) {
		fmt.Println("关闭", conn.GetID(), err, "IncrCount", srv.GetOnlineCount())
	})
	srv.RegConnectionReceivePacketEvent(func(srv *server.Server, conn *server.Conn, packet []byte) {
		conn.Write(packet)
	})
	if err := srv.Run(":9999"); err != nil {
		panic(err)
	}
}

```


</details>


***
#### func BindService(srv *Server, services ...Service)
<span id="BindService"></span>
> 绑定服务到特定 Server，被绑定的服务将会在 Server 初始化时执行 Service.OnInit 方法

**示例代码：**

```go

func ExampleBindService() {
	srv := server.New(server.NetworkNone, server.WithLimitLife(time.Second))
	server.BindService(srv, new(TestService))
	if err := srv.RunNone(); err != nil {
		panic(err)
	}
}

```

<details>
<summary>查看 / 收起单元测试</summary>


```go

func TestBindService(t *testing.T) {
	srv := server.New(server.NetworkNone, server.WithLimitLife(time.Second))
	server.BindService(srv, new(TestService))
	if err := srv.RunNone(); err != nil {
		t.Fatal(err)
	}
}

```


</details>


***
<span id="struct_Bot"></span>
### Bot `STRUCT`

```go
type Bot struct {
	conn   *Conn
	joined atomic.Bool
}
```
#### func (*Bot) JoinServer()
> 加入服务器
***
#### func (*Bot) LeaveServer()
> 离开服务器
***
#### func (*Bot) SetNetworkDelay(delay time.Duration, fluctuation time.Duration)
> 设置网络延迟和波动范围
>   - delay 延迟
>   - fluctuation 波动范围
***
#### func (*Bot) SetWriter(writer io.Writer)
> 设置写入器
***
#### func (*Bot) SendPacket(packet []byte)
> 发送数据包到服务器
***
#### func (*Bot) SendWSPacket(wst int, packet []byte)
> 发送 WebSocket 数据包到服务器
***
<span id="struct_BotOption"></span>
### BotOption `STRUCT`

```go
type BotOption func(bot *Bot)
```
<span id="struct_Conn"></span>
### Conn `STRUCT`
服务器连接单次消息的包装
```go
type Conn struct {
	*connection
	wst int
	ctx context.Context
}
```
#### func (*Conn) Ticker()  *timer.Ticker
> 获取定时器
***
#### func (*Conn) GetServer()  *Server
> 获取服务器
***
#### func (*Conn) GetOpenTime()  time.Time
> 获取连接打开时间
***
#### func (*Conn) GetOnlineTime()  time.Duration
> 获取连接在线时长
***
#### func (*Conn) GetWebsocketRequest()  *http.Request
> 获取websocket请求
***
#### func (*Conn) IsBot()  bool
> 是否是机器人连接
***
#### func (*Conn) RemoteAddr()  net.Addr
> 获取远程地址
***
#### func (*Conn) GetID()  string
> 获取连接ID
>   - 为远程地址的字符串形式
***
#### func (*Conn) GetIP()  string
> 获取连接IP
***
#### func (*Conn) IsClosed()  bool
> 是否已经关闭
***
#### func (*Conn) SetData(key any, value any)  *Conn
> 设置连接数据，该数据将在连接关闭前始终存在
***
#### func (*Conn) GetData(key any)  any
> 获取连接数据
***
#### func (*Conn) ViewData()  map[any]any
> 查看只读的连接数据
***
#### func (*Conn) SetMessageData(key any, value any)  *Conn
> 设置消息数据，该数据将在消息处理完成后释放
***
#### func (*Conn) GetMessageData(key any)  any
> 获取消息数据
***
#### func (*Conn) ReleaseData()  *Conn
> 释放数据
***
#### func (*Conn) IsWebsocket()  bool
> 是否是websocket连接
***
#### func (*Conn) GetWST()  int
> 获取本次 websocket 消息类型
>   - 默认将与发送类型相同
***
#### func (*Conn) SetWST(wst int)  *Conn
> 设置本次 websocket 消息类型
***
#### func (*Conn) PushAsyncMessage(caller func ()  error, callback func (err error), mark ...log.Field)
> 推送异步消息，该消息将通过 Server.PushShuntAsyncMessage 函数推送
>   - mark 为可选的日志标记，当发生异常时，将会在日志中进行体现
***
#### func (*Conn) PushUniqueAsyncMessage(name string, caller func ()  error, callback func (err error), mark ...log.Field)
> 推送唯一异步消息，该消息将通过 Server.PushUniqueShuntAsyncMessage 函数推送
>   - mark 为可选的日志标记，当发生异常时，将会在日志中进行体现
>   - 不同的是当上一个相同的 unique 消息未执行完成时，将会忽略该消息
***
#### func (*Conn) Write(packet []byte, callback ...func (err error))
> 向连接中写入数据
***
#### func (*Conn) Close(err ...error)
> 关闭连接
***
<span id="struct_ConsoleParams"></span>
### ConsoleParams `STRUCT`
控制台参数
```go
type ConsoleParams map[string][]string
```
#### func (ConsoleParams) Get(key string)  string
> 获取参数值
***
#### func (ConsoleParams) GetValues(key string)  []string
> 获取参数值
***
#### func (ConsoleParams) GetValueNum(key string)  int
> 获取参数值数量
***
#### func (ConsoleParams) Has(key string)  bool
> 是否存在参数
***
#### func (ConsoleParams) Add(key string, value string)
> 添加参数
***
#### func (ConsoleParams) Del(key string)
> 删除参数
***
#### func (ConsoleParams) Clear()
> 清空参数
***
<span id="struct_MessageReadyEventHandler"></span>
### MessageReadyEventHandler `STRUCT`

```go
type MessageReadyEventHandler func(srv *Server)
```
<span id="struct_Http"></span>
### Http `STRUCT`
基于 gin.Engine 包装的 http 服务器
```go
type Http[Context any] struct {
	srv *Server
	*gin.Engine
	*HttpRouter[Context]
}
```
#### func (*Http) Gin()  *gin.Engine
***
<span id="struct_HttpContext"></span>
### HttpContext `STRUCT`
基于 gin.Context 的 http 请求上下文
```go
type HttpContext struct {
	*gin.Context
}
```
#### func (*HttpContext) Gin()  *gin.Context
> 获取 gin.Context
***
#### func (*HttpContext) ReadTo(dest any)  error
> 读取请求数据到指定结构体，如果失败则返回错误
***
<span id="struct_HandlerFunc"></span>
### HandlerFunc `STRUCT`

```go
type HandlerFunc[Context any] func(ctx Context)
```
<span id="struct_ContextPacker"></span>
### ContextPacker `STRUCT`

```go
type ContextPacker[Context any] func(ctx *gin.Context) Context
```
<span id="struct_HttpRouter"></span>
### HttpRouter `STRUCT`

```go
type HttpRouter[Context any] struct {
	srv    *Server
	group  gin.IRouter
	packer ContextPacker[Context]
}
```
#### func (*HttpRouter) Handle(httpMethod string, relativePath string, handlers ...HandlerFunc[Context])  *HttpRouter[Context]
> 使用给定的路径和方法注册新的请求句柄和中间件
>   - 最后一个处理程序应该是真正的处理程序，其他处理程序应该是可以而且应该在不同路由之间共享的中间件。
***
#### func (*HttpRouter) POST(relativePath string, handlers ...HandlerFunc[Context])  *HttpRouter[Context]
> 是 Handle("POST", path, handlers) 的快捷方式
***
#### func (*HttpRouter) GET(relativePath string, handlers ...HandlerFunc[Context])  *HttpRouter[Context]
> 是 Handle("GET", path, handlers) 的快捷方式
***
#### func (*HttpRouter) DELETE(relativePath string, handlers ...HandlerFunc[Context])  *HttpRouter[Context]
> 是 Handle("DELETE", path, handlers) 的快捷方式
***
#### func (*HttpRouter) PATCH(relativePath string, handlers ...HandlerFunc[Context])  *HttpRouter[Context]
> 是 Handle("PATCH", path, handlers) 的快捷方式
***
#### func (*HttpRouter) PUT(relativePath string, handlers ...HandlerFunc[Context])  *HttpRouter[Context]
> 是 Handle("PUT", path, handlers) 的快捷方式
***
#### func (*HttpRouter) OPTIONS(relativePath string, handlers ...HandlerFunc[Context])  *HttpRouter[Context]
> 是 Handle("OPTIONS", path, handlers) 的快捷方式
***
#### func (*HttpRouter) HEAD(relativePath string, handlers ...HandlerFunc[Context])  *HttpRouter[Context]
> 是 Handle("HEAD", path, handlers) 的快捷方式
***
#### func (*HttpRouter) CONNECT(relativePath string, handlers ...HandlerFunc[Context])  *HttpRouter[Context]
> 是 Handle("CONNECT", path, handlers) 的快捷方式
***
#### func (*HttpRouter) TRACE(relativePath string, handlers ...HandlerFunc[Context])  *HttpRouter[Context]
> 是 Handle("TRACE", path, handlers) 的快捷方式
***
#### func (*HttpRouter) Any(relativePath string, handlers ...HandlerFunc[Context])  *HttpRouter[Context]
> 注册一个匹配所有 HTTP 方法的路由
>   - GET, POST, PUT, PATCH, HEAD, OPTIONS, DELETE, CONNECT, TRACE.
***
#### func (*HttpRouter) Match(methods []string, relativePath string, handlers ...HandlerFunc[Context])  *HttpRouter[Context]
> 注册一个匹配指定 HTTP 方法的路由
>   - GET, POST, PUT, PATCH, HEAD, OPTIONS, DELETE, CONNECT, TRACE.
***
#### func (*HttpRouter) StaticFile(relativePath string, filepath string)  *HttpRouter[Context]
> 注册单个路由以便为本地文件系统的单个文件提供服务。
>   - 例如: StaticFile("favicon.ico", "./resources/favicon.ico")
***
#### func (*HttpRouter) StaticFileFS(relativePath string, filepath string, fs http.FileSystem)  *HttpRouter[Context]
> 与 `StaticFile` 类似，但可以使用自定义的 `http.FileSystem` 代替。
>   - 例如: StaticFileFS("favicon.ico", "./resources/favicon.ico", Dir{".", false})
>   - 由于依赖于 gin.Engine 默认情况下使用：gin.Dir
***
#### func (*HttpRouter) Static(relativePath string, root string)  *HttpRouter[Context]
> 提供来自给定文件系统根目录的文件。
>   - 例如: Static("/static", "/var/www")
***
#### func (*HttpRouter) StaticFS(relativePath string, fs http.FileSystem)  *HttpRouter[Context]
> 与 `Static` 类似，但可以使用自定义的 `http.FileSystem` 代替。
>   - 例如: StaticFS("/static", Dir{"/var/www", false})
>   - 由于依赖于 gin.Engine 默认情况下使用：gin.Dir
***
#### func (*HttpRouter) Group(relativePath string, handlers ...HandlerFunc[Context])  *HttpRouter[Context]
> 创建一个新的路由组。您应该添加所有具有共同中间件的路由。
>   - 例如: v1 := slf.Group("/v1")
***
#### func (*HttpRouter) Use(middleware ...HandlerFunc[Context])  *HttpRouter[Context]
> 将中间件附加到路由组。
***
<span id="struct_HttpWrapperHandleFunc"></span>
### HttpWrapperHandleFunc `STRUCT`

```go
type HttpWrapperHandleFunc[CTX any] func(ctx CTX)
```
<span id="struct_HttpWrapper"></span>
### HttpWrapper `STRUCT`
http 包装器
```go
type HttpWrapper[CTX any] struct {
	server     *gin.Engine
	packHandle func(ctx *gin.Context) CTX
}
```
#### func (*HttpWrapper) Handle(httpMethod string, relativePath string, handlers ...HttpWrapperHandleFunc[CTX])  *HttpWrapper[CTX]
> 处理请求
***
#### func (*HttpWrapper) Use(middleware ...HttpWrapperHandleFunc[CTX])  *HttpWrapper[CTX]
> 使用中间件
***
#### func (*HttpWrapper) GET(relativePath string, handlers ...HttpWrapperHandleFunc[CTX])  *HttpWrapper[CTX]
> 注册 GET 请求
***
#### func (*HttpWrapper) POST(relativePath string, handlers ...HttpWrapperHandleFunc[CTX])  *HttpWrapper[CTX]
> 注册 POST 请求
***
#### func (*HttpWrapper) DELETE(relativePath string, handlers ...HttpWrapperHandleFunc[CTX])  *HttpWrapper[CTX]
> 注册 DELETE 请求
***
#### func (*HttpWrapper) PATCH(relativePath string, handlers ...HttpWrapperHandleFunc[CTX])  *HttpWrapper[CTX]
> 注册 PATCH 请求
***
#### func (*HttpWrapper) PUT(relativePath string, handlers ...HttpWrapperHandleFunc[CTX])  *HttpWrapper[CTX]
> 注册 PUT 请求
***
#### func (*HttpWrapper) OPTIONS(relativePath string, handlers ...HttpWrapperHandleFunc[CTX])  *HttpWrapper[CTX]
> 注册 OPTIONS 请求
***
#### func (*HttpWrapper) HEAD(relativePath string, handlers ...HttpWrapperHandleFunc[CTX])  *HttpWrapper[CTX]
> 注册 HEAD 请求
***
#### func (*HttpWrapper) Trace(relativePath string, handlers ...HttpWrapperHandleFunc[CTX])  *HttpWrapper[CTX]
> 注册 Trace 请求
***
#### func (*HttpWrapper) Connect(relativePath string, handlers ...HttpWrapperHandleFunc[CTX])  *HttpWrapper[CTX]
> 注册 Connect 请求
***
#### func (*HttpWrapper) Any(relativePath string, handlers ...HttpWrapperHandleFunc[CTX])  *HttpWrapper[CTX]
> 注册 Any 请求
***
#### func (*HttpWrapper) Match(methods []string, relativePath string, handlers ...HttpWrapperHandleFunc[CTX])  *HttpWrapper[CTX]
> 注册与您声明的指定方法相匹配的路由。
***
#### func (*HttpWrapper) StaticFile(relativePath string, filepath string)  *HttpWrapper[CTX]
> 注册 StaticFile 请求
***
#### func (*HttpWrapper) Static(relativePath string, root string)  *HttpWrapper[CTX]
> 注册 Static 请求
***
#### func (*HttpWrapper) StaticFS(relativePath string, fs http.FileSystem)  *HttpWrapper[CTX]
> 注册 StaticFS 请求
***
#### func (*HttpWrapper) Group(relativePath string, handlers ...HttpWrapperHandleFunc[CTX])  *HttpWrapperGroup[CTX]
> 创建一个新的路由组。您应该添加所有具有共同中间件的路由。
***
<span id="struct_HttpWrapperGroup"></span>
### HttpWrapperGroup `STRUCT`
http 包装器
```go
type HttpWrapperGroup[CTX any] struct {
	wrapper *HttpWrapper[CTX]
	group   *gin.RouterGroup
}
```
#### func (*HttpWrapperGroup) Handle(httpMethod string, relativePath string, handlers ...HttpWrapperHandleFunc[CTX])  *HttpWrapperGroup[CTX]
> 处理请求
***
#### func (*HttpWrapperGroup) Use(middleware ...HttpWrapperHandleFunc[CTX])  *HttpWrapperGroup[CTX]
> 使用中间件
***
#### func (*HttpWrapperGroup) GET(relativePath string, handlers ...HttpWrapperHandleFunc[CTX])  *HttpWrapperGroup[CTX]
> 注册 GET 请求
***
#### func (*HttpWrapperGroup) POST(relativePath string, handlers ...HttpWrapperHandleFunc[CTX])  *HttpWrapperGroup[CTX]
> 注册 POST 请求
***
#### func (*HttpWrapperGroup) DELETE(relativePath string, handlers ...HttpWrapperHandleFunc[CTX])  *HttpWrapperGroup[CTX]
> 注册 DELETE 请求
***
#### func (*HttpWrapperGroup) PATCH(relativePath string, handlers ...HttpWrapperHandleFunc[CTX])  *HttpWrapperGroup[CTX]
> 注册 PATCH 请求
***
#### func (*HttpWrapperGroup) PUT(relativePath string, handlers ...HttpWrapperHandleFunc[CTX])  *HttpWrapperGroup[CTX]
> 注册 PUT 请求
***
#### func (*HttpWrapperGroup) OPTIONS(relativePath string, handlers ...HttpWrapperHandleFunc[CTX])  *HttpWrapperGroup[CTX]
> 注册 OPTIONS 请求
***
#### func (*HttpWrapperGroup) Group(relativePath string, handlers ...HttpWrapperHandleFunc[CTX])  *HttpWrapperGroup[CTX]
> 创建分组
***
<span id="struct_MessageType"></span>
### MessageType `STRUCT`

```go
type MessageType byte
```
#### func (MessageType) String()  string
> 返回消息类型的字符串表示
***
<span id="struct_Message"></span>
### Message `STRUCT`
服务器消息
```go
type Message struct {
	dis              *dispatcher.Dispatcher[string, *Message]
	conn             *Conn
	err              error
	ordinaryHandler  func()
	exceptionHandler func() error
	errHandler       func(err error)
	marks            []log.Field
	packet           []byte
	producer         string
	name             string
	t                MessageType
}
```
#### func (*Message) GetProducer()  string
***
#### func (*Message) MessageType()  MessageType
> 返回消息类型
***
#### func (*Message) String()  string
> 返回消息的字符串表示
***
<span id="struct_MultipleServer"></span>
### MultipleServer `STRUCT`

```go
type MultipleServer struct {
	servers          []*Server
	addresses        []string
	exitEventHandles []func()
}
```
#### func (*MultipleServer) Run()
***
#### func (*MultipleServer) RegExitEvent(handle func ())
> 注册退出事件
***
#### func (*MultipleServer) OnExitEvent()
***
<span id="struct_Network"></span>
### Network `STRUCT`

```go
type Network string
```
#### func (Network) IsSocket()  bool
> 返回当前服务器的网络模式是否为 Socket 模式
***
<span id="struct_Option"></span>
### Option `STRUCT`

```go
type Option func(srv *Server)
```
<span id="struct_Server"></span>
### Server `STRUCT`
网络服务器
```go
type Server struct {
	*event
	*runtime
	*option
	*connMgr
	dispatcherMgr            *dispatcher.Manager[string, *Message]
	ginServer                *gin.Engine
	httpServer               *http.Server
	grpcServer               *grpc.Server
	gServer                  *gNet
	multiple                 *MultipleServer
	ants                     *ants.Pool
	messagePool              *hub.ObjectPool[*Message]
	ctx                      context.Context
	cancel                   context.CancelFunc
	systemSignal             chan os.Signal
	closeChannel             chan struct{}
	multipleRuntimeErrorChan chan error
	messageCounter           atomic.Int64
	addr                     string
	network                  Network
	closed                   uint32
	services                 []func()
}
```
#### func (*Server) Run(addr string) (err error)
> 使用特定地址运行服务器
>   - server.NetworkTcp (addr:":8888")
>   - server.NetworkTcp4 (addr:":8888")
>   - server.NetworkTcp6 (addr:":8888")
>   - server.NetworkUdp (addr:":8888")
>   - server.NetworkUdp4 (addr:":8888")
>   - server.NetworkUdp6 (addr:":8888")
>   - server.NetworkUnix (addr:"socketPath")
>   - server.NetworkHttp (addr:":8888")
>   - server.NetworkWebsocket (addr:":8888/ws")
>   - server.NetworkKcp (addr:":8888")
>   - server.NetworkNone (addr:"")
**示例代码：**

```go

func ExampleServer_Run() {
	srv := server.New(server.NetworkWebsocket, server.WithLimitLife(time.Millisecond))
	srv.RegConnectionReceivePacketEvent(func(srv *server.Server, conn *server.Conn, packet []byte) {
		conn.Write(packet)
	})
	if err := srv.Run(":9999"); err != nil {
		panic(err)
	}
}

```

***
#### func (*Server) IsSocket()  bool
> 是否是 Socket 模式
***
#### func (*Server) RunNone()  error
> 是 Run("") 的简写，仅适用于运行 NetworkNone 服务器
***
#### func (*Server) Context()  context.Context
> 获取服务器上下文
***
#### func (*Server) TimeoutContext(timeout time.Duration) ( context.Context,  context.CancelFunc)
> 获取服务器超时上下文，context.WithTimeout 的简写
***
#### func (*Server) Ticker()  *timer.Ticker
> 获取服务器定时器
***
#### func (*Server) Shutdown()
> 主动停止运行服务器
***
#### func (*Server) GRPCServer()  *grpc.Server
> 当网络类型为 NetworkGRPC 时将被允许获取 grpc 服务器，否则将会发生 panic
***
#### func (*Server) HttpRouter()  gin.IRouter
> 当网络类型为 NetworkHttp 时将被允许获取路由器进行路由注册，否则将会发生 panic
>   - 通过该函数注册的路由将无法在服务器关闭时正常等待请求结束
> 
> Deprecated: 从 Minotaur 0.0.29 开始，由于设计原因已弃用，该函数将直接返回 *gin.Server 对象，导致无法正常的对请求结束时进行处理
***
#### func (*Server) HttpServer()  *Http[*HttpContext]
> 替代 HttpRouter 的函数，返回一个 *Http[*HttpContext] 对象
>   - 通过该函数注册的路由将在服务器关闭时正常等待请求结束
>   - 如果需要自行包装 Context 对象，可以使用 NewHttpHandleWrapper 方法
***
#### func (*Server) GetMessageCount()  int64
> 获取当前服务器中消息的数量
***
#### func (*Server) UseShunt(conn *Conn, name string)
> 切换连接所使用的消息分流渠道，当分流渠道 name 不存在时将会创建一个新的分流渠道，否则将会加入已存在的分流渠道
>   - 默认情况下，所有连接都使用系统通道进行消息分发，当指定消息分流渠道且为分流消息类型时，将会使用指定的消息分流渠道进行消息分发
>   - 分流渠道会在连接断开时标记为驱逐状态，当分流渠道中的所有消息处理完毕且没有新连接使用时，将会被清除
***
#### func (*Server) HasShunt(name string)  bool
> 检查特定消息分流渠道是否存在
***
#### func (*Server) GetConnCurrShunt(conn *Conn)  string
> 获取连接当前所使用的消息分流渠道
***
#### func (*Server) GetShuntNum()  int
> 获取消息分流渠道数量
***
#### func (*Server) PushSystemMessage(handler func (), mark ...log.Field)
> 向服务器中推送 MessageTypeSystem 消息
>   - 系统消息仅包含一个可执行函数，将在系统分发器中执行
>   - mark 为可选的日志标记，当发生异常时，将会在日志中进行体现
***
#### func (*Server) PushAsyncMessage(caller func ()  error, callback func (err error), mark ...log.Field)
> 向服务器中推送 MessageTypeAsync 消息
>   - 异步消息将在服务器的异步消息队列中进行处理，处理完成 caller 的阻塞操作后，将会通过系统消息执行 callback 函数
>   - callback 函数将在异步消息处理完成后进行调用，无论过程是否产生 err，都将被执行，允许为 nil
>   - 需要注意的是，为了避免并发问题，caller 函数请仅处理阻塞操作，其他操作应该在 callback 函数中进行
>   - mark 为可选的日志标记，当发生异常时，将会在日志中进行体现
***
#### func (*Server) PushShuntAsyncMessage(conn *Conn, caller func ()  error, callback func (err error), mark ...log.Field)
> 向特定分发器中推送 MessageTypeAsync 消息，消息执行与 MessageTypeAsync 一致
>   - 需要注意的是，当未指定 UseShunt 时，将会通过 PushAsyncMessage 进行转发
>   - mark 为可选的日志标记，当发生异常时，将会在日志中进行体现
***
#### func (*Server) PushPacketMessage(conn *Conn, wst int, packet []byte, mark ...log.Field)
> 向服务器中推送 MessageTypePacket 消息
>   - 当存在 UseShunt 的选项时，将会根据选项中的 shuntMatcher 进行分发，否则将在系统分发器中处理消息
***
#### func (*Server) PushTickerMessage(name string, caller func (), mark ...log.Field)
> 向服务器中推送 MessageTypeTicker 消息
>   - 通过该函数推送定时消息，当消息触发时将在系统分发器中处理消息
>   - 可通过 timer.Ticker 或第三方定时器将执行函数(caller)推送到该消息中进行处理，可有效的避免线程安全问题
>   - 参数 name 仅用作标识该定时器名称
> 
> 定时消息执行不会有特殊的处理，仅标记为定时任务，也就是允许将各类函数通过该消息发送处理，但是并不建议
>   - mark 为可选的日志标记，当发生异常时，将会在日志中进行体现
***
#### func (*Server) PushShuntTickerMessage(conn *Conn, name string, caller func (), mark ...log.Field)
> 向特定分发器中推送 MessageTypeTicker 消息，消息执行与 MessageTypeTicker 一致
>   - 需要注意的是，当未指定 UseShunt 时，将会通过 PushTickerMessage 进行转发
>   - mark 为可选的日志标记，当发生异常时，将会在日志中进行体现
***
#### func (*Server) PushUniqueAsyncMessage(unique string, caller func ()  error, callback func (err error), mark ...log.Field)
> 向服务器中推送 MessageTypeAsync 消息，消息执行与 MessageTypeAsync 一致
>   - 不同的是当上一个相同的 unique 消息未执行完成时，将会忽略该消息
***
#### func (*Server) PushUniqueShuntAsyncMessage(conn *Conn, unique string, caller func ()  error, callback func (err error), mark ...log.Field)
> 向特定分发器中推送 MessageTypeAsync 消息，消息执行与 MessageTypeAsync 一致
>   - 需要注意的是，当未指定 UseShunt 时，将会通过系统分流渠道进行转发
>   - 不同的是当上一个相同的 unique 消息未执行完成时，将会忽略该消息
***
#### func (*Server) PushShuntMessage(conn *Conn, caller func (), mark ...log.Field)
> 向特定分发器中推送 MessageTypeShunt 消息，消息执行与 MessageTypeSystem 一致，不同的是将会在特定分发器中执行
***
#### func (*Server) GetDurationMessageCount()  int64
> 获取当前 WithMessageStatistics 设置的 duration 期间的消息量
***
#### func (*Server) GetDurationMessageCountByOffset(offset int)  int64
> 获取特定偏移次数的 WithMessageStatistics 设置的 duration 期间的消息量
>   - 该值小于 0 时，将与 GetDurationMessageCount 无异，否则将返回 +n 个期间的消息量，例如 duration 为 1 分钟，limit 为 10，那么 offset 为 1 的情况下，获取的则是上一分钟消息量
***
#### func (*Server) GetAllDurationMessageCount()  []int64
> 获取所有 WithMessageStatistics 设置的 duration 期间的消息量
***
#### func (*Server) HasMessageStatistics()  bool
> 是否了开启消息统计
***
<span id="struct_Service"></span>
### Service `INTERFACE`
兼容传统 service 设计模式的接口
```go
type Service interface {
	OnInit(srv *Server)
}
```
