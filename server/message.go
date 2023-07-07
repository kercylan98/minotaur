package server

import "runtime/debug"

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
}

const (
	MessageErrorActionNone     MessageErrorAction = iota // 错误消息类型操作：将不会被进行任何特殊处理，仅进行日志输出
	MessageErrorActionShutdown                           // 错误消息类型操作：当接收到该类型的操作时，服务器将执行 Server.Shutdown 函数
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

func (slf MessageType) String() string {
	return messageNames[slf]
}

func (slf MessageType) deconstructWebSocketPacket(attrs ...any) (conn *Conn, packet []byte, messageType int) {
	if len(attrs) != 3 {
		panic(ErrWebsocketMessageTypePacketAttrs)
	}
	var ok bool
	if conn, ok = attrs[0].(*Conn); !ok {
		panic(ErrWebsocketMessageTypePacketAttrs)
	}
	if packet, ok = attrs[1].([]byte); !ok {
		panic(ErrWebsocketMessageTypePacketAttrs)
	}
	if messageType, ok = attrs[2].(int); !ok {
		panic(ErrWebsocketMessageTypePacketAttrs)
	}
	return
}

func (slf MessageType) deconstructPacket(attrs ...any) (conn *Conn, packet []byte) {
	if len(attrs) != 2 {
		panic(ErrMessageTypePacketAttrs)
	}
	var ok bool
	if conn, ok = attrs[0].(*Conn); !ok {
		panic(ErrMessageTypePacketAttrs)
	}
	if packet, ok = attrs[1].([]byte); !ok {
		panic(ErrMessageTypePacketAttrs)
	}
	return
}

func (slf MessageType) deconstructError(attrs ...any) (err error, action MessageErrorAction, stack string) {
	if len(attrs) != 3 {
		panic(ErrMessageTypeErrorAttrs)
	}
	var ok bool
	if err, ok = attrs[0].(error); !ok {
		panic(ErrMessageTypeErrorAttrs)
	}
	if action, ok = attrs[1].(MessageErrorAction); !ok {
		panic(ErrMessageTypeErrorAttrs)
	}
	stack = attrs[2].(string)
	return
}

func (slf MessageType) deconstructCross(attrs ...any) (serverId int64, packet []byte) {
	if len(attrs) != 2 {
		panic(ErrMessageTypeCrossErrorAttrs)
	}
	var ok bool
	if serverId, ok = attrs[0].(int64); !ok {
		panic(ErrMessageTypeCrossErrorAttrs)
	}
	if packet, ok = attrs[1].([]byte); !ok {
		panic(ErrMessageTypeCrossErrorAttrs)
	}
	return
}

func (slf MessageType) deconstructTicker(attrs ...any) (caller func()) {
	if len(attrs) != 1 {
		panic(ErrMessageTypeTickerErrorAttrs)
	}
	var ok bool
	if caller, ok = attrs[0].(func()); !ok {
		panic(ErrMessageTypeTickerErrorAttrs)
	}
	return
}

// PushWebsocketPacketMessage 向特定服务器中推送 WebsocketPacket 消息
func PushWebsocketPacketMessage(srv *Server, conn *Conn, packet []byte, messageType int) {
	msg := srv.messagePool.Get()
	msg.t = MessageTypePacket
	msg.attrs = []any{conn, packet, messageType}
	srv.pushMessage(msg)
}

// PushPacketMessage 向特定服务器中推送 Packet 消息
func PushPacketMessage(srv *Server, conn *Conn, packet []byte) {
	msg := srv.messagePool.Get()
	msg.t = MessageTypePacket
	msg.attrs = []any{conn, packet}
	srv.pushMessage(msg)
}

// PushErrorMessage 向特定服务器中推送 Error 消息
func PushErrorMessage(srv *Server, err error, action MessageErrorAction) {
	msg := srv.messagePool.Get()
	msg.t = MessageTypeError
	msg.attrs = []any{err, action, string(debug.Stack())}
	srv.pushMessage(msg)
}

// PushCrossMessage 向特定服务器中推送 Cross 消息
func PushCrossMessage(srv *Server, crossName string, serverId int64, packet []byte) {
	if len(srv.cross) == 0 {
		return
	}
	_, exist := srv.cross[crossName]
	if !exist {
		return
	}
	msg := srv.messagePool.Get()
	msg.t = MessageTypeCross
	msg.attrs = []any{serverId, packet}
	srv.pushMessage(msg)
}
