package vivid

import (
	"github.com/kercylan98/minotaur/core"
)

type Message = core.Message

type RegulatoryMessage struct {
	Sender   ActorRef
	Receiver ActorRef
	Message  Message
}

type (
	onBindChildren struct {
		ChildrenRef ActorRef
	}
	OnLaunch     struct{}
	OnTerminate  struct{}
	OnTerminated struct {
		TerminatedActor ActorRef
	}
	OnRestart             struct{}
	OnRestarting          struct{}
	OnResumeMailbox       struct{}
	OnSuspendMailbox      struct{}
	OnPersistenceSnapshot struct{}
	TerminateGracefully   struct{}
	onSchedulerFunc       func()
)

var (
	onLaunch              OnLaunch
	onTerminate           OnTerminate
	onRestart             OnRestart
	onRestarting          OnRestarting
	onResumeMailbox       OnResumeMailbox
	onSuspendMailbox      OnSuspendMailbox
	onPersistenceSnapshot OnPersistenceSnapshot
	onTerminateGracefully TerminateGracefully
)
