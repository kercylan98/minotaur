package prc_test

import (
	"github.com/kercylan98/minotaur/experiment/internal/vivid/prc"
	"sync"
	"testing"
)

func TestShared(t *testing.T) {
	process := &TestSharedProcess{}
	process.Add(1)

	rc1 := prc.NewResourceController(":8080")
	rc2 := prc.NewResourceController(":8081")

	pid := prc.NewProcessId("/test", "127.0.0.1:8080")
	ref, _ := rc1.Register(pid, process)

	shared1 := prc.NewShared(rc1)
	shared2 := prc.NewShared(rc2)

	if err := shared1.Share(); err != nil {
		panic(err)
	}
	if err := shared2.Share(); err != nil {
		panic(err)
	}
	rc2.GetProcess(ref).DeliveryUserMessage(ref, nil, nil, pid)
	process.Wait()
}

type TestSharedProcess struct {
	wg sync.WaitGroup
}

func (p *TestSharedProcess) Initialize(rc *prc.ResourceController, id *prc.ProcessId) {

}

func (p *TestSharedProcess) DeliveryUserMessage(receiver, sender, forward *prc.ProcessRef, message prc.Message) {
	p.wg.Done()
}

func (p *TestSharedProcess) DeliverySystemMessage(receiver, sender, forward *prc.ProcessRef, message prc.Message) {
	p.wg.Done()
}

func (p *TestSharedProcess) IsTerminated() bool {
	return false
}

func (p *TestSharedProcess) Terminate(source *prc.ProcessRef) {

}
func (p *TestSharedProcess) Add(n int) {
	p.wg.Add(1)
}

func (p *TestSharedProcess) Wait() {
	p.wg.Wait()
}
