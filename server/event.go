package server

import (
	"go.uber.org/zap"
	"minotaur/utils/log"
	"minotaur/utils/runtimes"
	"reflect"
)

type ServerStartBeforeEventHandle func(srv *Server)
type ServerStartFinishEventHandle func(srv *Server)
type ConnectionReceivePacketEventHandle func(conn *Conn, packet []byte)
type ConnectionOpenedEventHandle func(conn *Conn)
type ConnectionClosedEventHandle func(conn *Conn)

type event struct {
	*Server
	serverStartBeforeEventHandles       []ServerStartBeforeEventHandle
	serverStartFinishEventHandles       []ServerStartFinishEventHandle
	connectionReceivePacketEventHandles []ConnectionReceivePacketEventHandle
	connectionOpenedEventHandles        []ConnectionOpenedEventHandle
	connectionClosedEventHandles        []ConnectionClosedEventHandle
}

// RegServerStartBeforeEvent 在服务器初始化完成启动前立刻执行被注册的事件处理函数
func (slf *event) RegServerStartBeforeEvent(handle ServerStartBeforeEventHandle) {
	slf.serverStartBeforeEventHandles = append(slf.serverStartBeforeEventHandles, handle)
	log.Info("Server", zap.String("RegEvent", runtimes.CurrentRunningFuncName()), zap.String("handle", reflect.TypeOf(handle).String()))
}

func (slf *event) OnServerStartBeforeEvent() {
	for _, handle := range slf.serverStartBeforeEventHandles {
		handle(slf.Server)
	}
}

// RegServerStartFinishEvent 在服务器启动完成时将立刻执行被注册的事件处理函数
func (slf *event) RegServerStartFinishEvent(handle ServerStartFinishEventHandle) {
	slf.serverStartFinishEventHandles = append(slf.serverStartFinishEventHandles, handle)
	log.Info("Server", zap.String("RegEvent", runtimes.CurrentRunningFuncName()), zap.String("handle", reflect.TypeOf(handle).String()))
}

func (slf *event) OnServerStartFinishEvent() {
	for _, handle := range slf.serverStartFinishEventHandles {
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
	log.Info("Server", zap.String("RegEvent", runtimes.CurrentRunningFuncName()), zap.String("handle", reflect.TypeOf(handle).String()))
}

func (slf *event) OnConnectionOpenedEvent(conn *Conn) {
	for _, handle := range slf.connectionOpenedEventHandles {
		handle(conn)
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
		handle(conn, packet)
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
