package fiber

import (
	gofiber "github.com/gofiber/fiber/v2"
	"github.com/kercylan98/minotaur/engine/vivid"
	"github.com/kercylan98/minotaur/engine/vivid/supervision"
	"github.com/kercylan98/minotaur/skeleton/pkg/application"
	"github.com/kercylan98/minotaur/skeleton/pkg/fiber"
	"github.com/kercylan98/minotaur/toolkit/log"
	"time"
)

var (
	_ application.Component = (*fiberComponent)(nil)
)

func NewFiberComponent(addr string) application.Component {
	return &fiberComponent{
		addr: addr,
	}
}

type fiberComponent struct {
	addr         string
	fiberHandler []func(fiberApp *fiber.App)
}

func (f *fiberComponent) OnStart(app *application.Context) {
	app.ActorSystem().ActorOfF(func() vivid.Actor {
		actor := &fiberActor{
			component: f,
			fiberApp: fiber.NewApp(app, gofiber.New(gofiber.Config{
				DisableStartupMessage: true,
			})),
		}

		for _, h := range f.fiberHandler {
			h(actor.fiberApp)
		}

		return actor
	}, func(descriptor *vivid.ActorDescriptor) {
		descriptor.WithNamePrefix("gofiber")
		descriptor.WithSupervisionStrategyProvider(supervision.FunctionalStrategyProvider(func() supervision.Strategy {
			return supervision.OneForOne(-1, time.Millisecond*100, time.Second, supervision.FunctionalDecide(func(record *supervision.AccidentRecord) supervision.Directive {
				return supervision.DirectiveRestart
			}))
		}))
	})
}

func (f *fiberComponent) RegisterFiberHandler(handlers ...func(fiberApp *fiber.App)) {
	f.fiberHandler = append(f.fiberHandler, handlers...)
}

type fiberActor struct {
	component *fiberComponent
	fiberApp  *fiber.App
}

func (f *fiberActor) OnReceive(ctx vivid.ActorContext) {
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

func (f *fiberActor) onLaunch(ctx vivid.ActorContext) {
	ctx.AwaitForward(ctx.Ref(), func() vivid.Message {
		return f.fiberApp.Listen(f.component.addr)
	})
}

func (f *fiberActor) onTerminated(ctx vivid.ActorContext, m *vivid.OnTerminated) {
	if !m.TerminatedActor.Equal(ctx.Ref()) {
		return
	}

	if err := f.fiberApp.ShutdownWithTimeout(time.Minute); err != nil {
		ctx.System().Logger().Error("gofiber", log.String("event", "shutdown failed"), log.Err(err))
	} else {
		ctx.System().Logger().Info("gofiber", log.String("event", "shutdown success"))
	}
}

func (f *fiberActor) onRestarted(ctx vivid.ActorContext, m *vivid.OnRestarted) {
	ctx.System().Logger().Warn("gofiber", log.String("event", "restarted"))
}
