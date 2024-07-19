package vivid

import (
	"github.com/kercylan98/minotaur/experiment/internal/vivid/future"
	"time"
)

// mixinSpawner 是一个混入类型接口，它定义了作为 Actor 的生成器需要满足的接口。
type mixinSpawner interface {
	// ActorOf 生成一个新的 Actor 实例，并以该实例作为其父 Actor。返回生成的 Actor 引用(ActorRef)
	//  - 该函数接收多个 ActorDescriptorConfigurator 参数，用于配置生成的 Actor 实例，当包含多个 ActorDescriptorConfigurator 参数时，它们的配置将会是向前覆盖的。
	//
	// 该函数不是并发安全的，你不应该在多个 goroutine 中同时调用 ActorOf 函数。
	ActorOf(provider ActorProvider, configurator ...ActorDescriptorConfigurator) ActorRef

	// ActorOfF 该函数是 ActorOf 的快捷方式，它提供了更为简便的使用方式，但是会额外创建一个切片并拷贝，用于 FunctionalActorDescriptorConfigurator 到 ActorDescriptorConfigurator 的转换。
	ActorOfF(provider FunctionalActorProvider, configurator ...FunctionalActorDescriptorConfigurator) ActorRef

	// Children 返回当前 Actor 的所有子 Actor 引用(ActorRef)。
	Children() []ActorRef
}

// mixinWorker 是一个混入类型接口，它定义了作为 Actor 工作者需要满足的接口。
type mixinWorker interface {
	// Terminate 终止目标 Actor。
	//  - 当 gracefully 参数为 true 时，会将终止消息作为用户级消息进行发送，在该消息之前的用户消息被处理完毕后升级为系统消息终止 Actor。
	Terminate(target ActorRef, gracefully bool)

	// ReportAbnormal 报告异常，该函数将触发事故向监管者传递
	ReportAbnormal(reason Message)
}

// mixinDeliver 是一个混入类型接口，它定义了作为 Actor 消息发送者需要满足的接口。
type mixinDeliver interface {
	// Tell 向指定的 Actor 引用(ActorRef) 发送消息。
	Tell(target ActorRef, message Message)

	// Ask 向目标 Actor 非阻塞地发送可被回复的消息，这个回复可能是无限期的
	Ask(target ActorRef, message Message)

	// FutureAsk 向目标 Actor 非阻塞地发送可被回复的消息，这个回复是有限期的，返回一个 future.Future 对象，可被用于获取响应消息
	//  - 当 timeout 参数为空时，将会使用默认的超时时间 DefaultFutureAskTimeout
	FutureAsk(target ActorRef, message Message, timeout ...time.Duration) future.Future

	// Broadcast 向所有子级 Actor 广播消息，广播消息是可以被回复的
	//  - 子级的子级不会收到广播消息
	Broadcast(message Message)
}

// mixinRecipient 是一个混入类型接口，它定义了作为 Actor 接收者需要满足的接口。
type mixinRecipient interface {
	// System 返回当前 Actor 所属的 Actor 系统。
	System() *ActorSystem

	// Ref 返回当前 Actor 的 Actor 引用(ActorRef)。
	Ref() ActorRef

	// Reply 向消息发送者回复消息。
	Reply(message Message)

	// Message 返回当前 Actor 接收到的消息。
	Message() Message

	// Sender 返回当前 Actor 接收到的消息的发送者。
	Sender() ActorRef
}
