package vivid

import (
	"github.com/kercylan98/minotaur/core"
	"github.com/kercylan98/minotaur/minotaur/vivid"
	"sync/atomic"
)

var (
	_ core.Process = &Process{}
)

func newProcess(address core.Address, mailbox Mailbox) *Process {
	return &Process{
		address: address,
		mailbox: mailbox,
	}
}

type Process struct {
	address core.Address
	mailbox Mailbox
	status  uint32
}

func (p *Process) GetAddress() core.Address {
	return p.address
}

func (p *Process) Deaden() bool {
	return atomic.LoadUint32(&p.status) == 1
}

func (p *Process) Dead() {
	atomic.StoreUint32(&p.status, 1)
}

func (p *Process) SendUserMessage(sender *core.ProcessRef, message vivid.Message) {
	p.mailbox.DeliveryUserMessage(message)
}

func (p *Process) SendSystemMessage(sender *core.ProcessRef, message vivid.Message) {
	p.mailbox.DeliverySystemMessage(message)
}

func (p *Process) Terminate(sender *core.ProcessRef) {
	p.SendSystemMessage(sender, onTerminate)
}
