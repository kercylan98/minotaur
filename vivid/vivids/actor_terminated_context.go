package vivids

// ActorTerminatedContext 是 Actor 销毁时的上下文
type ActorTerminatedContext interface {
	// GetActorId 获取 Actor 的 ID
	GetActorId() ActorId

	// GetActor 获取 Actor 原始对象
	//  - 常被用于类型断言进行不同 Actor 类型的处理
	GetActor() any

	// HasTerminatedMessage 判断是否有销毁消息
	HasTerminatedMessage() bool

	// GetTerminatedMessage 获取销毁消息
	GetTerminatedMessage() Message

	// Restart 以全新的状态重启 Actor，包括所有的子 Actor
	//  - 该函数将会一次执行 Actor.OnSaveSnapshot、 Actor.OnDestroy、 Actor.OnPreStart 三个函数来完成重启
	//  - 当重启过程中发生错误时将会通过 Actor.OnRecoverSnapshot 函数来恢复 Actor 的状态
	Restart() error

	// Recover 保留当前的状态恢复 Actor
	Recover() error

	// GetRestartNum 获取重启次数，该值将在重启成功后清零
	GetRestartNum() int
}
