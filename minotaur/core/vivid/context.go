package vivid

import (
	"github.com/kercylan98/minotaur/minotaur/vivid"
)

type ActorContext interface {
	basicContextCompose
	senderContextCompose
	receiverContextCompose
	spawnerContextCompose
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

	// Ref 获取当前 Actor 的引用
	Ref() ActorRef

	// System 获取当前 Actor 所在的 Actor 系统
	System() *ActorSystem
}

type senderContextCompose interface {
	// Tell 向目标 Actor 发送消息
	Tell(target ActorRef, message vivid.Message, options ...MessageOption)

	// Ask 向目标 Actor 非阻塞地发送可被回复的消息，这个回复可能是无限期的
	Ask(target ActorRef, message vivid.Message, options ...MessageOption)

	// FutureAsk 向目标 Actor 非阻塞地发送可被回复的消息，这个回复是有限期的，返回一个 Future 对象，可被用于获取响应消息
	FutureAsk(target ActorRef, message vivid.Message, options ...MessageOption) Future
}

type receiverContextCompose interface {
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

	// Terminate 通知目标 Actor 终止
	Terminate(target ActorRef)
}