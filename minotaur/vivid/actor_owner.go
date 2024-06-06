package vivid

// ActorOwner 提供了作为合格 Actor 所有者的接口定义
type ActorOwner interface {
	// GetSystem 获取 Actor 所属的 ActorSystem
	GetSystem() *ActorSystem

	// GetContext 获取 Actor 上下文
	GetContext() ActorContext
}
