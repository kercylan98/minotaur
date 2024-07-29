package vivid

import (
	"github.com/kercylan98/minotaur/engine/prc"
	"time"
)

var (
	onTerminate           = &OnTerminate{}
	onGracefullyTerminate = &OnTerminate{Gracefully: true}
	onLaunch              = new(OnLaunch)
	onRestart             = new(onRestartMessage)
	onRestarting          = new(OnRestarting)
	onRestarted           = new(OnRestarted)
	onSuspendMailbox      = new(onSuspendMailboxMessage)
	onResumeMailbox       = new(onResumeMailboxMessage)
	onPersistenceSnapshot = new(OnPersistenceSnapshot)
)

type (
	onSuspendMailboxMessage int8
	onResumeMailboxMessage  int8
	onRestartMessage        int8
	onSchedulerFunc         func()
)

type (
	// OnLaunch 在 Actor 启动时，将会作为第一条消息被处理，适用于初始化 Actor 状态等场景。
	OnLaunch int8

	// OnTerminated 当收到该消息时，说明 TerminatedActor 已经被终止，如果是自身，那么表示自身已被终止。
	OnTerminated struct {
		TerminatedActor ActorRef
	}

	// OnRestarting 在 Actor 由于意外情况被监管者执行重启策略时，将收到该消息，后续还会根据生命周期收到一系列消息，具体如下：
	//  - OnTerminate
	//  - OnTerminated
	//  - OnRestarted
	//  - OnLaunch
	OnRestarting int8

	// OnRestarted 当收到该消息时候即表示 Actor 的重启已完成，紧跟着将会收到 OnLaunch 消息。
	OnRestarted int8

	// OnPersistenceSnapshot 当 Actor 的事件数量超过持久化事件数量阈值时，将会触发快照的持久化，收到该消息时应主动调用 SaveSnapshot 函数保存快照。
	OnPersistenceSnapshot int8

	OnSlowProcess struct {
		Duration time.Duration // 处理耗时
		ActorRef ActorRef      // 慢处理 Actor 引用
	}
)

type Message = prc.Message
type MessageWrapper struct {
	Sender   ActorRef
	Receiver ActorRef
	Message  Message
}

func wrapMessage(sender ActorRef, receiver ActorRef, message Message) MessageWrapper {
	return MessageWrapper{
		Sender:   sender,
		Receiver: receiver,
		Message:  message,
	}
}

func unwrapMessage(wrapper Message) (sender ActorRef, receiver ActorRef, message Message) {
	w, ok := wrapper.(MessageWrapper)
	if !ok {
		return nil, nil, message
	}
	return w.Sender, w.Receiver, w.Message
}
