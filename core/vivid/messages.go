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
	OnLaunch     struct{}
	OnTerminate  struct{}
	OnTerminated struct {
		TerminatedActor ActorRef
	}
	OnRestart        struct{}
	OnRestarting     struct{}
	OnResumeMailbox  struct{}
	OnSuspendMailbox struct{}
)

var (
	onLaunch         OnLaunch
	onTerminate      OnTerminate
	onRestart        OnRestart
	onRestarting     OnRestarting
	onResumeMailbox  OnResumeMailbox
	onSuspendMailbox OnSuspendMailbox
)
