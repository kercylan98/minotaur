package gateway

import (
	"github.com/kercylan98/minotaur/server"
	"github.com/kercylan98/minotaur/server/client"
	"time"
)

// NewEndpoint 创建网关端点
func NewEndpoint(name, address string) *Endpoint {
	endpoint := &Endpoint{
		client:  client.NewWebsocket(address),
		name:    name,
		address: address,
	}
	endpoint.client.RegConnectionClosedEvent(endpoint.onConnectionClosed)
	endpoint.client.RegConnectionReceivePacketEvent(endpoint.onConnectionReceivePacket)
	return endpoint
}

// Endpoint 网关端点
type Endpoint struct {
	client  *client.Websocket // 端点客户端
	name    string            // 端点名称
	address string            // 端点地址
	state   float64           // 端点健康值（0为不可用，越高越优）
	offline bool              // 离线
}

// Offline 离线
func (slf *Endpoint) Offline() {
	slf.offline = true
}

// Connect 连接端点
func (slf *Endpoint) Connect() {
	for {
		var now = time.Now()
		if err := slf.client.Run(); err == nil {
			slf.state = 1 - (time.Since(now).Seconds() / 10)
			break
		}
		time.Sleep(100 * time.Millisecond)
	}
}

// Write 写入数据
func (slf *Endpoint) Write(packet server.Packet) {
	slf.client.Write(packet)
}

func (slf *Endpoint) onConnectionClosed(conn *client.Websocket, err any) {
	if !slf.offline {
		go slf.Connect()
	}
}

func (slf *Endpoint) onConnectionReceivePacket(conn *client.Websocket, packet server.Packet) {
	p := UnpackGatewayPacket(packet)
	packet.Data = p.Data
	conn.GetData(p.ConnID).(*server.Conn).Write(packet)
}
