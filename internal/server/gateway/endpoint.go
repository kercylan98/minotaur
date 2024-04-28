package gateway

import (
	"github.com/alphadose/haxmap"
	"github.com/kercylan98/minotaur/server"
	"github.com/kercylan98/minotaur/server/client"
	"github.com/kercylan98/minotaur/utils/log"
	"go.uber.org/atomic"
	"sync"
	"time"
)

// NewEndpoint 创建网关端点
func NewEndpoint(name string, cli *client.Client, options ...EndpointOption) *Endpoint {
	endpoint := &Endpoint{
		name:        name,
		address:     cli.GetServerAddr(),
		connections: haxmap.New[string, *server.Conn](),
		rci:         DefaultEndpointReconnectInterval,
		cps:         DefaultEndpointConnectionPoolSize,
	}
	for _, option := range options {
		option(endpoint)
	}
	for i := 0; i < endpoint.cps; i++ {
		endpoint.client = append(endpoint.client, client.CloneClient(cli))
	}
	if endpoint.evaluator == nil {
		endpoint.evaluator = func(costUnixNano float64) float64 {
			return 1 / (1 + 1.5*time.Duration(costUnixNano).Seconds())
		}
	}
	return endpoint
}

// Endpoint 网关端点
//   - 每一个端点均表示了一个目标服务，网关会将数据包转发到该端点，由该端点负责将数据包转发到目标服务。
//   - 每个端点会建立一个连接池，默认大小为 DefaultEndpointConnectionPoolSize，可通过 WithEndpointConnectionPoolSize 进行设置。
//   - 网关在转发数据包时会自行根据延迟维护端点健康值，端点健康值越高，网关越倾向于将数据包转发到该端点。
//   - 端点支持连接未中断前始终将数据包转发到特定端点，这样可以保证连接的状态维持。
//
// 连接池：
//   - 连接池大小决定了网关服务器与端点服务器建立的连接数，例如当连接池大小为 1 时，那么所有连接到该端点的客户端都会共用一个连接。
//   - 连接池的设计可以突破单机理论 65535 个 WebSocket 客户端的限制，适当的增大连接池大小可以提高网关服务器的承载能力。
type Endpoint struct {
	gateway     *Gateway
	client      []*client.Client                   // 端点客户端
	name        string                             // 端点名称
	address     string                             // 端点地址
	state       atomic.Float64                     // 端点健康值（0为不可用，越高越优）
	evaluator   func(costUnixNano float64) float64 // 端点健康值评估函数
	connections *haxmap.Map[string, *server.Conn]  // 被该端点转发的连接列表
	rci         time.Duration                      // 端点重连间隔
	cps         int                                // 端点连接池大小
}

// start 开始与目标服务端点建立连接
func (slf *Endpoint) start(cli *client.Client) {
	for {
		cur := time.Now().UnixNano()
		if err := cli.Run(); err == nil {
			slf.state.Swap(slf.evaluator(float64(time.Now().UnixNano() - cur)))
			break
		}
		if slf.rci > 0 {
			time.Sleep(slf.rci)
		} else {
			slf.state.Swap(0)
			break
		}
	}
}

// connect 连接端点
func (slf *Endpoint) connect(gateway *Gateway) {
	slf.gateway = gateway
	var least sync.WaitGroup
	var leastOnce sync.Once
	least.Add(1)
	for _, cli := range slf.client {
		go func(cli *client.Client) {
			cli.RegConnectionOpenedEvent(func(conn *client.Client) {
				slf.gateway.OnEndpointConnectOpenedEvent(slf.gateway, slf)
			})
			cli.RegConnectionClosedEvent(func(conn *client.Client, err any) {
				slf.gateway.OnEndpointConnectClosedEvent(slf.gateway, slf)
				slf.start(cli)
			})
			cli.RegConnectionReceivePacketEvent(func(conn *client.Client, wst int, packet []byte) {
				addr, sendTime, packet, err := UnmarshalGatewayInPacket(packet)
				if err != nil {
					log.Error("Endpoint", log.String("Action", "ReceivePacket"), log.String("Name", slf.name), log.String("Addr", slf.address), log.Err(err))
					return
				}
				slf.state.Swap(slf.evaluator(float64(time.Now().UnixNano() - sendTime)))
				c, ok := slf.connections.Get(addr)
				if !ok {
					log.Error("Endpoint", log.String("Action", "ReceivePacket"), log.String("Name", slf.name), log.String("Addr", slf.address), log.String("ConnAddr", addr), log.Err(ErrConnectionNotFount))
					return
				}
				c.SetWST(wst)
				slf.gateway.OnEndpointConnectReceivePacketEvent(slf.gateway, slf, c, packet)
			})
			slf.start(cli)
			leastOnce.Do(least.Done)
		}(cli)
	}
	least.Wait()
}

// GetName 获取端点名称
func (slf *Endpoint) GetName() string {
	return slf.name
}

// GetAddress 获取端点地址
func (slf *Endpoint) GetAddress() string {
	return slf.address
}

// GetState 获取端点健康值
func (slf *Endpoint) GetState() float64 {
	return slf.state.Load()
}

// Forward 转发数据包到该端点
//   - 端点在处理数据包时，应区分数据包为普通直连数据包还是网关数据包。可通过 UnmarshalGatewayOutPacket 进行数据包解析，当解析失败且无其他数据包协议时，可认为该数据包为普通直连数据包。
func (slf *Endpoint) Forward(conn *server.Conn, packet []byte, callback ...func(err error)) {
	var err error
	packet, err = MarshalGatewayOutPacket(conn.GetID(), packet)
	if err != nil {
		if len(callback) > 0 {
			callback[0](err)
		}
		return
	}

	var superior *client.Client
	for _, cli := range slf.client {
		if cli.IsConnected() {
			superior = cli
			break
		}
	}

	var cb = func(err error) {
		if len(callback) > 0 {
			callback[0](err)
		}
		if err != nil {
			slf.connections.Del(conn.GetID())
		} else {
			slf.connections.Set(conn.GetID(), conn)
			slf.gateway.cceLock.Lock()
			slf.gateway.cce[conn.GetID()] = slf
			slf.gateway.cceLock.Unlock()
		}
	}

	if conn.IsWebsocket() {
		superior.WriteWS(conn.GetWST(), packet, cb)
	} else {
		superior.Write(packet, cb)
	}
}
