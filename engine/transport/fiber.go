package transport

import (
	"github.com/gofiber/fiber/v2"
	"github.com/kercylan98/minotaur/engine/prc"
	"github.com/kercylan98/minotaur/engine/vivid"
	"github.com/kercylan98/minotaur/toolkit/log"
)

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

	// 自身初始化
	for _, service := range f.config.services {
		service.OnInit(ctx, f.app)
	}

	// 所有服务加载完成（可注入其他服务依赖）
	if err := f.app.hooks.onLoadedHook(); err != nil {
		ctx.System().Logger().Error("Fiber", log.String("hook", "loaded"), log.Err(err))
		ctx.Terminate(ctx.Ref(), false)
		return
	}

	// 所有服务挂载完成（可初始化需要依赖其他服务的内容）
	if err := f.app.hooks.onMountedHook(); err != nil {
		ctx.System().Logger().Error("Fiber", log.String("hook", "mounted"), log.Err(err))
		ctx.Terminate(ctx.Ref(), false)
		return
	}

	// 启动
	if err := f.app.hooks.onLaunchHook(); err != nil {
		ctx.System().Logger().Error("Fiber", log.String("hook", "launch"), log.Err(err))
		ctx.Terminate(ctx.Ref(), false)
		return
	}

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
