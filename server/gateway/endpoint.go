package gateway

import (
	"github.com/alphadose/haxmap"
	"github.com/kercylan98/minotaur/server"
	"github.com/kercylan98/minotaur/server/client"
	"github.com/kercylan98/minotaur/utils/log"
	"time"
)

var DefaultEndpointReconnectInterval = time.Second

// NewEndpoint 创建网关端点
func NewEndpoint(name string, cli *client.Client, options ...EndpointOption) *Endpoint {
	endpoint := &Endpoint{
		client:      cli,
		name:        name,
		address:     cli.GetServerAddr(),
		connections: haxmap.New[string, *server.Conn](),
		rci:         DefaultEndpointReconnectInterval,
	}
	for _, option := range options {
		option(endpoint)
	}
	if endpoint.evaluator == nil {
		endpoint.evaluator = func(costUnixNano float64) float64 {
			return 1 / (1 + 1.5*time.Duration(costUnixNano).Seconds())
		}
	}
	return endpoint
}

// Endpoint 网关端点
type Endpoint struct {
	gateway     *Gateway
	client      *client.Client                     // 端点客户端
	name        string                             // 端点名称
	address     string                             // 端点地址
	state       float64                            // 端点健康值（0为不可用，越高越优）
	evaluator   func(costUnixNano float64) float64 // 端点健康值评估函数
	connections *haxmap.Map[string, *server.Conn]  // 被该端点转发的连接列表
	rci         time.Duration                      // 端点重连间隔
}

// connect 连接端点
func (slf *Endpoint) connect(gateway *Gateway) {
	slf.gateway = gateway
	slf.client.RegConnectionOpenedEvent(func(conn *client.Client) {
		slf.gateway.OnEndpointConnectOpenedEvent(slf.gateway, slf)
	})
	slf.client.RegConnectionClosedEvent(func(conn *client.Client, err any) {
		slf.gateway.OnEndpointConnectClosedEvent(slf.gateway, slf)
		for {
			cur := time.Now().UnixNano()
			if err := slf.client.Run(); err == nil {
				slf.state = slf.evaluator(float64(time.Now().UnixNano() - cur))
				break
			}
			if slf.rci > 0 {
				time.Sleep(slf.rci)
			} else {
				slf.state = 0
				break
			}
		}
	})
	slf.client.RegConnectionReceivePacketEvent(func(conn *client.Client, wst int, packet []byte) {
		addr, sendTime, packet, err := UnmarshalGatewayInPacket(packet)
		if err != nil {
			log.Error("Endpoint", log.String("Action", "ReceivePacket"), log.String("Name", slf.name), log.String("Addr", slf.address), log.Err(err))
			return
		}
		slf.state = slf.evaluator(float64(time.Now().UnixNano() - sendTime))
		c, ok := slf.connections.Get(addr)
		if !ok {
			log.Error("Endpoint", log.String("Action", "ReceivePacket"), log.String("Name", slf.name), log.String("Addr", slf.address), log.String("ConnAddr", addr), log.Err(ErrConnectionNotFount))
			return
		}
		c.SetWST(wst)
		slf.gateway.OnEndpointConnectReceivePacketEvent(slf.gateway, slf, c, packet)
	})
	for {
		cur := time.Now().UnixNano()
		if err := slf.client.Run(); err == nil {
			slf.state = slf.evaluator(float64(time.Now().UnixNano() - cur))
			break
		}
		if slf.rci > 0 {
			time.Sleep(slf.rci)
		} else {
			slf.state = 0
			break
		}
	}
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
	return slf.state
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

	cb := func(err error) {
		if len(callback) > 0 {
			callback[0](err)
		}
		if err != nil {
			slf.connections.Del(conn.GetID())
		} else {
			slf.connections.Set(conn.GetID(), conn)
		}
	}
	if conn.IsWebsocket() {
		slf.client.WriteWS(conn.GetWST(), packet, cb)
	} else {
		slf.client.Write(packet, cb)
	}
}
