package server

import (
	"github.com/kercylan98/minotaur/utils/log"
	"github.com/kercylan98/minotaur/utils/runtimes"
	"go.uber.org/zap"
	"reflect"
)

type StartBeforeEventHandle func(srv *Server)
type StartFinishEventHandle func(srv *Server)
type ConnectionReceivePacketEventHandle func(srv *Server, conn *Conn, packet []byte)
type ConnectionReceiveWebsocketPacketEventHandle func(srv *Server, conn *Conn, packet []byte, messageType int)
type ConnectionOpenedEventHandle func(srv *Server, conn *Conn)
type ConnectionClosedEventHandle func(srv *Server, conn *Conn)
type ReceiveCrossPacketEventHandle func(srv *Server, senderServerId int64, packet []byte)

type event struct {
	*Server
	startBeforeEventHandles                      []StartBeforeEventHandle
	startFinishEventHandles                      []StartFinishEventHandle
	connectionReceivePacketEventHandles          []ConnectionReceivePacketEventHandle
	connectionReceiveWebsocketPacketEventHandles []ConnectionReceiveWebsocketPacketEventHandle
	connectionOpenedEventHandles                 []ConnectionOpenedEventHandle
	connectionClosedEventHandles                 []ConnectionClosedEventHandle
	receiveCrossPacketEventHandles               []ReceiveCrossPacketEventHandle
}

// RegStartBeforeEvent 在服务器初始化完成启动前立刻执行被注册的事件处理函数
func (slf *event) RegStartBeforeEvent(handle StartBeforeEventHandle) {
	slf.startBeforeEventHandles = append(slf.startBeforeEventHandles, handle)
	log.Info("Server", zap.String("RegEvent", runtimes.CurrentRunningFuncName()), zap.String("handle", reflect.TypeOf(handle).String()))
}

func (slf *event) OnStartBeforeEvent() {
	for _, handle := range slf.startBeforeEventHandles {
		handle(slf.Server)
	}
}

// RegStartFinishEvent 在服务器启动完成时将立刻执行被注册的事件处理函数
func (slf *event) RegStartFinishEvent(handle StartFinishEventHandle) {
	slf.startFinishEventHandles = append(slf.startFinishEventHandles, handle)
	log.Info("Server", zap.String("RegEvent", runtimes.CurrentRunningFuncName()), zap.String("handle", reflect.TypeOf(handle).String()))
}

func (slf *event) OnStartFinishEvent() {
	for _, handle := range slf.startFinishEventHandles {
		handle(slf.Server)
	}
}

// RegConnectionClosedEvent 在连接关闭后将立刻执行被注册的事件处理函数
func (slf *event) RegConnectionClosedEvent(handle ConnectionClosedEventHandle) {
	if slf.network == NetworkHttp {
		panic(ErrNetworkIncompatibleHttp)
	}
	slf.connectionClosedEventHandles = append(slf.connectionClosedEventHandles, handle)
	log.Info("Server", zap.String("RegEvent", runtimes.CurrentRunningFuncName()), zap.String("handle", reflect.TypeOf(handle).String()))
}

func (slf *event) OnConnectionClosedEvent(conn *Conn) {
	log.Debug("Server", zap.String("ConnectionClosed", conn.GetID()))
	for _, handle := range slf.connectionClosedEventHandles {
		handle(slf.Server, conn)
	}
}

// RegConnectionOpenedEvent 在连接打开后将立刻执行被注册的事件处理函数
func (slf *event) RegConnectionOpenedEvent(handle ConnectionOpenedEventHandle) {
	if slf.network == NetworkHttp {
		panic(ErrNetworkIncompatibleHttp)
	}
	slf.connectionOpenedEventHandles = append(slf.connectionOpenedEventHandles, handle)
	log.Info("Server", zap.String("RegEvent", runtimes.CurrentRunningFuncName()), zap.String("handle", reflect.TypeOf(handle).String()))
}

func (slf *event) OnConnectionOpenedEvent(conn *Conn) {
	if len(slf.diversionMessageChannels) == 0 {
		log.Debug("Server", zap.String("ConnectionOpened", conn.GetID()))
	} else {
		log.Debug("Server", zap.String("ConnectionOpened", conn.GetID()), zap.Int("Node", slf.diversionConsistency.PickNode(conn.GetID())))
	}
	for _, handle := range slf.connectionOpenedEventHandles {
		handle(slf.Server, conn)
	}
}

// RegConnectionReceivePacketEvent 在接收到数据包时将立刻执行被注册的事件处理函数
func (slf *event) RegConnectionReceivePacketEvent(handle ConnectionReceivePacketEventHandle) {
	if slf.network == NetworkHttp {
		panic(ErrNetworkIncompatibleHttp)
	}
	if slf.network == NetworkWebsocket {
		panic(ErrPleaseUseWebsocketHandle)
	}
	slf.connectionReceivePacketEventHandles = append(slf.connectionReceivePacketEventHandles, handle)
	log.Info("Server", zap.String("RegEvent", runtimes.CurrentRunningFuncName()), zap.String("handle", reflect.TypeOf(handle).String()))
}

func (slf *event) OnConnectionReceivePacketEvent(conn *Conn, packet []byte) {
	for _, handle := range slf.connectionReceivePacketEventHandles {
		handle(slf.Server, conn, packet)
	}
}

// RegConnectionReceiveWebsocketPacketEvent 在接收到Websocket数据包时将立刻执行被注册的事件处理函数
func (slf *event) RegConnectionReceiveWebsocketPacketEvent(handle ConnectionReceiveWebsocketPacketEventHandle) {
	if slf.network != NetworkWebsocket {
		panic(ErrPleaseUseOrdinaryPacketHandle)
	}
	slf.connectionReceiveWebsocketPacketEventHandles = append(slf.connectionReceiveWebsocketPacketEventHandles, handle)
	log.Info("Server", zap.String("RegEvent", runtimes.CurrentRunningFuncName()), zap.String("handle", reflect.TypeOf(handle).String()))
}

func (slf *event) OnConnectionReceiveWebsocketPacketEvent(conn *Conn, packet []byte, messageType int) {
	for _, handle := range slf.connectionReceiveWebsocketPacketEventHandles {
		handle(slf.Server, conn, packet, messageType)
	}
}

// RegReceiveCrossPacketEvent 在接收到跨服数据包时将立即执行被注册的事件处理函数
func (slf *event) RegReceiveCrossPacketEvent(handle ReceiveCrossPacketEventHandle) {
	slf.receiveCrossPacketEventHandles = append(slf.receiveCrossPacketEventHandles, handle)
	log.Info("Server", zap.String("RegEvent", runtimes.CurrentRunningFuncName()), zap.String("handle", reflect.TypeOf(handle).String()))
}

func (slf *event) OnReceiveCrossPacketEvent(serverId int64, packet []byte) {
	for _, handle := range slf.receiveCrossPacketEventHandles {
		handle(slf.Server, serverId, packet)
	}
}

func (slf *event) check() {
	switch slf.network {
	case NetworkHttp, NetworkGRPC:
	default:
		switch slf.network {
		case NetworkWebsocket:
			if len(slf.connectionReceiveWebsocketPacketEventHandles) == 0 {
				log.Warn("Server", zap.String("ConnectionReceiveWebsocketPacketEvent", "invalid server, no packets processed"))
			}
		default:
			if len(slf.connectionReceivePacketEventHandles) == 0 {
				log.Warn("Server", zap.String("ConnectionReceivePacketEvent", "invalid server, no packets processed"))
			}
		}
	}

	if len(slf.receiveCrossPacketEventHandles) > 0 && slf.cross == nil {
		log.Warn("Server", zap.String("ReceiveCrossPacketEvent", "invalid server, not register cross server"))
	}

}
