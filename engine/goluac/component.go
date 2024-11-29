package goluac

import (
	"github.com/kercylan98/minotaur/engine/vivid"
)

func NewComponent() Component {
	return &component{}
}

type Component interface {
	vivid.Component
	vivid.ActorReceiveMessageCaptureComponent
}

type component struct {
}

func (c *component) OnInitialize(actorSystem *vivid.ActorSystem) error {
	return nil
}

func (c *component) OnActorReceiveMessageCapture(ctx vivid.ActorContext) (abort bool) {
	switch m := ctx.Message().(type) {
	case *LuaMessage:
		return c.onLuaMessage(ctx, m)
	case *vivid.OnTerminated:
		c.onTerminated(ctx, m)
	}

	return
}

func (c *component) onTerminated(ctx vivid.ActorContext, m *vivid.OnTerminated) {
	ac, exist := ctx.GetValue(actorContextKey).(*actorContext)
	if !exist {
		return
	}
	ac.lua.Close()
}

func (c *component) onLuaMessage(ctx vivid.ActorContext, m *LuaMessage) bool {
	ac, exist := ctx.GetValue(actorContextKey).(*actorContext)
	if !exist {
		return false
	}
	ac.OnReceive(m)
	return true
}
