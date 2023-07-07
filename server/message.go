package server

import (
	"runtime/debug"
)

const (
	// MessageTypePacket 数据包消息类型：该类型的数据将被发送到 ConnectionReceivePacketEvent 进行处理
	MessageTypePacket MessageType = iota

	// MessageTypeError 错误消息类型：根据不同的错误状态，将交由 Server 进行统一处理
	MessageTypeError

	// MessageTypeCross 跨服消息类型：将被推送到跨服的 Cross 实现中进行处理
	MessageTypeCross

	// MessageTypeTicker 定时器消息类型
	MessageTypeTicker

	// MessageTypeAsync 异步消息类型
	MessageTypeAsync
)

var messageNames = map[MessageType]string{
	MessageTypePacket: "MessageTypePacket",
	MessageTypeError:  "MessageTypeError",
	MessageTypeCross:  "MessageTypeCross",
	MessageTypeTicker: "MessageTypeTicker",
	MessageTypeAsync:  "MessageTypeAsync",
}

const (
	MessageErrorActionNone     MessageErrorAction = iota // 错误消息类型操作：将不会被进行任何特殊处理，仅进行日志输出
	MessageErrorActionShutdown                           // 错误消息类型操作：当接收到该类型的操作时，服务器将执行 Server.shutdown 函数
)

var messageErrorActionNames = map[MessageErrorAction]string{
	MessageErrorActionNone:     "None",
	MessageErrorActionShutdown: "shutdown",
}

type (
	// MessageType 消息类型
	MessageType byte

	// MessageErrorAction 错误消息类型操作
	MessageErrorAction byte
)

func (slf MessageErrorAction) String() string {
	return messageErrorActionNames[slf]
}

type Message struct {
	t     MessageType
	attrs []any
}

func (slf MessageType) String() string {
	return messageNames[slf]
}

// PushPacketMessage 向特定服务器中推送 MessageTypePacket 消息
func PushPacketMessage(srv *Server, conn *Conn, packet []byte) {
	msg := srv.messagePool.Get()
	msg.t = MessageTypePacket
	msg.attrs = []any{conn, packet}
	srv.pushMessage(msg)
}

// PushErrorMessage 向特定服务器中推送 MessageTypeError 消息
func PushErrorMessage(srv *Server, err error, action MessageErrorAction) {
	msg := srv.messagePool.Get()
	msg.t = MessageTypeError
	msg.attrs = []any{err, action, string(debug.Stack())}
	srv.pushMessage(msg)
}

// PushCrossMessage 向特定服务器中推送 MessageTypeCross 消息
func PushCrossMessage(srv *Server, crossName string, serverId int64, packet []byte) {
	if serverId == srv.id {
		msg := srv.messagePool.Get()
		msg.t = MessageTypeCross
		msg.attrs = []any{serverId, packet}
		srv.pushMessage(msg)
	} else {
		if len(srv.cross) == 0 {
			return
		}
		cross, exist := srv.cross[crossName]
		if !exist {
			return
		}
		_ = cross.PushMessage(serverId, packet)
	}
}

// PushTickerMessage 向特定服务器中推送 MessageTypeTicker 消息
func PushTickerMessage(srv *Server, caller func()) {
	msg := srv.messagePool.Get()
	msg.t = MessageTypeTicker
	msg.attrs = []any{caller}
	srv.pushMessage(msg)
}

// PushAsyncMessage 向特定服务器中推送 MessageTypeAsync 消息
func PushAsyncMessage(srv *Server, caller func() error, callback ...func(err error)) {
	msg := srv.messagePool.Get()
	msg.t = MessageTypeAsync
	msg.attrs = []any{caller, callback}
	srv.pushMessage(msg)
}
