package server

import (
	"github.com/kercylan98/minotaur/utils/hash"
	"github.com/kercylan98/minotaur/utils/log"
	"github.com/kercylan98/minotaur/utils/super"
)

const (
	// MessageTypePacket 数据包消息类型：该类型的数据将被发送到 ConnectionReceivePacketEvent 进行处理
	MessageTypePacket MessageType = iota + 1

	// MessageTypeError 错误消息类型：根据不同的错误状态，将交由 Server 进行统一处理
	MessageTypeError

	// MessageTypeTicker 定时器消息类型
	MessageTypeTicker

	// MessageTypeShuntTicker 分流定时器消息类型
	MessageTypeShuntTicker

	// MessageTypeAsync 异步消息类型
	MessageTypeAsync

	// MessageTypeAsyncCallback 异步回调消息类型
	MessageTypeAsyncCallback

	// MessageTypeShuntAsync 分流异步消息类型
	MessageTypeShuntAsync

	// MessageTypeShuntAsyncCallback 分流异步回调消息类型
	MessageTypeShuntAsyncCallback

	// MessageTypeUniqueAsync 唯一异步消息类型
	MessageTypeUniqueAsync

	// MessageTypeUniqueAsyncCallback 唯一异步回调消息类型
	MessageTypeUniqueAsyncCallback

	// MessageTypeUniqueShuntAsync 唯一分流异步消息类型
	MessageTypeUniqueShuntAsync

	// MessageTypeUniqueShuntAsyncCallback 唯一分流异步回调消息类型
	MessageTypeUniqueShuntAsyncCallback

	// MessageTypeSystem 系统消息类型
	MessageTypeSystem

	// MessageTypeShunt 普通分流消息类型
	MessageTypeShunt
)

var messageNames = map[MessageType]string{
	MessageTypePacket:                   "MessageTypePacket",
	MessageTypeError:                    "MessageTypeError",
	MessageTypeTicker:                   "MessageTypeTicker",
	MessageTypeShuntTicker:              "MessageTypeShuntTicker",
	MessageTypeAsync:                    "MessageTypeAsync",
	MessageTypeAsyncCallback:            "MessageTypeAsyncCallback",
	MessageTypeShuntAsync:               "MessageTypeShuntAsync",
	MessageTypeShuntAsyncCallback:       "MessageTypeShuntAsyncCallback",
	MessageTypeUniqueAsync:              "MessageTypeUniqueAsync",
	MessageTypeUniqueAsyncCallback:      "MessageTypeUniqueAsyncCallback",
	MessageTypeUniqueShuntAsync:         "MessageTypeUniqueShuntAsync",
	MessageTypeUniqueShuntAsyncCallback: "MessageTypeUniqueShuntAsyncCallback",
	MessageTypeSystem:                   "MessageTypeSystem",
	MessageTypeShunt:                    "MessageTypeShunt",
}

const (
	MessageErrorActionNone     MessageErrorAction = iota + 1 // 错误消息类型操作：将不会被进行任何特殊处理，仅进行日志输出
	MessageErrorActionShutdown                               // 错误消息类型操作：当接收到该类型的操作时，服务器将执行 Server.shutdown 函数
)

var messageErrorActionNames = map[MessageErrorAction]string{
	MessageErrorActionNone:     "None",
	MessageErrorActionShutdown: "Shutdown",
}

type (
	// MessageType 消息类型
	MessageType byte

	// MessageErrorAction 错误消息类型操作
	MessageErrorAction byte
)

// HasMessageType 检查是否存在指定的消息类型
func HasMessageType(mt MessageType) bool {
	return hash.Exist(messageNames, mt)
}

func (slf MessageErrorAction) String() string {
	return messageErrorActionNames[slf]
}

// Message 服务器消息
type Message struct {
	conn             *Conn
	ordinaryHandler  func()
	exceptionHandler func() error
	errHandler       func(err error)
	packet           []byte
	err              error
	name             string
	t                MessageType
	errAction        MessageErrorAction
	marks            []log.Field
}

// reset 重置消息结构体
func (slf *Message) reset() {
	slf.conn = nil
	slf.ordinaryHandler = nil
	slf.exceptionHandler = nil
	slf.errHandler = nil
	slf.packet = nil
	slf.err = nil
	slf.name = ""
	slf.t = 0
	slf.errAction = 0
	slf.marks = nil
}

// MessageType 返回消息类型
func (slf *Message) MessageType() MessageType {
	return slf.t
}

// String 返回消息的字符串表示
func (slf *Message) String() string {
	var info = struct {
		Type   string `json:"type,omitempty"`
		Name   string `json:"name,omitempty"`
		Packet string `json:"packet,omitempty"`
	}{
		Type:   slf.t.String(),
		Name:   slf.name,
		Packet: string(slf.packet),
	}

	return string(super.MarshalJSON(info))
}

// String 返回消息类型的字符串表示
func (slf MessageType) String() string {
	return messageNames[slf]
}

// castToPacketMessage 将消息转换为数据包消息
func (slf *Message) castToPacketMessage(conn *Conn, packet []byte, mark ...log.Field) *Message {
	slf.t, slf.conn, slf.packet, slf.marks = MessageTypePacket, conn, packet, mark
	return slf
}

// castToTickerMessage 将消息转换为定时器消息
func (slf *Message) castToTickerMessage(name string, caller func(), mark ...log.Field) *Message {
	slf.t, slf.name, slf.ordinaryHandler, slf.marks = MessageTypeTicker, name, caller, mark
	return slf
}

// castToShuntTickerMessage 将消息转换为分发器定时器消息
func (slf *Message) castToShuntTickerMessage(conn *Conn, name string, caller func(), mark ...log.Field) *Message {
	slf.t, slf.conn, slf.name, slf.ordinaryHandler, slf.marks = MessageTypeShuntTicker, slf.conn, name, caller, mark
	return slf
}

// castToAsyncMessage 将消息转换为异步消息
func (slf *Message) castToAsyncMessage(caller func() error, callback func(err error), mark ...log.Field) *Message {
	slf.t, slf.exceptionHandler, slf.errHandler, slf.marks = MessageTypeAsync, caller, callback, mark
	return slf
}

// castToAsyncCallbackMessage 将消息转换为异步回调消息
func (slf *Message) castToAsyncCallbackMessage(err error, caller func(err error), mark ...log.Field) *Message {
	slf.t, slf.err, slf.errHandler, slf.marks = MessageTypeAsyncCallback, err, caller, mark
	return slf
}

// castToShuntAsyncMessage 将消息转换为分流异步消息
func (slf *Message) castToShuntAsyncMessage(conn *Conn, caller func() error, callback func(err error), mark ...log.Field) *Message {
	slf.t, slf.conn, slf.exceptionHandler, slf.errHandler, slf.marks = MessageTypeShuntAsync, conn, caller, callback, mark
	return slf
}

// castToShuntAsyncCallbackMessage 将消息转换为分流异步回调消息
func (slf *Message) castToShuntAsyncCallbackMessage(conn *Conn, err error, caller func(err error), mark ...log.Field) *Message {
	slf.t, slf.conn, slf.err, slf.errHandler, slf.marks = MessageTypeShuntAsyncCallback, conn, err, caller, mark
	return slf
}

// castToUniqueAsyncMessage 将消息转换为唯一异步消息
func (slf *Message) castToUniqueAsyncMessage(unique string, caller func() error, callback func(err error), mark ...log.Field) *Message {
	slf.t, slf.name, slf.exceptionHandler, slf.errHandler, slf.marks = MessageTypeUniqueAsync, unique, caller, callback, mark
	return slf
}

// castToUniqueAsyncCallbackMessage 将消息转换为唯一异步回调消息
func (slf *Message) castToUniqueAsyncCallbackMessage(unique string, err error, caller func(err error), mark ...log.Field) *Message {
	slf.t, slf.name, slf.err, slf.errHandler, slf.marks = MessageTypeUniqueAsyncCallback, unique, err, caller, mark
	return slf
}

// castToUniqueShuntAsyncMessage 将消息转换为唯一分流异步消息
func (slf *Message) castToUniqueShuntAsyncMessage(conn *Conn, unique string, caller func() error, callback func(err error), mark ...log.Field) *Message {
	slf.t, slf.conn, slf.name, slf.exceptionHandler, slf.errHandler, slf.marks = MessageTypeUniqueShuntAsync, conn, unique, caller, callback, mark
	return slf
}

// castToUniqueShuntAsyncCallbackMessage 将消息转换为唯一分流异步回调消息
func (slf *Message) castToUniqueShuntAsyncCallbackMessage(conn *Conn, unique string, err error, caller func(err error), mark ...log.Field) *Message {
	slf.t, slf.conn, slf.name, slf.err, slf.errHandler, slf.marks = MessageTypeUniqueShuntAsyncCallback, conn, unique, err, caller, mark
	return slf
}

// castToSystemMessage 将消息转换为系统消息
func (slf *Message) castToSystemMessage(caller func(), mark ...log.Field) *Message {
	slf.t, slf.ordinaryHandler, slf.marks = MessageTypeSystem, caller, mark
	return slf
}

// castToErrorMessage 将消息转换为错误消息
func (slf *Message) castToErrorMessage(err error, action MessageErrorAction, mark ...log.Field) *Message {
	slf.t, slf.err, slf.errAction, slf.marks = MessageTypeError, err, action, mark
	return slf
}

// castToShuntMessage 将消息转换为分流消息
func (slf *Message) castToShuntMessage(conn *Conn, caller func(), mark ...log.Field) *Message {
	slf.t, slf.conn, slf.ordinaryHandler, slf.marks = MessageTypeShunt, conn, caller, mark
	return slf
}
