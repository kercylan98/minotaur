package vivid

import "github.com/kercylan98/minotaur/experiment/internal/vivid/prc"

const (
	onSuspendMailbox      onSuspendMailboxMessage = 0
	onResumeMailbox       onResumeMailboxMessage  = 0
	onLaunch              OnLaunch                = 0
	onTerminate           OnTerminate             = 0
	onGracefullyTerminate OnTerminate             = 1
	onRestart             onRestartMessage        = 0
	onRestarting          OnRestarting            = 0
	onRestarted           OnRestarted             = 0
)

type (
	onSuspendMailboxMessage int8
	onResumeMailboxMessage  int8
	onRestartMessage        int8
)

type (
	// OnLaunch 在 Actor 启动时，将会作为第一条消息被处理，适用于初始化 Actor 状态等场景。
	OnLaunch int8

	// OnTerminate 在 Actor 处理完该消息后，将会被终止，适用于释放 Actor 资源等场景。
	OnTerminate int8

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
