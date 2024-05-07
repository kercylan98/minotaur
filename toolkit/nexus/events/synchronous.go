package events

import (
	"context"
	"github.com/kercylan98/minotaur/toolkit/constraints"
	"github.com/kercylan98/minotaur/toolkit/nexus"
	"runtime/debug"
	"time"
)

type (
	SynchronousHandler func(context.Context)
)

func Synchronous[I constraints.Ordered, T comparable](handler SynchronousHandler, opts ...*nexus.EventOptions) nexus.Event[I, T] {
	var opt *nexus.EventOptions
	var stack []byte
	if len(opts) > 0 {
		opt = opts[0]
		for i := 1; i < len(opts); i++ {
			opt.Apply(opts[i])
		}
		if opt != nil && opt.LowHandlerTrace && opt.LowHandlerTraceHandler != nil {
			stack = debug.Stack()
		}
	}
	return &synchronous[I, T]{
		handler: handler,
		opt:     opt,
		stack:   stack,
	}
}

type synchronous[I constraints.Ordered, T comparable] struct {
	ctx     context.Context
	handler SynchronousHandler
	opt     *nexus.EventOptions
	stack   []byte
}

func (s *synchronous[I, T]) OnInitialize(ctx context.Context, broker nexus.Broker[I, T]) {
	s.ctx = ctx
}

func (s *synchronous[I, T]) OnPublished(topic T, queue nexus.Queue[I, T]) {

}

func (s *synchronous[I, T]) OnProcess(topic T, queue nexus.Queue[I, T], startAt time.Time) {
	defer func() {
		var cost = time.Since(startAt)
		if s.opt != nil && cost >= s.opt.LowHandlerThreshold {
			if s.opt.LowHandlerTrace && s.opt.LowHandlerTraceHandler != nil {
				s.opt.LowHandlerTraceHandler(cost, s.stack)
			} else if s.opt.LowHandlerThresholdHandler != nil {
				s.opt.LowHandlerThresholdHandler(cost)
			}
		}
	}()

	s.handler(s.ctx)
}

func (s *synchronous[I, T]) OnProcessed(topic T, queue nexus.Queue[I, T], endAt time.Time) {
}
