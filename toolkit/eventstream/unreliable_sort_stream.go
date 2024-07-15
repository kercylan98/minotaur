package eventstream

import (
	"sync"
	"sync/atomic"
)

func NewUnreliableSortStream() *UnreliableSortStream {
	return &UnreliableSortStream{}
}

var _ Stream = &UnreliableSortStream{}
var _ Subscription = &UnreliableSortStreamSubscription{}

// UnreliableSortStream 这是一个并发安全的事件流，在订阅事件流后将会收到所有传入该流的事件，事件的执行是在发布者的 goroutine 中执行的
//   - 正常情况下，订阅的执行是按照订阅的顺序执行的，当发生取消订阅的时候，顺序可能会发生变化
//   - 事件流的订阅会在生产者的 goroutine 中执行，订阅方在订阅中更改自身的状态是不安全的
type UnreliableSortStream struct {
	rw            sync.RWMutex
	subscriptions []*UnreliableSortStreamSubscription // 当前所有的订阅信息
}

func (s *UnreliableSortStream) Subscribe(handler Handler) Subscription {
	if handler == nil {
		return nil
	}

	sub := &UnreliableSortStreamSubscription{
		handler: handler,
		active:  1,
	}

	s.rw.Lock()
	defer s.rw.Unlock()

	sub.id = int32(len(s.subscriptions))
	s.subscriptions = append(s.subscriptions, sub)

	return sub
}

func (s *UnreliableSortStream) Unsubscribe(subscription Subscription) {
	sub, ok := subscription.(*UnreliableSortStreamSubscription)
	if !ok {
		return
	}

	if atomic.CompareAndSwapUint32(&sub.active, 1, 0) {
		s.rw.Lock()
		defer s.rw.Unlock()

		lastIdx := len(s.subscriptions) - 1
		s.subscriptions[sub.id] = s.subscriptions[lastIdx]
		s.subscriptions[sub.id].id = sub.id
		s.subscriptions = s.subscriptions[:lastIdx]

		if len(s.subscriptions) == 0 {
			s.subscriptions = nil
		}
	}
}

func (s *UnreliableSortStream) Publish(event Event) {
	s.rw.RLock()
	subs := make([]*UnreliableSortStreamSubscription, len(s.subscriptions))
	copy(subs, s.subscriptions)
	s.rw.RUnlock()

	for _, sub := range subs {
		sub.handler(event)
	}
}

func (s *UnreliableSortStream) Length() int {
	s.rw.RLock()
	defer s.rw.RUnlock()

	return len(s.subscriptions)
}

type UnreliableSortStreamSubscription struct {
	id      int32
	handler Handler
	active  uint32
}
