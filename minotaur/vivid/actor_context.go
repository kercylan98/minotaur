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

	// Become 切换 Actor 在面对特定消息的行为。通过调用 BehaviorOf 函数可以创建一个特定消息类型的行为，将 Actor 的消息处理逻辑切换为指定的新行为。新的行为会覆盖当前的行为，直到下次调用 Become 或 UnBecome 为止
	Become(behavior Behavior)

	// UnBecome 恢复 Actor 在面对特性消息的行为为之前的行为，多次调用 UnBecome 会依次恢复之前的行为，直到没有行为为止
	UnBecome(message Message)

	// ActorOf 创建一个 Actor 并返回 ActorRef
	//  - ActorOfO 对象可通过 OfO 函数快速创建
	ActorOf(ofo ActorOfO) ActorRef

	// Terminate 销毁当前 Actor，该操作将会触发 Actor 的 OnTerminate 生命周期
	Terminate()
}

type _ActorContext struct {
	*_internalActorContext
	*_ActorCore
	behaviors map[reflect.Type][]Behavior
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

func (c *_ActorContext) Become(behavior Behavior) {
	if c.behaviors == nil {
		c.behaviors = make(map[reflect.Type][]Behavior)
	}

	messageType := behavior.getMessageType()
	c.behaviors[messageType] = append(c.behaviors[messageType], behavior)
}

func (c *_ActorContext) UnBecome(message Message) {
	if c.behaviors == nil {
		return
	}

	messageType := reflect.TypeOf(message)
	if behaviors, ok := c.behaviors[messageType]; ok {
		if len(behaviors) > 0 {
			c.behaviors[messageType] = behaviors[:len(behaviors)-1]
		}
		if len(behaviors) == 0 {
			delete(c.behaviors, messageType)
		}
	}
}

func (c *_ActorContext) ActorOf(ofo ActorOfO) ActorRef {
	return ofo.generate(c)
}

func (c *_ActorContext) Terminate() {
	c.Tell(OnTerminate{})
}
