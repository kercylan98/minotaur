package vivid

import (
	core2 "github.com/kercylan98/minotaur/core"
	"github.com/kercylan98/minotaur/toolkit/log"
)

var _ DeadLetter = &deadLetterProcess{}

type DeadLetterEvent struct {
	Sender   core2.Address
	Receiver core2.Address
	Message  core2.Message
}

type DeadLetter interface {
	core2.Process
}

type deadLetterProcess struct {
	ref ActorRef
}

func (d *deadLetterProcess) Ref() ActorRef {
	return d.ref
}

func (d *deadLetterProcess) GetAddress() core2.Address {
	return core2.NewAddress("", "system", "dead_letter", 0, "")
}

func (d *deadLetterProcess) SendUserMessage(sender *core2.ProcessRef, message core2.Message) {
	switch m := message.(type) {
	case DeadLetterEvent:
		log.Warn("DeadLetter", log.String("sender", m.Sender.String()), log.String("receiver", m.Receiver.String()), log.Any("message", m.Message))
	default:
		log.Warn("DeadLetter", log.String("sender", sender.Address().String()), log.Any("message", message))
	}
}

func (d *deadLetterProcess) SendSystemMessage(sender *core2.ProcessRef, message core2.Message) {
	d.SendUserMessage(sender, message)
}

func (d *deadLetterProcess) Terminate(ref *core2.ProcessRef) {

}
