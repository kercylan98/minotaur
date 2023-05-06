package server

import (
	"go.uber.org/zap"
	"minotaur/utils/log"
	"minotaur/utils/runtimes"
	"reflect"
)

type StartBeforeEventHandle func(srv *Server)
type StartFinishEventHandle func(srv *Server)
type ConnectionReceivePacketEventHandle func(srv *Server, conn *Conn, packet []byte)
type ConnectionOpenedEventHandle func(srv *Server, conn *Conn)
type ConnectionClosedEventHandle func(srv *Server, conn *Conn)

type event struct {
	*Server
	startBeforeEventHandles             []StartBeforeEventHandle
	startFinishEventHandles             []StartFinishEventHandle
	connectionReceivePacketEventHandles []ConnectionReceivePacketEventHandle
	connectionOpenedEventHandles        []ConnectionOpenedEventHandle
	connectionClosedEventHandles        []ConnectionClosedEventHandle
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
	log.Debug("Server", zap.String("ConnectionOpened", conn.GetID()))
	for _, handle := range slf.connectionOpenedEventHandles {
		handle(slf.Server, conn)
	}
}

// RegConnectionReceivePacketEvent 在接收到数据包时将立刻执行被注册的事件处理函数
func (slf *event) RegConnectionReceivePacketEvent(handle ConnectionReceivePacketEventHandle) {
	if slf.network == NetworkHttp {
		panic(ErrNetworkIncompatibleHttp)
	}
	slf.connectionReceivePacketEventHandles = append(slf.connectionReceivePacketEventHandles, handle)
	log.Info("Server", zap.String("RegEvent", runtimes.CurrentRunningFuncName()), zap.String("handle", reflect.TypeOf(handle).String()))
}

func (slf *event) OnConnectionReceivePacketEvent(conn *Conn, packet []byte) {
	for _, handle := range slf.connectionReceivePacketEventHandles {
		handle(slf.Server, conn, packet)
	}
}

func (slf *event) check() {
	switch slf.network {
	case NetworkHttp, NetworkGRPC:
	default:
		if len(slf.connectionReceivePacketEventHandles) == 0 {
			log.Warn("Server", zap.String("ConnectionReceivePacketEvent", "Invalid server, no packets processed"))
		}
	}

}
