package session

import (
	"github.com/kercylan98/minotaur/engine/socket"
	"github.com/kercylan98/minotaur/engine/vivid"
	"github.com/kercylan98/minotaur/skeleton/pkg/components"
)

type HandlerFunc func(ctx vivid.ActorContext, socket socket.Socket, reader *Reader)

func New[HandleFunc any](router components.RouterComponent[HandleFunc]) socket.Actor {
	return &conn[HandleFunc]{
		router: router,
	}
}

type conn[HandleFunc any] struct {
	router components.RouterComponent[HandleFunc]
}

func (c *conn[HandleFunc]) OnOpened(ctx vivid.ActorContext, socket socket.Socket) {

}

func (c *conn[HandleFunc]) OnClose(ctx vivid.ActorContext, socket socket.Socket, err error) {

}

func (c *conn[HandleFunc]) OnReceive(ctx vivid.ActorContext) {

}

func (c *conn[HandleFunc]) OnPacket(ctx vivid.ActorContext, socket socket.Socket, packet *socket.Packet) {

}
