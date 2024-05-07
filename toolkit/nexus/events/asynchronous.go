package events

import (
	"context"
	"errors"
	"fmt"
	"github.com/kercylan98/minotaur/toolkit/constraints"
	"github.com/kercylan98/minotaur/toolkit/nexus"
	"runtime/debug"
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
	opts ...*nexus.EventOptions,
) nexus.Event[I, T] {
	var opt *nexus.EventOptions
	var stack []byte
	if len(opts) > 0 {
		opt = opts[0]
		for i := 1; i < len(opts); i++ {
			opt.Apply(opts[i])
		}
		if opt != nil && ((opt.LowHandlerTrace && opt.LowHandlerTraceHandler != nil) || (opt.DeadLockThreshold > 0 && opt.DeadLockThresholdHandler != nil)) {
			stack = debug.Stack()
		}
		if opt != nil && len(opt.ParentStack) > 0 {
			stack = append([]byte(fmt.Sprintf("parent stack: \n%s\n", opt.ParentStack)), stack...)
			opt.ParentStack = nil
		}
	}
	m := &asynchronous[I, T]{
		actuator: actuator,
		handler:  handler,
		callback: callback,
		opt:      opt,
		stack:    stack,
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
	opt      *nexus.EventOptions
	stack    []byte
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
		if s.opt != nil && s.opt.DeadLockThreshold > 0 && s.opt.DeadLockThresholdHandler != nil {
			ctx, cancel := context.WithTimeout(s.ctx, s.opt.DeadLockThreshold)
			defer cancel()
			go func(ctx context.Context) {
				select {
				case <-ctx.Done():
					if errors.Is(ctx.Err(), context.DeadlineExceeded) {
						defer func() {
							if err := recover(); err != nil {
								fmt.Println(fmt.Errorf("event dead lock panic: %v\n%s", err, debug.Stack()))
								return
							}
						}()
						s.opt.DeadLockThresholdHandler(s.stack)
					}
				}
			}(ctx)
		}

		var err error
		defer func() {
			queue.IncrementCustomMessageCount(topic, -1)

			var cost = time.Since(startAt)
			if s.opt != nil && s.opt.LowHandlerThreshold > 0 && cost >= s.opt.LowHandlerThreshold {
				if s.opt.LowHandlerTrace && s.opt.LowHandlerTraceHandler != nil {
					s.opt.LowHandlerTraceHandler(cost, s.stack)
				} else if s.opt.LowHandlerThresholdHandler != nil {
					s.opt.LowHandlerThresholdHandler(cost)
				}
			}

			if recoverErr := recover(); recoverErr != nil {
				switch recoverErr.(type) {
				case error:
					err = recoverErr.(error)
				default:
					err = fmt.Errorf("asynchronous panic: %v", recoverErr)
				}
			}

			s.broker.Publish(topic, Synchronous[I, T](
				func(ctx context.Context) {
					if s.callback != nil {
						s.callback(ctx, err)
					}
				},
				s.opt.WithParentStack(s.stack),
			))
		}()
		if s.handler != nil {
			err = s.handler(s.ctx)
		}
	})
}

func (s *asynchronous[I, T]) OnProcessed(topic T, queue nexus.Queue[I, T], endAt time.Time) {

}
