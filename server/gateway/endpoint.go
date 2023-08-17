package gateway

import (
	"github.com/kercylan98/minotaur/server"
	"github.com/kercylan98/minotaur/server/client"
	"github.com/kercylan98/minotaur/utils/super"
	"time"
)

// NewEndpoint 创建网关端点
func NewEndpoint(name, address string, options ...EndpointOption) *Endpoint {
	endpoint := &Endpoint{
		client:  client.NewWebsocket(address),
		name:    name,
		address: address,
	}
	for _, option := range options {
		option(endpoint)
	}
	if endpoint.evaluator == nil {
		endpoint.evaluator = func(costUnixNano float64) float64 {
			return 1 / (1 + 1.5*time.Duration(costUnixNano).Seconds())
		}
	}
	endpoint.client.RegConnectionClosedEvent(endpoint.onConnectionClosed)
	endpoint.client.RegConnectionReceivePacketEvent(endpoint.onConnectionReceivePacket)
	return endpoint
}

// Endpoint 网关端点
type Endpoint struct {
	client    *client.Websocket                  // 端点客户端
	name      string                             // 端点名称
	address   string                             // 端点地址
	state     float64                            // 端点健康值（0为不可用，越高越优）
	offline   bool                               // 离线
	evaluator func(costUnixNano float64) float64 // 端点健康值评估函数
}

// Offline 离线
func (slf *Endpoint) Offline() {
	slf.offline = true
}

// Connect 连接端点
func (slf *Endpoint) Connect() {
	for {
		cur := time.Now().UnixNano()
		if err := slf.client.Run(); err == nil {
			slf.state = slf.evaluator(float64(time.Now().UnixNano() - cur))
			break
		}
		time.Sleep(100 * time.Millisecond)
	}
}

// Write 写入数据
func (slf *Endpoint) Write(packet server.Packet) {
	slf.client.Write(packet)
}

// onConnectionClosed 与端点连接断开事件
func (slf *Endpoint) onConnectionClosed(conn *client.Websocket, err any) {
	if !slf.offline {
		go slf.Connect()
	}
}

// onConnectionReceivePacket 接收到来自端点的数据包事件
func (slf *Endpoint) onConnectionReceivePacket(conn *client.Websocket, packet server.Packet) {
	var gp server.GP
	if err := super.UnmarshalJSON(packet.Data[:len(packet.Data)-1], &gp); err != nil {
		panic(err)
	}
	cur := time.Now().UnixNano()
	slf.state = slf.evaluator(float64(cur - gp.T))
	conn.GetData(gp.C).(*server.Conn).Write(server.NewWSPacket(gp.WT, gp.D))
}
