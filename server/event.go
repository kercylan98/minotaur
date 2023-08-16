package server

import (
	"fmt"
	"github.com/kercylan98/minotaur/utils/log"
	"github.com/kercylan98/minotaur/utils/runtimes"
	"reflect"
	"runtime/debug"
	"sync"
	"time"
)

type StartBeforeEventHandle func(srv *Server)
type StartFinishEventHandle func(srv *Server)
type StopEventHandle func(srv *Server)
type ConnectionReceivePacketEventHandle func(srv *Server, conn *Conn, packet Packet)
type ConnectionOpenedEventHandle func(srv *Server, conn *Conn)
type ConnectionClosedEventHandle func(srv *Server, conn *Conn, err any)
type ReceiveCrossPacketEventHandle func(srv *Server, senderServerId int64, packet []byte)
type MessageErrorEventHandle func(srv *Server, message *Message, err error)
type MessageLowExecEventHandle func(srv *Server, message *Message, cost time.Duration)
type ConsoleCommandEventHandle func(srv *Server)
type ConnectionOpenedAfterEventHandle func(srv *Server, conn *Conn)
type ConnectionWritePacketBeforeEventHandle func(srv *Server, conn *Conn, packet Packet) Packet
type ShuntChannelCreatedEventHandle func(srv *Server, guid int64)
type ShuntChannelClosedEventHandle func(srv *Server, guid int64)
type ConnectionPacketPreprocessEventHandle func(srv *Server, conn *Conn, packet []byte, abort func(), usePacket func(newPacket []byte))

type event struct {
	*Server
	startBeforeEventHandles                []StartBeforeEventHandle
	startFinishEventHandles                []StartFinishEventHandle
	stopEventHandles                       []StopEventHandle
	connectionReceivePacketEventHandles    []ConnectionReceivePacketEventHandle
	connectionOpenedEventHandles           []ConnectionOpenedEventHandle
	connectionClosedEventHandles           []ConnectionClosedEventHandle
	receiveCrossPacketEventHandles         []ReceiveCrossPacketEventHandle
	messageErrorEventHandles               []MessageErrorEventHandle
	messageLowExecEventHandles             []MessageLowExecEventHandle
	connectionOpenedAfterEventHandles      []ConnectionOpenedAfterEventHandle
	connectionWritePacketBeforeHandles     []ConnectionWritePacketBeforeEventHandle
	shuntChannelCreatedEventHandles        []ShuntChannelCreatedEventHandle
	shuntChannelClosedEventHandles         []ShuntChannelClosedEventHandle
	connectionPacketPreprocessEventHandles []ConnectionPacketPreprocessEventHandle

	consoleCommandEventHandles        map[string][]ConsoleCommandEventHandle
	consoleCommandEventHandleInitOnce sync.Once
}

// RegStopEvent 服务器停止时将立即执行被注册的事件处理函数
func (slf *event) RegStopEvent(handle StopEventHandle) {
	slf.stopEventHandles = append(slf.stopEventHandles, handle)
	log.Info("Server", log.String("RegEvent", runtimes.CurrentRunningFuncName()), log.String("handle", reflect.TypeOf(handle).String()))
}

func (slf *event) OnStopEvent() {
	for _, handle := range slf.stopEventHandles {
		handle(slf.Server)
	}
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
	log.Info("Server", log.String("RegEvent", runtimes.CurrentRunningFuncName()), log.String("handle", reflect.TypeOf(handle).String()))
}

func (slf *event) OnConsoleCommandEvent(command string) {
	PushSystemMessage(slf.Server, func() {
		handles, exist := slf.consoleCommandEventHandles[command]
		if !exist {
			switch command {
			case "exit", "quit", "close", "shutdown", "EXIT", "QUIT", "CLOSE", "SHUTDOWN":
				log.Info("Console", log.String("Receive", command), log.String("Action", "Shutdown"))
				slf.Server.shutdown(nil)
				return
			}
			log.Warn("Server", log.String("Command", "unregistered"))
		} else {
			for _, handle := range handles {
				handle(slf.Server)
			}
		}
	}, "ConsoleCommandEvent")
}

// RegStartBeforeEvent 在服务器初始化完成启动前立刻执行被注册的事件处理函数
func (slf *event) RegStartBeforeEvent(handle StartBeforeEventHandle) {
	slf.startBeforeEventHandles = append(slf.startBeforeEventHandles, handle)
	log.Info("Server", log.String("RegEvent", runtimes.CurrentRunningFuncName()), log.String("handle", reflect.TypeOf(handle).String()))
}

func (slf *event) OnStartBeforeEvent() {
	defer func() {
		if err := recover(); err != nil {
			log.Error("Server", log.String("OnStartBeforeEvent", fmt.Sprintf("%v", err)))
			debug.PrintStack()
		}
	}()
	for _, handle := range slf.startBeforeEventHandles {
		handle(slf.Server)
	}
}

// RegStartFinishEvent 在服务器启动完成时将立刻执行被注册的事件处理函数
func (slf *event) RegStartFinishEvent(handle StartFinishEventHandle) {
	slf.startFinishEventHandles = append(slf.startFinishEventHandles, handle)
	log.Info("Server", log.String("RegEvent", runtimes.CurrentRunningFuncName()), log.String("handle", reflect.TypeOf(handle).String()))
}

func (slf *event) OnStartFinishEvent() {
	PushSystemMessage(slf.Server, func() {
		for _, handle := range slf.startFinishEventHandles {
			handle(slf.Server)
		}
	}, "StartFinishEvent")
}

// RegConnectionClosedEvent 在连接关闭后将立刻执行被注册的事件处理函数
func (slf *event) RegConnectionClosedEvent(handle ConnectionClosedEventHandle) {
	if slf.network == NetworkHttp {
		panic(ErrNetworkIncompatibleHttp)
	}
	slf.connectionClosedEventHandles = append(slf.connectionClosedEventHandles, handle)
	log.Info("Server", log.String("RegEvent", runtimes.CurrentRunningFuncName()), log.String("handle", reflect.TypeOf(handle).String()))
}

func (slf *event) OnConnectionClosedEvent(conn *Conn, err any) {
	PushSystemMessage(slf.Server, func() {
		for _, handle := range slf.connectionClosedEventHandles {
			handle(slf.Server, conn, err)
		}
		conn.Close()
		slf.Server.online.Delete(conn.GetID())
	}, "ConnectionClosedEvent")
}

// RegConnectionOpenedEvent 在连接打开后将立刻执行被注册的事件处理函数
func (slf *event) RegConnectionOpenedEvent(handle ConnectionOpenedEventHandle) {
	if slf.network == NetworkHttp {
		panic(ErrNetworkIncompatibleHttp)
	}
	slf.connectionOpenedEventHandles = append(slf.connectionOpenedEventHandles, handle)
	log.Info("Server", log.String("RegEvent", runtimes.CurrentRunningFuncName()), log.String("handle", reflect.TypeOf(handle).String()))
}

func (slf *event) OnConnectionOpenedEvent(conn *Conn) {
	PushSystemMessage(slf.Server, func() {
		slf.Server.online.Set(conn.GetID(), conn)
		for _, handle := range slf.connectionOpenedEventHandles {
			handle(slf.Server, conn)
		}
	}, "ConnectionOpenedEvent")
}

// RegConnectionReceivePacketEvent 在接收到数据包时将立刻执行被注册的事件处理函数
func (slf *event) RegConnectionReceivePacketEvent(handle ConnectionReceivePacketEventHandle) {
	if slf.network == NetworkHttp {
		panic(ErrNetworkIncompatibleHttp)
	}
	slf.connectionReceivePacketEventHandles = append(slf.connectionReceivePacketEventHandles, handle)
	log.Info("Server", log.String("RegEvent", runtimes.CurrentRunningFuncName()), log.String("handle", reflect.TypeOf(handle).String()))
}

func (slf *event) OnConnectionReceivePacketEvent(conn *Conn, packet Packet) {
	for _, handle := range slf.connectionReceivePacketEventHandles {
		handle(slf.Server, conn, packet)
	}
}

// RegReceiveCrossPacketEvent 在接收到跨服数据包时将立即执行被注册的事件处理函数
func (slf *event) RegReceiveCrossPacketEvent(handle ReceiveCrossPacketEventHandle) {
	slf.receiveCrossPacketEventHandles = append(slf.receiveCrossPacketEventHandles, handle)
	log.Info("Server", log.String("RegEvent", runtimes.CurrentRunningFuncName()), log.String("handle", reflect.TypeOf(handle).String()))
}

func (slf *event) OnReceiveCrossPacketEvent(serverId int64, packet []byte) {
	for _, handle := range slf.receiveCrossPacketEventHandles {
		handle(slf.Server, serverId, packet)
	}
}

// RegMessageErrorEvent 在处理消息发生错误时将立即执行被注册的事件处理函数
func (slf *event) RegMessageErrorEvent(handle MessageErrorEventHandle) {
	slf.messageErrorEventHandles = append(slf.messageErrorEventHandles, handle)
	log.Info("Server", log.String("RegEvent", runtimes.CurrentRunningFuncName()), log.String("handle", reflect.TypeOf(handle).String()))
}

func (slf *event) OnMessageErrorEvent(message *Message, err error) {
	PushSystemMessage(slf.Server, func() {
		for _, handle := range slf.messageErrorEventHandles {
			handle(slf.Server, message, err)
		}
	}, "MessageErrorEvent")
}

// RegMessageLowExecEvent 在处理消息缓慢时将立即执行被注册的事件处理函数
func (slf *event) RegMessageLowExecEvent(handle MessageLowExecEventHandle) {
	slf.messageLowExecEventHandles = append(slf.messageLowExecEventHandles, handle)
	log.Info("Server", log.String("RegEvent", runtimes.CurrentRunningFuncName()), log.String("handle", reflect.TypeOf(handle).String()))
}

func (slf *event) OnMessageLowExecEvent(message *Message, cost time.Duration) {
	PushSystemMessage(slf.Server, func() {
		for _, handle := range slf.messageLowExecEventHandles {
			handle(slf.Server, message, cost)
		}
	}, "MessageLowExecEvent")
}

// RegConnectionOpenedAfterEvent 在连接打开事件处理完成后将立刻执行被注册的事件处理函数
func (slf *event) RegConnectionOpenedAfterEvent(handle ConnectionOpenedAfterEventHandle) {
	if slf.network == NetworkHttp {
		panic(ErrNetworkIncompatibleHttp)
	}
	slf.connectionOpenedAfterEventHandles = append(slf.connectionOpenedAfterEventHandles, handle)
	log.Info("Server", log.String("RegEvent", runtimes.CurrentRunningFuncName()), log.String("handle", reflect.TypeOf(handle).String()))
}

func (slf *event) OnConnectionOpenedAfterEvent(conn *Conn) {
	PushSystemMessage(slf.Server, func() {
		for _, handle := range slf.connectionOpenedAfterEventHandles {
			handle(slf.Server, conn)
		}
	}, "ConnectionOpenedAfterEvent")
}

// RegConnectionWritePacketBeforeEvent 在发送数据包前将立刻执行被注册的事件处理函数
func (slf *event) RegConnectionWritePacketBeforeEvent(handle ConnectionWritePacketBeforeEventHandle) {
	if slf.network == NetworkHttp {
		panic(ErrNetworkIncompatibleHttp)
	}
	slf.connectionWritePacketBeforeHandles = append(slf.connectionWritePacketBeforeHandles, handle)
	log.Info("Server", log.String("RegEvent", runtimes.CurrentRunningFuncName()), log.String("handle", reflect.TypeOf(handle).String()))
}

func (slf *event) OnConnectionWritePacketBeforeEvent(conn *Conn, packet Packet) (newPacket Packet) {
	if len(slf.connectionWritePacketBeforeHandles) == 0 {
		return packet
	}
	newPacket = packet
	for _, handle := range slf.connectionWritePacketBeforeHandles {
		newPacket = handle(slf.Server, conn, packet)
	}
	return newPacket
}

// RegShuntChannelCreatedEvent 在分流通道创建时将立刻执行被注册的事件处理函数
func (slf *event) RegShuntChannelCreatedEvent(handle ShuntChannelCreatedEventHandle) {
	slf.shuntChannelCreatedEventHandles = append(slf.shuntChannelCreatedEventHandles, handle)
	log.Info("Server", log.String("RegEvent", runtimes.CurrentRunningFuncName()), log.String("handle", reflect.TypeOf(handle).String()))
}

func (slf *event) OnShuntChannelCreatedEvent(guid int64) {
	PushSystemMessage(slf.Server, func() {
		for _, handle := range slf.shuntChannelCreatedEventHandles {
			handle(slf.Server, guid)
		}
	}, "ShuntChannelCreatedEvent")
}

// RegShuntChannelCloseEvent 在分流通道关闭时将立刻执行被注册的事件处理函数
func (slf *event) RegShuntChannelCloseEvent(handle ShuntChannelClosedEventHandle) {
	slf.shuntChannelClosedEventHandles = append(slf.shuntChannelClosedEventHandles, handle)
	log.Info("Server", log.String("RegEvent", runtimes.CurrentRunningFuncName()), log.String("handle", reflect.TypeOf(handle).String()))
}

func (slf *event) OnShuntChannelClosedEvent(guid int64) {
	PushSystemMessage(slf.Server, func() {
		for _, handle := range slf.shuntChannelClosedEventHandles {
			handle(slf.Server, guid)
		}
	}, "ShuntChannelCloseEvent")
}

// RegConnectionPacketPreprocessEvent 在接收到数据包后将立刻执行被注册的事件处理函数
//   - 预处理函数可以用于对数据包进行预处理，如解密、解压缩等
//   - 在调用 abort() 后，将不会再调用后续的预处理函数，也不会调用 OnConnectionReceivePacketEvent 函数
//   - 在调用 usePacket() 后，将使用新的数据包，而不会使用原始数据包，同时阻止后续的预处理函数的调用
//
// 场景：
//   - 数据包格式校验
//   - 数据包分包等情况处理
func (slf *event) RegConnectionPacketPreprocessEvent(handle ConnectionPacketPreprocessEventHandle) {
	slf.connectionPacketPreprocessEventHandles = append(slf.connectionPacketPreprocessEventHandles, handle)
	log.Info("Server", log.String("RegEvent", runtimes.CurrentRunningFuncName()), log.String("handle", reflect.TypeOf(handle).String()))
}

func (slf *event) OnConnectionPacketPreprocessEvent(conn *Conn, packet []byte, usePacket func(newPacket []byte)) bool {
	if len(slf.connectionPacketPreprocessEventHandles) == 0 {
		return false
	}
	var abort = false
	for _, handle := range slf.connectionPacketPreprocessEventHandles {
		handle(slf.Server, conn, packet, func() { abort = true }, usePacket)
		if abort {
			return abort
		}
	}
	return abort
}

func (slf *event) check() {
	switch slf.network {
	case NetworkHttp, NetworkGRPC:
	default:
		if len(slf.connectionReceivePacketEventHandles) == 0 {
			log.Warn("Server", log.String("ConnectionReceivePacketEvent", "invalid server, no packets processed"))
		}
	}

	if len(slf.receiveCrossPacketEventHandles) > 0 && slf.cross == nil {
		log.Warn("Server", log.String("ReceiveCrossPacketEvent", "invalid server, not register cross server"))
	}

}
