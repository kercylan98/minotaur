package pulse

import (
	"github.com/kercylan98/minotaur/minotaur/vivid"
	"math"
	"reflect"
	"sync/atomic"
)

// NewPulse 创建一个新的事件总线
func NewPulse(system *vivid.ActorSystem, options ...*vivid.ActorOptions[*EventBusActor]) Pulse {
	pulse := Pulse{
		busActor: vivid.ActorOf[*EventBusActor](system, options...),
		eventSeq: new(atomic.Int64),
	}
	return pulse
}

// Pulse 基于 vivid.Actor 模型实现的事件总线
type Pulse struct {
	busActor vivid.ActorRef
	eventSeq *atomic.Int64 // 全局事件序号，仅标记事件产生顺序
}

func (p *Pulse) Subscribe(subscribeId SubscribeId, subscriber Subscriber, event Event, options ...SubscribeOption) {
	opts := new(SubscribeOptions).apply(options)
	p.busActor.Tell(SubscribeMessage{
		Producer:        opts.Producer,
		Subscriber:      subscriber,
		Event:           reflect.TypeOf(event),
		SubscribeId:     subscribeId,
		Priority:        opts.Priority,
		PriorityTimeout: opts.PriorityTimeout,
	}, vivid.WithPriority(math.MinInt64), vivid.WithInstantly(true))
}

func (p *Pulse) Unsubscribe(subscribeId SubscribeId) {
	p.busActor.Tell(UnsubscribeMessage{
		SubscribeId: subscribeId,
	}, vivid.WithPriority(math.MinInt64), vivid.WithInstantly(true))
}

func (p *Pulse) Publish(producer Producer, event Event) {
	p.busActor.Tell(PublishMessage{
		Producer: producer,
		Event:    event,
	}, vivid.WithPriority(p.eventSeq.Add(1)))
}
