package stream

import (
	"github.com/gofiber/contrib/websocket"
	"github.com/kercylan98/minotaur/engine/vivid"
)

type fiberWebSocketWrapper struct {
	*websocket.Conn
	system    *vivid.ActorSystem
	streamRef vivid.ActorRef
	closed    chan struct{}
}

func (f *fiberWebSocketWrapper) OnWriterCreated(writer Writer) {
	f.system.Tell(f.streamRef, f.Conn)
}

func (f *fiberWebSocketWrapper) Write(packet *Packet) error {
	return f.Conn.WriteMessage(packet.Context().(int), packet.Data())
}

func (f *fiberWebSocketWrapper) Close() error {
	return f.Conn.WriteMessage(websocket.CloseMessage, nil)
}
