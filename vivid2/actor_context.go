package vivid

type ActorContext interface {
	internalActorContext
	// GetId 获取当前 ActorContext 的 ID
	GetId() ActorId

	// GetSystem 获取当前 ActorContext 所属的 ActorSystem
	GetSystem() *ActorSystem
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
