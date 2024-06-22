package vivid

import (
	"github.com/google/uuid"
	"github.com/kercylan98/minotaur/minotaur/core"
	"sync"
	"sync/atomic"
	"time"
)

var (
	_ core.Process = &future{}
	_ Future       = &future{}
)

func generateFutureAddress(system *ActorSystem) core.Address {
	return core.NewAddress("", system.name, "", 0, "/future/"+uuid.NewString())
}

func NewFuture(system *ActorSystem, timeout time.Duration) Future {
	f := &future{
		actorSystem: system,
		address:     generateFutureAddress(system),
		cond:        sync.NewCond(new(sync.Mutex)),
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
	// Ref 返回该 Future 的 ActorRef
	Ref() ActorRef

	// Result 阻塞并等待结果
	Result() (Message, error)

	// Wait 不关注结果，只等待是否完成
	Wait() error
}

type future struct {
	actorSystem *ActorSystem
	address     core.Address
	ref         ActorRef
	status      uint32
	cond        *sync.Cond
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

	f.timer = time.AfterFunc(timeout, f.onTimeout)
}

func (f *future) Result() (Message, error) {
	f.cond.L.Lock()
	defer f.cond.L.Unlock()
	for !f.Deaden() {
		f.cond.Wait()
	}
	return f.result, f.err
}

func (f *future) Wait() error {
	_, err := f.Result()
	return err
}

func (f *future) GetAddress() core.Address {
	return f.address
}

func (f *future) Deaden() bool {
	return atomic.LoadUint32(&f.status) == 1
}

func (f *future) Dead() {
	atomic.StoreUint32(&f.status, 1)
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
	f.Terminate(nil)
	f.err = ErrFutureTimeout
}

func (f *future) Terminate(_ *core.ProcessRef) {
	f.cond.L.Lock()
	if f.Deaden() {
		f.cond.L.Unlock()
		return
	}

	if f.timer != nil {
		f.timer.Stop()
	}
	f.Dead()
	f.actorSystem.processes.Unregister(f.ref)

	f.cond.L.Unlock()
	f.cond.Broadcast()
}
