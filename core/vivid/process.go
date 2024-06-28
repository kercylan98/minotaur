package vivid

import (
	core2 "github.com/kercylan98/minotaur/core"
	"github.com/kercylan98/minotaur/minotaur/vivid"
	"sync/atomic"
)

var (
	_ core2.Process = &Process{}
)

func NewProcess(address core2.Address, mailbox Mailbox) *Process {
	return &Process{
		address: address,
		mailbox: mailbox,
	}
}

type Process struct {
	address core2.Address
	mailbox Mailbox
	status  uint32
}

func (p *Process) GetAddress() core2.Address {
	return p.address
}

func (p *Process) Deaden() bool {
	return atomic.LoadUint32(&p.status) == 1
}

func (p *Process) Dead() {
	atomic.StoreUint32(&p.status, 1)
}

func (p *Process) SendUserMessage(sender *core2.ProcessRef, message vivid.Message) {
	p.mailbox.DeliveryUserMessage(message)
}

func (p *Process) SendSystemMessage(sender *core2.ProcessRef, message vivid.Message) {
	p.mailbox.DeliverySystemMessage(message)
}

func (p *Process) Terminate(sender *core2.ProcessRef) {
	p.SendSystemMessage(sender, onTerminate)
}
