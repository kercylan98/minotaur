package vivid

// ActorBehavior 是 Actor 的行为，用于处理 Actor 接收到的消息
type ActorBehavior[T any] func(ctx MessageContext) error

// ActorBehaviorExecutor 是 Actor 的行为执行器
type ActorBehaviorExecutor func() error

func (f ActorBehaviorExecutor) Execute() error {
	return f()
}
