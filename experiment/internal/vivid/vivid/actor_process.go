package vivid

import (
	"github.com/kercylan98/minotaur/experiment/internal/vivid/prc"
	"github.com/kercylan98/minotaur/experiment/internal/vivid/vivid/mailbox"
	"sync/atomic"
)

var _ prc.Process = (*actorProcess)(nil)

func newActorProcess(mailbox mailbox.Mailbox) *actorProcess {
	ap := &actorProcess{
		mailbox: mailbox,
	}
	return ap
}

type actorProcess struct {
	ref        ActorRef
	mailbox    mailbox.Mailbox
	terminated atomic.Bool
}

func (a *actorProcess) Initialize(rc *prc.ResourceController, id *prc.ProcessId) {
	a.ref = prc.NewProcessRef(id)
}

func (a *actorProcess) DeliveryUserMessage(sender, forward *prc.ProcessRef, message prc.Message) {
	a.delivery(sender, forward, message, a.mailbox.DeliveryUserMessage)

}

func (a *actorProcess) DeliverySystemMessage(sender, forward *prc.ProcessRef, message prc.Message) {
	a.delivery(sender, forward, message, a.mailbox.DeliverySystemMessage)
}

func (a *actorProcess) delivery(sender, forward *prc.ProcessRef, message prc.Message, delivery func(message prc.Message)) {
	if forward != nil {
		// 在使用 future.Future 的情况下，forward 将会是 Future 的引用，此刻将 Future 作为发送方进行包装，以便回复消息时能正确发送到 Future 对象进程
		sender = forward
	}

	switch message.(type) {
	case onSuspendMailboxMessage:
		a.mailbox.Suspend()
	case onResumeMailboxMessage:
		a.mailbox.Resume()
	default:
		delivery(wrapMessage(sender, a.ref, message))
	}
}

func (a *actorProcess) IsTerminated() bool {
	return a.terminated.Load()
}

func (a *actorProcess) Terminate(source *prc.ProcessRef) {
	// 交由资源控制器调用
	a.terminated.Store(true)
}
