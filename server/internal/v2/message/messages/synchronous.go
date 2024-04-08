package messages

import (
	"context"
	"github.com/kercylan98/minotaur/server/internal/v2/message"
	"github.com/kercylan98/minotaur/server/internal/v2/queue"
)

type (
	SynchronousHandler[P message.Producer, Q message.Queue, B message.Broker[P, Q]] func(context.Context, B)
)

func Synchronous[P message.Producer, Q message.Queue, B message.Broker[P, Q]](
	broker B, producer P, queue Q,
	handler SynchronousHandler[P, Q, B],
) message.Message[P, Q] {
	return &synchronous[P, Q, B]{
		broker:   broker,
		producer: producer,
		queue:    queue,
		handler:  handler,
	}
}

type synchronous[P message.Producer, Q message.Queue, B message.Broker[P, Q]] struct {
	broker   B
	producer P
	queue    Q
	ctx      context.Context
	handler  SynchronousHandler[P, Q, B]
}

func (s *synchronous[P, Q, B]) OnPublished(controller queue.Controller) {

}

func (s *synchronous[P, Q, B]) OnProcessed(controller queue.Controller) {

}

func (s *synchronous[P, Q, B]) OnInitialize(ctx context.Context) {
	s.ctx = ctx
}

func (s *synchronous[P, Q, B]) OnProcess() {
	s.handler(s.ctx, s.broker)
}

func (s *synchronous[P, Q, B]) GetProducer() P {
	return s.producer
}

func (s *synchronous[P, Q, B]) GetQueue() Q {
	return s.queue
}
