package pulse

import (
	"github.com/kercylan98/minotaur/minotaur/vivid"
	"reflect"
	"sort"
	"time"
)

type (
	// SubscribeMessage 订阅事件消息
	SubscribeMessage struct {
		SubscribeId     SubscribeId    // 订阅 ID
		Event           EventType      // 订阅的事件
		Subscriber      Subscriber     // 事件订阅者
		Producer        *Producer      // 生产者过滤
		Priority        *int64         // 优先级，值越小优先级越高
		PriorityTimeout *time.Duration // 优先级事件超时时间
	}

	// UnsubscribeMessage 取消订阅事件消息
	UnsubscribeMessage struct {
		SubscribeId SubscribeId // 订阅 ID
	}

	// PublishMessage 发布事件消息
	PublishMessage struct {
		Producer Producer // 事件生产者
		Event    Event    // 事件
	}
)

// EventBusActor 基于 Actor 模型实现的事件总线
type EventBusActor struct {
	subscribes     map[EventType]map[SubscribeId]subscribeInfo
	subscribeTypes map[SubscribeId]EventType
}

func (e *EventBusActor) OnReceive(ctx vivid.MessageContext) {
	switch m := ctx.GetMessage().(type) {
	case vivid.OnOptionApply[*EventBusActor]:
		m.Options.WithMailboxFactory(vivid.PriorityMailboxFactoryId)
	case vivid.OnPreStart:
		e.onPreStart()
	case SubscribeMessage:
		e.onSubscribe(m)
	case UnsubscribeMessage:
		e.onUnsubscribe(m)
	case PublishMessage:
		e.onPublish(ctx, m)
	}
}

func (e *EventBusActor) onPreStart() {
	e.subscribes = make(map[EventType]map[SubscribeId]subscribeInfo)
	e.subscribeTypes = make(map[SubscribeId]EventType)
}

func (e *EventBusActor) onSubscribe(m SubscribeMessage) {
	info := subscribeInfo{
		subscribeId: m.SubscribeId,
		subscriber:  m.Subscriber,
		priority:    m.Priority,
		producer:    m.Producer,
	}

	subscribeMap, exists := e.subscribes[m.Event]
	if !exists {
		subscribeMap = make(map[SubscribeId]subscribeInfo)
		e.subscribes[m.Event] = subscribeMap
	}

	subscribeMap[m.SubscribeId] = info
}

func (e *EventBusActor) onUnsubscribe(m UnsubscribeMessage) {
	eventType := e.subscribeTypes[m.SubscribeId]
	if eventType == nil {
		return
	}

	subscribeMap, exists := e.subscribes[eventType]
	if !exists {
		return
	}

	delete(subscribeMap, m.SubscribeId)
	delete(e.subscribeTypes, m.SubscribeId)
	if len(subscribeMap) == 0 {
		delete(e.subscribes, eventType)
	}
}

func (e *EventBusActor) onPublish(ctx vivid.MessageContext, m PublishMessage) {
	subscribeMap, exists := e.subscribes[reflect.TypeOf(m.Event)]
	if !exists {
		return
	}

	var producerActorId = vivid.GetActorIdByActorRef(m.Producer)
	var subscribeNoPriorityList []subscribeInfo
	var subscribePriorityList []subscribeInfo
	for _, info := range subscribeMap {
		if info.producer != nil && producerActorId != m.Producer.Id() {
			continue
		}
		if info.priority == nil {
			subscribeNoPriorityList = append(subscribeNoPriorityList, info)
		} else {
			subscribePriorityList = append(subscribePriorityList, info)
		}
	}

	// 优先处理无优先级的订阅者
	for _, info := range subscribeNoPriorityList {
		info.subscriber.Tell(m.Event)
	}

	if len(subscribePriorityList) != 0 {
		// 按照优先级排序，值越小优先级越高
		sort.Slice(subscribePriorityList, func(i, j int) bool {
			return *subscribePriorityList[i].priority < *subscribePriorityList[j].priority
		})

		// 处理优先级订阅者
		vivid.ActorOf(ctx, vivid.NewActorOptions[*priorityEventActor]().WithConstruct(func() *priorityEventActor {
			return &priorityEventActor{
				subscribes: subscribePriorityList,
			}
		}())).Tell(priorityEventMessage{event: m.Event})
	}
}
