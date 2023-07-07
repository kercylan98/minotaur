package server

import (
	"encoding/json"
	"fmt"
	"github.com/kercylan98/minotaur/utils/str"
	"reflect"
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
	MessageErrorActionShutdown: "Shutdown",
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

func (slf *Message) String() string {
	var attrs = make([]any, 0, len(slf.attrs))
	for _, attr := range slf.attrs {
		if reflect.TypeOf(attr).Kind() == reflect.Func {
			continue
		}
		attrs = append(attrs, attr)
	}
	raw, _ := json.Marshal(attrs)
	s := string(raw)
	if s == str.None {
		s = "NoneAttr"
	}
	return fmt.Sprintf("[%s] %s", slf.t, s)
}

func (slf MessageType) String() string {
	return messageNames[slf]
}

// PushPacketMessage 向特定服务器中推送 MessageTypePacket 消息
func PushPacketMessage(srv *Server, conn *Conn, packet []byte, mark ...any) {
	msg := srv.messagePool.Get()
	msg.t = MessageTypePacket
	msg.attrs = append([]any{conn, packet}, mark...)
	srv.pushMessage(msg)
}

// PushErrorMessage 向特定服务器中推送 MessageTypeError 消息
func PushErrorMessage(srv *Server, err error, action MessageErrorAction, mark ...any) {
	msg := srv.messagePool.Get()
	msg.t = MessageTypeError
	msg.attrs = append([]any{err, action, string(debug.Stack())}, mark...)
	srv.pushMessage(msg)
}

// PushCrossMessage 向特定服务器中推送 MessageTypeCross 消息
func PushCrossMessage(srv *Server, crossName string, serverId int64, packet []byte, mark ...any) {
	if serverId == srv.id {
		msg := srv.messagePool.Get()
		msg.t = MessageTypeCross
		msg.attrs = append([]any{serverId, packet}, mark...)
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
func PushTickerMessage(srv *Server, caller func(), mark ...any) {
	msg := srv.messagePool.Get()
	msg.t = MessageTypeTicker
	msg.attrs = append([]any{caller}, mark...)
	srv.pushMessage(msg)
}

// PushAsyncMessage 向特定服务器中推送 MessageTypeAsync 消息
func PushAsyncMessage(srv *Server, caller func() error, callback func(err error), mark ...any) {
	msg := srv.messagePool.Get()
	msg.t = MessageTypeAsync
	msg.attrs = append([]any{caller, callback, string(debug.Stack())}, mark...)
	srv.pushMessage(msg)
}
