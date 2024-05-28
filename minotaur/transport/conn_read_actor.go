package transport

import (
	"github.com/kercylan98/minotaur/minotaur/pulse"
	"github.com/kercylan98/minotaur/minotaur/vivid"
)

type (
	connReceivePacketMessage struct {
		Packet Packet
	}
)

type (
	ConnReceiveEvent struct {
		Conn   Conn
		Packet Packet
	}
)

type ConnReadActor struct {
	actor    vivid.ActorRef
	conn     Conn
	eventBus *pulse.Pulse
}

func (c *ConnReadActor) OnReceive(ctx vivid.MessageContext) {
	switch m := ctx.GetMessage().(type) {
	case vivid.OnPreStart:
		c.actor = ctx.GetReceiver()
	case connReceivePacketMessage:
		c.onConnReceiveMessage(ctx, m)
	}
}

func (c *ConnReadActor) onConnReceiveMessage(ctx vivid.MessageContext, m connReceivePacketMessage) {
	c.eventBus.Publish(c.actor, ConnReceiveEvent{Conn: c.conn, Packet: m.Packet})
}
