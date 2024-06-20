package vivid

import (
	"reflect"
	"sort"
	"time"
)

type (
	// SubscribeMessage 订阅事件消息
	SubscribeMessage struct {
		Event           EventType      // 订阅的事件
		Subscriber      Subscriber     // 事件订阅者
		Producer        *Producer      // 生产者过滤
		Priority        *int64         // 优先级，值越小优先级越高
		PriorityTimeout *time.Duration // 优先级事件超时时间
	}

	// UnsubscribeMessage 取消订阅事件消息
	UnsubscribeMessage struct {
		Event      EventType  // 取消订阅的事件
		Subscriber Subscriber // 事件订阅者
	}

	// PublishMessage 发布事件消息
	PublishMessage struct {
		Producer Producer // 事件生产者
		Event    Event    // 事件
	}
)

// EventBusActor 基于 Actor 模型实现的事件总线
type EventBusActor struct {
	subscribes map[EventType]map[Subscriber]subscribeInfo // 不同事件类型的订阅者
}

func (e *EventBusActor) OnReceive(ctx MessageContext) {
	switch m := ctx.GetMessage().(type) {
	case OnInit[*EventBusActor]:
		m.Options.WithMailboxFactory(PriorityMailboxFactoryId)
	case OnBoot:
		e.onBoot()
	case SubscribeMessage:
		e.onSubscribe(m)
	case UnsubscribeMessage:
		e.onUnsubscribe(m)
	case PublishMessage:
		e.onPublish(ctx, m)
	}
}

func (e *EventBusActor) onBoot() {
	e.subscribes = make(map[EventType]map[Subscriber]subscribeInfo)
}

func (e *EventBusActor) onSubscribe(m SubscribeMessage) {
	info := subscribeInfo{
		subscriber:      m.Subscriber,
		priority:        m.Priority,
		producer:        m.Producer,
		priorityTimeout: m.PriorityTimeout,
	}

	subscribeMap, exists := e.subscribes[m.Event]
	if !exists {
		subscribeMap = make(map[Subscriber]subscribeInfo)
		e.subscribes[m.Event] = subscribeMap
	}

	subscribeMap[m.Subscriber] = info
}

func (e *EventBusActor) onUnsubscribe(m UnsubscribeMessage) {
	subscribe := e.subscribes[m.Event]
	if subscribe == nil {
		return
	}

	delete(subscribe, m.Subscriber)
}

func (e *EventBusActor) onPublish(ctx MessageContext, m PublishMessage) {
	// 获取订阅者集合
	subscribeMap, exists := e.subscribes[reflect.TypeOf(m.Event)]
	if !exists {
		return
	}

	// 获取生产者 ActorId
	var producerActorId = m.Producer.Id()
	var subscribeNoPriorityList []subscribeInfo
	var subscribePriorityList []subscribeInfo
	for _, info := range subscribeMap {
		if info.producer != nil && producerActorId != m.Producer.Id() {
			continue // 过滤生产者，部分订阅者只接收指定生产者的事件
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
		eventActor := ActorOf(ctx, NewActorOptions[*priorityEventActor]().WithConstruct(func() *priorityEventActor {
			return &priorityEventActor{
				subscribes: subscribePriorityList,
			}
		}()))
		eventActor.Tell(priorityEventMessage{event: m.Event})
		eventActor.Tell(OnTerminate{})
	}
}
