package transport

import (
	"github.com/kercylan98/minotaur/minotaur/vivid"
	"net"
)

type ConnActor struct {
	Conn             net.Conn
	Writer           vivid.ActorRef
	Typed            vivid.TypedActorRef[ConnActorTyped]
	TerminateHandler ConnTerminateHandler
}

func (c *ConnActor) OnReceive(ctx vivid.MessageContext) {
	switch m := ctx.GetMessage().(type) {
	case vivid.OnBoot:
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
	c.Typed = vivid.Typed[ConnActorTyped](ctx.GetRef(), &ConnActorTypedImpl{
		ConnActorRef:       ctx.GetRef(),
		ConnWriterActorRef: c.Writer,
	})
}
