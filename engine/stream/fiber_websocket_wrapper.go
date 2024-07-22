package stream

import "github.com/gofiber/contrib/websocket"

type fiberWebSocketWrapper struct {
	*websocket.Conn
}

func (f *fiberWebSocketWrapper) Write(packet *Packet) error {
	return f.WriteMessage(packet.Context().(int), packet.Data())
}
