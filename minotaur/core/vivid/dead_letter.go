package vivid

import (
	"github.com/kercylan98/minotaur/minotaur/core"
	"github.com/kercylan98/minotaur/toolkit/log"
)

var _ DeadLetter = &deadLetterProcess{}

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

func (d *deadLetterProcess) Deaden() bool {
	return false
}

func (d *deadLetterProcess) Dead() {

}

func (d *deadLetterProcess) SendUserMessage(sender *core.ProcessRef, message core.Message) {
	log.Error("DeadLetter", log.String("sender", sender.Address().String()), log.Any("message", message))
}

func (d *deadLetterProcess) SendSystemMessage(sender *core.ProcessRef, message core.Message) {

}

func (d *deadLetterProcess) Terminate(ref *core.ProcessRef) {

}
