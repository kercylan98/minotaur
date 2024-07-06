package transport

import (
	"github.com/gofiber/fiber/v2"
	"github.com/kercylan98/minotaur/core/vivid"
	"github.com/kercylan98/minotaur/core/vivid/supervisorstategy"
	"reflect"
)

func NewFiber(addr string) *Fiber {
	return &Fiber{
		addr: addr,
		app:  fiber.New(),
	}
}

type Fiber struct {
	support  *vivid.ModuleSupport
	app      *fiber.App
	addr     string
	services []FiberService
}

func (n *Fiber) BindService(services ...FiberService) *Fiber {
	n.services = append(n.services, services...)
	return n
}

func (n *Fiber) OnLoad(support *vivid.ModuleSupport, hasTransportModule bool) {
	// field load
	n.support = support

	// init kit
	kit := &FiberKit{
		app:         n.app,
		actorSystem: support.System(),
	}

	// init services
	for _, service := range n.services {
		service.OnInit(kit)
	}

	// init actor
	actorType := reflect.TypeOf((*fiberActor)(nil)).Elem().Name()
	kit.ownerRef = n.support.System().ActorOf(func() vivid.Actor {
		return newFiberActor(n, kit, n.addr)
	}, func(options *vivid.ActorOptions) {
		options.WithNamePrefix(actorType)
		options.WithSupervisorStrategy(supervisorstategy.OneForOne(func(reason, message vivid.Message) vivid.Directive {
			return vivid.DirectiveStop
		}, 0))
	})
}
