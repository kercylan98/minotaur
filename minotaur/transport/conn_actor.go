package transport

import (
	"github.com/kercylan98/minotaur/minotaur/vivid"
	"net"
)

type ConnActor struct {
	vivid.ActorRef
	Conn             net.Conn
	Writer           vivid.ActorRef
	Typed            ConnActorTyped
	TerminateHandler ConnTerminateHandler
}

func (c *ConnActor) OnReceive(ctx vivid.MessageContext) {
	switch m := ctx.GetMessage().(type) {
	case vivid.OnBoot:
		c.ActorRef = ctx
	case vivid.OnActorTyped[ConnActorTyped]:
		c.Typed = m.Typed
	case ConnectionInitMessage:
		c.onInit(ctx, m)
	case ConnectionSetPacketHandlerMessage:
		behavior := vivid.BehaviorOf(func(ctx vivid.MessageContext, packet ConnectionReactPacketMessage) {
			m.Handler(ctx, c.Typed, packet)
		})
		ctx.Become(behavior, true)
	case ConnectionSetTerminateHandlerMessage:
		c.TerminateHandler = m.Handler
	case ConnectionLoadModMessage:
		ctx.LoadMod(m.Mods...)
	case ConnectionUnloadModMessage:
		ctx.UnloadMod(m.Mods...)
	case ConnectionApplyModMessage:
		ctx.ApplyMod()
	case ConnectionSetZombieTimeoutMessage:
		ctx.SetIdleTimeout(m.Timeout)
	case vivid.OnTerminate:
		if c.TerminateHandler != nil {
			c.TerminateHandler(ctx, c.Typed, m)
		}
		_ = c.Conn.Close()
	}
}

func (c *ConnActor) onInit(ctx vivid.MessageContext, m ConnectionInitMessage) {
	c.Conn = m.Conn
	c.Writer = ctx.ActorOf(vivid.OfO(func(options *vivid.ActorOptions[*ConnWriteActor]) {
		options.WithName("writer")
		options.WithInit(func(actor *ConnWriteActor) {
			actor.Writer = m.Writer
		})

		options.WithSupervisor(func(message, reason vivid.Message) vivid.Directive {
			// 关闭连接及其写入器
			ctx.Stop()
			return vivid.DirectiveStop
		})
	}))
	if m.ActorHook != nil {
		m.ActorHook(c)
	}
}
