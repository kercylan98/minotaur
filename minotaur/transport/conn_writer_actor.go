package transport

import (
	"github.com/kercylan98/minotaur/minotaur/vivid"
)

type ConnWriter func(packet Packet) error

type ConnWriteActor struct {
	Writer ConnWriter
}

func (c *ConnWriteActor) OnReceive(ctx vivid.MessageContext) {
	switch m := ctx.GetMessage().(type) {
	case Packet:
		c.onConnWriteMessage(ctx, m)
	}
}

func (c *ConnWriteActor) onConnWriteMessage(ctx vivid.MessageContext, m Packet) {
	if err := c.Writer(m); err != nil {
		panic(err)
	}
}
