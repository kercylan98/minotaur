package pulse

import "reflect"

// eventSubscribeMessage 订阅事件消息
type eventSubscribeMessage struct {
	producer    Producer     // 事件生产者
	subscriber  Subscriber   // 事件订阅者
	event       reflect.Type // 订阅的事件
	subscribeId SubscribeId  // 订阅 ID
	priority    int          // 优先级，值越小优先级越高
}

// eventUnsubscribeMessage 取消订阅事件消息
type eventUnsubscribeMessage struct {
	producer    Producer     // 事件生产者
	subscriber  Subscriber   // 事件订阅者
	event       reflect.Type // 订阅的事件
	subscribeId SubscribeId  // 订阅 ID
}

// eventPublishMessage 发布事件消息
type eventPublishMessage struct {
	producer Producer // 事件生产者
	event    Event    // 事件
}
