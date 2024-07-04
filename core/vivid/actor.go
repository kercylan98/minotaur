package vivid

import "github.com/kercylan98/minotaur/core"

type (
	// ActorId 是一个类型别名，它表示一个 Actor 的唯一标识符。
	// 这个类型别名将 Address 类型重命名为 ActorId，以提高代码的可读性和可维护性。
	ActorId = core.Address

	// ActorOptionDefiner 是一个函数类型，用于定义 Actor 的选项。
	// 这个类型表示一个函数，该函数接收一个指向 ActorOptions 结构体的指针作为参数。
	// 当调用这个函数时，它会对传入的 ActorOptions 进行配置。
	ActorOptionDefiner func(options *ActorOptions)

	// ActorProducer 是一个函数类型，用于生成一个新的 Actor 实例。
	// 这个类型表示一个无参数的函数，返回一个实现了 Actor 接口的对象。
	ActorProducer func() Actor

	// FunctionalActor 是 OnReceiveFunc 类型的别名。
	// 这个类型表示一个函数，该函数接收一个 ActorContext 作为参数。
	//   - 该类型常被用于定义无状态的 Actor。
	FunctionalActor = OnReceiveFunc

	// OnReceiveFunc 是一个函数类型，用于定义 Actor 的消息处理函数。
	// 这个类型表示一个函数，该函数接收一个 ActorContext 作为参数，并处理接收到的消息。
	OnReceiveFunc func(ctx ActorContext)
)

// Actor 接口表示一个可以接收消息的实体。
// 这个接口定义了一个方法 OnReceive，它接收一个 ActorContext 作为参数。
type Actor interface {
	// OnReceive 是 Actor 接口的核心方法。
	// 当 Actor 接收到一条消息时，这个方法会被调用。
	// ctx 是 ActorContext 类型的参数，它提供了与 Actor 的上下文信息，
	// 如消息的发送者、消息的内容等。
	OnReceive(ctx ActorContext)
}

// OnReceive 是 OnReceiveFunc 类型的方法。它的作用与 Actor 接口的 Actor.OnReceive 方法相同。
func (f OnReceiveFunc) OnReceive(ctx ActorContext) {
	f(ctx)
}
