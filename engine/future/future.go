package future

import (
	"github.com/kercylan98/minotaur/engine/prc"
	"sync"
	"sync/atomic"
	"time"
)

type Future[M prc.Message] interface {
	// Ref 返回该 Future 的 ActorRef
	Ref() *prc.ProcessRef

	// Result 阻塞地等待结果
	Result() (M, error)

	// OnlyResult 阻塞地等待结果，不关心错误，如果发送错误将会返回空指针
	OnlyResult() M

	// AssertResult 阻塞地等待结果，当发生错误时将会引发 panic
	AssertResult() M

	// Wait 阻塞的等待结果，该方式不关心结果，仅关心是否成功
	Wait() error

	// AssertWait 阻塞的等待结果，该方式不关心结果，仅关心是否成功，当发生错误时将会引发 panic
	AssertWait()

	// Forward 将结果转发给其他的 ActorRef
	Forward(refs ...*prc.ProcessRef)

	// Close 提前关闭
	Close(reason error)

	// AwaitForward 异步地等待阻塞结束后向目标 Actor 转发消息
	AwaitForward(ref *prc.ProcessRef, f func() M)
}

func New[M prc.Message](rc *prc.ResourceController, id *prc.ProcessId, timeout time.Duration) Future[M] {
	fp := &futureProcess[M]{
		done:    make(chan struct{}),
		timeout: timeout,
	}

	fp.ref, _ = rc.Register(id, fp)
	return fp
}

type futureProcess[M prc.Message] struct {
	rc            *prc.ResourceController
	ref           *prc.ProcessRef
	timer         *time.Timer
	done          chan struct{}
	message       any
	err           error
	timeout       time.Duration
	forwards      []*prc.ProcessRef
	closed        atomic.Bool
	forwardsMutex sync.Mutex
}

func (f *futureProcess[M]) OnlyResult() (m M) {
	result, err := f.Result()
	if err != nil {
		return m
	}
	return result
}

func (f *futureProcess[M]) AwaitForward(ref *prc.ProcessRef, asyncFunc func() M) {
	f.Forward(ref)
	go func() {
		if reason := recover(); reason != nil {
			f.rc.GetProcess(ref).DeliveryUserMessage(ref, f.ref, nil, reason)
		}
		m := asyncFunc()
		f.rc.GetProcess(ref).DeliveryUserMessage(ref, f.ref, nil, m)
	}()
}

func (f *futureProcess[M]) Ref() *prc.ProcessRef {
	return f.ref
}

func (f *futureProcess[M]) Result() (M, error) {
	<-f.done
	return f.message.(M), f.err
}

func (f *futureProcess[M]) AssertResult() M {
	result, err := f.Result()
	if err != nil {
		panic(err)
	}
	return result
}

func (f *futureProcess[M]) Wait() error {
	_, err := f.Result()
	return err
}

func (f *futureProcess[M]) AssertWait() {
	if err := f.Wait(); err != nil {
		panic(err)
	}
}

func (f *futureProcess[M]) Forward(refs ...*prc.ProcessRef) {
	f.forwardsMutex.Lock()
	defer f.forwardsMutex.Unlock()
	f.forwards = append(f.forwards, refs...)
	if f.closed.Load() {
		f.execForward()
	}
}

func (f *futureProcess[M]) Initialize(rc *prc.ResourceController, id *prc.ProcessId) {
	f.rc = rc
	if f.timeout > 0 {
		f.timer = time.AfterFunc(f.timeout, func() {
			f.Close(ErrorFutureTimeout)
		})
	}
}

func (f *futureProcess[M]) DeliveryUserMessage(receiver, sender, forward *prc.ProcessRef, message prc.Message) {
	if f.closed.Load() {
		return
	}

	switch m := message.(type) {
	case error:
		f.Close(m)
	default:
		f.message = message
		f.Close(nil)
	}
}

func (f *futureProcess[M]) DeliverySystemMessage(receiver, sender, forward *prc.ProcessRef, message prc.Message) {
	f.DeliveryUserMessage(receiver, sender, forward, message)
}

func (f *futureProcess[M]) IsTerminated() bool {
	return f.closed.Load()
}

func (f *futureProcess[M]) Terminate(source *prc.ProcessRef) {
	// 不做什么
}

func (f *futureProcess[M]) Close(reason error) {
	if !f.closed.CompareAndSwap(false, true) {
		return
	}
	f.err = reason
	close(f.done)
	if f.timer != nil {
		f.timer.Stop()
	}
	f.rc.Unregister(f.ref, f.ref)
	f.forwardsMutex.Lock()
	defer f.forwardsMutex.Unlock()
	f.execForward()
}

func (f *futureProcess[M]) execForward() {
	if len(f.forwards) == 0 {
		return
	}

	var m prc.Message
	if f.err != nil {
		m = f.err
	}

	for _, ref := range f.forwards {
		f.rc.GetProcess(ref).DeliveryUserMessage(ref, f.ref, ref, m)
	}
	f.forwards = nil
}
