package transport

import (
	"net"
)

type (
	// ServerLaunchMessage 服务器启动消息，服务器在收到该消息后将开始运行，运行过程是异步的
	ServerLaunchMessage struct {
		Network Network // 网络接口
	}

	// ServerShutdownMessage 服务器关闭消息，服务器在收到该消息后将停止运行
	ServerShutdownMessage struct{}

	ServerConnOpenedMessage struct {
		conn   net.Conn
		writer ConnWriter
	}

	ServerConnClosedMessage struct {
		conn net.Conn
	}
)
