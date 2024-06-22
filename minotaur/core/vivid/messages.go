package vivid

import "github.com/kercylan98/minotaur/minotaur/core"

type Message = core.Message

type regulatoryMessages struct {
	Sender  ActorRef
	Message Message
}

type (
	OnBoot       struct{}
	OnTerminate  struct{}
	OnTerminated struct { // 需要迁移到 protobuf
		TerminatedActor ActorRef
	}
)

var (
	onBoot      OnBoot
	onTerminate OnTerminate
)
