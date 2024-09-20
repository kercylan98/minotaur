package socket

// gorillaSocket 是一个接口，它定义了 github.com/gorilla/websocket 的 WebSocket 实现
type gorillaSocket interface {
	ReadMessage() (messageType int, p []byte, err error)
	WriteMessage(messageType int, data []byte) error
	Close() error
}

// ProduceGorillaSocket 创建一个基于 github.com/gorilla/websocket 的 Socket
func ProduceGorillaSocket(factory Factory, socket gorillaSocket, actor Actor) {
	c := factory.Produce(actor, func(packet []byte, ctx any) error {
		return socket.WriteMessage(ctx.(int), packet)
	}, func() error {
		if err := socket.WriteMessage(internalWebSocketCloseMessageType, nil); err != nil {
			return err
		}
		return socket.Close()
	})

	defer c.Close()
	for {
		typ, bytes, err := socket.ReadMessage()
		if err != nil {
			c.Close(err)
			break
		}

		c.React(bytes, typ)
	}
}
