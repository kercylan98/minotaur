package eventstream

type Event = any

type Handler func(event Event)

type EventStream interface {
	// Subscribe 订阅事件流
	Subscribe(handler Handler) Subscription

	// Unsubscribe 取消订阅事件流
	Unsubscribe(subscription Subscription)

	// Publish 发布事件
	Publish(event Event)

	// Length 获取当前事件流的订阅数量
	Length() int
}

// Subscription 事件流订阅信息
type Subscription interface {
}
