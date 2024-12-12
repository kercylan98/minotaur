package session

import (
	"github.com/kercylan98/minotaur/engine/socket"
	"github.com/kercylan98/minotaur/engine/vivid"
)

func New() socket.Actor {
	return &sessionActor{}
}

type sessionActor struct {
}

func (s *sessionActor) OnOpened(ctx vivid.ActorContext, socket socket.Socket) {

}

func (s *sessionActor) OnClose(ctx vivid.ActorContext, socket socket.Socket, err error) {

}

func (s *sessionActor) OnPacket(ctx vivid.ActorContext, socket socket.Socket, packet *socket.Packet) {

}

func (s *sessionActor) OnReceive(ctx vivid.ActorContext) {

}
