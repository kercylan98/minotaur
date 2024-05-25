package pulse

import (
	"github.com/kercylan98/minotaur/vivid"
	"math"
	"reflect"
	"sync/atomic"
)

var eventSeq atomic.Int64 // 全局事件序号，仅标记事件产生顺序

// Subscribe 订阅消息总线中来自特定生产者的特定事件，该事件可在 SubscribeId 不同的情况下重复订阅
//   - 由于订阅的过程是异步的，订阅不会立即生效，而是在下一个事件循环中生效
func Subscribe[T Event](eventBus vivid.ActorRef, producer Producer, subscriber Subscriber, subscribeId SubscribeId) {
	eventBus.Tell(eventSubscribeMessage{
		producer:    producer,
		subscriber:  subscriber,
		event:       reflect.TypeOf((*T)(nil)).Elem(),
		subscribeId: subscribeId,
		priority:    0,
	}, vivid.WithPriority(math.MinInt64), vivid.WithInstantly(true))
}

// Unsubscribe 取消订阅消息总线中来自特定生产者的特定事件
//   - 由于取消订阅的过程是异步的，取消订阅不会立即生效，而是在下一个事件循环中生效，例如可能期望在收到第一个事件后取消订阅，但实际上可能会收到多个事件后才取消订阅。这是由于在取消订阅的过程中已经产生了多个事件并已经投递到了订阅者的邮箱中。
//   - 如要确保取消订阅的实时性，可在订阅者中实现过滤器。
func Unsubscribe[T Event](eventBus vivid.ActorRef, producer Producer, subscriber Subscriber, subscribeId SubscribeId) {
	eventBus.Tell(eventUnsubscribeMessage{
		producer:    producer,
		subscriber:  subscriber,
		event:       reflect.TypeOf((*T)(nil)).Elem(),
		subscribeId: subscribeId,
	}, vivid.WithPriority(math.MinInt64), vivid.WithInstantly(true))
}

// Publish 发布事件到消息总线，消息总线会将事件投递给所有订阅者
func Publish(eventBus vivid.ActorRef, producer Producer, event Event) {
	eventBus.Tell(eventPublishMessage{
		producer: producer,
		event:    event,
	}, vivid.WithPriority(eventSeq.Add(1)))
}
