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
	Parent() ActorRef
	Ref() ActorRef
	System() *ActorSystem
}

type senderContextCompose interface {
	// Tell 向目标 Actor 发送消息
	Tell(target ActorRef, message vivid.Message)

	// Ask 向目标 Actor 非阻塞地发送可被回复的消息，这个回复可能是无限期的
	Ask(target ActorRef, message vivid.Message)

	// AgentAsk 向目标 Actor 非阻塞地发送可被回复的消息，这个回复可能是有限期的，在收到回复时将会由代理 Actor 进行处理
	AgentAsk(target ActorRef, message vivid.Message, agent ActorRef)

	// FutureAsk 向目标 Actor 非阻塞地发送可被回复的消息，这个回复是有限期的，返回一个 Future 对象，可被用于获取响应消息
	FutureAsk(target ActorRef, message vivid.Message) Future
}

type receiverContextCompose interface {
	Message() Message

	Reply(message Message)
}

type spawnerContextCompose interface {
	// ActorOf 以该上下文为父级创建一个新的 Actor，返回新 Actor 的引用
	ActorOf(producer ActorProducer) ActorRef

	// Terminate 通知目标 Actor 终止
	Terminate(target ActorRef)
}
