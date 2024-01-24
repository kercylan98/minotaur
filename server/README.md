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
|[LoadData](#LoadData)|加载绑定的服务器数据
|[BindData](#BindData)|绑定数据到特定服务器
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
|`STRUCT`|[Network](#struct_Network)|服务器运行的网络模式
|`STRUCT`|[Option](#struct_Option)|暂无描述...
|`STRUCT`|[Server](#struct_Server)|网络服务器
|`INTERFACE`|[Service](#struct_Service)|兼容传统 service 设计模式的接口，通过该接口可以实现更简洁、更具有可读性的服务绑定

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

**示例代码：**

服务器在启动时将阻塞 1s，模拟了慢消息的过程，这时候如果通过 RegMessageLowExecEvent 函数注册过慢消息事件，将会收到该事件的消息
  - 该示例中，将在收到慢消息时关闭服务器


```go

func ExampleWithLowMessageDuration() {
	srv := server.New(server.NetworkNone, server.WithLowMessageDuration(time.Second))
	srv.RegStartFinishEvent(func(srv *server.Server) {
		time.Sleep(time.Second)
	})
	srv.RegMessageLowExecEvent(func(srv *server.Server, message *server.Message, cost time.Duration) {
		srv.Shutdown()
		fmt.Println(times.GetSecond(cost))
	})
	if err := srv.RunNone(); err != nil {
		panic(err)
	}
}

```

<details>
<summary>查看 / 收起单元测试</summary>


```go

func TestWithLowMessageDuration(t *testing.T) {
	var cases = []struct {
		name     string
		duration time.Duration
	}{{name: "TestWithLowMessageDuration", duration: server.DefaultLowMessageDuration}, {name: "TestWithLowMessageDuration_Zero", duration: 0}, {name: "TestWithLowMessageDuration_Negative", duration: -server.DefaultAsyncLowMessageDuration}}
	for _, c := range cases {
		c := c
		t.Run(c.name, func(t *testing.T) {
			networks := server.GetNetworks()
			for i := 0; i < len(networks); i++ {
				low := false
				network := networks[i]
				srv := server.New(network, server.WithLowMessageDuration(c.duration))
				srv.RegMessageLowExecEvent(func(srv *server.Server, message *server.Message, cost time.Duration) {
					low = true
					srv.Shutdown()
				})
				srv.RegStartFinishEvent(func(srv *server.Server) {
					if c.duration <= 0 {
						srv.Shutdown()
						return
					}
					time.Sleep(server.DefaultLowMessageDuration)
				})
				var lis string
				switch network {
				case server.NetworkNone, server.NetworkUnix:
					lis = "addr"
				default:
					lis = fmt.Sprintf(":%d", random.UsablePort())
				}
				if err := srv.Run(lis); err != nil {
					t.Fatalf("%s run error: %s", network, err)
				}
				if !low && c.duration > 0 {
					t.Fatalf("%s low message not exec", network)
				}
			}
		})
	}
}

```


</details>


***
#### func WithAsyncLowMessageDuration(duration time.Duration) Option
<span id="WithAsyncLowMessageDuration"></span>
> 通过指定异步消息的慢消息时长的方式创建服务器，当消息处理时间超过指定时长时，将会输出 WARN 类型的日志
>   - 默认值为 DefaultAsyncLowMessageDuration
>   - 当 duration <= 0 时，表示关闭慢消息检测

**示例代码：**

服务器在启动时将发布一条阻塞 1s 的异步消息，模拟了慢消息的过程，这时候如果通过 RegMessageLowExecEvent 函数注册过慢消息事件，将会收到该事件的消息
  - 该示例中，将在收到慢消息时关闭服务器


```go

func ExampleWithAsyncLowMessageDuration() {
	srv := server.New(server.NetworkNone, server.WithAsyncLowMessageDuration(time.Second))
	srv.RegStartFinishEvent(func(srv *server.Server) {
		srv.PushAsyncMessage(func() error {
			time.Sleep(time.Second)
			return nil
		}, nil)
	})
	srv.RegMessageLowExecEvent(func(srv *server.Server, message *server.Message, cost time.Duration) {
		srv.Shutdown()
		fmt.Println(times.GetSecond(cost))
	})
	if err := srv.RunNone(); err != nil {
		panic(err)
	}
}

```

<details>
<summary>查看 / 收起单元测试</summary>


```go

func TestWithAsyncLowMessageDuration(t *testing.T) {
	var cases = []struct {
		name     string
		duration time.Duration
	}{{name: "TestWithAsyncLowMessageDuration", duration: time.Millisecond * 100}, {name: "TestWithAsyncLowMessageDuration_Zero", duration: 0}, {name: "TestWithAsyncLowMessageDuration_Negative", duration: -server.DefaultAsyncLowMessageDuration}}
	for _, c := range cases {
		t.Run(c.name, func(t *testing.T) {
			networks := server.GetNetworks()
			for i := 0; i < len(networks); i++ {
				low := false
				network := networks[i]
				srv := server.New(network, server.WithAsyncLowMessageDuration(c.duration))
				srv.RegMessageLowExecEvent(func(srv *server.Server, message *server.Message, cost time.Duration) {
					low = true
					srv.Shutdown()
				})
				srv.RegStartFinishEvent(func(srv *server.Server) {
					if c.duration <= 0 {
						srv.Shutdown()
						return
					}
					srv.PushAsyncMessage(func() error {
						time.Sleep(c.duration)
						return nil
					}, nil)
				})
				var lis string
				switch network {
				case server.NetworkNone, server.NetworkUnix:
					lis = fmt.Sprintf("%s%d", "addr", random.Int(0, 9999))
				default:
					lis = fmt.Sprintf(":%d", random.UsablePort())
				}
				if err := srv.Run(lis); err != nil {
					t.Fatalf("%s run error: %s", network, err)
				}
				if !low && c.duration > 0 {
					t.Fatalf("%s low message not exec", network)
				}
			}
		})
	}
}

```


</details>


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

该案例将创建一个简单的 WebSocket 服务器，如果需要更多的服务器类型可参考 [` Network `](#struct_Network) 部分
  - server.WithLimitLife(time.Millisecond) 通常不是在正常开发应该使用的，在这里只是为了让服务器在启动完成后的 1 毫秒后自动关闭

该案例的输出结果为 true


```go

func ExampleNew() {
	srv := server.New(server.NetworkWebsocket, server.WithLimitLife(time.Millisecond))
	fmt.Println(srv != nil)
}

```

<details>
<summary>查看 / 收起单元测试</summary>


该单元测试用于测试以不同的基本参数创建服务器是否存在异常


```go

func TestNew(t *testing.T) {
	var cases = []struct {
		name        string
		network     server.Network
		addr        string
		shouldPanic bool
	}{{name: "TestNew_Unknown", addr: "", network: "Unknown", shouldPanic: true}, {name: "TestNew_None", addr: "", network: server.NetworkNone, shouldPanic: false}, {name: "TestNew_None_Addr", addr: "addr", network: server.NetworkNone, shouldPanic: false}, {name: "TestNew_Tcp_AddrEmpty", addr: "", network: server.NetworkTcp, shouldPanic: true}, {name: "TestNew_Tcp_AddrIllegal", addr: "addr", network: server.NetworkTcp, shouldPanic: true}, {name: "TestNew_Tcp_Addr", addr: ":9999", network: server.NetworkTcp, shouldPanic: false}, {name: "TestNew_Tcp4_AddrEmpty", addr: "", network: server.NetworkTcp4, shouldPanic: true}, {name: "TestNew_Tcp4_AddrIllegal", addr: "addr", network: server.NetworkTcp4, shouldPanic: true}, {name: "TestNew_Tcp4_Addr", addr: ":9999", network: server.NetworkTcp4, shouldPanic: false}, {name: "TestNew_Tcp6_AddrEmpty", addr: "", network: server.NetworkTcp6, shouldPanic: true}, {name: "TestNew_Tcp6_AddrIllegal", addr: "addr", network: server.NetworkTcp6, shouldPanic: true}, {name: "TestNew_Tcp6_Addr", addr: ":9999", network: server.NetworkTcp6, shouldPanic: false}, {name: "TestNew_Udp_AddrEmpty", addr: "", network: server.NetworkUdp, shouldPanic: true}, {name: "TestNew_Udp_AddrIllegal", addr: "addr", network: server.NetworkUdp, shouldPanic: true}, {name: "TestNew_Udp_Addr", addr: ":9999", network: server.NetworkUdp, shouldPanic: false}, {name: "TestNew_Udp4_AddrEmpty", addr: "", network: server.NetworkUdp4, shouldPanic: true}, {name: "TestNew_Udp4_AddrIllegal", addr: "addr", network: server.NetworkUdp4, shouldPanic: true}, {name: "TestNew_Udp4_Addr", addr: ":9999", network: server.NetworkUdp4, shouldPanic: false}, {name: "TestNew_Udp6_AddrEmpty", addr: "", network: server.NetworkUdp6, shouldPanic: true}, {name: "TestNew_Udp6_AddrIllegal", addr: "addr", network: server.NetworkUdp6, shouldPanic: true}, {name: "TestNew_Udp6_Addr", addr: ":9999", network: server.NetworkUdp6, shouldPanic: false}, {name: "TestNew_Unix_AddrEmpty", addr: "", network: server.NetworkUnix, shouldPanic: true}, {name: "TestNew_Unix_AddrIllegal", addr: "addr", network: server.NetworkUnix, shouldPanic: true}, {name: "TestNew_Unix_Addr", addr: "addr", network: server.NetworkUnix, shouldPanic: false}, {name: "TestNew_Websocket_AddrEmpty", addr: "", network: server.NetworkWebsocket, shouldPanic: true}, {name: "TestNew_Websocket_AddrIllegal", addr: "addr", network: server.NetworkWebsocket, shouldPanic: true}, {name: "TestNew_Websocket_Addr", addr: ":9999/ws", network: server.NetworkWebsocket, shouldPanic: false}, {name: "TestNew_Http_AddrEmpty", addr: "", network: server.NetworkHttp, shouldPanic: true}, {name: "TestNew_Http_AddrIllegal", addr: "addr", network: server.NetworkHttp, shouldPanic: true}, {name: "TestNew_Http_Addr", addr: ":9999", network: server.NetworkHttp, shouldPanic: false}, {name: "TestNew_Kcp_AddrEmpty", addr: "", network: server.NetworkKcp, shouldPanic: true}, {name: "TestNew_Kcp_AddrIllegal", addr: "addr", network: server.NetworkKcp, shouldPanic: true}, {name: "TestNew_Kcp_Addr", addr: ":9999", network: server.NetworkKcp, shouldPanic: false}, {name: "TestNew_GRPC_AddrEmpty", addr: "", network: server.NetworkGRPC, shouldPanic: true}, {name: "TestNew_GRPC_AddrIllegal", addr: "addr", network: server.NetworkGRPC, shouldPanic: true}, {name: "TestNew_GRPC_Addr", addr: ":9999", network: server.NetworkGRPC, shouldPanic: false}}
	for _, c := range cases {
		t.Run(c.name, func(t *testing.T) {
			defer func() {
				if err := super.RecoverTransform(recover()); err != nil && !c.shouldPanic {
					debug.PrintStack()
					t.Fatal("not should panic, err:", err)
				}
			}()
			if err := server.New(c.network, server.WithLimitLife(time.Millisecond*10)).Run(""); err != nil {
				panic(err)
			}
		})
	}
}

```


</details>


***
#### func LoadData\[T any\](srv *Server, name string, data any) T
<span id="LoadData"></span>
> 加载绑定的服务器数据

***
#### func BindData(srv *Server, name string, data any)
<span id="BindData"></span>
> 绑定数据到特定服务器

***
#### func BindService(srv *Server, services ...Service)
<span id="BindService"></span>
> 绑定服务到特定 Server，被绑定的服务将会在 Server 初始化时执行 Service.OnInit 方法

**示例代码：**

这个案例中我们将 `TestService` 绑定到了 `srv` 服务器中，当服务器启动时，将会对 `TestService` 进行初始化

其中 `TestService` 的定义如下：
```go

	type TestService struct{}

	func (ts *TestService) OnInit(srv *server.Server) {
		srv.RegStartFinishEvent(onStartFinish)

		srv.RegStopEvent(func(srv *server.Server) {
			fmt.Println("server stop")
		})
	}

	func (ts *TestService) onStartFinish(srv *server.Server) {
		fmt.Println("server start finish")
	}

```

可以看出，在服务初始化时，该服务向服务器注册了启动完成事件及停止事件。这是我们推荐的编码方式，这样编码有以下好处：
  - 具备可控制的初始化顺序，避免 init 产生的各种顺序导致的问题，如配置还未加载完成，即开始进行数据库连接等操作
  - 可以方便的将不同的服务拆分到不同的包中进行管理
  - 当不需要某个服务时，可以直接删除该服务的绑定，而不需要修改其他代码
  - ...


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
	var cases = []struct{ name string }{{name: "TestBindService"}}
	for _, c := range cases {
		t.Run(c.name, func(t *testing.T) {
			srv := server.New(server.NetworkNone, server.WithLimitLife(time.Millisecond))
			server.BindService(srv, new(TestService))
			if err := srv.RunNone(); err != nil {
				t.Fatal(err)
			}
		})
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
<span id="struct_Bot_JoinServer"></span>

#### func (*Bot) JoinServer()
> 加入服务器

***
<span id="struct_Bot_LeaveServer"></span>

#### func (*Bot) LeaveServer()
> 离开服务器

***
<span id="struct_Bot_SetNetworkDelay"></span>

#### func (*Bot) SetNetworkDelay(delay time.Duration, fluctuation time.Duration)
> 设置网络延迟和波动范围
>   - delay 延迟
>   - fluctuation 波动范围

***
<span id="struct_Bot_SetWriter"></span>

#### func (*Bot) SetWriter(writer io.Writer)
> 设置写入器

***
<span id="struct_Bot_SendPacket"></span>

#### func (*Bot) SendPacket(packet []byte)
> 发送数据包到服务器

***
<span id="struct_Bot_SendWSPacket"></span>

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
<span id="struct_Conn_Ticker"></span>

#### func (*Conn) Ticker()  *timer.Ticker
> 获取定时器

***
<span id="struct_Conn_GetServer"></span>

#### func (*Conn) GetServer()  *Server
> 获取服务器

***
<span id="struct_Conn_GetOpenTime"></span>

#### func (*Conn) GetOpenTime()  time.Time
> 获取连接打开时间

***
<span id="struct_Conn_GetOnlineTime"></span>

#### func (*Conn) GetOnlineTime()  time.Duration
> 获取连接在线时长

***
<span id="struct_Conn_GetWebsocketRequest"></span>

#### func (*Conn) GetWebsocketRequest()  *http.Request
> 获取websocket请求

***
<span id="struct_Conn_IsBot"></span>

#### func (*Conn) IsBot()  bool
> 是否是机器人连接

***
<span id="struct_Conn_RemoteAddr"></span>

#### func (*Conn) RemoteAddr()  net.Addr
> 获取远程地址

***
<span id="struct_Conn_GetID"></span>

#### func (*Conn) GetID()  string
> 获取连接ID
>   - 为远程地址的字符串形式

***
<span id="struct_Conn_GetIP"></span>

#### func (*Conn) GetIP()  string
> 获取连接IP

***
<span id="struct_Conn_IsClosed"></span>

#### func (*Conn) IsClosed()  bool
> 是否已经关闭

***
<span id="struct_Conn_SetData"></span>

#### func (*Conn) SetData(key any, value any)  *Conn
> 设置连接数据，该数据将在连接关闭前始终存在

***
<span id="struct_Conn_GetData"></span>

#### func (*Conn) GetData(key any)  any
> 获取连接数据

***
<span id="struct_Conn_ViewData"></span>

#### func (*Conn) ViewData()  map[any]any
> 查看只读的连接数据

***
<span id="struct_Conn_SetMessageData"></span>

#### func (*Conn) SetMessageData(key any, value any)  *Conn
> 设置消息数据，该数据将在消息处理完成后释放

***
<span id="struct_Conn_GetMessageData"></span>

#### func (*Conn) GetMessageData(key any)  any
> 获取消息数据

***
<span id="struct_Conn_ReleaseData"></span>

#### func (*Conn) ReleaseData()  *Conn
> 释放数据

***
<span id="struct_Conn_IsWebsocket"></span>

#### func (*Conn) IsWebsocket()  bool
> 是否是websocket连接

***
<span id="struct_Conn_GetWST"></span>

#### func (*Conn) GetWST()  int
> 获取本次 websocket 消息类型
>   - 默认将与发送类型相同

***
<span id="struct_Conn_SetWST"></span>

#### func (*Conn) SetWST(wst int)  *Conn
> 设置本次 websocket 消息类型

***
<span id="struct_Conn_PushAsyncMessage"></span>

#### func (*Conn) PushAsyncMessage(caller func ()  error, callback func (err error), mark ...log.Field)
> 推送异步消息，该消息将通过 Server.PushShuntAsyncMessage 函数推送
>   - mark 为可选的日志标记，当发生异常时，将会在日志中进行体现

***
<span id="struct_Conn_PushUniqueAsyncMessage"></span>

#### func (*Conn) PushUniqueAsyncMessage(name string, caller func ()  error, callback func (err error), mark ...log.Field)
> 推送唯一异步消息，该消息将通过 Server.PushUniqueShuntAsyncMessage 函数推送
>   - mark 为可选的日志标记，当发生异常时，将会在日志中进行体现
>   - 不同的是当上一个相同的 unique 消息未执行完成时，将会忽略该消息

***
<span id="struct_Conn_Write"></span>

#### func (*Conn) Write(packet []byte, callback ...func (err error))
> 向连接中写入数据

***
<span id="struct_Conn_Close"></span>

#### func (*Conn) Close(err ...error)
> 关闭连接

***
<span id="struct_ConsoleParams"></span>
### ConsoleParams `STRUCT`
控制台参数
```go
type ConsoleParams map[string][]string
```
<span id="struct_ConsoleParams_Get"></span>

#### func (ConsoleParams) Get(key string)  string
> 获取参数值

***
<span id="struct_ConsoleParams_GetValues"></span>

#### func (ConsoleParams) GetValues(key string)  []string
> 获取参数值

***
<span id="struct_ConsoleParams_GetValueNum"></span>

#### func (ConsoleParams) GetValueNum(key string)  int
> 获取参数值数量

***
<span id="struct_ConsoleParams_Has"></span>

#### func (ConsoleParams) Has(key string)  bool
> 是否存在参数

***
<span id="struct_ConsoleParams_Add"></span>

#### func (ConsoleParams) Add(key string, value string)
> 添加参数

***
<span id="struct_ConsoleParams_Del"></span>

#### func (ConsoleParams) Del(key string)
> 删除参数

***
<span id="struct_ConsoleParams_Clear"></span>

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
<span id="struct_Http_Gin"></span>

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
<span id="struct_HttpContext_Gin"></span>

#### func (*HttpContext) Gin()  *gin.Context
> 获取 gin.Context

***
<span id="struct_HttpContext_ReadTo"></span>

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
<span id="struct_HttpRouter_Handle"></span>

#### func (*HttpRouter) Handle(httpMethod string, relativePath string, handlers ...HandlerFunc[Context])  *HttpRouter[Context]
> 使用给定的路径和方法注册新的请求句柄和中间件
>   - 最后一个处理程序应该是真正的处理程序，其他处理程序应该是可以而且应该在不同路由之间共享的中间件。

***
<span id="struct_HttpRouter_POST"></span>

#### func (*HttpRouter) POST(relativePath string, handlers ...HandlerFunc[Context])  *HttpRouter[Context]
> 是 Handle("POST", path, handlers) 的快捷方式

***
<span id="struct_HttpRouter_GET"></span>

#### func (*HttpRouter) GET(relativePath string, handlers ...HandlerFunc[Context])  *HttpRouter[Context]
> 是 Handle("GET", path, handlers) 的快捷方式

***
<span id="struct_HttpRouter_DELETE"></span>

#### func (*HttpRouter) DELETE(relativePath string, handlers ...HandlerFunc[Context])  *HttpRouter[Context]
> 是 Handle("DELETE", path, handlers) 的快捷方式

***
<span id="struct_HttpRouter_PATCH"></span>

#### func (*HttpRouter) PATCH(relativePath string, handlers ...HandlerFunc[Context])  *HttpRouter[Context]
> 是 Handle("PATCH", path, handlers) 的快捷方式

***
<span id="struct_HttpRouter_PUT"></span>

#### func (*HttpRouter) PUT(relativePath string, handlers ...HandlerFunc[Context])  *HttpRouter[Context]
> 是 Handle("PUT", path, handlers) 的快捷方式

***
<span id="struct_HttpRouter_OPTIONS"></span>

#### func (*HttpRouter) OPTIONS(relativePath string, handlers ...HandlerFunc[Context])  *HttpRouter[Context]
> 是 Handle("OPTIONS", path, handlers) 的快捷方式

***
<span id="struct_HttpRouter_HEAD"></span>

#### func (*HttpRouter) HEAD(relativePath string, handlers ...HandlerFunc[Context])  *HttpRouter[Context]
> 是 Handle("HEAD", path, handlers) 的快捷方式

***
<span id="struct_HttpRouter_CONNECT"></span>

#### func (*HttpRouter) CONNECT(relativePath string, handlers ...HandlerFunc[Context])  *HttpRouter[Context]
> 是 Handle("CONNECT", path, handlers) 的快捷方式

***
<span id="struct_HttpRouter_TRACE"></span>

#### func (*HttpRouter) TRACE(relativePath string, handlers ...HandlerFunc[Context])  *HttpRouter[Context]
> 是 Handle("TRACE", path, handlers) 的快捷方式

***
<span id="struct_HttpRouter_Any"></span>

#### func (*HttpRouter) Any(relativePath string, handlers ...HandlerFunc[Context])  *HttpRouter[Context]
> 注册一个匹配所有 HTTP 方法的路由
>   - GET, POST, PUT, PATCH, HEAD, OPTIONS, DELETE, CONNECT, TRACE.

***
<span id="struct_HttpRouter_Match"></span>

#### func (*HttpRouter) Match(methods []string, relativePath string, handlers ...HandlerFunc[Context])  *HttpRouter[Context]
> 注册一个匹配指定 HTTP 方法的路由
>   - GET, POST, PUT, PATCH, HEAD, OPTIONS, DELETE, CONNECT, TRACE.

***
<span id="struct_HttpRouter_StaticFile"></span>

#### func (*HttpRouter) StaticFile(relativePath string, filepath string)  *HttpRouter[Context]
> 注册单个路由以便为本地文件系统的单个文件提供服务。
>   - 例如: StaticFile("favicon.ico", "./resources/favicon.ico")

***
<span id="struct_HttpRouter_StaticFileFS"></span>

#### func (*HttpRouter) StaticFileFS(relativePath string, filepath string, fs http.FileSystem)  *HttpRouter[Context]
> 与 `StaticFile` 类似，但可以使用自定义的 `http.FileSystem` 代替。
>   - 例如: StaticFileFS("favicon.ico", "./resources/favicon.ico", Dir{".", false})
>   - 由于依赖于 gin.Engine 默认情况下使用：gin.Dir

***
<span id="struct_HttpRouter_Static"></span>

#### func (*HttpRouter) Static(relativePath string, root string)  *HttpRouter[Context]
> 提供来自给定文件系统根目录的文件。
>   - 例如: Static("/static", "/var/www")

***
<span id="struct_HttpRouter_StaticFS"></span>

#### func (*HttpRouter) StaticFS(relativePath string, fs http.FileSystem)  *HttpRouter[Context]
> 与 `Static` 类似，但可以使用自定义的 `http.FileSystem` 代替。
>   - 例如: StaticFS("/static", Dir{"/var/www", false})
>   - 由于依赖于 gin.Engine 默认情况下使用：gin.Dir

***
<span id="struct_HttpRouter_Group"></span>

#### func (*HttpRouter) Group(relativePath string, handlers ...HandlerFunc[Context])  *HttpRouter[Context]
> 创建一个新的路由组。您应该添加所有具有共同中间件的路由。
>   - 例如: v1 := slf.Group("/v1")

***
<span id="struct_HttpRouter_Use"></span>

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
<span id="struct_HttpWrapper_Handle"></span>

#### func (*HttpWrapper) Handle(httpMethod string, relativePath string, handlers ...HttpWrapperHandleFunc[CTX])  *HttpWrapper[CTX]
> 处理请求

***
<span id="struct_HttpWrapper_Use"></span>

#### func (*HttpWrapper) Use(middleware ...HttpWrapperHandleFunc[CTX])  *HttpWrapper[CTX]
> 使用中间件

***
<span id="struct_HttpWrapper_GET"></span>

#### func (*HttpWrapper) GET(relativePath string, handlers ...HttpWrapperHandleFunc[CTX])  *HttpWrapper[CTX]
> 注册 GET 请求

***
<span id="struct_HttpWrapper_POST"></span>

#### func (*HttpWrapper) POST(relativePath string, handlers ...HttpWrapperHandleFunc[CTX])  *HttpWrapper[CTX]
> 注册 POST 请求

***
<span id="struct_HttpWrapper_DELETE"></span>

#### func (*HttpWrapper) DELETE(relativePath string, handlers ...HttpWrapperHandleFunc[CTX])  *HttpWrapper[CTX]
> 注册 DELETE 请求

***
<span id="struct_HttpWrapper_PATCH"></span>

#### func (*HttpWrapper) PATCH(relativePath string, handlers ...HttpWrapperHandleFunc[CTX])  *HttpWrapper[CTX]
> 注册 PATCH 请求

***
<span id="struct_HttpWrapper_PUT"></span>

#### func (*HttpWrapper) PUT(relativePath string, handlers ...HttpWrapperHandleFunc[CTX])  *HttpWrapper[CTX]
> 注册 PUT 请求

***
<span id="struct_HttpWrapper_OPTIONS"></span>

#### func (*HttpWrapper) OPTIONS(relativePath string, handlers ...HttpWrapperHandleFunc[CTX])  *HttpWrapper[CTX]
> 注册 OPTIONS 请求

***
<span id="struct_HttpWrapper_HEAD"></span>

#### func (*HttpWrapper) HEAD(relativePath string, handlers ...HttpWrapperHandleFunc[CTX])  *HttpWrapper[CTX]
> 注册 HEAD 请求

***
<span id="struct_HttpWrapper_Trace"></span>

#### func (*HttpWrapper) Trace(relativePath string, handlers ...HttpWrapperHandleFunc[CTX])  *HttpWrapper[CTX]
> 注册 Trace 请求

***
<span id="struct_HttpWrapper_Connect"></span>

#### func (*HttpWrapper) Connect(relativePath string, handlers ...HttpWrapperHandleFunc[CTX])  *HttpWrapper[CTX]
> 注册 Connect 请求

***
<span id="struct_HttpWrapper_Any"></span>

#### func (*HttpWrapper) Any(relativePath string, handlers ...HttpWrapperHandleFunc[CTX])  *HttpWrapper[CTX]
> 注册 Any 请求

***
<span id="struct_HttpWrapper_Match"></span>

#### func (*HttpWrapper) Match(methods []string, relativePath string, handlers ...HttpWrapperHandleFunc[CTX])  *HttpWrapper[CTX]
> 注册与您声明的指定方法相匹配的路由。

***
<span id="struct_HttpWrapper_StaticFile"></span>

#### func (*HttpWrapper) StaticFile(relativePath string, filepath string)  *HttpWrapper[CTX]
> 注册 StaticFile 请求

***
<span id="struct_HttpWrapper_Static"></span>

#### func (*HttpWrapper) Static(relativePath string, root string)  *HttpWrapper[CTX]
> 注册 Static 请求

***
<span id="struct_HttpWrapper_StaticFS"></span>

#### func (*HttpWrapper) StaticFS(relativePath string, fs http.FileSystem)  *HttpWrapper[CTX]
> 注册 StaticFS 请求

***
<span id="struct_HttpWrapper_Group"></span>

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
<span id="struct_HttpWrapperGroup_Handle"></span>

#### func (*HttpWrapperGroup) Handle(httpMethod string, relativePath string, handlers ...HttpWrapperHandleFunc[CTX])  *HttpWrapperGroup[CTX]
> 处理请求

***
<span id="struct_HttpWrapperGroup_Use"></span>

#### func (*HttpWrapperGroup) Use(middleware ...HttpWrapperHandleFunc[CTX])  *HttpWrapperGroup[CTX]
> 使用中间件

***
<span id="struct_HttpWrapperGroup_GET"></span>

#### func (*HttpWrapperGroup) GET(relativePath string, handlers ...HttpWrapperHandleFunc[CTX])  *HttpWrapperGroup[CTX]
> 注册 GET 请求

***
<span id="struct_HttpWrapperGroup_POST"></span>

#### func (*HttpWrapperGroup) POST(relativePath string, handlers ...HttpWrapperHandleFunc[CTX])  *HttpWrapperGroup[CTX]
> 注册 POST 请求

***
<span id="struct_HttpWrapperGroup_DELETE"></span>

#### func (*HttpWrapperGroup) DELETE(relativePath string, handlers ...HttpWrapperHandleFunc[CTX])  *HttpWrapperGroup[CTX]
> 注册 DELETE 请求

***
<span id="struct_HttpWrapperGroup_PATCH"></span>

#### func (*HttpWrapperGroup) PATCH(relativePath string, handlers ...HttpWrapperHandleFunc[CTX])  *HttpWrapperGroup[CTX]
> 注册 PATCH 请求

***
<span id="struct_HttpWrapperGroup_PUT"></span>

#### func (*HttpWrapperGroup) PUT(relativePath string, handlers ...HttpWrapperHandleFunc[CTX])  *HttpWrapperGroup[CTX]
> 注册 PUT 请求

***
<span id="struct_HttpWrapperGroup_OPTIONS"></span>

#### func (*HttpWrapperGroup) OPTIONS(relativePath string, handlers ...HttpWrapperHandleFunc[CTX])  *HttpWrapperGroup[CTX]
> 注册 OPTIONS 请求

***
<span id="struct_HttpWrapperGroup_Group"></span>

#### func (*HttpWrapperGroup) Group(relativePath string, handlers ...HttpWrapperHandleFunc[CTX])  *HttpWrapperGroup[CTX]
> 创建分组

***
<span id="struct_MessageType"></span>
### MessageType `STRUCT`

```go
type MessageType byte
```
<span id="struct_MessageType_String"></span>

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
<span id="struct_Message_GetProducer"></span>

#### func (*Message) GetProducer()  string

***
<span id="struct_Message_MessageType"></span>

#### func (*Message) MessageType()  MessageType
> 返回消息类型

***
<span id="struct_Message_String"></span>

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
<span id="struct_MultipleServer_Run"></span>

#### func (*MultipleServer) Run()

***
<span id="struct_MultipleServer_RegExitEvent"></span>

#### func (*MultipleServer) RegExitEvent(handle func ())
> 注册退出事件

***
<span id="struct_MultipleServer_OnExitEvent"></span>

#### func (*MultipleServer) OnExitEvent()

***
<span id="struct_Network"></span>
### Network `STRUCT`
服务器运行的网络模式
  - 根据不同的网络模式，服务器将会产生不同的行为，该类型将在服务器创建时候指定

服务器支持的网络模式如下：
  - NetworkNone 该模式下不监听任何网络端口，仅开启消息队列，适用于纯粹的跨服服务器等情况
  - NetworkTcp 该模式下将会监听 TCP 协议的所有地址，包括 IPv4 和 IPv6
  - NetworkTcp4 该模式下将会监听 TCP 协议的 IPv4 地址
  - NetworkTcp6 该模式下将会监听 TCP 协议的 IPv6 地址
  - NetworkUdp 该模式下将会监听 UDP 协议的所有地址，包括 IPv4 和 IPv6
  - NetworkUdp4 该模式下将会监听 UDP 协议的 IPv4 地址
  - NetworkUdp6 该模式下将会监听 UDP 协议的 IPv6 地址
  - NetworkUnix 该模式下将会监听 Unix 协议的地址
  - NetworkHttp 该模式下将会监听 HTTP 协议的地址
  - NetworkWebsocket 该模式下将会监听 Websocket 协议的地址
  - NetworkKcp 该模式下将会监听 KCP 协议的地址
  - NetworkGRPC 该模式下将会监听 GRPC 协议的地址
```go
type Network string
```
<span id="struct_Network_IsSocket"></span>

#### func (Network) IsSocket()  bool
> 返回当前服务器的网络模式是否为 Socket 模式，目前为止仅有如下几种模式为 Socket 模式：
>   - NetworkTcp
>   - NetworkTcp4
>   - NetworkTcp6
>   - NetworkUdp
>   - NetworkUdp4
>   - NetworkUdp6
>   - NetworkUnix
>   - NetworkKcp
>   - NetworkWebsocket

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
	data                     map[string]any
	messageCounter           atomic.Int64
	addr                     string
	network                  Network
	closed                   uint32
	services                 []func()
}
```
<span id="struct_Server_LoadData"></span>

#### func (*Server) LoadData(name string, data any)  any
> 加载绑定的服务器数据

***
<span id="struct_Server_BindData"></span>

#### func (*Server) BindData(name string, data any)
> 绑定数据到特定服务器

***
<span id="struct_Server_Run"></span>

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

该案例将创建一个简单的 WebSocket 服务器并启动监听 `:9999/` 作为 WebSocket 监听地址，如果需要更多的服务器类型可参考 [` Network `](#struct_Network) 部分
  - 当服务器启动失败后，将会返回错误信息并触发 panic
  - server.WithLimitLife(time.Millisecond) 通常不是在正常开发应该使用的，在这里只是为了让服务器在启动完成后的 1 毫秒后自动关闭


```go

func ExampleServer_Run() {
	srv := server.New(server.NetworkWebsocket, server.WithLimitLife(time.Millisecond))
	if err := srv.Run(":9999"); err != nil {
		panic(err)
	}
}

```

***
<span id="struct_Server_IsSocket"></span>

#### func (*Server) IsSocket()  bool
> 通过执行 Network.IsSocket 函数检查该服务器是否是 Socket 模式

**示例代码：**

该案例将创建两个不同类型的服务器，其中 WebSocket 是一个 Socket 服务器，而 Http 是一个非 Socket 服务器

可知案例输出结果为：
  - true
  - false


```go

func ExampleServer_IsSocket() {
	srv1 := server.New(server.NetworkWebsocket)
	fmt.Println(srv1.IsSocket())
	srv2 := server.New(server.NetworkHttp)
	fmt.Println(srv2.IsSocket())
}

```

<details>
<summary>查看 / 收起单元测试</summary>


这个测试检查了各个类型的服务器是否为 Socket 模式。如需查看为 Socket 模式的网络类型，请参考 [` Network.IsSocket` ](#struct_Network_IsSocket)


```go

func TestServer_IsSocket(t *testing.T) {
	var cases = []struct {
		name    string
		network server.Network
		expect  bool
	}{{name: "TestServer_IsSocket_None", network: server.NetworkNone, expect: false}, {name: "TestServer_IsSocket_Tcp", network: server.NetworkTcp, expect: true}, {name: "TestServer_IsSocket_Tcp4", network: server.NetworkTcp4, expect: true}, {name: "TestServer_IsSocket_Tcp6", network: server.NetworkTcp6, expect: true}, {name: "TestServer_IsSocket_Udp", network: server.NetworkUdp, expect: true}, {name: "TestServer_IsSocket_Udp4", network: server.NetworkUdp4, expect: true}, {name: "TestServer_IsSocket_Udp6", network: server.NetworkUdp6, expect: true}, {name: "TestServer_IsSocket_Unix", network: server.NetworkUnix, expect: true}, {name: "TestServer_IsSocket_Http", network: server.NetworkHttp, expect: false}, {name: "TestServer_IsSocket_Websocket", network: server.NetworkWebsocket, expect: true}, {name: "TestServer_IsSocket_Kcp", network: server.NetworkKcp, expect: true}, {name: "TestServer_IsSocket_GRPC", network: server.NetworkGRPC, expect: false}}
	for _, c := range cases {
		t.Run(c.name, func(t *testing.T) {
			s := server.New(c.network)
			if s.IsSocket() != c.expect {
				t.Fatalf("expect: %v, got: %v", c.expect, s.IsSocket())
			}
		})
	}
}

```


</details>


***
<span id="struct_Server_RunNone"></span>

#### func (*Server) RunNone()  error
> 是 Run("") 的简写，仅适用于运行 NetworkNone 服务器

**示例代码：**

RunNone 函数并没有特殊的意义，该函数内部调用了 `srv.Run("")` 函数，仅是一个语法糖，用来表示服务器不需要监听任何地址


```go

func ExampleServer_RunNone() {
	srv := server.New(server.NetworkNone)
	if err := srv.RunNone(); err != nil {
		panic(err)
	}
}

```

***
<span id="struct_Server_Context"></span>

#### func (*Server) Context()  context.Context
> 获取服务器上下文

***
<span id="struct_Server_TimeoutContext"></span>

#### func (*Server) TimeoutContext(timeout time.Duration) ( context.Context,  context.CancelFunc)
> 获取服务器超时上下文，context.WithTimeout 的简写

***
<span id="struct_Server_Ticker"></span>

#### func (*Server) Ticker()  *timer.Ticker
> 获取服务器定时器

***
<span id="struct_Server_Shutdown"></span>

#### func (*Server) Shutdown()
> 主动停止运行服务器

***
<span id="struct_Server_GRPCServer"></span>

#### func (*Server) GRPCServer()  *grpc.Server
> 当网络类型为 NetworkGRPC 时将被允许获取 grpc 服务器，否则将会发生 panic

***
<span id="struct_Server_HttpRouter"></span>

#### func (*Server) HttpRouter()  gin.IRouter
> 当网络类型为 NetworkHttp 时将被允许获取路由器进行路由注册，否则将会发生 panic
>   - 通过该函数注册的路由将无法在服务器关闭时正常等待请求结束
> 
> Deprecated: 从 Minotaur 0.0.29 开始，由于设计原因已弃用，该函数将直接返回 *gin.Server 对象，导致无法正常的对请求结束时进行处理

***
<span id="struct_Server_HttpServer"></span>

#### func (*Server) HttpServer()  *Http[*HttpContext]
> 替代 HttpRouter 的函数，返回一个 *Http[*HttpContext] 对象
>   - 通过该函数注册的路由将在服务器关闭时正常等待请求结束
>   - 如果需要自行包装 Context 对象，可以使用 NewHttpHandleWrapper 方法

***
<span id="struct_Server_GetMessageCount"></span>

#### func (*Server) GetMessageCount()  int64
> 获取当前服务器中消息的数量

***
<span id="struct_Server_UseShunt"></span>

#### func (*Server) UseShunt(conn *Conn, name string)
> 切换连接所使用的消息分流渠道，当分流渠道 name 不存在时将会创建一个新的分流渠道，否则将会加入已存在的分流渠道
>   - 默认情况下，所有连接都使用系统通道进行消息分发，当指定消息分流渠道且为分流消息类型时，将会使用指定的消息分流渠道进行消息分发
>   - 分流渠道会在连接断开时标记为驱逐状态，当分流渠道中的所有消息处理完毕且没有新连接使用时，将会被清除

***
<span id="struct_Server_HasShunt"></span>

#### func (*Server) HasShunt(name string)  bool
> 检查特定消息分流渠道是否存在

***
<span id="struct_Server_GetConnCurrShunt"></span>

#### func (*Server) GetConnCurrShunt(conn *Conn)  string
> 获取连接当前所使用的消息分流渠道

***
<span id="struct_Server_GetShuntNum"></span>

#### func (*Server) GetShuntNum()  int
> 获取消息分流渠道数量

***
<span id="struct_Server_PushSystemMessage"></span>

#### func (*Server) PushSystemMessage(handler func (), mark ...log.Field)
> 向服务器中推送 MessageTypeSystem 消息
>   - 系统消息仅包含一个可执行函数，将在系统分发器中执行
>   - mark 为可选的日志标记，当发生异常时，将会在日志中进行体现

***
<span id="struct_Server_PushAsyncMessage"></span>

#### func (*Server) PushAsyncMessage(caller func ()  error, callback func (err error), mark ...log.Field)
> 向服务器中推送 MessageTypeAsync 消息
>   - 异步消息将在服务器的异步消息队列中进行处理，处理完成 caller 的阻塞操作后，将会通过系统消息执行 callback 函数
>   - callback 函数将在异步消息处理完成后进行调用，无论过程是否产生 err，都将被执行，允许为 nil
>   - 需要注意的是，为了避免并发问题，caller 函数请仅处理阻塞操作，其他操作应该在 callback 函数中进行
>   - mark 为可选的日志标记，当发生异常时，将会在日志中进行体现

***
<span id="struct_Server_PushShuntAsyncMessage"></span>

#### func (*Server) PushShuntAsyncMessage(conn *Conn, caller func ()  error, callback func (err error), mark ...log.Field)
> 向特定分发器中推送 MessageTypeAsync 消息，消息执行与 MessageTypeAsync 一致
>   - 需要注意的是，当未指定 UseShunt 时，将会通过 PushAsyncMessage 进行转发
>   - mark 为可选的日志标记，当发生异常时，将会在日志中进行体现

***
<span id="struct_Server_PushPacketMessage"></span>

#### func (*Server) PushPacketMessage(conn *Conn, wst int, packet []byte, mark ...log.Field)
> 向服务器中推送 MessageTypePacket 消息
>   - 当存在 UseShunt 的选项时，将会根据选项中的 shuntMatcher 进行分发，否则将在系统分发器中处理消息

***
<span id="struct_Server_PushTickerMessage"></span>

#### func (*Server) PushTickerMessage(name string, caller func (), mark ...log.Field)
> 向服务器中推送 MessageTypeTicker 消息
>   - 通过该函数推送定时消息，当消息触发时将在系统分发器中处理消息
>   - 可通过 timer.Ticker 或第三方定时器将执行函数(caller)推送到该消息中进行处理，可有效的避免线程安全问题
>   - 参数 name 仅用作标识该定时器名称
> 
> 定时消息执行不会有特殊的处理，仅标记为定时任务，也就是允许将各类函数通过该消息发送处理，但是并不建议
>   - mark 为可选的日志标记，当发生异常时，将会在日志中进行体现

***
<span id="struct_Server_PushShuntTickerMessage"></span>

#### func (*Server) PushShuntTickerMessage(conn *Conn, name string, caller func (), mark ...log.Field)
> 向特定分发器中推送 MessageTypeTicker 消息，消息执行与 MessageTypeTicker 一致
>   - 需要注意的是，当未指定 UseShunt 时，将会通过 PushTickerMessage 进行转发
>   - mark 为可选的日志标记，当发生异常时，将会在日志中进行体现

***
<span id="struct_Server_PushUniqueAsyncMessage"></span>

#### func (*Server) PushUniqueAsyncMessage(unique string, caller func ()  error, callback func (err error), mark ...log.Field)
> 向服务器中推送 MessageTypeAsync 消息，消息执行与 MessageTypeAsync 一致
>   - 不同的是当上一个相同的 unique 消息未执行完成时，将会忽略该消息

***
<span id="struct_Server_PushUniqueShuntAsyncMessage"></span>

#### func (*Server) PushUniqueShuntAsyncMessage(conn *Conn, unique string, caller func ()  error, callback func (err error), mark ...log.Field)
> 向特定分发器中推送 MessageTypeAsync 消息，消息执行与 MessageTypeAsync 一致
>   - 需要注意的是，当未指定 UseShunt 时，将会通过系统分流渠道进行转发
>   - 不同的是当上一个相同的 unique 消息未执行完成时，将会忽略该消息

***
<span id="struct_Server_PushShuntMessage"></span>

#### func (*Server) PushShuntMessage(conn *Conn, caller func (), mark ...log.Field)
> 向特定分发器中推送 MessageTypeShunt 消息，消息执行与 MessageTypeSystem 一致，不同的是将会在特定分发器中执行

***
<span id="struct_Server_GetDurationMessageCount"></span>

#### func (*Server) GetDurationMessageCount()  int64
> 获取当前 WithMessageStatistics 设置的 duration 期间的消息量

***
<span id="struct_Server_GetDurationMessageCountByOffset"></span>

#### func (*Server) GetDurationMessageCountByOffset(offset int)  int64
> 获取特定偏移次数的 WithMessageStatistics 设置的 duration 期间的消息量
>   - 该值小于 0 时，将与 GetDurationMessageCount 无异，否则将返回 +n 个期间的消息量，例如 duration 为 1 分钟，limit 为 10，那么 offset 为 1 的情况下，获取的则是上一分钟消息量

***
<span id="struct_Server_GetAllDurationMessageCount"></span>

#### func (*Server) GetAllDurationMessageCount()  []int64
> 获取所有 WithMessageStatistics 设置的 duration 期间的消息量

***
<span id="struct_Server_HasMessageStatistics"></span>

#### func (*Server) HasMessageStatistics()  bool
> 是否了开启消息统计

***
<span id="struct_Service"></span>
### Service `INTERFACE`
兼容传统 service 设计模式的接口，通过该接口可以实现更简洁、更具有可读性的服务绑定
  - 在这之前，我们在实现功能上会将 Server 进行全局存储，之后通过 init 函数进行初始化，这样的顺序是不可控的。
```go
type Service interface {
	OnInit(srv *Server)
}
```
