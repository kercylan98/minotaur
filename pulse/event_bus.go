package pulse

import (
	"github.com/kercylan98/minotaur/vivid"
)

// EventBus 基于 Actor 模型实现的事件总线
type EventBus struct {
	//rw              sync.RWMutex                // 保护生产者列表的读写锁
	expectProducers map[vivid.ActorId]*producer // 期望的生产者
}

func (e *EventBus) OnReceive(ctx vivid.MessageContext) {
	switch m := ctx.GetMessage().(type) {
	case vivid.OnPreStart:
		e.onStart()
	case eventSubscribeMessage:
		e.onSubscribe(m)
	case eventUnsubscribeMessage:
		e.onUnsubscribe(m)
	case eventPublishMessage:
		e.onPublish(m)
	}
}

func (e *EventBus) onStart() {
	e.expectProducers = make(map[vivid.ActorId]*producer)
}

func (e *EventBus) onSubscribe(m eventSubscribeMessage) {
	//e.rw.Lock()
	producerActorId := vivid.GetActorIdByActorRef(m.producer)
	producer, exists := e.expectProducers[producerActorId]
	if !exists {
		producer = newProducer(m.producer)
		e.expectProducers[producerActorId] = producer
	}
	//e.rw.Unlock()

	producer.subscribe(m)
}

func (e *EventBus) onUnsubscribe(m eventUnsubscribeMessage) {
	//e.rw.Lock()
	//defer e.rw.Unlock()
	producerActorId := vivid.GetActorIdByActorRef(m.producer)
	producer, exists := e.expectProducers[producerActorId]
	if !exists {
		return
	}

	if producer.unsubscribe(m) {
		delete(e.expectProducers, producerActorId)
	}
}

func (e *EventBus) onPublish(m eventPublishMessage) {
	//e.rw.RLock()
	producerActorId := vivid.GetActorIdByActorRef(m.producer)
	producer, exists := e.expectProducers[producerActorId]
	if !exists {
		//e.rw.RUnlock()
		return
	}
	//e.rw.RUnlock()

	producer.publish(m)
}
