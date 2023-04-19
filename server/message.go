package server

const (
	// MessageTypePacket 数据包消息类型：该类型的数据将被发送到 ConnectionReceivePacketEvent 进行处理
	//  - *server.Conn
	//  - []byte
	MessageTypePacket MessageType = iota

	// MessageTypeError 错误消息类型：根据不同的错误状态，将交由 Server 进行统一处理
	//  - error
	//  - server.MessageErrorAction
	MessageTypeError
)

var messageNames = map[MessageType]string{
	MessageTypePacket: "MessageTypePacket",
	MessageTypeError:  "MessageTypeError",
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

type message struct {
	t     MessageType
	attrs []any
}

func (slf MessageType) String() string {
	return messageNames[slf]
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

func (slf MessageType) deconstructError(attrs ...any) (err error, action MessageErrorAction) {
	if len(attrs) != 2 {
		panic(ErrMessageTypeErrorAttrs)
	}
	var ok bool
	if err, ok = attrs[0].(error); !ok {
		panic(ErrMessageTypeErrorAttrs)
	}
	if action, ok = attrs[1].(MessageErrorAction); !ok {
		panic(ErrMessageTypeErrorAttrs)
	}
	return
}
