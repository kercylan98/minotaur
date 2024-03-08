package modular

import (
	"fmt"
	"github.com/kercylan98/minotaur/utils/log"
	"sync"
)

var application *modular

func init() {
	application = &modular{}
}

type modular struct {
	registerServices []Service
	services         []*service
}

// RegisterServices 注册服务
func (m *modular) RegisterServices(s ...Service) {
	application.registerServices = append(application.registerServices, s...)
}

// Run 运行模块化应用程序
func Run() {
	m := application
	var names = make(map[string]bool)
	for i := 0; i < len(m.registerServices); i++ {
		s := newService(m.registerServices[i])
		if names[s.name] {
			panic(fmt.Errorf("service %s is already registered", s.name))
		}
		names[s.name] = true
		m.services = append(m.services, s)
	}

	// OnInit
	for i := 0; i < len(m.services); i++ {
		s := m.services[i]
		s.instance.OnInit()
		log.Info(fmt.Sprintf("service %s initialized", s.name))
	}

	// OnPreload
	for i := 0; i < len(m.services); i++ {
		s := m.services[i]
		s.instance.OnPreload()
		log.Info(fmt.Sprintf("service %s preloaded", s.name))
	}

	// OnMount
	for i := 0; i < len(m.services); i++ {
		s := m.services[i]
		s.instance.OnMount()
	}

	// OnBlock
	var wait = new(sync.WaitGroup)
	for i := 0; i < len(m.services); i++ {
		s := m.services[i]
		if block, ok := s.instance.(Block); ok {
			wait.Add(1)
			go func(wait *sync.WaitGroup) {
				defer wait.Done()
				block.OnBlock()
			}(wait)
		}
	}
	wait.Wait()
	log.Info("all services exited")
}
