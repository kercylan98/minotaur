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
	mailbox    mailbox.Mailbox
	terminated atomic.Bool
}

func (a *actorProcess) Initialize(rc *prc.ResourceController, id *prc.ProcessId) {
	// 不重要
}

func (a *actorProcess) DeliveryUserMessage(sender, forward *prc.ProcessRef, message prc.Message) {
	a.delivery(message, a.mailbox.DeliveryUserMessage)

}

func (a *actorProcess) DeliverySystemMessage(sender, forward *prc.ProcessRef, message prc.Message) {
	a.delivery(message, a.mailbox.DeliverySystemMessage)
}

func (a *actorProcess) delivery(message prc.Message, delivery func(message prc.Message)) {
	switch message.(type) {
	case onSuspendMailboxMessage:
		a.mailbox.Suspend()
	case onResumeMailboxMessage:
		a.mailbox.Resume()
	default:
		delivery(message)
	}
}

func (a *actorProcess) IsTerminated() bool {
	return a.terminated.Load()
}

func (a *actorProcess) Terminate(source *prc.ProcessRef) {
	// 交由资源控制器调用
	a.terminated.Store(true)
}
