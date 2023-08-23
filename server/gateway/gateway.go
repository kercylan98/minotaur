package gateway

import (
	"github.com/kercylan98/minotaur/server"
	"math"
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
	slf.srv.RegConnectionOpenedEvent(slf.onConnectionOpened, math.MinInt)
	slf.srv.RegConnectionReceivePacketEvent(slf.onConnectionReceivePacket, math.MinInt)
	return slf.srv.Run(addr)
}

// Shutdown 关闭网关
func (slf *Gateway) Shutdown() {
	slf.srv.Shutdown()
}

func (slf *Gateway) onConnectionOpened(srv *server.Server, conn *server.Conn) {
	endpoint, err := slf.GetEndpoint("test", conn.GetID())
	if err != nil {
		conn.Close()
		return
	}
	endpoint.Link(conn)
}

// onConnectionReceivePacket 连接接收数据包事件
func (slf *Gateway) onConnectionReceivePacket(srv *server.Server, conn *server.Conn, packet []byte) {
	endpoint, err := slf.GetEndpoint("test", conn.GetID())
	if err != nil {
		conn.Close()
		return
	}
	packet, err = MarshalGatewayOutPacket(conn.GetID(), packet)
	if err != nil {
		conn.Close()
		return
	}
	if conn.IsWebsocket() {
		endpoint.WriteWS(conn.GetWST(), packet)
	} else {
		endpoint.Write(packet)
	}
}
