package vivid

// ActorContext 是 Actor 的上下文，包含了 Actor 的基本信息以及 Actor 的操作方法
type ActorContext interface {
	basicContextCompose
	senderContextCompose
	receiverContextCompose
	spawnerContextCompose
	persistentContextCompose
}

type SenderContext interface {
	basicContextCompose
	senderContextCompose
}

type ReceiverContext interface {
	basicContextCompose
	receiverContextCompose
}

type SpawnerContext interface {
	spawnerContextCompose
}

type basicContextCompose interface {
	// Parent 获取当前 Actor 的父级 Actor
	Parent() ActorRef

	// Children 获取当前 Actor 的子级 Actor
	Children() []ActorRef

	// Ref 获取当前 Actor 的引用
	Ref() ActorRef

	// System 获取当前 Actor 所在的 Actor 系统
	System() *ActorSystem

	// DeadLetter 获取当前 Actor 系统的死信队列
	DeadLetter() DeadLetter
}

type senderContextCompose interface {
	// Tell 向目标 Actor 发送消息
	Tell(target ActorRef, message Message, options ...MessageOption)

	// Ask 向目标 Actor 非阻塞地发送可被回复的消息，这个回复可能是无限期的
	Ask(target ActorRef, message Message, options ...MessageOption)

	// FutureAsk 向目标 Actor 非阻塞地发送可被回复的消息，这个回复是有限期的，返回一个 Future 对象，可被用于获取响应消息
	FutureAsk(target ActorRef, message Message, options ...MessageOption) Future

	// AwaitForward 异步地等待阻塞结束后向目标 Actor 转发消息，收到的消息类型将是 FutureForwardMessage
	AwaitForward(target ActorRef, blockFunc func() Message)
}

type receiverContextCompose interface {
	// Sender 获取当前消息的发送人
	Sender() ActorRef

	// Message 获取当前 Actor 接收到的消息
	Message() Message

	// Reply 回复消息
	Reply(message Message)

	// BehaviorOf 生成一个行为
	BehaviorOf() Behavior
}

type spawnerContextCompose interface {
	// ActorOf 以该上下文为父级创建一个新的 Actor，返回新 Actor 的引用
	ActorOf(producer ActorProducer, options ...ActorOptionDefiner) ActorRef

	// KindOf 以该上下文为父级创建一个新的 Actor，返回新 Actor 的引用，该 Actor 的类型由 kind 指定
	//  - 当显示的指定了 parent 时，将会以 parent 为父级创建 Actor，父级可能位于远端
	KindOf(kind Kind, parent ...ActorRef) ActorRef

	// Terminate 通知目标 Actor 立即终止
	Terminate(target ActorRef)

	// TerminateGracefully 通知目标 Actor 立即终止，但是不会立即终止，而是在之前的用户消息处理完毕后终止
	TerminateGracefully(target ActorRef)
}

type persistentContextCompose interface {
	// PersistSnapshot 持久化当前快照
	PersistSnapshot(snapshot Message)

	// StatusChanged 记录导致状态变更的事件
	StatusChanged(event Message)
}
