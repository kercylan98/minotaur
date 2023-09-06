package server

import (
	"context"
	"encoding/json"
	"fmt"
	"github.com/kercylan98/minotaur/utils/hash"
	"reflect"
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

	// MessageTypeAsyncCallback 异步回调消息类型
	MessageTypeAsyncCallback

	// MessageTypeSystem 系统消息类型
	MessageTypeSystem
)

var messageNames = map[MessageType]string{
	MessageTypePacket:        "MessageTypePacket",
	MessageTypeError:         "MessageTypeError",
	MessageTypeCross:         "MessageTypeCross",
	MessageTypeTicker:        "MessageTypeTicker",
	MessageTypeAsync:         "MessageTypeAsync",
	MessageTypeAsyncCallback: "MessageTypeAsyncCallback",
	MessageTypeSystem:        "MessageTypeSystem",
}

const (
	MessageErrorActionNone     MessageErrorAction = iota + 1 // 错误消息类型操作：将不会被进行任何特殊处理，仅进行日志输出
	MessageErrorActionShutdown                               // 错误消息类型操作：当接收到该类型的操作时，服务器将执行 Server.shutdown 函数
)

var messageErrorActionNames = map[MessageErrorAction]string{
	MessageErrorActionNone:     "None",
	MessageErrorActionShutdown: "Shutdown",
}

var (
	messagePacketVisualization = func(packet []byte) string {
		return string(packet)
	}
)

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
	t     MessageType // 消息类型
	attrs []any       // 消息属性
}

// MessageType 返回消息类型
func (slf *Message) MessageType() MessageType {
	return slf.t
}

// String 返回消息的字符串表示
func (slf *Message) String() string {
	var attrs = make([]any, 0, len(slf.attrs))
	for _, attr := range slf.attrs {
		if reflect.TypeOf(attr).Kind() == reflect.Func {
			continue
		}
		attrs = append(attrs, attr)
	}
	var s string
	switch slf.t {
	case MessageTypePacket:
		if len(attrs) > 1 {
			s = messagePacketVisualization(attrs[1].([]byte))
		}
	default:
		if len(slf.attrs) == 0 {
			s = "NoneAttr"
		} else {
			raw, _ := json.Marshal(attrs)
			s = string(raw)
		}
	}

	return fmt.Sprintf("[%s] %s", slf.t, s)
}

// String 返回消息类型的字符串表示
func (slf MessageType) String() string {
	return messageNames[slf]
}

// GetPacketMessageAttrs 获取消息中的数据包属性
func (slf *Message) GetPacketMessageAttrs() (conn *Conn, packet []byte) {
	conn = slf.attrs[0].(*Conn)
	packet = slf.attrs[1].([]byte)
	return
}

// PushPacketMessage 向特定服务器中推送 MessageTypePacket 消息
func PushPacketMessage(srv *Server, conn *Conn, wst int, packet []byte, mark ...any) {
	msg := srv.messagePool.Get()
	msg.t = MessageTypePacket
	msg.attrs = append([]any{&Conn{ctx: context.WithValue(conn.ctx, contextKeyWST, wst), connection: conn.connection}, packet}, mark...)
	srv.pushMessage(msg)
}

// GetErrorMessageAttrs 获取消息中的错误属性
func (slf *Message) GetErrorMessageAttrs() (err error, action MessageErrorAction) {
	err = slf.attrs[0].(error)
	action = slf.attrs[1].(MessageErrorAction)
	return
}

// PushErrorMessage 向特定服务器中推送 MessageTypeError 消息
func PushErrorMessage(srv *Server, err error, action MessageErrorAction, mark ...any) {
	msg := srv.messagePool.Get()
	msg.t = MessageTypeError
	msg.attrs = append([]any{err, action}, mark...)
	srv.pushMessage(msg)
}

// GetCrossMessageAttrs 获取消息中的跨服属性
func (slf *Message) GetCrossMessageAttrs() (serverId int64, packet []byte) {
	serverId = slf.attrs[0].(int64)
	packet = slf.attrs[1].([]byte)
	return
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

// GetTickerMessageAttrs 获取消息中的定时器属性
func (slf *Message) GetTickerMessageAttrs() (caller func()) {
	caller = slf.attrs[0].(func())
	return
}

// PushTickerMessage 向特定服务器中推送 MessageTypeTicker 消息
func PushTickerMessage(srv *Server, caller func(), mark ...any) {
	msg := srv.messagePool.Get()
	msg.t = MessageTypeTicker
	msg.attrs = append([]any{caller}, mark...)
	srv.pushMessage(msg)
}

// GetAsyncMessageAttrs 获取消息中的异步消息属性
func (slf *Message) GetAsyncMessageAttrs() (caller func() error, callback func(err error), hasCallback bool) {
	caller = slf.attrs[0].(func() error)
	callback, hasCallback = slf.attrs[1].(func(err error))
	return
}

// PushAsyncMessage 向特定服务器中推送 MessageTypeAsync 消息
//   - 异步消息将在服务器的异步消息队列中进行处理，处理完成 caller 的阻塞操作后，将会通过系统消息执行 callback 函数
//   - callback 函数将在异步消息处理完成后进行调用，无论过程是否产生 err，都将被执行，允许为 nil
//   - 需要注意的是，为了避免并发问题，caller 函数请仅处理阻塞操作，其他操作应该在 callback 函数中进行
//
// 在通过 WithShunt 使用分流服务器时，异步消息不会转换到分流通道中进行处理。依旧需要注意上方第三条
func PushAsyncMessage(srv *Server, caller func() error, callback func(err error), mark ...any) {
	msg := srv.messagePool.Get()
	msg.t = MessageTypeAsync
	msg.attrs = append([]any{caller, callback}, mark...)
	srv.pushMessage(msg)
}

// GetSystemMessageAttrs 获取消息中的系统消息属性
func (slf *Message) GetSystemMessageAttrs() (handle func()) {
	handle = slf.attrs[0].(func())
	return
}

// PushSystemMessage 向特定服务器中推送 MessageTypeSystem 消息
func PushSystemMessage(srv *Server, handle func(), mark ...any) {
	msg := srv.messagePool.Get()
	msg.t = MessageTypeSystem
	msg.attrs = append([]any{handle}, mark...)
	srv.pushMessage(msg)
}

// SetMessagePacketVisualizer 设置消息可视化函数
//   - 消息可视化将在慢消息等情况用于打印，使用自定消息可视化函数可以便于开发者进行调试
//   - 默认的消息可视化函数将直接返回消息的字符串表示
func SetMessagePacketVisualizer(handle func(packet []byte) string) {
	messagePacketVisualization = handle
}
