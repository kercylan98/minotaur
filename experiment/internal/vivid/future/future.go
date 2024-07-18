package future

import (
	"github.com/kercylan98/minotaur/experiment/internal/vivid/prc"
	"sync"
	"sync/atomic"
	"time"
)

type Future interface {
	// Ref 返回该 Future 的 ActorRef
	Ref() *prc.ProcessRef

	// Result 阻塞地等待结果
	Result() (prc.Message, error)

	// AssertResult 阻塞地等待结果，当发生错误时将会引发 panic
	AssertResult() prc.Message

	// Wait 阻塞的等待结果，该方式不关心结果，仅关心是否成功
	Wait() error

	// AssertWait 阻塞的等待结果，该方式不关心结果，仅关心是否成功，当发生错误时将会引发 panic
	AssertWait()

	// Forward 将结果转发给其他的 ActorRef
	Forward(refs ...*prc.ProcessRef)

	// Close 提前关闭
	Close(reason error)
}

func New(rc *prc.ResourceController, id *prc.ProcessId, timeout time.Duration) Future {
	fp := &futureProcess{
		done:    make(chan struct{}),
		timeout: timeout,
	}

	var exist bool
	fp.ref, exist = rc.Register(id, fp)
	if exist {
		panic("future process already exist")
	}
	return fp
}

type futureProcess struct {
	rc            *prc.ResourceController
	ref           *prc.ProcessRef
	done          chan struct{}
	closed        atomic.Bool
	timeout       time.Duration
	message       prc.Message
	err           error
	timer         *time.Timer
	forwards      []*prc.ProcessRef
	forwardsMutex sync.Mutex
}

func (f *futureProcess) Ref() *prc.ProcessRef {
	return f.ref
}

func (f *futureProcess) Result() (prc.Message, error) {
	<-f.done
	return f.message, f.err
}

func (f *futureProcess) AssertResult() prc.Message {
	result, err := f.Result()
	if err != nil {
		panic(err)
	}
	return result
}

func (f *futureProcess) Wait() error {
	_, err := f.Result()
	return err
}

func (f *futureProcess) AssertWait() {
	if err := f.Wait(); err != nil {
		panic(err)
	}
}

func (f *futureProcess) Forward(refs ...*prc.ProcessRef) {
	f.forwardsMutex.Lock()
	defer f.forwardsMutex.Unlock()
	f.forwards = append(f.forwards, refs...)
	if f.closed.Load() {
		f.execForward()
	}
}

func (f *futureProcess) Initialize(rc *prc.ResourceController, id *prc.ProcessId) {
	f.rc = rc
	if f.timeout > 0 {
		f.timer = time.AfterFunc(f.timeout, func() {
			f.Close(ErrorFutureTimeout)
		})
	}
}

func (f *futureProcess) DeliveryUserMessage(sender, forward *prc.ProcessRef, message prc.Message) {
	if f.closed.Load() {
		return
	}
	switch m := f.message.(type) {
	case error:
		f.Close(m)
	default:
		f.message = message
		f.Close(nil)
	}
}

func (f *futureProcess) DeliverySystemMessage(sender, forward *prc.ProcessRef, message prc.Message) {
	f.DeliveryUserMessage(sender, forward, message)
}

func (f *futureProcess) IsTerminated() bool {
	return f.closed.Load()
}

func (f *futureProcess) Terminate(source *prc.ProcessRef) {
	close(f.done)
}

func (f *futureProcess) Close(reason error) {
	if !f.closed.CompareAndSwap(false, true) {
		return
	}
	f.err = reason
	f.timer.Stop()
	f.rc.Unregister(f.ref, f.ref)
	f.forwardsMutex.Lock()
	defer f.forwardsMutex.Unlock()
	f.execForward()
}

func (f *futureProcess) execForward() {
	if len(f.forwards) == 0 {
		return
	}

	var m = f.message
	if f.err != nil {
		m = f.err
	}

	for _, ref := range f.forwards {
		f.rc.GetProcess(ref).DeliveryUserMessage(f.ref, ref, m)
	}
	f.forwards = nil
}
