package transport

import (
	"github.com/kercylan98/minotaur/minotaur/vivid"
	"github.com/kercylan98/minotaur/toolkit/log"
	"net"
)

type (
	connWriteMessage struct {
		Packet Packet
	}
)

type ConnWriter func(packet Packet) error

type ConnWriteActor struct {
	conn   net.Conn
	writer ConnWriter
}

func (c *ConnWriteActor) OnReceive(ctx vivid.MessageContext) {
	switch m := ctx.GetMessage().(type) {
	case connWriteMessage:
		c.onConnWriteMessage(ctx, m)
	}
}

func (c *ConnWriteActor) onConnWriteMessage(ctx vivid.MessageContext, m connWriteMessage) {
	if err := c.writer(m.Packet); err != nil {
		log.Error("ConnActor write error: %v", err)
	}
}
