package session

import (
	"github.com/kercylan98/minotaur/engine/socket"
	"github.com/kercylan98/minotaur/engine/vivid"
)

func New() socket.Actor {
	return &conn{}
}

type conn struct {
}

func (c *conn) OnOpened(ctx vivid.ActorContext, socket socket.Socket) {

}

func (c *conn) OnClose(ctx vivid.ActorContext, socket socket.Socket, err error) {

}

func (c *conn) OnReceive(ctx vivid.ActorContext) {

}

func (c *conn) OnPacket(ctx vivid.ActorContext, socket socket.Socket, packet *socket.Packet) {

}
