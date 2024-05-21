package vivids

// Actor 是 Actor 模型的接口，该接口用于定义一个 Actor
type Actor interface {
	// OnPreStart 在 Actor 启动之前执行的逻辑，适用于对 Actor 状态的初始化
	OnPreStart(ctx ActorContext) (err error)

	// OnReceived 当 Actor 接收到消息时执行的逻辑，返回值的错误用于表示消息处理是否成功，若需要响应错误信息则应通过 MessageContext.Reply 函数进行回复
	OnReceived(ctx MessageContext) error

	// OnDestroy 当 Actor 被要求销毁时将会调用该函数，需要在该函数中释放 Actor 的资源
	//  - 该函数会在重启前被调用，被用于重置 Actor 的状态
	OnDestroy(ctx ActorContext) (err error)

	// OnSaveSnapshot 当 Actor 被要求保存快照时将会调用该函数
	OnSaveSnapshot(ctx ActorContext) (snapshot []byte, err error)

	// OnRecoverSnapshot 当 Actor 被要求恢复快照时将会调用该函数
	OnRecoverSnapshot(ctx ActorContext, snapshot []byte) (err error)

	// OnChildTerminated 当 Actor 的子 Actor 被销毁时将会调用该函数
	OnChildTerminated(ctx ActorContext, child ActorTerminatedContext)

	// OnEvent 当 Actor 接收到事件时将会调用该函数
	OnEvent(ctx ActorContext, event Message) (err error)
}
