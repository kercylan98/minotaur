package server

import (
	"fmt"
	"github.com/kercylan98/minotaur/utils/log"
	"github.com/kercylan98/minotaur/utils/runtimes"
	"github.com/kercylan98/minotaur/utils/slice"
	"reflect"
	"runtime/debug"
	"sync"
	"time"
)

type StartBeforeEventHandle func(srv *Server)
type StartFinishEventHandle func(srv *Server)
type StopEventHandle func(srv *Server)
type ConnectionReceivePacketEventHandle func(srv *Server, conn *Conn, packet []byte)
type ConnectionOpenedEventHandle func(srv *Server, conn *Conn)
type ConnectionClosedEventHandle func(srv *Server, conn *Conn, err any)
type ReceiveCrossPacketEventHandle func(srv *Server, senderServerId int64, packet []byte)
type MessageErrorEventHandle func(srv *Server, message *Message, err error)
type MessageLowExecEventHandle func(srv *Server, message *Message, cost time.Duration)
type ConsoleCommandEventHandle func(srv *Server)
type ConnectionOpenedAfterEventHandle func(srv *Server, conn *Conn)
type ConnectionWritePacketBeforeEventHandle func(srv *Server, conn *Conn, packet []byte) []byte
type ShuntChannelCreatedEventHandle func(srv *Server, guid int64)
type ShuntChannelClosedEventHandle func(srv *Server, guid int64)
type ConnectionPacketPreprocessEventHandle func(srv *Server, conn *Conn, packet []byte, abort func(), usePacket func(newPacket []byte))
type MessageExecBeforeEventHandle func(srv *Server, message *Message) bool
type MessageReadyEventHandle func(srv *Server)

func newEvent(srv *Server) *event {
	return &event{
		Server:                                 srv,
		startBeforeEventHandles:                slice.NewPriority[StartBeforeEventHandle](),
		startFinishEventHandles:                slice.NewPriority[StartFinishEventHandle](),
		stopEventHandles:                       slice.NewPriority[StopEventHandle](),
		connectionReceivePacketEventHandles:    slice.NewPriority[ConnectionReceivePacketEventHandle](),
		connectionOpenedEventHandles:           slice.NewPriority[ConnectionOpenedEventHandle](),
		connectionClosedEventHandles:           slice.NewPriority[ConnectionClosedEventHandle](),
		receiveCrossPacketEventHandles:         slice.NewPriority[ReceiveCrossPacketEventHandle](),
		messageErrorEventHandles:               slice.NewPriority[MessageErrorEventHandle](),
		messageLowExecEventHandles:             slice.NewPriority[MessageLowExecEventHandle](),
		connectionOpenedAfterEventHandles:      slice.NewPriority[ConnectionOpenedAfterEventHandle](),
		connectionWritePacketBeforeHandles:     slice.NewPriority[ConnectionWritePacketBeforeEventHandle](),
		shuntChannelCreatedEventHandles:        slice.NewPriority[ShuntChannelCreatedEventHandle](),
		shuntChannelClosedEventHandles:         slice.NewPriority[ShuntChannelClosedEventHandle](),
		connectionPacketPreprocessEventHandles: slice.NewPriority[ConnectionPacketPreprocessEventHandle](),
		messageExecBeforeEventHandles:          slice.NewPriority[MessageExecBeforeEventHandle](),
		messageReadyEventHandles:               slice.NewPriority[MessageReadyEventHandle](),
	}
}

type event struct {
	*Server
	startBeforeEventHandles                *slice.Priority[StartBeforeEventHandle]
	startFinishEventHandles                *slice.Priority[StartFinishEventHandle]
	stopEventHandles                       *slice.Priority[StopEventHandle]
	connectionReceivePacketEventHandles    *slice.Priority[ConnectionReceivePacketEventHandle]
	connectionOpenedEventHandles           *slice.Priority[ConnectionOpenedEventHandle]
	connectionClosedEventHandles           *slice.Priority[ConnectionClosedEventHandle]
	receiveCrossPacketEventHandles         *slice.Priority[ReceiveCrossPacketEventHandle]
	messageErrorEventHandles               *slice.Priority[MessageErrorEventHandle]
	messageLowExecEventHandles             *slice.Priority[MessageLowExecEventHandle]
	connectionOpenedAfterEventHandles      *slice.Priority[ConnectionOpenedAfterEventHandle]
	connectionWritePacketBeforeHandles     *slice.Priority[ConnectionWritePacketBeforeEventHandle]
	shuntChannelCreatedEventHandles        *slice.Priority[ShuntChannelCreatedEventHandle]
	shuntChannelClosedEventHandles         *slice.Priority[ShuntChannelClosedEventHandle]
	connectionPacketPreprocessEventHandles *slice.Priority[ConnectionPacketPreprocessEventHandle]
	messageExecBeforeEventHandles          *slice.Priority[MessageExecBeforeEventHandle]
	messageReadyEventHandles               *slice.Priority[MessageReadyEventHandle]

	consoleCommandEventHandles        map[string]*slice.Priority[ConsoleCommandEventHandle]
	consoleCommandEventHandleInitOnce sync.Once
}

// RegStopEvent 服务器停止时将立即执行被注册的事件处理函数
func (slf *event) RegStopEvent(handle StopEventHandle, priority ...int) {
	slf.stopEventHandles.Append(handle, slice.GetValue(priority, 0))
	log.Info("Server", log.String("RegEvent", runtimes.CurrentRunningFuncName()), log.String("handle", reflect.TypeOf(handle).String()))
}

func (slf *event) OnStopEvent() {
	slf.stopEventHandles.RangeValue(func(index int, value StopEventHandle) bool {
		value(slf.Server)
		return true
	})
}

// RegConsoleCommandEvent 控制台收到指令时将立即执行被注册的事件处理函数
//   - 默认将注册 "exit", "quit", "close", "shutdown", "EXIT", "QUIT", "CLOSE", "SHUTDOWN" 指令作为关闭服务器的指令
//   - 可通过注册默认指令进行默认行为的覆盖
func (slf *event) RegConsoleCommandEvent(command string, handle ConsoleCommandEventHandle, priority ...int) {
	slf.consoleCommandEventHandleInitOnce.Do(func() {
		slf.consoleCommandEventHandles = map[string]*slice.Priority[ConsoleCommandEventHandle]{}
		go func() {
			for {
				var input string
				_, _ = fmt.Scanln(&input)
				slf.OnConsoleCommandEvent(input)
			}
		}()
	})
	list, exist := slf.consoleCommandEventHandles[command]
	if !exist {
		list = slice.NewPriority[ConsoleCommandEventHandle]()
		slf.consoleCommandEventHandles[command] = list
	}
	list.Append(handle, slice.GetValue(priority, 0))
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
			handles.RangeValue(func(index int, value ConsoleCommandEventHandle) bool {
				value(slf.Server)
				return true
			})
		}
	}, "ConsoleCommandEvent")
}

// RegStartBeforeEvent 在服务器初始化完成启动前立刻执行被注册的事件处理函数
func (slf *event) RegStartBeforeEvent(handle StartBeforeEventHandle, priority ...int) {
	slf.startBeforeEventHandles.Append(handle, slice.GetValue(priority, 0))
	log.Info("Server", log.String("RegEvent", runtimes.CurrentRunningFuncName()), log.String("handle", reflect.TypeOf(handle).String()))
}

func (slf *event) OnStartBeforeEvent() {
	defer func() {
		if err := recover(); err != nil {
			log.Error("Server", log.String("OnStartBeforeEvent", fmt.Sprintf("%v", err)))
			debug.PrintStack()
		}
	}()
	slf.startBeforeEventHandles.RangeValue(func(index int, value StartBeforeEventHandle) bool {
		value(slf.Server)
		return true
	})
}

// RegStartFinishEvent 在服务器启动完成时将立刻执行被注册的事件处理函数
//   - 需要注意该时刻服务器已经启动完成，但是还有可能未开始处理消息，客户端有可能无法连接，如果需要在消息处理器准备就绪后执行，请使用 RegMessageReadyEvent 函数
func (slf *event) RegStartFinishEvent(handle StartFinishEventHandle, priority ...int) {
	slf.startFinishEventHandles.Append(handle, slice.GetValue(priority, 0))
	log.Info("Server", log.String("RegEvent", runtimes.CurrentRunningFuncName()), log.String("handle", reflect.TypeOf(handle).String()))
}

func (slf *event) OnStartFinishEvent() {
	PushSystemMessage(slf.Server, func() {
		slf.startFinishEventHandles.RangeValue(func(index int, value StartFinishEventHandle) bool {
			value(slf.Server)
			return true
		})
	}, "StartFinishEvent")
}

// RegConnectionClosedEvent 在连接关闭后将立刻执行被注册的事件处理函数
func (slf *event) RegConnectionClosedEvent(handle ConnectionClosedEventHandle, priority ...int) {
	if slf.network == NetworkHttp {
		panic(ErrNetworkIncompatibleHttp)
	}
	slf.connectionClosedEventHandles.Append(handle, slice.GetValue(priority, 0))
	log.Info("Server", log.String("RegEvent", runtimes.CurrentRunningFuncName()), log.String("handle", reflect.TypeOf(handle).String()))
}

func (slf *event) OnConnectionClosedEvent(conn *Conn, err any) {
	PushSystemMessage(slf.Server, func() {
		slf.connectionClosedEventHandles.RangeValue(func(index int, value ConnectionClosedEventHandle) bool {
			value(slf.Server, conn, err)
			return true
		})
		conn.Close()
		slf.Server.online.Delete(conn.GetID())
	}, "ConnectionClosedEvent")
}

// RegConnectionOpenedEvent 在连接打开后将立刻执行被注册的事件处理函数
func (slf *event) RegConnectionOpenedEvent(handle ConnectionOpenedEventHandle, priority ...int) {
	if slf.network == NetworkHttp {
		panic(ErrNetworkIncompatibleHttp)
	}
	slf.connectionOpenedEventHandles.Append(handle, slice.GetValue(priority, 0))
	log.Info("Server", log.String("RegEvent", runtimes.CurrentRunningFuncName()), log.String("handle", reflect.TypeOf(handle).String()))
}

func (slf *event) OnConnectionOpenedEvent(conn *Conn) {
	PushSystemMessage(slf.Server, func() {
		slf.Server.online.Set(conn.GetID(), conn)
		slf.connectionOpenedEventHandles.RangeValue(func(index int, value ConnectionOpenedEventHandle) bool {
			value(slf.Server, conn)
			return true
		})
	}, "ConnectionOpenedEvent")
}

// RegConnectionReceivePacketEvent 在接收到数据包时将立刻执行被注册的事件处理函数
func (slf *event) RegConnectionReceivePacketEvent(handle ConnectionReceivePacketEventHandle, priority ...int) {
	if slf.network == NetworkHttp {
		panic(ErrNetworkIncompatibleHttp)
	}
	slf.connectionReceivePacketEventHandles.Append(handle, slice.GetValue(priority, 0))
	log.Info("Server", log.String("RegEvent", runtimes.CurrentRunningFuncName()), log.String("handle", reflect.TypeOf(handle).String()))
}

func (slf *event) OnConnectionReceivePacketEvent(conn *Conn, packet []byte) {
	slf.connectionReceivePacketEventHandles.RangeValue(func(index int, value ConnectionReceivePacketEventHandle) bool {
		value(slf.Server, conn, packet)
		return true
	})
}

// RegReceiveCrossPacketEvent 在接收到跨服数据包时将立即执行被注册的事件处理函数
func (slf *event) RegReceiveCrossPacketEvent(handle ReceiveCrossPacketEventHandle, priority ...int) {
	slf.receiveCrossPacketEventHandles.Append(handle, slice.GetValue(priority, 0))
	log.Info("Server", log.String("RegEvent", runtimes.CurrentRunningFuncName()), log.String("handle", reflect.TypeOf(handle).String()))
}

func (slf *event) OnReceiveCrossPacketEvent(serverId int64, packet []byte) {
	slf.receiveCrossPacketEventHandles.RangeValue(func(index int, value ReceiveCrossPacketEventHandle) bool {
		value(slf.Server, serverId, packet)
		return true
	})
}

// RegMessageErrorEvent 在处理消息发生错误时将立即执行被注册的事件处理函数
func (slf *event) RegMessageErrorEvent(handle MessageErrorEventHandle, priority ...int) {
	slf.messageErrorEventHandles.Append(handle, slice.GetValue(priority, 0))
	log.Info("Server", log.String("RegEvent", runtimes.CurrentRunningFuncName()), log.String("handle", reflect.TypeOf(handle).String()))
}

func (slf *event) OnMessageErrorEvent(message *Message, err error) {
	PushSystemMessage(slf.Server, func() {
		slf.messageErrorEventHandles.RangeValue(func(index int, value MessageErrorEventHandle) bool {
			value(slf.Server, message, err)
			return true
		})
	}, "MessageErrorEvent")
}

// RegMessageLowExecEvent 在处理消息缓慢时将立即执行被注册的事件处理函数
func (slf *event) RegMessageLowExecEvent(handle MessageLowExecEventHandle, priority ...int) {
	slf.messageLowExecEventHandles.Append(handle, slice.GetValue(priority, 0))
	log.Info("Server", log.String("RegEvent", runtimes.CurrentRunningFuncName()), log.String("handle", reflect.TypeOf(handle).String()))
}

func (slf *event) OnMessageLowExecEvent(message *Message, cost time.Duration) {
	PushSystemMessage(slf.Server, func() {
		slf.messageLowExecEventHandles.RangeValue(func(index int, value MessageLowExecEventHandle) bool {
			value(slf.Server, message, cost)
			return true
		})
	}, "MessageLowExecEvent")
}

// RegConnectionOpenedAfterEvent 在连接打开事件处理完成后将立刻执行被注册的事件处理函数
func (slf *event) RegConnectionOpenedAfterEvent(handle ConnectionOpenedAfterEventHandle, priority ...int) {
	if slf.network == NetworkHttp {
		panic(ErrNetworkIncompatibleHttp)
	}
	slf.connectionOpenedAfterEventHandles.Append(handle, slice.GetValue(priority, 0))
	log.Info("Server", log.String("RegEvent", runtimes.CurrentRunningFuncName()), log.String("handle", reflect.TypeOf(handle).String()))
}

func (slf *event) OnConnectionOpenedAfterEvent(conn *Conn) {
	PushSystemMessage(slf.Server, func() {
		slf.connectionOpenedAfterEventHandles.RangeValue(func(index int, value ConnectionOpenedAfterEventHandle) bool {
			value(slf.Server, conn)
			return true
		})
	}, "ConnectionOpenedAfterEvent")
}

// RegConnectionWritePacketBeforeEvent 在发送数据包前将立刻执行被注册的事件处理函数
func (slf *event) RegConnectionWritePacketBeforeEvent(handle ConnectionWritePacketBeforeEventHandle, priority ...int) {
	if slf.network == NetworkHttp {
		panic(ErrNetworkIncompatibleHttp)
	}
	slf.connectionWritePacketBeforeHandles.Append(handle, slice.GetValue(priority, 0))
	log.Info("Server", log.String("RegEvent", runtimes.CurrentRunningFuncName()), log.String("handle", reflect.TypeOf(handle).String()))
}

func (slf *event) OnConnectionWritePacketBeforeEvent(conn *Conn, packet []byte) (newPacket []byte) {
	if slf.connectionWritePacketBeforeHandles.Len() == 0 {
		return packet
	}
	newPacket = packet
	slf.connectionWritePacketBeforeHandles.RangeValue(func(index int, value ConnectionWritePacketBeforeEventHandle) bool {
		newPacket = value(slf.Server, conn, newPacket)
		return true
	})
	return newPacket
}

// RegShuntChannelCreatedEvent 在分流通道创建时将立刻执行被注册的事件处理函数
func (slf *event) RegShuntChannelCreatedEvent(handle ShuntChannelCreatedEventHandle, priority ...int) {
	slf.shuntChannelCreatedEventHandles.Append(handle, slice.GetValue(priority, 0))
	log.Info("Server", log.String("RegEvent", runtimes.CurrentRunningFuncName()), log.String("handle", reflect.TypeOf(handle).String()))
}

func (slf *event) OnShuntChannelCreatedEvent(guid int64) {
	PushSystemMessage(slf.Server, func() {
		slf.shuntChannelCreatedEventHandles.RangeValue(func(index int, value ShuntChannelCreatedEventHandle) bool {
			value(slf.Server, guid)
			return true
		})
	}, "ShuntChannelCreatedEvent")
}

// RegShuntChannelCloseEvent 在分流通道关闭时将立刻执行被注册的事件处理函数
func (slf *event) RegShuntChannelCloseEvent(handle ShuntChannelClosedEventHandle, priority ...int) {
	slf.shuntChannelClosedEventHandles.Append(handle, slice.GetValue(priority, 0))
	log.Info("Server", log.String("RegEvent", runtimes.CurrentRunningFuncName()), log.String("handle", reflect.TypeOf(handle).String()))
}

func (slf *event) OnShuntChannelClosedEvent(guid int64) {
	PushSystemMessage(slf.Server, func() {
		slf.shuntChannelClosedEventHandles.RangeValue(func(index int, value ShuntChannelClosedEventHandle) bool {
			value(slf.Server, guid)
			return true
		})
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
func (slf *event) RegConnectionPacketPreprocessEvent(handle ConnectionPacketPreprocessEventHandle, priority ...int) {
	slf.connectionPacketPreprocessEventHandles.Append(handle, slice.GetValue(priority, 0))
	log.Info("Server", log.String("RegEvent", runtimes.CurrentRunningFuncName()), log.String("handle", reflect.TypeOf(handle).String()))
}

func (slf *event) OnConnectionPacketPreprocessEvent(conn *Conn, packet []byte, usePacket func(newPacket []byte)) bool {
	if slf.connectionPacketPreprocessEventHandles.Len() == 0 {
		return false
	}
	var abort = false
	slf.connectionPacketPreprocessEventHandles.RangeValue(func(index int, value ConnectionPacketPreprocessEventHandle) bool {
		value(slf.Server, conn, packet, func() { abort = true }, usePacket)
		if abort {
			return false
		}
		return true
	})
	return abort
}

// RegMessageExecBeforeEvent 在处理消息前将立刻执行被注册的事件处理函数
//   - 当返回 true 时，将继续执行后续的消息处理函数，否则将不会执行后续的消息处理函数，并且该消息将被丢弃
//
// 适用于限流等场景
func (slf *event) RegMessageExecBeforeEvent(handle MessageExecBeforeEventHandle, priority ...int) {
	slf.messageExecBeforeEventHandles.Append(handle, slice.GetValue(priority, 0))
	log.Info("Server", log.String("RegEvent", runtimes.CurrentRunningFuncName()), log.String("handle", reflect.TypeOf(handle).String()))
}

func (slf *event) OnMessageExecBeforeEvent(message *Message) bool {
	if slf.messageExecBeforeEventHandles.Len() == 0 {
		return true
	}
	var result = true
	defer func() {
		if err := recover(); err != nil {
			log.Error("Server", log.String("OnMessageExecBeforeEvent", fmt.Sprintf("%v", err)))
			debug.PrintStack()
		}
	}()
	slf.messageExecBeforeEventHandles.RangeValue(func(index int, value MessageExecBeforeEventHandle) bool {
		result = value(slf.Server, message)
		return result
	})
	return result
}

// RegMessageReadyEvent 在服务器消息处理器准备就绪时立即执行被注册的事件处理函数
func (slf *event) RegMessageReadyEvent(handle MessageReadyEventHandle, priority ...int) {
	slf.messageReadyEventHandles.Append(handle, slice.GetValue(priority, 0))
}

func (slf *event) OnMessageReadyEvent() {
	if slf.messageReadyEventHandles.Len() == 0 {
		return
	}
	defer func() {
		if err := recover(); err != nil {
			log.Error("Server", log.String("OnMessageReadyEvent", fmt.Sprintf("%v", err)))
			debug.PrintStack()
		}
	}()
	slf.messageReadyEventHandles.RangeValue(func(index int, value MessageReadyEventHandle) bool {
		value(slf.Server)
		return true
	})
}

func (slf *event) check() {
	switch slf.network {
	case NetworkHttp, NetworkGRPC, NetworkNone:
	default:
		if slf.connectionReceivePacketEventHandles.Len() == 0 {
			log.Warn("Server", log.String("ConnectionReceivePacketEvent", "invalid server, no packets processed"))
		}
	}

	if slf.receiveCrossPacketEventHandles.Len() > 0 && slf.cross == nil {
		log.Warn("Server", log.String("ReceiveCrossPacketEvent", "invalid server, not register cross server"))
	}

}
