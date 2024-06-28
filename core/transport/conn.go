package transport

import (
	"github.com/kercylan98/minotaur/core/vivid"
	"github.com/kercylan98/minotaur/toolkit/log"
)

type ConnWriter func(packet Packet) error

type conn struct {
	writer ConnWriter
}

func (c *conn) OnReceive(ctx vivid.ActorContext) {
	switch m := ctx.Message().(type) {
	case vivid.OnLaunch:

	case Packet:
		c.onWritePacket(ctx, m)
	}
}

func (c *conn) onWritePacket(ctx vivid.ActorContext, m Packet) {
	if err := c.writer(m); err != nil {
		log.Error("WritePacket", log.Err(err))
		ctx.Terminate(ctx.Ref())
	}
}
