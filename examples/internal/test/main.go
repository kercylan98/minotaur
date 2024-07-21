package main

import (
	"fmt"
	"github.com/kercylan98/minotaur/engine/transport"
	"github.com/kercylan98/minotaur/engine/vivid"
	"time"
)

type TestServiceExpose interface {
	transport.FiberService

	Hi()
}

type TestService struct {
}

func (*TestService) Hi() {
	fmt.Println("hi")
}

func (t *TestService) OnInit(ctx vivid.ActorContext, app *transport.FiberApp) {
	app.ExposeService(transport.FiberExpose[TestServiceExpose](t))
	v := app.ProvideService(transport.FiberProvide[TestServiceExpose]())
	v.(TestServiceExpose).Hi()
	fmt.Println("init")
}

func main() {
	system := vivid.NewActorSystem()

	system.ActorOf(vivid.FunctionalActorProvider(func() vivid.Actor {
		return transport.NewFiber(":8888", transport.FunctionalFiberConfigurator(func(configuration *transport.FiberConfiguration) {
			configuration.WithServices(new(TestService))
		}))
	}))

	time.Sleep(time.Hour)
}
