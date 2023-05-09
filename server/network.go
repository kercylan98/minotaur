package server

type Network string

const (
	NetworkTCP  Network = "tcp"
	NetworkTCP4 Network = "tcp4"
	NetworkTCP6 Network = "tcp6"
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
