package transport

import (
	"github.com/gofiber/fiber/v2"
	"github.com/kercylan98/minotaur/engine/prc"
	"github.com/kercylan98/minotaur/engine/vivid"
)

func NewFiber(addr prc.PhysicalAddress, configurator ...FiberConfigurator) *Fiber {
	f := &Fiber{
		config: newFiberConfiguration(),
		app:    fiber.New(),
		addr:   addr,
	}

	for _, c := range configurator {
		c.Configure(f.config)
	}

	for _, service := range f.config.services {
		service.OnInit(f.app)
	}

	return f
}

type Fiber struct {
	config *FiberConfiguration
	app    *fiber.App
	addr   string
}

func (f *Fiber) OnReceive(ctx vivid.ActorContext) {
	switch m := ctx.Message().(type) {
	case *vivid.OnLaunch:
		f.onLaunch(ctx)
	case *vivid.OnTerminate:
		f.onTerminate(ctx)
	case error:
		f.onError(ctx, m)
	}
}

func (f *Fiber) onLaunch(ctx vivid.ActorContext) {
	ctx.Future().AwaitForward(ctx.Ref(), func() vivid.Message {
		return f.app.Listen(f.addr)
	})
}

func (f *Fiber) onError(ctx vivid.ActorContext, err error) {
	ctx.ReportAbnormal(err)
}

func (f *Fiber) onTerminate(ctx vivid.ActorContext) {
	_ = f.app.Shutdown()
}
