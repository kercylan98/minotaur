package server

import (
	"github.com/kercylan98/minotaur/vivid/vivids"
	"net"
)

type Network interface {
	vivids.Actor
}

type (
	// NetworkConnectionOpenedMessage 网络连接打开消息
	NetworkConnectionOpenedMessage struct {
		net.Conn
		ConnectionWriter
	}

	// NetworkConnectionClosedMessage 网络连接关闭消息
	NetworkConnectionClosedMessage struct {
		Conn  vivids.ActorRef
		Error error
	}

	// NetworkConnectionReceivedMessage 网络连接接收消息
	NetworkConnectionReceivedMessage struct {
		net.Conn
		Packet
	}

	// NetworkConnectionAsyncWriteErrorMessage 网络连接异步写入失败消息
	NetworkConnectionAsyncWriteErrorMessage struct {
		Conn
		Packet
		Error error
	}
)
