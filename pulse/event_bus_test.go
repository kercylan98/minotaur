package pulse_test

import (
	"fmt"
	"github.com/kercylan98/minotaur/pulse"
	"github.com/kercylan98/minotaur/vivid"
	"testing"
)

type TestActor struct {
	eventBus vivid.ActorRef
}

func (t *TestActor) OnReceive(ctx vivid.MessageContext) {
	switch m := ctx.GetMessage().(type) {
	case TestActorBindEventBus:
		t.eventBus = m.EventBus
	case vivid.ActorRef:
		pulse.Subscribe[string](t.eventBus, m, ctx.GetReceiver(), "event-string")
	case string:
		fmt.Println("receive event:", m)
	case int:
		// producer message
		pulse.Publish(t.eventBus, ctx.GetReceiver(), fmt.Sprintf("event-%d", m))
	}
}

type TestActorBindEventBus struct {
	EventBus vivid.ActorRef
}

func TestNewEventBus(t *testing.T) {
	system := vivid.NewActorSystem("test")

	eventBus := vivid.ActorOf[*pulse.EventBus](&system, vivid.NewActorOptions[*pulse.EventBus]().WithName("event_bus"))
	producer := vivid.ActorOf[*TestActor](&system, vivid.NewActorOptions[*TestActor]().WithName("producer"))
	subscriber := vivid.ActorOf[*TestActor](&system, vivid.NewActorOptions[*TestActor]().WithName("subscriber"))

	producer.Tell(TestActorBindEventBus{EventBus: eventBus})
	subscriber.Tell(TestActorBindEventBus{EventBus: eventBus})
	subscriber.Tell(producer)

	for i := 0; i < 100; i++ {
		producer.Tell(i)
	}

	system.Shutdown()
}
