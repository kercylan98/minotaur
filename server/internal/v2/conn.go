package server

import (
	"context"
	"github.com/kercylan98/minotaur/server/internal/v2/dispatcher"
	"net"
)

type ConnWriter func(packet Packet) error

type Conn interface {
}

func newConn(ctx context.Context, c net.Conn, connWriter ConnWriter) Conn {
	return &conn{
		conn:   c,
		writer: connWriter,
		actor:  dispatcher.NewActor[Packet](ctx, handler),
	}
}

type conn struct {
	conn   net.Conn
	writer ConnWriter
	actor  *dispatcher.Actor[Packet]
}
