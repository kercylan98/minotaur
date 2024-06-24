package vivid

import (
	"github.com/kercylan98/minotaur/minotaur/core"
)

type Message = core.Message

type RegulatoryMessage struct {
	Sender  ActorRef
	Message Message
}

// _ 开头仅为系统消息，非 _ 开头可能为用户消息和系统消息

type (
	OnLaunch      struct{}
	OnTerminate   struct{}
	_OnTerminated struct { // 需要迁移到 protobuf
		TerminatedActor ActorRef
	}
	_OnRestart   struct{}
	OnRestarting struct{}
)

var (
	onLaunch     OnLaunch
	onTerminate  OnTerminate
	onRestart    _OnRestart
	onRestarting OnRestarting
)
