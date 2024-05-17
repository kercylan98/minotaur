package vivids

import "context"

type ActorContext interface {
	context.Context

	// GetActorId 获取 Actor 的 ID
	GetActorId() ActorId

	// GetActor 获取 Actor 的引用
	GetActor() Query

	// GetParentActor 获取父 Actor 的引用
	GetParentActor() ActorRef

	// Future 创建一个 Future 对象，用于异步获取 Actor 的返回值
	Future(handler func() Message) Future

	// ActorOf 创建一个 Actor，该 Actor 是当前 Actor 的子 Actor
	ActorOf(actor Actor, opts ...*ActorOptions) (ActorRef, error)

	// NotifyTerminated 当 Actor 主动销毁时，务必调用该函数，以便在整个 Actor 系统中得到完整的释放
	NotifyTerminated(v ...Message)
}
