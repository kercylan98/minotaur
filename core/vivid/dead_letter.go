package vivid

import (
	"github.com/kercylan98/minotaur/core"
	"github.com/kercylan98/minotaur/toolkit/log"
)

var _ DeadLetter = &deadLetterProcess{}

type DeadLetterEvent struct {
	Sender   core.Address
	Receiver core.Address
	Message  core.Message
}

type DeadLetter interface {
	core.Process
}

type deadLetterProcess struct {
	ref ActorRef
}

func (d *deadLetterProcess) Ref() ActorRef {
	return d.ref
}

func (d *deadLetterProcess) GetAddress() core.Address {
	return core.NewAddress("", "system", "dead_letter", 0, "")
}

func (d *deadLetterProcess) SendUserMessage(sender *core.ProcessRef, message core.Message) {
	switch m := message.(type) {
	case DeadLetterEvent:
		log.Warn("DeadLetter", log.String("sender", m.Sender.String()), log.String("receiver", m.Receiver.String()), log.Any("message", m.Message))
	default:
		log.Warn("DeadLetter", log.String("sender", sender.Address().String()), log.Any("message", message))
	}
}

func (d *deadLetterProcess) SendSystemMessage(sender *core.ProcessRef, message core.Message) {
	d.SendUserMessage(sender, message)
}

func (d *deadLetterProcess) Terminate(ref *core.ProcessRef) {

}
