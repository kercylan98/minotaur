package pulse

import (
	"github.com/kercylan98/minotaur/vivid"
	"reflect"
)

func Subscribe[T Event](eventBus vivid.ActorRef, producer Producer, subscriber Subscriber, subscribeId SubscribeId) {
	eventBus.Tell(eventSubscribeMessage{
		producer:    producer,
		subscriber:  subscriber,
		event:       reflect.TypeOf((*T)(nil)).Elem(),
		subscribeId: subscribeId,
		priority:    0,
	})
}

func Unsubscribe[T Event](eventBus vivid.ActorRef, producer Producer, subscriber Subscriber, subscribeId SubscribeId) {
	eventBus.Tell(eventUnsubscribeMessage{
		producer:    producer,
		subscriber:  subscriber,
		event:       reflect.TypeOf((*T)(nil)).Elem(),
		subscribeId: subscribeId,
	})
}

func Publish(eventBus vivid.ActorRef, producer Producer, event Event) {
	eventBus.Tell(eventPublishMessage{
		producer: producer,
		event:    event,
	})
}
