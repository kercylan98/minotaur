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
	OnLaunch     struct{} // Actor 收到的第一条消息，表明 Actor 已经准备就绪，可以处理消息。该阶段通常被用于初始化 Actor 的运行时状态。
	OnTerminate  struct{} // 当收到该消息时表明 Actor 在处理完该消息后将被销毁。该阶段通常被用于释放或持久化 Actor 的运行时状态。
	OnTerminated struct { // 当 Actor 已被销毁时将会收到该消息，通常用于记录日志、计数等情况。
		TerminatedActor ActorRef
	}
	OnRestart             struct{}
	OnRestarting          struct{} // 当 Actor 重启时将会收到该消息，在重启完成后将会收到 OnLaunch 消息。
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

func wrapRegulatoryMessage(sender, receiver ActorRef, message Message) RegulatoryMessage {
	switch m := message.(type) {
	case RegulatoryMessage:
		return m
	default:
		return RegulatoryMessage{
			Sender:   sender,
			Receiver: receiver,
			Message:  message,
		}
	}
}

func unwrapRegulatoryMessage(message Message) (sender, receiver ActorRef, msg Message) {
	switch m := message.(type) {
	case RegulatoryMessage:
		return m.Sender, m.Receiver, m.Message
	default:
		return nil, nil, message
	}
}
