package vivid

// ActorContext 针对 Actor 的上下文，该上下文暴露给 Actor 自身使用，但不提供外部自行实现
//   - 上下文代表了 Actor 完整的生命周期，该上下文将在 Actor 的生命周期中一直存在
type ActorContext interface {
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
}

type _ActorContext struct {
	*_internalActorContext
	*_ActorCore
}

func (c *_ActorContext) actorOf(actor Actor, opt any) ActorRef {
	return nil
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
