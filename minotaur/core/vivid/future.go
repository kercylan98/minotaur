package vivid

import (
	"github.com/kercylan98/minotaur/minotaur/core"
	"github.com/kercylan98/minotaur/toolkit/convert"
	"sync"
	"sync/atomic"
	"time"
)

const (
	futurePrefix = "/future/"
)

var (
	_ core.Process = &future{}
	_ Future       = &future{}
)

var futurePool = sync.Pool{
	New: func() interface{} {
		return &future{}
	},
}

func NewFuture(system *ActorSystem, timeout time.Duration) Future {
	f := futurePool.Get().(*future)
	f.actorSystem = system
	f.address = core.NewAddress("", system.name, "", 0, futurePrefix+convert.Uint64ToString(system.nextFutureId.Add(1)))
	atomic.StoreUint32(&f.ok, 0)
	f.result = nil
	f.err = nil
	if cap(f.done) == 0 {
		f.done = make(chan struct{}, 1)
	} else {
		select {
		case <-f.done:
		default:
		}
	}
	ref, exist := system.processes.Register(f)
	if exist {
		panic("future process already exist")
	}
	f.ref = ref
	f.bindTimeout(timeout)
	return f
}

type Future interface {
	Ref() ActorRef
	Result() (Message, error)
	Wait() error
}

type future struct {
	actorSystem *ActorSystem
	address     core.Address
	ref         ActorRef
	ok          uint32
	done        chan struct{}
	result      Message
	err         error
	timer       *time.Timer
}

func (f *future) Ref() ActorRef {
	return f.ref
}

func (f *future) bindTimeout(timeout time.Duration) {
	if timeout <= 0 {
		return
	}
	if f.timer == nil {
		f.timer = time.AfterFunc(timeout, f.onTimeout)
	} else {
		f.timer.Reset(timeout)
	}
}

func (f *future) Result() (Message, error) {
	select {
	case <-f.done:
		return f.result, f.err
	}
}

func (f *future) Wait() error {
	_, err := f.Result()
	return err
}

func (f *future) GetAddress() core.Address {
	return f.address
}

func (f *future) Deaden() bool {
	return atomic.LoadUint32(&f.ok) == 1
}

func (f *future) Dead() {
	// No-op
}

func (f *future) SendUserMessage(sender *core.ProcessRef, message core.Message) {
	f.result = message
	f.Terminate(nil)
}

func (f *future) SendSystemMessage(sender *core.ProcessRef, message core.Message) {
	f.result = message
	f.Terminate(nil)
}

func (f *future) onTimeout() {
	f.err = ErrFutureTimeout
	f.Terminate(nil)
}

func (f *future) Terminate(_ *core.ProcessRef) {
	if !atomic.CompareAndSwapUint32(&f.ok, 0, 1) {
		return
	}

	if f.timer != nil {
		f.timer.Stop()
	}
	f.actorSystem.processes.Unregister(f.ref)
	select {
	case f.done <- struct{}{}:
	default:
	}
	futurePool.Put(f)
}
