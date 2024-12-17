package controller

import (
	"github.com/kercylan98/minotaur/engine/socket"
	"github.com/kercylan98/minotaur/engine/vivid"
)

func newWebSocketActor() *webSocketActor {
	return &webSocketActor{}
}

type webSocketActor struct {
}

func (w *webSocketActor) OnOpened(ctx vivid.ActorContext, socket socket.Socket) {

}

func (w *webSocketActor) OnClose(ctx vivid.ActorContext, socket socket.Socket, err error) {

}

func (w *webSocketActor) OnPacket(ctx vivid.ActorContext, socket socket.Socket, packet *socket.Packet) {
	socket.WritePacket(packet)
}

func (w *webSocketActor) OnReceive(ctx vivid.ActorContext) {

}
