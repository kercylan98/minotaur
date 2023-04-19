package server

const (
	MessageTypePacket MessageType = iota
	MessageTypeError
)

var messageNames = map[MessageType]string{
	MessageTypePacket: "MessageTypePacket",
	MessageTypeError:  "MessageTypeError",
}

const (
	MessageErrorActionNone MessageErrorAction = iota
	MessageErrorActionShutdown
)

var messageErrorActionNames = map[MessageErrorAction]string{
	MessageErrorActionNone:     "None",
	MessageErrorActionShutdown: "Shutdown",
}

type (
	MessageType        byte
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
