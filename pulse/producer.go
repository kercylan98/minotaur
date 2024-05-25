package pulse

import (
	"github.com/kercylan98/minotaur/vivid"
	"reflect"
	"sort"
	"sync"
)

type Producer interface {
	vivid.ActorRef
}

type (
	// 内部生产者
	producer struct {
		Producer                                       // 生产者
		rw       sync.RWMutex                          // 保护事件列表的读写锁
		events   map[reflect.Type]*producerSubscribers // 事件列表
	}

	// 生产者订阅者列表
	producerSubscribers struct {
		rw           sync.RWMutex                  // 保护订阅者列表的读写锁
		priorities   map[SubscribeId]int           // 事件优先级
		subscribeIds map[SubscribeId]vivid.ActorId // 订阅 ID
		subscribers  map[vivid.ActorId]Subscriber  // 订阅者
		subscribeNum map[vivid.ActorId]int         // 订阅数量
	}
)

// newProducer 创建一个新的内部生产者
func newProducer(p Producer) *producer {
	return &producer{
		Producer: p,
		events:   make(map[reflect.Type]*producerSubscribers),
	}
}

func (p *producer) subscribe(message eventSubscribeMessage) {
	p.rw.Lock()
	events, exists := p.events[message.event]
	if !exists {
		events = &producerSubscribers{
			priorities:   make(map[SubscribeId]int),
			subscribeIds: make(map[SubscribeId]vivid.ActorId),
			subscribers:  make(map[vivid.ActorId]Subscriber),
			subscribeNum: make(map[vivid.ActorId]int),
		}
		p.events[message.event] = events
	}
	p.rw.Unlock()

	events.rw.Lock()
	defer events.rw.Unlock()

	subscriberActorId := vivid.GetActorIdByActorRef(message.subscriber)

	if _, exists := events.subscribeIds[message.subscribeId]; !exists {
		events.subscribeNum[subscriberActorId]++
	}

	events.priorities[message.subscribeId] = message.priority
	events.subscribeIds[message.subscribeId] = subscriberActorId
	events.subscribers[subscriberActorId] = message.subscriber
}

func (p *producer) unsubscribe(message eventUnsubscribeMessage) (empty bool) {
	p.rw.Lock()
	defer p.rw.Unlock()
	events, exists := p.events[message.event]

	if !exists {
		return
	}

	events.rw.Lock()
	defer events.rw.Unlock()

	subscriberActorId := vivid.GetActorIdByActorRef(message.subscriber)

	if _, exists := events.subscribeIds[message.subscribeId]; exists {
		events.subscribeNum[subscriberActorId]--
		if events.subscribeNum[subscriberActorId] == 0 {
			delete(events.subscribers, subscriberActorId)
			empty = true
		}
	}

	delete(events.priorities, message.subscribeId)
	delete(events.subscribeIds, message.subscribeId)

	return
}

func (p *producer) publish(message eventPublishMessage) {
	eventType := reflect.TypeOf(message.event)

	p.rw.RLock()
	events, exists := p.events[eventType]
	p.rw.RUnlock()

	if !exists {
		return
	}

	events.rw.RLock()
	defer events.rw.RUnlock()

	// 优先级排序
	var subscribeIds = make([]SubscribeId, 0, len(events.priorities))
	for subscribeId := range events.subscribeIds {
		subscribeIds = append(subscribeIds, subscribeId)
	}

	sort.Slice(subscribeIds, func(i, j int) bool {
		return events.priorities[subscribeIds[i]] < events.priorities[subscribeIds[j]] // 优先级越小优先级越高
	})

	// 发布事件
	for _, subscribeId := range subscribeIds {
		subscriber := events.subscribers[events.subscribeIds[subscribeId]]
		subscriber.Tell(message.event)
	}
}
