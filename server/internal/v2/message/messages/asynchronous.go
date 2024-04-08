package messages

import (
	"context"
	"github.com/kercylan98/minotaur/server/internal/v2/message"
	"github.com/kercylan98/minotaur/server/internal/v2/queue"
)

type (
	AsynchronousActuator[P message.Producer, Q message.Queue, B message.Broker[P, Q]]        func(context.Context, B, func(context.Context, B)) // 负责执行异步消息的执行器
	AsynchronousHandler[P message.Producer, Q message.Queue, B message.Broker[P, Q]]         func(context.Context, B) error                     // 异步消息逻辑处理器
	AsynchronousCallbackHandler[P message.Producer, Q message.Queue, B message.Broker[P, Q]] func(context.Context, B, error)                    // 异步消息回调处理器
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
// 返回值为一个实现了 Message 接口的异步消息实例。
func Asynchronous[P message.Producer, Q message.Queue, B message.Broker[P, Q]](
	broker B, producer P, queue Q,
	actuator AsynchronousActuator[P, Q, B],
	handler AsynchronousHandler[P, Q, B],
	callback AsynchronousCallbackHandler[P, Q, B],
) message.Message[P, Q] {
	m := &asynchronous[P, Q, B]{
		broker:   broker,
		producer: producer,
		queue:    queue,
		actuator: actuator,
		handler:  handler,
		callback: callback,
	}
	if m.actuator == nil {
		m.actuator = func(ctx context.Context, b B, f func(context.Context, B)) {
			go f(ctx, b)
		}
	}

	return m
}

type asynchronous[P message.Producer, Q message.Queue, B message.Broker[P, Q]] struct {
	broker   B
	producer P
	queue    Q
	ctx      context.Context
	actuator AsynchronousActuator[P, Q, B]
	handler  AsynchronousHandler[P, Q, B]
	callback AsynchronousCallbackHandler[P, Q, B]
}

func (s *asynchronous[P, Q, B]) OnPublished(controller queue.Controller) {
	controller.IncrementCustomMessageCount(1)
}

func (s *asynchronous[P, Q, B]) OnProcessed(controller queue.Controller) {
	controller.IncrementCustomMessageCount(-1)
}

func (s *asynchronous[P, Q, B]) OnInitialize(ctx context.Context) {
	s.ctx = ctx
}

func (s *asynchronous[P, Q, B]) OnProcess() {
	s.actuator(s.ctx, s.broker, func(ctx context.Context, broker B) {
		var err error
		if s.handler != nil {
			err = s.handler(s.ctx, s.broker)
		}

		broker.PublishMessage(Synchronous(broker, s.producer, s.queue, func(ctx context.Context, broker B) {
			if s.callback != nil {
				s.callback(ctx, broker, err)
			}
		}))
	})
}

func (s *asynchronous[P, Q, B]) GetProducer() P {
	return s.producer
}

func (s *asynchronous[P, Q, B]) GetQueue() Q {
	return s.queue
}
