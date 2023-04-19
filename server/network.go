package server

type Network string

const (
	NetworkTCP       Network = "tcp"
	NetworkTCP4      Network = "tcp4"
	NetworkTCP6      Network = "tcp6"
	NetworkUdp       Network = "udp"
	NetworkUdp4      Network = "udp4"
	NetworkUdp6      Network = "udp6"
	NetworkUnix      Network = "unix"
	NetworkHttp      Network = "http"
	NetworkWebsocket Network = "websocket"
)
