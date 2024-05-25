package pulse_test

import (
	"fmt"
	"github.com/kercylan98/minotaur/pulse"
	"github.com/kercylan98/minotaur/vivid"
	"testing"
	"time"
)

var tester *testing.T
var receiveEventCount int
var produceEventCount int

type TestActor struct {
	eventBus vivid.ActorRef
	producer vivid.ActorRef
}

func (t *TestActor) OnReceive(ctx vivid.MessageContext) {
	switch m := ctx.GetMessage().(type) {
	case TestActorBindEventBus:
		t.eventBus = m.EventBus
	case vivid.ActorRef:
		t.producer = m
		pulse.Subscribe[string](t.eventBus, m, ctx.GetReceiver(), "event-string")
		tester.Log("subscribe event")
		m.Tell(true)
	case string:
		receiveEventCount++
		tester.Log("receive event:", m)
		pulse.Unsubscribe[string](t.eventBus, t.producer, ctx.GetReceiver(), "event-string")
		tester.Log("unsubscribe event")
	case int:
		// producer message
		tester.Log("produce event:", m)
		produceEventCount++
		pulse.Publish(t.eventBus, ctx.GetReceiver(), fmt.Sprintf("event-%d", m))
	}
}

type TestActorBindEventBus struct {
	EventBus vivid.ActorRef
}

func TestNewEventBus(t *testing.T) {
	tester = t
	system := vivid.NewActorSystem("test")

	eventBus := vivid.ActorOf[*pulse.EventBus](&system, vivid.NewActorOptions[*pulse.EventBus]().WithName("event_bus"))
	producer := vivid.ActorOf[*TestActor](&system, vivid.NewActorOptions[*TestActor]().WithName("producer"))
	subscriber := vivid.ActorOf[*TestActor](&system, vivid.NewActorOptions[*TestActor]().WithName("subscriber"))

	producer.Tell(TestActorBindEventBus{EventBus: eventBus})
	subscriber.Tell(TestActorBindEventBus{EventBus: eventBus})
	subscriber.Tell(producer)

	for i := 0; i < 10; i++ {
		producer.Tell(i)
	}

	time.Sleep(time.Second)
	system.Shutdown()

	t.Log("produce event count:", produceEventCount)
	t.Log("receive event count:", receiveEventCount)
}
