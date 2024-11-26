package vivid

// Component 是对于 ActorSystem 的扩展
type Component interface {
	// OnInitialize 在 ActorSystem 初始化之后被调用
	OnInitialize(actorSystem *ActorSystem) error
}

// ShutdownComponent 是对于 ActorSystem 的扩展，它允许在 ActorSystem 关闭后执行一些清理工作
type ShutdownComponent interface {
	Component

	// OnShutdown 在 ActorSystem 关闭后被调用，它的执行阶段位于最末期
	OnShutdown(actorSystem *ActorSystem) error
}

// ActorSpawnBeforeComponent 可以在 Actor 创建之前对其提供者及描述符进行捕获
type ActorSpawnBeforeComponent interface {
	Component

	// OnActorSpawnBefore 在 Actor 创建之前被调用
	OnActorSpawnBefore(actorSystem *ActorSystem, provider ActorProvider, descriptor *ActorDescriptor)
}
