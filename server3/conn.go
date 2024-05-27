package server

import (
	"github.com/kercylan98/minotaur/vivid"
	"net"
)

type ConnWriter func(packet Packet) error

type conn struct {
	conn   net.Conn
	writer ConnWriter
	actor  vivid.ActorRef
}

func (c *conn) OnReceive(ctx vivid.MessageContext) {
	switch m := ctx.GetMessage().(type) {
	case vivid.OnPreStart:
		c.actor = ctx.GetReceiver()
	case onConnectionReceivedMessage:
		c.onConnectionReceivedMessage(ctx, m)
	}
}

func (c *conn) onConnectionReceivedMessage(ctx vivid.MessageContext, m onConnectionReceivedMessage) {
	packet := m.Packet
	if err := c.writer(packet); err != nil {
		c.conn.Close()
		c.actor.Tell(onConnectionClosedTellMessage{conn: c.conn})
	}
}
