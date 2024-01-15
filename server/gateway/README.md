# Gateway

[![Go doc](https://img.shields.io/badge/go.dev-reference-brightgreen?logo=go&logoColor=white&style=flat)](https://pkg.go.dev/github.com/kercylan98/minotaur)
![](https://img.shields.io/badge/Email-kercylan@gmail.com-green.svg?style=flat)

gateway 是用于处理服务器消息的网关模块，适用于对客户端消息进行处理、转发的情况。


## 目录导航
列出了该 `package` 下所有的函数及类型定义，可通过目录导航进行快捷跳转 ❤️
<details>
<summary>展开 / 折叠目录导航</summary>


> 包级函数定义

|函数名称|描述
|:--|:--
|[NewEndpoint](#NewEndpoint)|创建网关端点
|[WithEndpointStateEvaluator](#WithEndpointStateEvaluator)|设置端点健康值评估函数
|[WithEndpointConnectionPoolSize](#WithEndpointConnectionPoolSize)|设置端点连接池大小
|[WithEndpointReconnectInterval](#WithEndpointReconnectInterval)|设置端点重连间隔
|[NewGateway](#NewGateway)|基于 server.Server 创建 Gateway 网关服务器
|[WithEndpointSelector](#WithEndpointSelector)|设置端点选择器
|[MarshalGatewayOutPacket](#MarshalGatewayOutPacket)|将数据包转换为网关出网数据包
|[UnmarshalGatewayOutPacket](#UnmarshalGatewayOutPacket)|将网关出网数据包转换为数据包
|[MarshalGatewayInPacket](#MarshalGatewayInPacket)|将数据包转换为网关入网数据包
|[UnmarshalGatewayInPacket](#UnmarshalGatewayInPacket)|将网关入网数据包转换为数据包


> 类型定义

|类型|名称|描述
|:--|:--|:--
|`STRUCT`|[Endpoint](#struct_Endpoint)|网关端点
|`STRUCT`|[EndpointOption](#struct_EndpointOption)|网关端点选项
|`STRUCT`|[ConnectionOpenedEventHandle](#struct_ConnectionOpenedEventHandle)|暂无描述...
|`STRUCT`|[EndpointSelector](#struct_EndpointSelector)|暂无描述...
|`STRUCT`|[Gateway](#struct_Gateway)|基于 server.Server 实现的网关服务器
|`STRUCT`|[Option](#struct_Option)|网关选项
|`INTERFACE`|[Scanner](#struct_Scanner)|端点扫描器

</details>


***
## 详情信息
#### func NewEndpoint(name string, cli *client.Client, options ...EndpointOption) *Endpoint
<span id="NewEndpoint"></span>
> 创建网关端点

***
#### func WithEndpointStateEvaluator(evaluator func (costUnixNano float64)  float64) EndpointOption
<span id="WithEndpointStateEvaluator"></span>
> 设置端点健康值评估函数

***
#### func WithEndpointConnectionPoolSize(size int) EndpointOption
<span id="WithEndpointConnectionPoolSize"></span>
> 设置端点连接池大小
>   - 默认为 DefaultEndpointConnectionPoolSize
>   - 端点连接池大小决定了网关服务器与端点服务器建立的连接数，如果 <= 0 则会使用默认值
>   - 在网关服务器中，多个客户端在发送消息到端点服务器时，会共用一个连接，适当的增大连接池大小可以提高网关服务器的承载能力

***
#### func WithEndpointReconnectInterval(interval time.Duration) EndpointOption
<span id="WithEndpointReconnectInterval"></span>
> 设置端点重连间隔
>   - 默认为 DefaultEndpointReconnectInterval
>   - 端点在连接失败后会在该间隔后重连，如果 <= 0 则不会重连

***
#### func NewGateway(srv *server.Server, scanner Scanner, options ...Option) *Gateway
<span id="NewGateway"></span>
> 基于 server.Server 创建 Gateway 网关服务器

***
#### func WithEndpointSelector(selector EndpointSelector) Option
<span id="WithEndpointSelector"></span>
> 设置端点选择器
>   - 默认情况下，网关会随机选择一个端点作为目标，如果需要自定义端点选择器，可以通过该选项设置

***
#### func MarshalGatewayOutPacket(addr string, packet []byte) ([]byte,  error)
<span id="MarshalGatewayOutPacket"></span>
> 将数据包转换为网关出网数据包
>   - | identifier(4) | ipv4(4) | port(2) | packet |

***
#### func UnmarshalGatewayOutPacket(data []byte) (addr string, packet []byte, err error)
<span id="UnmarshalGatewayOutPacket"></span>
> 将网关出网数据包转换为数据包
>   - | identifier(4) | ipv4(4) | port(2) | packet |

***
#### func MarshalGatewayInPacket(addr string, currentTime int64, packet []byte) ([]byte,  error)
<span id="MarshalGatewayInPacket"></span>
> 将数据包转换为网关入网数据包
>   - | ipv4(4) | port(2) | cost(4) | packet |

***
#### func UnmarshalGatewayInPacket(data []byte) (addr string, sendTime int64, packet []byte, err error)
<span id="UnmarshalGatewayInPacket"></span>
> 将网关入网数据包转换为数据包
>   - | ipv4(4) | port(2) | cost(4) | packet |

***
<span id="struct_Endpoint"></span>
### Endpoint `STRUCT`
网关端点
  - 每一个端点均表示了一个目标服务，网关会将数据包转发到该端点，由该端点负责将数据包转发到目标服务。
  - 每个端点会建立一个连接池，默认大小为 DefaultEndpointConnectionPoolSize，可通过 WithEndpointConnectionPoolSize 进行设置。
  - 网关在转发数据包时会自行根据延迟维护端点健康值，端点健康值越高，网关越倾向于将数据包转发到该端点。
  - 端点支持连接未中断前始终将数据包转发到特定端点，这样可以保证连接的状态维持。

连接池：
  - 连接池大小决定了网关服务器与端点服务器建立的连接数，例如当连接池大小为 1 时，那么所有连接到该端点的客户端都会共用一个连接。
  - 连接池的设计可以突破单机理论 65535 个 WebSocket 客户端的限制，适当的增大连接池大小可以提高网关服务器的承载能力。
```go
type Endpoint struct {
	gateway     *Gateway
	client      []*client.Client
	name        string
	address     string
	state       atomic.Float64
	evaluator   func(costUnixNano float64) float64
	connections *haxmap.Map[string, *server.Conn]
	rci         time.Duration
	cps         int
}
```
<span id="struct_Endpoint_GetName"></span>

#### func (*Endpoint) GetName()  string
> 获取端点名称

***
<span id="struct_Endpoint_GetAddress"></span>

#### func (*Endpoint) GetAddress()  string
> 获取端点地址

***
<span id="struct_Endpoint_GetState"></span>

#### func (*Endpoint) GetState()  float64
> 获取端点健康值

***
<span id="struct_Endpoint_Forward"></span>

#### func (*Endpoint) Forward(conn *server.Conn, packet []byte, callback ...func (err error))
> 转发数据包到该端点
>   - 端点在处理数据包时，应区分数据包为普通直连数据包还是网关数据包。可通过 UnmarshalGatewayOutPacket 进行数据包解析，当解析失败且无其他数据包协议时，可认为该数据包为普通直连数据包。

***
<span id="struct_EndpointOption"></span>
### EndpointOption `STRUCT`
网关端点选项
```go
type EndpointOption func(endpoint *Endpoint)
```
<span id="struct_ConnectionOpenedEventHandle"></span>
### ConnectionOpenedEventHandle `STRUCT`

```go
type ConnectionOpenedEventHandle func(gateway *Gateway, conn *server.Conn)
```
<span id="struct_EndpointSelector"></span>
### EndpointSelector `STRUCT`

```go
type EndpointSelector func(endpoints []*Endpoint) *Endpoint
```
<span id="struct_Gateway"></span>
### Gateway `STRUCT`
基于 server.Server 实现的网关服务器
  - 网关服务器是一个特殊的服务器，它会通过扫描器扫描端点列表，然后连接到端点列表中的所有端点，当端点连接成功后，网关服务器会将客户端的连接数据转发到端点服务器
  - 由于该网关为多个客户端共享一个端点的连接，所以不会受限于单机 65535 个端口的限制
  - 需要注意的是，网关可以通过扫描器得到每个端点需要建立多少个连接，该连接总数将受限与 65535 个端口的限制

一些特性：
  - 支持 server.Server 所支持的所有 Socket 网络类型
  - 支持将客户端网络类型进行不同的转换，例如：客户端使用 Websocket 连接，但是网关服务器可以将其转换为 TCP 端点的连接
  - 支持客户端消息绑定，在客户端未断开连接的情况下，可以将客户端的连接绑定到某个端点，这样该客户端的所有消息都会转发到该端点
  - 根据端点延迟实时调整端点状态评分，根据评分选择最优的端点，默认评分算法为：1 / (1 + 1.5 * ${DelaySeconds})
```go
type Gateway struct {
	*events
	srv     *server.Server
	scanner Scanner
	es      map[string]map[string]*Endpoint
	esm     sync.Mutex
	ess     EndpointSelector
	closed  bool
	running bool
	cce     map[string]*Endpoint
	cceLock sync.RWMutex
}
```
<span id="struct_Gateway_Run"></span>

#### func (*Gateway) Run(addr string)  error
> 运行网关

<details>
<summary>查看 / 收起单元测试</summary>


```go

func TestGateway_Run(t *testing.T) {
	gw := gateway.NewGateway(server.New(server.NetworkWebsocket, server.WithDeadlockDetect(time.Second*3)), new(Scanner))
	gw.RegConnectionReceivePacketEventHandle(func(gateway *gateway.Gateway, conn *server.Conn, packet []byte) {
		endpoint, err := gateway.GetConnEndpoint("test", conn)
		if err == nil {
			endpoint.Forward(conn, packet)
		}
	})
	gw.RegEndpointConnectReceivePacketEventHandle(func(gateway *gateway.Gateway, endpoint *gateway.Endpoint, conn *server.Conn, packet []byte) {
		conn.Write(packet)
	})
	if err := gw.Run(":8888"); err != nil {
		panic(err)
	}
}

```


</details>


***
<span id="struct_Gateway_Shutdown"></span>

#### func (*Gateway) Shutdown()
> 关闭网关

***
<span id="struct_Gateway_Server"></span>

#### func (*Gateway) Server()  *server.Server
> 获取网关服务器核心

***
<span id="struct_Gateway_GetEndpoint"></span>

#### func (*Gateway) GetEndpoint(name string) ( *Endpoint,  error)
> 获取一个可用的端点
>   - name: 端点名称

***
<span id="struct_Gateway_GetConnEndpoint"></span>

#### func (*Gateway) GetConnEndpoint(name string, conn *server.Conn) ( *Endpoint,  error)
> 获取一个可用的端点，如果客户端已经连接到了某个端点，将优先返回该端点
>   - 当连接到的端点不可用或没有连接记录时，效果同 GetEndpoint 相同
>   - 当连接行为为有状态时，推荐使用该方法

***
<span id="struct_Gateway_SwitchEndpoint"></span>

#### func (*Gateway) SwitchEndpoint(source *Endpoint, dest *Endpoint)
> 将端点端点的所有连接切换到另一个端点

***
<span id="struct_Option"></span>
### Option `STRUCT`
网关选项
```go
type Option func(gateway *Gateway)
```
<span id="struct_Scanner"></span>
### Scanner `INTERFACE`
端点扫描器
```go
type Scanner interface {
	GetEndpoints() ([]*Endpoint, error)
	GetInterval() time.Duration
}
```
<span id="struct_Scanner_GetEndpoints"></span>

#### func (*Scanner) GetEndpoints() ( []*gateway.Endpoint,  error)

***
<span id="struct_Scanner_GetInterval"></span>

#### func (*Scanner) GetInterval()  time.Duration

***
