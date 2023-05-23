package server

import (
	"fmt"
	"github.com/kercylan98/minotaur/utils/log"
	"github.com/kercylan98/minotaur/utils/runtimes"
	"go.uber.org/zap"
	"reflect"
	"sync"
	"time"
)

type StartBeforeEventHandle func(srv *Server)
type StartFinishEventHandle func(srv *Server)
type ConnectionReceivePacketEventHandle func(srv *Server, conn *Conn, packet []byte)
type ConnectionReceiveWebsocketPacketEventHandle func(srv *Server, conn *Conn, packet []byte, messageType int)
type ConnectionOpenedEventHandle func(srv *Server, conn *Conn)
type ConnectionClosedEventHandle func(srv *Server, conn *Conn)
type ReceiveCrossPacketEventHandle func(srv *Server, senderServerId int64, packet []byte)
type MessageErrorEventHandle func(srv *Server, message *Message, err error)
type MessageLowExecEventHandle func(srv *Server, message *Message, cost time.Duration)
type ConsoleCommandEventHandle func(srv *Server)

type event struct {
	*Server
	startBeforeEventHandles                      []StartBeforeEventHandle
	startFinishEventHandles                      []StartFinishEventHandle
	connectionReceivePacketEventHandles          []ConnectionReceivePacketEventHandle
	connectionReceiveWebsocketPacketEventHandles []ConnectionReceiveWebsocketPacketEventHandle
	connectionOpenedEventHandles                 []ConnectionOpenedEventHandle
	connectionClosedEventHandles                 []ConnectionClosedEventHandle
	receiveCrossPacketEventHandles               []ReceiveCrossPacketEventHandle
	messageErrorEventHandles                     []MessageErrorEventHandle
	messageLowExecEventHandles                   []MessageLowExecEventHandle

	consoleCommandEventHandles map[string][]ConsoleCommandEventHandle

	consoleCommandEventHandleInitOnce sync.Once
}

// RegConsoleCommandEvent 控制台收到指令时将立即执行被注册的事件处理函数
//   - 默认将注册 "exit", "quit", "close", "shutdown", "EXIT", "QUIT", "CLOSE", "SHUTDOWN" 指令作为关闭服务器的指令
//   - 可通过注册默认指令进行默认行为的覆盖
func (slf *event) RegConsoleCommandEvent(command string, handle ConsoleCommandEventHandle) {
	slf.consoleCommandEventHandleInitOnce.Do(func() {
		slf.consoleCommandEventHandles = map[string][]ConsoleCommandEventHandle{}
		go func() {
			for {
				var input string
				_, _ = fmt.Scanln(&input)
				slf.OnConsoleCommandEvent(input)
			}
		}()
	})
	slf.consoleCommandEventHandles[command] = append(slf.consoleCommandEventHandles[command], handle)
	log.Info("Server", zap.String("RegEvent", runtimes.CurrentRunningFuncName()), zap.String("handle", reflect.TypeOf(handle).String()))
}

func (slf *event) OnConsoleCommandEvent(command string) {
	handles, exist := slf.consoleCommandEventHandles[command]
	if !exist {
		switch command {
		case "exit", "quit", "close", "shutdown", "EXIT", "QUIT", "CLOSE", "SHUTDOWN":
			log.Info("Console", zap.String("Receive", command), zap.String("Action", "Shutdown"))
			slf.Server.Shutdown(nil)
			return
		}
		log.Warn("Server", zap.String("Command", "unregistered"))
	} else {
		for _, handle := range handles {
			handle(slf.Server)
		}
	}
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
	conn.Close()
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

// RegMessageErrorEvent 在处理消息发生错误时将立即执行被注册的事件处理函数
func (slf *event) RegMessageErrorEvent(handle MessageErrorEventHandle) {
	slf.messageErrorEventHandles = append(slf.messageErrorEventHandles, handle)
	log.Info("Server", zap.String("RegEvent", runtimes.CurrentRunningFuncName()), zap.String("handle", reflect.TypeOf(handle).String()))
}

func (slf *event) OnMessageErrorEvent(message *Message, err error) {
	for _, handle := range slf.messageErrorEventHandles {
		handle(slf.Server, message, err)
	}
}

// RegMessageLowExecEvent 在处理消息缓慢时将立即执行被注册的事件处理函数
func (slf *event) RegMessageLowExecEvent(handle MessageLowExecEventHandle) {
	slf.messageLowExecEventHandles = append(slf.messageLowExecEventHandles, handle)
	log.Info("Server", zap.String("RegEvent", runtimes.CurrentRunningFuncName()), zap.String("handle", reflect.TypeOf(handle).String()))
}

func (slf *event) OnMessageLowExecEvent(message *Message, cost time.Duration) {
	for _, handle := range slf.messageLowExecEventHandles {
		handle(slf.Server, message, cost)
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
