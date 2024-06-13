package transport

import (
	"github.com/kercylan98/minotaur/minotaur/vivid"
	"net"
)

type ConnActor struct {
	Conn   net.Conn
	Writer ConnWriter
}

func (c *ConnActor) OnReceive(ctx vivid.MessageContext) {
	switch m := ctx.GetMessage().(type) {
	case vivid.OnBoot:
	case ConnectionInitMessage:
		c.Conn, c.Writer = m.Conn, m.Writer
		ctx.Become(vivid.BehaviorOf[ConnectionReactPacketMessage](c.onReactPacket))
	case ConnectionBecomeReactPacketMessage:
		ctx.Become(m.Behavior)
	case ConnectionWritePacketMessage:
		if err := c.Writer(m.Packet); err != nil {
			ctx.Stop()
		}
	case ConnectionLoadModMessage:
		ctx.LoadMod(m.Mods...)
	case ConnectionUnloadModMessage:
		ctx.UnloadMod(m.Mods...)
	case ConnectionApplyModMessage:
		ctx.ApplyMod()
	case vivid.OnTerminate:
		_ = c.Conn.Close()
	}
}

func (c *ConnActor) onReactPacket(ctx vivid.MessageContext, message ConnectionReactPacketMessage) {
	ctx.Publish(ConnectionReceivePacketEvent{})
}
