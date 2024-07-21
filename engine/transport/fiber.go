package transport

import (
	"github.com/gofiber/fiber/v2"
	"github.com/kercylan98/minotaur/engine/prc"
	"github.com/kercylan98/minotaur/engine/vivid"
)

// TODO: 使用体验待优化

func NewFiber(addr string, configurator ...FiberConfigurator) *Fiber {
	f := &Fiber{
		config: newFiberConfiguration(),
		addr:   addr,
	}

	for _, c := range configurator {
		c.Configure(f.config)
	}

	return f
}

type Fiber struct {
	config *FiberConfiguration
	app    *FiberApp
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
	f.app = newFiberApp(fiber.New())

	for _, service := range f.config.services {
		service.OnInit(ctx, f.app)
	}

	// 如果 service 互相依赖

	ctx.Future().AwaitForward(ctx.Ref(), func() prc.Message {
		return f.app.Listen(f.addr)
	})
}

func (f *Fiber) onError(ctx vivid.ActorContext, err error) {
	ctx.ReportAbnormal(err)
}

func (f *Fiber) onTerminate(ctx vivid.ActorContext) {
	_ = f.app.Shutdown()
}
