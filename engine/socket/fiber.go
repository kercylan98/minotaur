package socket

// fiberSocketV2 是一个接口，它定义了 Fiber 的 WebSocket 实现
type fiberSocketV2 interface {
	ReadMessage() (messageType int, p []byte, err error)
	WriteMessage(messageType int, data []byte) error
	Close() error
}

// ProduceFiberSocketV2 创建一个基于 Fiber 的 Socket，它使用 FiberSocket 作为 Fiber WebSocket 实现的接口，并使用 factory 来创建 Socket
//   - 该函数仅支持基于 fiber v2 实现的 "github.com/gofiber/contrib/websocket" 生成的 WebSocket 连接
func ProduceFiberSocketV2(factory Factory, socket fiberSocketV2, actor Actor) {
	c := factory.Produce(actor, func(packet []byte, ctx any) error {
		return socket.WriteMessage(ctx.(int), packet)
	}, func() error {
		if err := socket.WriteMessage(8, nil); err != nil {
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
