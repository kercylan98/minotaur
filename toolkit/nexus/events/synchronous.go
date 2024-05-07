package events

import (
	"context"
	"github.com/kercylan98/minotaur/toolkit/constraints"
	"github.com/kercylan98/minotaur/toolkit/nexus"
	"time"
)

type (
	SynchronousHandler func(context.Context)
)

func Synchronous[I constraints.Ordered, T comparable](handler SynchronousHandler) nexus.Event[I, T] {
	return &synchronous[I, T]{
		handler: handler,
	}
}

type synchronous[I constraints.Ordered, T comparable] struct {
	ctx     context.Context
	handler SynchronousHandler
}

func (s *synchronous[I, T]) OnInitialize(ctx context.Context, broker nexus.Broker[I, T]) {
	s.ctx = ctx
}

func (s *synchronous[I, T]) OnPublished(topic T, queue nexus.Queue[I, T]) {

}

func (s *synchronous[I, T]) OnProcess(topic T, queue nexus.Queue[I, T], startAt time.Time) {
	s.handler(s.ctx)
}

func (s *synchronous[I, T]) OnProcessed(topic T, queue nexus.Queue[I, T], endAt time.Time) {
}
