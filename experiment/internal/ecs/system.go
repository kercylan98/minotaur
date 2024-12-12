package ecs

type Lifecycle uint8

const (
	OnInit    Lifecycle = iota // 初始化启动阶段
	OnRunning                  // 运行阶段
)

type System interface {
	OnLifecycle(world World, lifecycle Lifecycle)

	OnUpdate(world World)
}
