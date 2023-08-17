package server

import "github.com/kercylan98/minotaur/utils/slice"

type Network string

const (
	// NetworkNone 该模式下不监听任何网络端口，仅开启消息队列，适用于纯粹的跨服服务器等情况
	NetworkNone Network = "none"
	NetworkTcp  Network = "tcp"
	NetworkTcp4 Network = "tcp4"
	NetworkTcp6 Network = "tcp6"
	NetworkUdp  Network = "udp"
	NetworkUdp4 Network = "udp4"
	NetworkUdp6 Network = "udp6"
	NetworkUnix Network = "unix"
	NetworkHttp Network = "http"
	// NetworkWebsocket 该模式下需要获取url参数值时，可以通过连接的GetData函数获取
	//  - 当有多个同名参数时，获取到的值为切片类型
	NetworkWebsocket Network = "websocket"
	NetworkKcp       Network = "kcp"
	NetworkGRPC      Network = "grpc"
)

var (
	networks = []Network{
		NetworkNone, NetworkTcp, NetworkTcp4, NetworkTcp6, NetworkUdp, NetworkUdp4, NetworkUdp6, NetworkUnix, NetworkHttp, NetworkWebsocket, NetworkKcp, NetworkGRPC,
	}
)

// GetNetworks 获取所有支持的网络模式
func GetNetworks() []Network {
	return slice.Copy(networks)
}
