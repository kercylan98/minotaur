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
}

type senderContextCompose interface {
	Tell(target ActorRef, message vivid.Message)
}

type receiverContextCompose interface {
	Message() Message
}

type spawnerContextCompose interface {
	// ActorOf 以该上下文为父级创建一个新的 Actor，返回新 Actor 的引用
	ActorOf(producer ActorProducer) ActorRef

	// Terminate 通知目标 Actor 终止
	Terminate(target ActorRef)
}
