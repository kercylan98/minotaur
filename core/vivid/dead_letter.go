package vivid

import (
	"github.com/kercylan98/minotaur/core"
	"github.com/kercylan98/minotaur/toolkit/eventstream"
	"github.com/kercylan98/minotaur/toolkit/log"
)

var _ DeadLetter = &deadLetterProcess{}

type DeadLetter interface {
	core.Process

	Ref() ActorRef

	Subscribe(handler func(event DeadLetterEvent)) eventstream.Subscription

	Unsubscribe(subscription eventstream.Subscription)
}

type deadLetterProcess struct {
	system *ActorSystem
	ref    ActorRef
}

func (d *deadLetterProcess) Subscribe(handler func(event DeadLetterEvent)) eventstream.Subscription {
	return d.system.eventStream.Subscribe(func(event interface{}) {
		if e, ok := event.(DeadLetterEvent); ok {
			handler(e)
		}
	})
}

func (d *deadLetterProcess) Unsubscribe(subscription eventstream.Subscription) {
	d.system.eventStream.Unsubscribe(subscription)
}

func (d *deadLetterProcess) Ref() ActorRef {
	return d.ref
}

func (d *deadLetterProcess) GetAddress() core.Address {
	return core.NewAddress("", "system", "dead_letter", 0, "")
}

func (d *deadLetterProcess) SendUserMessage(sender *core.ProcessRef, message core.Message) {
	s, r, m := unwrapRegulatoryMessage(message)
	d.system.eventStream.Publish(DeadLetterEvent{
		Sender:   s,
		Receiver: r,
		Message:  m,
	})

	d.system.opts.LoggerProvider().Warn("DeadLetter", log.String("sender", sender.Address().String()), log.Any("message", message))
}

func (d *deadLetterProcess) SendSystemMessage(sender *core.ProcessRef, message core.Message) {
	d.SendUserMessage(sender, message)
}

func (d *deadLetterProcess) Terminate(ref *core.ProcessRef) {

}
