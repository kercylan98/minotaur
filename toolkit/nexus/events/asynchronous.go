package events

import (
	"context"
	"github.com/kercylan98/minotaur/toolkit/constraints"
	"github.com/kercylan98/minotaur/toolkit/nexus"
	"time"
)

type (
	AsynchronousActuator        func(context.Context, func(context.Context)) // 负责执行异步消息的执行器
	AsynchronousHandler         func(context.Context) error                  // 异步消息逻辑处理器
	AsynchronousCallbackHandler func(context.Context, error)                 // 异步消息回调处理器
)

// Asynchronous 创建一个异步消息实例，并指定相应的处理器。
// 该函数接收以下参数：
//   - broker：消息所属的 Broker 实例。
//   - actuator：异步消息的执行器，负责执行异步消息的逻辑，当该参数为空时，将会使用默认的 go func()。
//   - handler：异步消息的逻辑处理器，用于执行实际的异步消息处理逻辑，可选参数。
//   - callback：异步消息的回调处理器，处理消息处理完成后的回调操作，可选参数。
//   - afterHandler：异步消息执行完成后的处理器，用于进行后续的处理操作，可选参数。
//
// 该函数除了 handler，其他所有处理器均为同步执行
//
// 返回值为一个实现了 Event 接口的异步消息实例。
func Asynchronous[I constraints.Ordered, T comparable](
	actuator AsynchronousActuator,
	handler AsynchronousHandler,
	callback AsynchronousCallbackHandler,
) nexus.Event[I, T] {
	m := &asynchronous[I, T]{
		actuator: actuator,
		handler:  handler,
		callback: callback,
	}
	if m.actuator == nil {
		m.actuator = func(ctx context.Context, f func(context.Context)) {
			go f(ctx)
		}
	}

	return m
}

type asynchronous[I constraints.Ordered, T comparable] struct {
	ctx      context.Context
	broker   nexus.Broker[I, T]
	actuator AsynchronousActuator
	handler  AsynchronousHandler
	callback AsynchronousCallbackHandler
}

func (s *asynchronous[I, T]) OnInitialize(ctx context.Context, broker nexus.Broker[I, T]) {
	s.ctx = ctx
	s.broker = broker
}

func (s *asynchronous[I, T]) OnPublished(topic T, queue nexus.Queue[I, T]) {
	queue.IncrementCustomMessageCount(topic, 1)
}

func (s *asynchronous[I, T]) OnProcess(topic T, queue nexus.Queue[I, T], startAt time.Time) {
	s.actuator(s.ctx, func(ctx context.Context) {
		var err error
		if s.handler != nil {
			err = s.handler(s.ctx)
		}

		s.broker.Publish(topic, Synchronous[I, T](func(ctx context.Context) {
			if s.callback != nil {
				s.callback(ctx, err)
			}
		}))
	})
}

func (s *asynchronous[I, T]) OnProcessed(topic T, queue nexus.Queue[I, T], endAt time.Time) {
	queue.IncrementCustomMessageCount(topic, -1)
}
