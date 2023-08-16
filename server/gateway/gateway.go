package gateway

import (
	"github.com/kercylan98/minotaur/server"
	"github.com/kercylan98/minotaur/utils/super"
)

// NewGateway 基于 server.Server 创建网关服务器
func NewGateway(srv *server.Server, options ...Option) *Gateway {
	gateway := &Gateway{
		srv:             srv,
		EndpointManager: NewEndpointManager(),
	}
	for _, option := range options {
		option(gateway)
	}
	return gateway
}

// Gateway 网关
type Gateway struct {
	*EndpointManager                // 端点管理器
	srv              *server.Server // 网关服务器核心
}

// Run 运行网关
func (slf *Gateway) Run(addr string) error {
	slf.srv.RegConnectionOpenedEvent(slf.onConnectionOpened)
	slf.srv.RegConnectionReceivePacketEvent(slf.onConnectionReceivePacket)
	return slf.srv.Run(addr)
}

// Shutdown 关闭网关
func (slf *Gateway) Shutdown() {
	slf.srv.Shutdown()
}

// onConnectionOpened 连接打开事件
func (slf *Gateway) onConnectionOpened(srv *server.Server, conn *server.Conn) {
	endpoint, err := slf.GetEndpoint("test", conn)
	if err != nil {
		conn.Close()
		return
	}
	endpoint.client.SetData(conn.GetID(), conn)
	conn.SetData("endpoint", endpoint)
}

// onConnectionReceivePacket 连接接收数据包事件
func (slf *Gateway) onConnectionReceivePacket(srv *server.Server, conn *server.Conn, packet server.Packet) {
	conn.GetData("endpoint").(*Endpoint).Write(PackGatewayPacket(conn.GetID(), packet.WebsocketType, packet.Data))
}

// PackGatewayPacket 打包网关数据包
func PackGatewayPacket(connID string, websocketType int, data []byte) server.Packet {
	var gatewayPacket = Packet{
		ConnID:        connID,
		WebsocketType: websocketType,
		Data:          data,
	}
	return server.Packet{
		WebsocketType: websocketType,
		Data:          super.MarshalJSON(&gatewayPacket),
	}
}

// UnpackGatewayPacket 解包网关数据包
func UnpackGatewayPacket(packet server.Packet) Packet {
	var gatewayPacket Packet
	if err := super.UnmarshalJSON(packet.Data, &gatewayPacket); err != nil {
		panic(err)
	}
	return gatewayPacket
}
