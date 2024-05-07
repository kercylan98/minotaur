package nexus

import (
	"context"
	"github.com/kercylan98/minotaur/toolkit/constraints"
	"time"
)

type Event[I constraints.Ordered, T comparable] interface {
	// OnInitialize 消息初始化
	OnInitialize(ctx context.Context, broker Broker[I, T])

	// OnPublished 消息发布成功
	OnPublished(topic T, queue Queue[I, T])

	// OnProcess 消息开始处理
	OnProcess(topic T, queue Queue[I, T], startAt time.Time)

	// OnProcessed 消息处理完成
	OnProcessed(topic T, queue Queue[I, T], endAt time.Time)
}

type EventExecutor func()
type EventHandler[T comparable] func(topic T, event EventExecutor)
type EventFinisher[I constraints.Ordered, T comparable] func(topic T, last bool)

type EventInfo[I constraints.Ordered, T comparable] interface {
	GetTopic() T
	Exec(
		handler EventHandler[T],
		finisher EventFinisher[I, T],
	)
}

func (e EventExecutor) Exec() {
	e()
}
