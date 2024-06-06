package vivid

import (
	"context"
	"reflect"
)

// ActorContext 针对 Actor 的上下文，该上下文暴露给 Actor 自身使用，但不提供外部自行实现
//   - 上下文代表了 Actor 完整的生命周期，该上下文将在 Actor 的生命周期中一直存在
type ActorContext interface {
	context.Context
	internalActorContext
	ActorRef
	// GetId 获取当前 ActorContext 的 ID
	GetId() ActorId

	// GetSystem 获取当前 ActorContext 所属的 ActorSystem
	GetSystem() *ActorSystem

	// GetActor 获取当前 ActorContext 的 Actor 对象
	GetActor() Actor

	// GetParent 获取当前上下文的父 Actor 的引用
	GetParent() ActorRef

	// BindBehavior 动态的绑定一个行为，当行为存在时，Actor.OnReceive 将会被覆盖
	BindBehavior(behavior Behavior)

	// UnbindBehavior 解绑一个已绑定的行为
	UnbindBehavior(message Message)

	// ActorOf 创建一个 Actor 并返回 ActorRef
	//  - ActorOfO 对象可通过 OfO 函数快速创建
	ActorOf(ofo ActorOfO) ActorRef
}

type _ActorContext struct {
	*_internalActorContext
	*_ActorCore
	behaviors map[reflect.Type]Behavior // 行为
}

func (c *_ActorContext) GetId() ActorId {
	return c.id
}

func (c *_ActorContext) GetSystem() *ActorSystem {
	return c.system
}

func (c *_ActorContext) GetActor() Actor {
	return c.Actor
}

func (c *_ActorContext) GetParent() ActorRef {
	return c.parent
}

func (c *_ActorContext) BindBehavior(behavior Behavior) {
	if c.behaviors == nil {
		c.behaviors = make(map[reflect.Type]Behavior)
	}

	c.behaviors[behavior.getMessageType()] = behavior
}

func (c *_ActorContext) UnbindBehavior(message Message) {
	if c.behaviors == nil {
		return
	}

	delete(c.behaviors, reflect.TypeOf(message))
}

func (c *_ActorContext) ActorOf(ofo ActorOfO) ActorRef {
	return ofo.generate(c)
}
