package goluac

import (
	"github.com/kercylan98/minotaur/engine/vivid"
	"github.com/kercylan98/minotaur/toolkit"
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
	case *vivid.OnLaunch, *vivid.OnTerminate, *vivid.OnRestarted, *vivid.OnRestarting, *vivid.OnSlowProcess:
	case *LuaMessage:
		return c.onLuaMessage(ctx, m)
	case *vivid.OnTerminated:
		c.onTerminated(ctx, m)
	default:
		return c.tryCastToLuaMessage(ctx, m)
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
	ac.OnReceive(m, true)
	return true
}

func (c *component) tryCastToLuaMessage(ctx vivid.ActorContext, m vivid.Message) bool {
	ac, exist := ctx.GetValue(actorContextKey).(*actorContext)
	if !exist {
		return false
	}

	data, err := toolkit.MarshalJSONE(m)
	if err != nil {
		return false
	}

	message := &LuaMessage{Data: data}
	if err = ac.createMessageCache(message); err != nil {
		return false
	}

	ac.OnReceive(message, false)
	return false
}
