package modular

import (
	"fmt"
	"github.com/kercylan98/minotaur/utils/log"
	"reflect"
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
	var tvm = make(map[reflect.Type]reflect.Value)
	for i := 0; i < len(m.registerServices); i++ {
		s := newService(m.registerServices[i])
		if names[s.name] {
			panic(fmt.Errorf("service %s is already registered", s.name))
		}
		names[s.name] = true
		tvm[s.vof.Type()] = s.vof
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
		log.Info(fmt.Sprintf("service %s mounted", s.name))
	}
}
