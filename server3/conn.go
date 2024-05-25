package server

import (
	"github.com/kercylan98/minotaur/vivid"
	"net"
)

type ConnWriter func(packet Packet) error

type conn struct {
	conn   net.Conn
	writer ConnWriter
}

func (c *conn) OnReceive(ctx vivid.MessageContext) {
	
}
