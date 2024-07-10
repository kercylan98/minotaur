package main

import (
	"fmt"
	"github.com/kercylan98/minotaur/core/transport"
	"github.com/kercylan98/minotaur/core/vivid"
	"github.com/kercylan98/minotaur/toolkit"
	"github.com/kercylan98/minotaur/toolkit/router"
)

type Reader func(dst any)

type Message struct {
	Code int
	Data []byte
}

type WebSocketService struct {
	router *router.Multistage[func(conn *transport.Conn, reader Reader)]
}

func (f *WebSocketService) OnInit(kit *transport.FiberKit) {
	f.router = router.NewMultistage[func(conn *transport.Conn, reader Reader)]()

	kit.WebSocket("/ws").ConnectionPacketHook(f.onConnectionPacket)

	f.router.Route(1, f.onPrint)
	f.router.Sub(1).Route(1, f.onPrint)
	f.router.Register(1, 2).Bind(f.onPrint)
}

func (f *WebSocketService) onConnectionPacket(kit *transport.FiberKit, ctx *transport.FiberContext, conn *transport.Conn, packet transport.Packet) error {
	var msg Message
	toolkit.UnmarshalJSON(packet.GetBytes(), &msg)

	handler := f.router.Match(msg.Code)
	if handler != nil {
		handler(conn, func(dst any) {
			toolkit.UnmarshalJSON(msg.Data, dst)
		})
	}

	return nil
}

func (f *WebSocketService) onPrint(conn *transport.Conn, reader Reader) {
	var str string
	reader(&str)
	fmt.Println(str)
}

func main() {
	vivid.NewActorSystem(func(options *vivid.ActorSystemOptions) {
		options.WithModule(transport.NewFiber(":8877").BindService(new(WebSocketService)))
	})
}
