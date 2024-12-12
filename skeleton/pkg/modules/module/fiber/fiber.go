package fiber

import (
	gofiber "github.com/gofiber/fiber/v2"
	"github.com/kercylan98/minotaur/engine/vivid"
	"github.com/kercylan98/minotaur/engine/vivid/supervision"
	"github.com/kercylan98/minotaur/skeleton/pkg/application"
	"github.com/kercylan98/minotaur/skeleton/pkg/modules"
	"github.com/kercylan98/minotaur/toolkit/log"
	"time"
)

func NewFiber(addr string, fiberHandler ...func(fiberApp *gofiber.App)) modules.Fiber {
	return &fiber{
		addr:         addr,
		fiberHandler: fiberHandler,
	}
}

type fiber struct {
	// CONFIG FIELDS

	addr         string
	fiberHandler []func(fiberApp *gofiber.App)

	// RUNTIME FIELDS

	fiberApp *gofiber.App
}

func (f *fiber) Name() string {
	return "fiber"
}

func (f *fiber) Setup(app *application.Context) error {
	app.ActorSystem().ActorOfF(func() vivid.Actor {
		actor := &fiber{
			fiberApp: gofiber.New(gofiber.Config{
				DisableStartupMessage: true,
			}),
			addr: f.addr,
		}

		for _, h := range f.fiberHandler {
			h(actor.fiberApp)
		}

		return actor
	}, func(descriptor *vivid.ActorDescriptor) {
		descriptor.WithNamePrefix("fiber")
		descriptor.WithSupervisionStrategyProvider(supervision.FunctionalStrategyProvider(func() supervision.Strategy {
			return supervision.OneForOne(-1, time.Millisecond*100, time.Second, supervision.FunctionalDecide(func(record *supervision.AccidentRecord) supervision.Directive {
				return supervision.DirectiveRestart
			}))
		}))
	})
	return nil
}

func (f *fiber) OnReceive(ctx vivid.ActorContext) {
	switch m := ctx.Message().(type) {
	case *vivid.OnLaunch:
		f.onLaunch(ctx)
	case *vivid.OnRestarted:
		f.onRestarted(ctx, m)
	case *vivid.OnTerminated:
		f.onTerminated(ctx, m)
	case error:
		ctx.ReportAbnormal(m)
	}
}

func (f *fiber) onLaunch(ctx vivid.ActorContext) {
	ctx.AwaitForward(ctx.Ref(), func() vivid.Message {
		return f.fiberApp.Listen(f.addr)
	})
}

func (f *fiber) onTerminated(ctx vivid.ActorContext, m *vivid.OnTerminated) {
	if !m.TerminatedActor.Equal(ctx.Ref()) {
		return
	}

	if err := f.fiberApp.ShutdownWithTimeout(time.Minute); err != nil {
		ctx.System().Logger().Error("fiber", log.String("event", "shutdown failed"), log.Err(err))
	} else {
		ctx.System().Logger().Info("fiber", log.String("event", "shutdown success"))
	}
}

func (f *fiber) onRestarted(ctx vivid.ActorContext, m *vivid.OnRestarted) {
	ctx.System().Logger().Warn("fiber", log.String("event", "restarted"))
}
