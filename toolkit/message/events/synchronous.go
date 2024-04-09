package events

import (
	"context"
	"github.com/kercylan98/minotaur/toolkit/message"
	"time"
)

type (
	SynchronousHandler func(context.Context)
)

func Synchronous[I, T comparable](handler SynchronousHandler) message.Event[I, T] {
	return &synchronous[I, T]{
		handler: handler,
	}
}

type synchronous[I, T comparable] struct {
	ctx     context.Context
	handler SynchronousHandler
}

func (s *synchronous[I, T]) OnInitialize(ctx context.Context, broker message.Broker[I, T]) {
	s.ctx = ctx
}

func (s *synchronous[I, T]) OnPublished(topic T, queue message.Queue[I, T]) {

}

func (s *synchronous[I, T]) OnProcess(topic T, queue message.Queue[I, T], startAt time.Time) {
	s.handler(s.ctx)
}

func (s *synchronous[I, T]) OnProcessed(topic T, queue message.Queue[I, T], endAt time.Time) {
}
