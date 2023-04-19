package server

import (
	"go.uber.org/zap"
	"minotaur/utils/log"
)

type ReceivePacketEventHandle func(conn *Conn, packet []byte)
type ConnectionOpenedEventHandle func(conn *Conn)
type ConnectionClosedEventHandle func(conn *Conn)

type event struct {
	*Server
	receivePacketEventHandles    []ReceivePacketEventHandle
	connectionOpenedEventHandles []ConnectionOpenedEventHandle
	connectionClosedEventHandles []ConnectionClosedEventHandle
}

// RegConnectionClosedEvent 在连接关闭后将立刻执行被注册的事件处理函数
func (slf *event) RegConnectionClosedEvent(handle ConnectionClosedEventHandle) {
	if slf.network == NetworkHttp {
		panic(ErrNetworkIncompatibleHttp)
	}
	slf.connectionClosedEventHandles = append(slf.connectionClosedEventHandles, handle)
}

func (slf *event) OnConnectionClosedEvent(conn *Conn) {
	for _, handle := range slf.connectionClosedEventHandles {
		handle(conn)
	}
}

// RegConnectionOpenedEvent 在连接打开后将立刻执行被注册的事件处理函数
func (slf *event) RegConnectionOpenedEvent(handle ConnectionOpenedEventHandle) {
	if slf.network == NetworkHttp {
		panic(ErrNetworkIncompatibleHttp)
	}
	slf.connectionOpenedEventHandles = append(slf.connectionOpenedEventHandles, handle)
}

func (slf *event) OnConnectionOpenedEvent(conn *Conn) {
	for _, handle := range slf.connectionOpenedEventHandles {
		handle(conn)
	}
}

// RegReceivePacketEvent 在接收到数据包时将立刻执行被注册的事件处理函数
func (slf *event) RegReceivePacketEvent(handle ReceivePacketEventHandle) {
	if slf.network == NetworkHttp {
		panic(ErrNetworkIncompatibleHttp)
	}
	slf.receivePacketEventHandles = append(slf.receivePacketEventHandles, handle)
}

func (slf *event) OnReceivePacketEvent(conn *Conn, packet []byte) {
	for _, handle := range slf.receivePacketEventHandles {
		handle(conn, packet)
	}
}

func (slf *event) check() {
	if len(slf.receivePacketEventHandles) == 0 {
		log.Warn("Server", zap.String("ReceivePacketEvent", "Invalid server, no packets processed"))
	}
}
