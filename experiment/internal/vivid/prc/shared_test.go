package prc_test

import (
	"fmt"
	"github.com/kercylan98/minotaur/experiment/internal/vivid/prc"
	"sync"
	"testing"
)

type Process struct {
	sync.WaitGroup
}

func (p *Process) Initialize(rc *prc.ResourceController, id *prc.ProcessId) {

}

func (p *Process) DeliveryUserMessage(receiver, sender, forward *prc.ProcessRef, message prc.Message) {
	fmt.Println("received", message)
	p.Done()
}

func (p *Process) DeliverySystemMessage(receiver, sender, forward *prc.ProcessRef, message prc.Message) {
	fmt.Println("received", message)
	p.Done()
}

func (p *Process) IsTerminated() bool {
	return false
}

func (p *Process) Terminate(source *prc.ProcessRef) {

}

func TestShared(t *testing.T) {
	rc1 := prc.NewResourceController("127.0.0.1:8080")
	pid := prc.NewProcessId("/test", "127.0.0.1:8080")
	process := &Process{}
	process.Add(1)
	rc1.Register(pid, process)
	rc2 := prc.NewResourceController("127.0.0.1:8081")

	shared1 := prc.NewShared(rc1)
	shared2 := prc.NewShared(rc2)

	if err := shared1.Share(); err != nil {
		panic(err)
	}
	if err := shared2.Share(); err != nil {
		panic(err)
	}
	rc2.GetProcess(prc.NewProcessRef(pid)).DeliveryUserMessage(prc.NewProcessRef(pid), nil, nil, pid)
	process.Wait()
}
