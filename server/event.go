package server

import (
	"fmt"
	"github.com/kercylan98/minotaur/utils/log"
	"github.com/kercylan98/minotaur/utils/runtimes"
	"github.com/kercylan98/minotaur/utils/slice"
	"golang.org/x/crypto/ssh/terminal"
	"net/url"
	"os"
	"reflect"
	"runtime/debug"
	"strings"
	"sync"
	"time"
)

type (
	StartBeforeEventHandler                 func(srv *Server)
	StartFinishEventHandler                 func(srv *Server)
	StopEventHandler                        func(srv *Server)
	ConnectionReceivePacketEventHandler     func(srv *Server, conn *Conn, packet []byte)
	ConnectionOpenedEventHandler            func(srv *Server, conn *Conn)
	ConnectionClosedEventHandler            func(srv *Server, conn *Conn, err any)
	MessageErrorEventHandler                func(srv *Server, message *Message, err error)
	MessageLowExecEventHandler              func(srv *Server, message *Message, cost time.Duration)
	ConsoleCommandEventHandler              func(srv *Server, command string, params ConsoleParams)
	ConnectionOpenedAfterEventHandler       func(srv *Server, conn *Conn)
	ConnectionWritePacketBeforeEventHandler func(srv *Server, conn *Conn, packet []byte) []byte
	ShuntChannelCreatedEventHandler         func(srv *Server, guid int64)
	ShuntChannelClosedEventHandler          func(srv *Server, guid int64)
	ConnectionPacketPreprocessEventHandler  func(srv *Server, conn *Conn, packet []byte, abort func(), usePacket func(newPacket []byte))
	MessageExecBeforeEventHandler           func(srv *Server, message *Message) bool
	MessageReadyEventHandler                func(srv *Server)
	OnDeadlockDetectEventHandler            func(srv *Server, message *Message)
)

func newEvent(srv *Server) *event {
	return &event{
		Server:                                  srv,
		startBeforeEventHandlers:                slice.NewPriority[StartBeforeEventHandler](),
		startFinishEventHandlers:                slice.NewPriority[StartFinishEventHandler](),
		stopEventHandlers:                       slice.NewPriority[StopEventHandler](),
		connectionReceivePacketEventHandlers:    slice.NewPriority[ConnectionReceivePacketEventHandler](),
		connectionOpenedEventHandlers:           slice.NewPriority[ConnectionOpenedEventHandler](),
		connectionClosedEventHandlers:           slice.NewPriority[ConnectionClosedEventHandler](),
		messageErrorEventHandlers:               slice.NewPriority[MessageErrorEventHandler](),
		messageLowExecEventHandlers:             slice.NewPriority[MessageLowExecEventHandler](),
		connectionOpenedAfterEventHandlers:      slice.NewPriority[ConnectionOpenedAfterEventHandler](),
		connectionWritePacketBeforeHandlers:     slice.NewPriority[ConnectionWritePacketBeforeEventHandler](),
		shuntChannelCreatedEventHandlers:        slice.NewPriority[ShuntChannelCreatedEventHandler](),
		shuntChannelClosedEventHandlers:         slice.NewPriority[ShuntChannelClosedEventHandler](),
		connectionPacketPreprocessEventHandlers: slice.NewPriority[ConnectionPacketPreprocessEventHandler](),
		messageExecBeforeEventHandlers:          slice.NewPriority[MessageExecBeforeEventHandler](),
		messageReadyEventHandlers:               slice.NewPriority[MessageReadyEventHandler](),
		deadlockDetectEventHandlers:             slice.NewPriority[OnDeadlockDetectEventHandler](),
	}
}

type event struct {
	*Server
	startBeforeEventHandlers                *slice.Priority[StartBeforeEventHandler]
	startFinishEventHandlers                *slice.Priority[StartFinishEventHandler]
	stopEventHandlers                       *slice.Priority[StopEventHandler]
	connectionReceivePacketEventHandlers    *slice.Priority[ConnectionReceivePacketEventHandler]
	connectionOpenedEventHandlers           *slice.Priority[ConnectionOpenedEventHandler]
	connectionClosedEventHandlers           *slice.Priority[ConnectionClosedEventHandler]
	messageErrorEventHandlers               *slice.Priority[MessageErrorEventHandler]
	messageLowExecEventHandlers             *slice.Priority[MessageLowExecEventHandler]
	connectionOpenedAfterEventHandlers      *slice.Priority[ConnectionOpenedAfterEventHandler]
	connectionWritePacketBeforeHandlers     *slice.Priority[ConnectionWritePacketBeforeEventHandler]
	shuntChannelCreatedEventHandlers        *slice.Priority[ShuntChannelCreatedEventHandler]
	shuntChannelClosedEventHandlers         *slice.Priority[ShuntChannelClosedEventHandler]
	connectionPacketPreprocessEventHandlers *slice.Priority[ConnectionPacketPreprocessEventHandler]
	messageExecBeforeEventHandlers          *slice.Priority[MessageExecBeforeEventHandler]
	messageReadyEventHandlers               *slice.Priority[MessageReadyEventHandler]
	deadlockDetectEventHandlers             *slice.Priority[OnDeadlockDetectEventHandler]

	consoleCommandEventHandlers        map[string]*slice.Priority[ConsoleCommandEventHandler]
	consoleCommandEventHandlerInitOnce sync.Once
}

// RegStopEvent 服务器停止时将立即执行被注册的事件处理函数
func (slf *event) RegStopEvent(handler StopEventHandler, priority ...int) {
	slf.stopEventHandlers.Append(handler, slice.GetValue(priority, 0))
	log.Info("Server", log.String("RegEvent", runtimes.CurrentRunningFuncName()), log.String("handler", reflect.TypeOf(handler).String()))
}

func (slf *event) OnStopEvent() {
	slf.stopEventHandlers.RangeValue(func(index int, value StopEventHandler) bool {
		value(slf.Server)
		return true
	})
}

// RegConsoleCommandEvent 控制台收到指令时将立即执行被注册的事件处理函数
//   - 默认将注册 "exit", "quit", "close", "shutdown", "EXIT", "QUIT", "CLOSE", "SHUTDOWN" 指令作为关闭服务器的指令
//   - 可通过注册默认指令进行默认行为的覆盖
func (slf *event) RegConsoleCommandEvent(command string, handler ConsoleCommandEventHandler, priority ...int) {
	fd := int(os.Stdin.Fd())
	if !terminal.IsTerminal(fd) {
		log.Info("Server", log.String("RegEvent", runtimes.CurrentRunningFuncName()), log.String("ignore", "system not terminal"))
		return
	}

	slf.consoleCommandEventHandlerInitOnce.Do(func() {
		slf.consoleCommandEventHandlers = map[string]*slice.Priority[ConsoleCommandEventHandler]{}
		go func() {
			for {
				var input string
				_, _ = fmt.Scanln(&input)
				c2p := strings.SplitN(input, "?", 2)
				if len(c2p) == 1 {
					c2p = append(c2p, "")
				}
				slf.OnConsoleCommandEvent(c2p[0], c2p[1])
			}
		}()
	})
	list, exist := slf.consoleCommandEventHandlers[command]
	if !exist {
		list = slice.NewPriority[ConsoleCommandEventHandler]()
		slf.consoleCommandEventHandlers[command] = list
	}
	list.Append(handler, slice.GetValue(priority, 0))
	log.Info("Server", log.String("RegEvent", runtimes.CurrentRunningFuncName()), log.String("handler", reflect.TypeOf(handler).String()))
}

func (slf *event) OnConsoleCommandEvent(command string, paramsStr string) {
	slf.PushSystemMessage(func() {
		handles, exist := slf.consoleCommandEventHandlers[command]
		if !exist {
			switch command {
			case "exit", "quit", "close", "shutdown", "EXIT", "QUIT", "CLOSE", "SHUTDOWN":
				log.Info("Console", log.String("Receive", command), log.String("Action", "Shutdown"))
				slf.Server.shutdown(nil)
				return
			}
			log.Warn("Server", log.String("Command", "unregistered"))
		} else {
			v, err := url.ParseQuery(paramsStr)
			if err != nil {
				log.Error("ConsoleCommandEvent", log.String("command", command), log.String("params", paramsStr), log.Err(err))
				return
			}
			var params = make(ConsoleParams)
			for key, value := range v {
				params[key] = value
			}
			handles.RangeValue(func(index int, value ConsoleCommandEventHandler) bool {
				value(slf.Server, command, params)
				return true
			})
		}
	}, log.String("Event", "OnConsoleCommandEvent"))
}

// RegStartBeforeEvent 在服务器初始化完成启动前立刻执行被注册的事件处理函数
func (slf *event) RegStartBeforeEvent(handler StartBeforeEventHandler, priority ...int) {
	slf.startBeforeEventHandlers.Append(handler, slice.GetValue(priority, 0))
	log.Info("Server", log.String("RegEvent", runtimes.CurrentRunningFuncName()), log.String("handler", reflect.TypeOf(handler).String()))
}

func (slf *event) OnStartBeforeEvent() {
	defer func() {
		if err := recover(); err != nil {
			log.Error("Server", log.String("OnStartBeforeEvent", fmt.Sprintf("%v", err)))
			debug.PrintStack()
		}
	}()
	slf.startBeforeEventHandlers.RangeValue(func(index int, value StartBeforeEventHandler) bool {
		value(slf.Server)
		return true
	})
}

// RegStartFinishEvent 在服务器启动完成时将立刻执行被注册的事件处理函数
//   - 需要注意该时刻服务器已经启动完成，但是还有可能未开始处理消息，客户端有可能无法连接，如果需要在消息处理器准备就绪后执行，请使用 RegMessageReadyEvent 函数
func (slf *event) RegStartFinishEvent(handler StartFinishEventHandler, priority ...int) {
	slf.startFinishEventHandlers.Append(handler, slice.GetValue(priority, 0))
	log.Info("Server", log.String("RegEvent", runtimes.CurrentRunningFuncName()), log.String("handler", reflect.TypeOf(handler).String()))
}

func (slf *event) OnStartFinishEvent() {
	slf.PushSystemMessage(func() {
		slf.startFinishEventHandlers.RangeValue(func(index int, value StartFinishEventHandler) bool {
			value(slf.Server)
			return true
		})
	}, log.String("Event", "OnStartFinishEvent"))
	if slf.Server.limitLife > 0 {
		go func() {
			time.Sleep(slf.Server.limitLife)
			slf.Shutdown()
		}()
	}
}

// RegConnectionClosedEvent 在连接关闭后将立刻执行被注册的事件处理函数
func (slf *event) RegConnectionClosedEvent(handler ConnectionClosedEventHandler, priority ...int) {
	if slf.network == NetworkHttp {
		panic(ErrNetworkIncompatibleHttp)
	}
	slf.connectionClosedEventHandlers.Append(handler, slice.GetValue(priority, 0))
	log.Info("Server", log.String("RegEvent", runtimes.CurrentRunningFuncName()), log.String("handler", reflect.TypeOf(handler).String()))
}

func (slf *event) OnConnectionClosedEvent(conn *Conn, err any) {
	slf.PushSystemMessage(func() {
		slf.Server.online.Delete(conn.GetID())
		slf.connectionClosedEventHandlers.RangeValue(func(index int, value ConnectionClosedEventHandler) bool {
			value(slf.Server, conn, err)
			return true
		})
	}, log.String("Event", "OnConnectionClosedEvent"))
}

// RegConnectionOpenedEvent 在连接打开后将立刻执行被注册的事件处理函数
//   - 该阶段的事件将会在系统消息中进行处理，不适合处理耗时操作
func (slf *event) RegConnectionOpenedEvent(handler ConnectionOpenedEventHandler, priority ...int) {
	if slf.network == NetworkHttp {
		panic(ErrNetworkIncompatibleHttp)
	}
	slf.connectionOpenedEventHandlers.Append(handler, slice.GetValue(priority, 0))
	log.Info("Server", log.String("RegEvent", runtimes.CurrentRunningFuncName()), log.String("handler", reflect.TypeOf(handler).String()))
}

func (slf *event) OnConnectionOpenedEvent(conn *Conn) {
	slf.PushSystemMessage(func() {
		slf.Server.online.Set(conn.GetID(), conn)
		slf.connectionOpenedEventHandlers.RangeValue(func(index int, value ConnectionOpenedEventHandler) bool {
			value(slf.Server, conn)
			return true
		})
		slf.OnConnectionOpenedAfterEvent(conn)
	}, log.String("Event", "OnConnectionOpenedEvent"))
}

// RegConnectionReceivePacketEvent 在接收到数据包时将立刻执行被注册的事件处理函数
func (slf *event) RegConnectionReceivePacketEvent(handler ConnectionReceivePacketEventHandler, priority ...int) {
	if slf.network == NetworkHttp {
		panic(ErrNetworkIncompatibleHttp)
	}
	slf.connectionReceivePacketEventHandlers.Append(handler, slice.GetValue(priority, 0))
	log.Info("Server", log.String("RegEvent", runtimes.CurrentRunningFuncName()), log.String("handler", reflect.TypeOf(handler).String()))
}

func (slf *event) OnConnectionReceivePacketEvent(conn *Conn, packet []byte) {
	if slf.Server.runtime.packetWarnSize > 0 && len(packet) > slf.Server.runtime.packetWarnSize {
		log.Warn("Server", log.String("OnConnectionReceivePacketEvent", fmt.Sprintf("packet size %d > %d", len(packet), slf.Server.runtime.packetWarnSize)))
	}
	slf.connectionReceivePacketEventHandlers.RangeValue(func(index int, value ConnectionReceivePacketEventHandler) bool {
		value(slf.Server, conn, packet)
		return true
	})
}

// RegMessageErrorEvent 在处理消息发生错误时将立即执行被注册的事件处理函数
func (slf *event) RegMessageErrorEvent(handler MessageErrorEventHandler, priority ...int) {
	slf.messageErrorEventHandlers.Append(handler, slice.GetValue(priority, 0))
	log.Info("Server", log.String("RegEvent", runtimes.CurrentRunningFuncName()), log.String("handler", reflect.TypeOf(handler).String()))
}

func (slf *event) OnMessageErrorEvent(message *Message, err error) {
	if slf.messageErrorEventHandlers.Len() == 0 {
		return
	}
	defer func() {
		if err := recover(); err != nil {
			log.Error("Server", log.String("OnMessageErrorEvent", messageNames[message.t]), log.Any("Error", err))
			debug.PrintStack()
		}
	}()
	slf.messageErrorEventHandlers.RangeValue(func(index int, value MessageErrorEventHandler) bool {
		value(slf.Server, message, err)
		return true
	})
}

// RegMessageLowExecEvent 在处理消息缓慢时将立即执行被注册的事件处理函数
func (slf *event) RegMessageLowExecEvent(handler MessageLowExecEventHandler, priority ...int) {
	slf.messageLowExecEventHandlers.Append(handler, slice.GetValue(priority, 0))
	log.Info("Server", log.String("RegEvent", runtimes.CurrentRunningFuncName()), log.String("handler", reflect.TypeOf(handler).String()))
}

func (slf *event) OnMessageLowExecEvent(message *Message, cost time.Duration) {
	if slf.messageLowExecEventHandlers.Len() == 0 {
		return
	}
	// 慢消息不再占用消息通道
	slf.messageLowExecEventHandlers.RangeValue(func(index int, value MessageLowExecEventHandler) bool {
		value(slf.Server, message, cost)
		return true
	})
}

// RegConnectionOpenedAfterEvent 在连接打开事件处理完成后将立刻执行被注册的事件处理函数
//   - 该阶段事件将会转到对应消息分流渠道中进行处理
func (slf *event) RegConnectionOpenedAfterEvent(handler ConnectionOpenedAfterEventHandler, priority ...int) {
	if slf.network == NetworkHttp {
		panic(ErrNetworkIncompatibleHttp)
	}
	slf.connectionOpenedAfterEventHandlers.Append(handler, slice.GetValue(priority, 0))
	log.Info("Server", log.String("RegEvent", runtimes.CurrentRunningFuncName()), log.String("handler", reflect.TypeOf(handler).String()))
}

func (slf *event) OnConnectionOpenedAfterEvent(conn *Conn) {
	slf.PushShuntMessage(conn, func() {
		slf.connectionOpenedAfterEventHandlers.RangeValue(func(index int, value ConnectionOpenedAfterEventHandler) bool {
			value(slf.Server, conn)
			return true
		})
	}, log.String("Event", "OnConnectionOpenedAfterEvent"))
}

// RegConnectionWritePacketBeforeEvent 在发送数据包前将立刻执行被注册的事件处理函数
func (slf *event) RegConnectionWritePacketBeforeEvent(handler ConnectionWritePacketBeforeEventHandler, priority ...int) {
	if slf.network == NetworkHttp {
		panic(ErrNetworkIncompatibleHttp)
	}
	slf.connectionWritePacketBeforeHandlers.Append(handler, slice.GetValue(priority, 0))
	log.Info("Server", log.String("RegEvent", runtimes.CurrentRunningFuncName()), log.String("handler", reflect.TypeOf(handler).String()))
}

func (slf *event) OnConnectionWritePacketBeforeEvent(conn *Conn, packet []byte) (newPacket []byte) {
	if slf.connectionWritePacketBeforeHandlers.Len() == 0 {
		return packet
	}
	newPacket = packet
	slf.connectionWritePacketBeforeHandlers.RangeValue(func(index int, value ConnectionWritePacketBeforeEventHandler) bool {
		newPacket = value(slf.Server, conn, newPacket)
		return true
	})
	return newPacket
}

// RegShuntChannelCreatedEvent 在分流通道创建时将立刻执行被注册的事件处理函数
func (slf *event) RegShuntChannelCreatedEvent(handler ShuntChannelCreatedEventHandler, priority ...int) {
	slf.shuntChannelCreatedEventHandlers.Append(handler, slice.GetValue(priority, 0))
	log.Info("Server", log.String("RegEvent", runtimes.CurrentRunningFuncName()), log.String("handler", reflect.TypeOf(handler).String()))
}

func (slf *event) OnShuntChannelCreatedEvent(guid int64) {
	slf.PushSystemMessage(func() {
		slf.shuntChannelCreatedEventHandlers.RangeValue(func(index int, value ShuntChannelCreatedEventHandler) bool {
			value(slf.Server, guid)
			return true
		})
	}, log.String("Event", "OnShuntChannelCreatedEvent"))
}

// RegShuntChannelCloseEvent 在分流通道关闭时将立刻执行被注册的事件处理函数
func (slf *event) RegShuntChannelCloseEvent(handler ShuntChannelClosedEventHandler, priority ...int) {
	slf.shuntChannelClosedEventHandlers.Append(handler, slice.GetValue(priority, 0))
	log.Info("Server", log.String("RegEvent", runtimes.CurrentRunningFuncName()), log.String("handler", reflect.TypeOf(handler).String()))
}

func (slf *event) OnShuntChannelClosedEvent(guid int64) {
	slf.PushSystemMessage(func() {
		slf.shuntChannelClosedEventHandlers.RangeValue(func(index int, value ShuntChannelClosedEventHandler) bool {
			value(slf.Server, guid)
			return true
		})
	}, log.String("Event", "OnShuntChannelClosedEvent"))
}

// RegConnectionPacketPreprocessEvent 在接收到数据包后将立刻执行被注册的事件处理函数
//   - 预处理函数可以用于对数据包进行预处理，如解密、解压缩等
//   - 在调用 abort() 后，将不会再调用后续的预处理函数，也不会调用 OnConnectionReceivePacketEvent 函数
//   - 在调用 usePacket() 后，将使用新的数据包，而不会使用原始数据包，同时阻止后续的预处理函数的调用
//
// 场景：
//   - 数据包格式校验
//   - 数据包分包等情况处理
func (slf *event) RegConnectionPacketPreprocessEvent(handler ConnectionPacketPreprocessEventHandler, priority ...int) {
	slf.connectionPacketPreprocessEventHandlers.Append(handler, slice.GetValue(priority, 0))
	log.Info("Server", log.String("RegEvent", runtimes.CurrentRunningFuncName()), log.String("handler", reflect.TypeOf(handler).String()))
}

func (slf *event) OnConnectionPacketPreprocessEvent(conn *Conn, packet []byte, usePacket func(newPacket []byte)) bool {
	if slf.connectionPacketPreprocessEventHandlers.Len() == 0 {
		return false
	}
	var abort = false
	slf.connectionPacketPreprocessEventHandlers.RangeValue(func(index int, value ConnectionPacketPreprocessEventHandler) bool {
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
func (slf *event) RegMessageExecBeforeEvent(handler MessageExecBeforeEventHandler, priority ...int) {
	slf.messageExecBeforeEventHandlers.Append(handler, slice.GetValue(priority, 0))
	log.Info("Server", log.String("RegEvent", runtimes.CurrentRunningFuncName()), log.String("handler", reflect.TypeOf(handler).String()))
}

func (slf *event) OnMessageExecBeforeEvent(message *Message) bool {
	if slf.messageExecBeforeEventHandlers.Len() == 0 {
		return true
	}
	var result = true
	defer func() {
		if err := recover(); err != nil {
			log.Error("Server", log.String("OnMessageExecBeforeEvent", fmt.Sprintf("%v", err)))
			debug.PrintStack()
		}
	}()
	slf.messageExecBeforeEventHandlers.RangeValue(func(index int, value MessageExecBeforeEventHandler) bool {
		result = value(slf.Server, message)
		return result
	})
	return result
}

// RegMessageReadyEvent 在服务器消息处理器准备就绪时立即执行被注册的事件处理函数
func (slf *event) RegMessageReadyEvent(handler MessageReadyEventHandler, priority ...int) {
	slf.messageReadyEventHandlers.Append(handler, slice.GetValue(priority, 0))
}

func (slf *event) OnMessageReadyEvent() {
	if slf.messageReadyEventHandlers.Len() == 0 {
		return
	}
	defer func() {
		if err := recover(); err != nil {
			log.Error("Server", log.String("OnMessageReadyEvent", fmt.Sprintf("%v", err)))
			debug.PrintStack()
		}
	}()
	slf.messageReadyEventHandlers.RangeValue(func(index int, value MessageReadyEventHandler) bool {
		value(slf.Server)
		return true
	})
}

// RegDeadlockDetectEvent 在死锁检测触发时立即执行被注册的事件处理函数
func (slf *event) RegDeadlockDetectEvent(handler OnDeadlockDetectEventHandler, priority ...int) {
	slf.deadlockDetectEventHandlers.Append(handler, slice.GetValue(priority, 0))
}

func (slf *event) OnDeadlockDetectEvent(message *Message) {
	if slf.deadlockDetectEventHandlers.Len() == 0 {
		return
	}
	defer func() {
		if err := recover(); err != nil {
			log.Error("Server", log.String("OnDeadlockDetectEvent", fmt.Sprintf("%v", err)))
			debug.PrintStack()
		}
	}()
	slf.deadlockDetectEventHandlers.RangeValue(func(index int, value OnDeadlockDetectEventHandler) bool {
		value(slf.Server, message)
		return true
	})
}

func (slf *event) check() {
	switch slf.network {
	case NetworkHttp, NetworkGRPC, NetworkNone:
	default:
		if slf.connectionReceivePacketEventHandlers.Len() == 0 {
			log.Warn("Server", log.String("ConnectionReceivePacketEvent", "invalid server, no packets processed"))
		}
	}
}
