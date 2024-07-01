package vivid

import (
	"github.com/kercylan98/minotaur/core"
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

func NewFuture(system *ActorSystem, timeout time.Duration) Future {
	systemAddress := system.processes.Address()
	f := &future{
		actorSystem: system,
		address:     core.NewAddress(systemAddress.Network(), system.name, systemAddress.Host(), systemAddress.Port(), futurePrefix+convert.Uint64ToString(system.nextFutureId.Add(1))),
		done:        make(chan struct{}),
	}
	f.forwards = f.forwards[:0]
	f.done = make(chan struct{})

	ref, exist := system.processes.Register(f)
	if exist {
		panic("future process already exist")
	}
	f.ref = ref
	if timeout > 0 {
		time.AfterFunc(timeout, f.onTimeout)
	}
	return f
}

type Future interface {
	Ref() ActorRef
	Result() (Message, error)
	Wait() error
	Forward(refs ...ActorRef)
}

type FutureForwardMessage struct {
	Message Message
	Error   error
}

type future struct {
	actorSystem   *ActorSystem
	address       core.Address
	ref           ActorRef
	ok            uint32
	done          chan struct{}
	result        Message
	err           error
	timer         *time.Timer
	forwards      []ActorRef
	forwardsMutex sync.Mutex
}

func (f *future) Ref() ActorRef {
	return f.ref
}

func (f *future) Forward(refs ...ActorRef) {
	f.forwardsMutex.Lock()
	defer f.forwardsMutex.Unlock()
	ok := atomic.LoadUint32(&f.ok) == 1
	f.forwards = append(f.forwards, refs...)
	if ok {
		f.execForward()
	}
}

func (f *future) execForward() {
	if len(f.forwards) == 0 {
		return
	}

	msg := FutureForwardMessage{
		Message: f.result, Error: f.err,
	}

	for _, ref := range f.forwards {
		f.actorSystem.sendUserMessage(f.ref, ref, msg)
	}
	f.forwards = nil
}

func (f *future) Result() (Message, error) {
	<-f.done
	return f.result, f.err
}

func (f *future) Wait() error {
	_, err := f.Result()
	return err
}

func (f *future) GetAddress() core.Address {
	return f.address
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
	if err, ok := f.result.(error); ok {
		f.err = err
		f.result = nil
	}

	f.actorSystem.processes.Unregister(f.ref)
	close(f.done)
	f.forwardsMutex.Lock()
	defer f.forwardsMutex.Unlock()
	f.execForward()
}
