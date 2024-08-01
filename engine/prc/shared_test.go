package prc_test

import (
	"errors"
	"github.com/kercylan98/minotaur/engine/prc"
	"sync"
	"testing"
)

func TestSharedRuntimeError(t *testing.T) {
	t.Run("restart", func(t *testing.T) {
		wg := new(sync.WaitGroup)
		wg.Add(2)
		rc1 := prc.NewResourceController(prc.FunctionalResourceControllerConfigurator(func(config *prc.ResourceControllerConfiguration) {
			config.WithPhysicalAddress("127.0.0.1:8080")
		}))
		shared1 := prc.NewShared(rc1, prc.FunctionalSharedConfigurator(func(config *prc.SharedConfiguration) {
			config.WithRuntimeErrorHandler(prc.FunctionalErrorPolicyDecisionHandler(func(err error) prc.SharedPolicyDecision {
				return prc.SharedPolicyDecisionRestart
			}))

			config.WithSharedHook(prc.FunctionalSharedStartHook(func() {
				wg.Done()
			}))
		}))
		if err := shared1.Share(); err != nil {
			panic(err)
		}
		shared1.Close(errors.New("test err"))
		wg.Wait()
		shared1.Close()
	})

	t.Run("stop", func(t *testing.T) {
		c := 0
		rc1 := prc.NewResourceController(prc.FunctionalResourceControllerConfigurator(func(config *prc.ResourceControllerConfiguration) {
			config.WithPhysicalAddress("127.0.0.1:8080")
		}))
		shared1 := prc.NewShared(rc1, prc.FunctionalSharedConfigurator(func(config *prc.SharedConfiguration) {
			config.WithRuntimeErrorHandler(prc.FunctionalErrorPolicyDecisionHandler(func(err error) prc.SharedPolicyDecision {
				return prc.SharedPolicyDecisionStop
			}))

			config.WithSharedHook(prc.FunctionalSharedStartHook(func() {
				c++
			}))
		}))
		if err := shared1.Share(); err != nil {
			panic(err)
		}
		shared1.Close(errors.New("test err"))
		if c != 1 {
			panic(c)
		}
	})

	t.Run("panic", func(t *testing.T) {
		defer func() {
			if r := recover(); r == nil {
				panic("not panic")
			}
		}()
		rc1 := prc.NewResourceController(prc.FunctionalResourceControllerConfigurator(func(config *prc.ResourceControllerConfiguration) {
			config.WithPhysicalAddress("127.0.0.1:8080")
		}))
		shared1 := prc.NewShared(rc1)
		if err := shared1.Share(); err != nil {
			panic(err)
		}
		shared1.Close(errors.New("test err"))

	})
}

func TestShared(t *testing.T) {
	process := &TestSharedProcess{}
	messageNum := 1000
	process.Add(messageNum)

	rc1 := prc.NewResourceController(prc.FunctionalResourceControllerConfigurator(func(config *prc.ResourceControllerConfiguration) {
		config.WithPhysicalAddress("127.0.0.1:8080")
	}))
	rc2 := prc.NewResourceController(prc.FunctionalResourceControllerConfigurator(func(config *prc.ResourceControllerConfiguration) {
		config.WithPhysicalAddress("127.0.0.1:8081")
	}))

	pid := prc.NewProcessId("127.0.0.1:8080", "/test")
	ref, _ := rc1.Register(pid, process)

	shared1 := prc.NewShared(rc1)
	shared2 := prc.NewShared(rc2)

	if err := shared1.Share(); err != nil {
		panic(err)
	}
	if err := shared2.Share(); err != nil {
		panic(err)
	}
	for i := 0; i < messageNum; i++ {
		rc2.GetProcess(ref).DeliveryUserMessage(ref, nil, nil, pid)
	}
	process.Wait()
}

type TestSharedProcess struct {
	wg sync.WaitGroup
}

func (p *TestSharedProcess) Initialize(rc *prc.ResourceController, id *prc.ProcessId) {

}

func (p *TestSharedProcess) DeliveryUserMessage(receiver, sender, forward *prc.ProcessId, message prc.Message) {
	p.wg.Done()
}

func (p *TestSharedProcess) DeliverySystemMessage(receiver, sender, forward *prc.ProcessId, message prc.Message) {
	p.DeliveryUserMessage(receiver, sender, forward, message)
}

func (p *TestSharedProcess) IsTerminated() bool {
	return false
}

func (p *TestSharedProcess) Terminate(source *prc.ProcessId) {

}
func (p *TestSharedProcess) Add(n int) {
	p.wg.Add(n)
}

func (p *TestSharedProcess) Wait() {
	p.wg.Wait()
}
