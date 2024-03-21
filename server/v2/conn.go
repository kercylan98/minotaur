package server

import (
	"context"
	"github.com/kercylan98/minotaur/server/v2/actor"
	"net"
)

type Conn interface {
}

func newConn(ctx context.Context, c net.Conn, connWriter ConnWriter, handler actor.MessageHandler[Packet]) Conn {
	return &conn{
		conn:   c,
		writer: connWriter,
		actor:  actor.NewActor[Packet](ctx, handler),
	}
}

type conn struct {
	conn   net.Conn
	writer ConnWriter
	actor  *actor.Actor[Packet]
}
