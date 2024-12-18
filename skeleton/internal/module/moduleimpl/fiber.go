package moduleimpl

import (
	gofiber "github.com/gofiber/fiber/v2"
	"github.com/kercylan98/minotaur/engine/vivid"
	"github.com/kercylan98/minotaur/engine/vivid/supervision"
	"github.com/kercylan98/minotaur/skeleton/internal/module"
	"github.com/kercylan98/minotaur/skeleton/pkg/application"
	"github.com/kercylan98/minotaur/skeleton/pkg/fiber"
	"time"
)

var _ module.FiberModule = (*FiberModule)(nil)

type FiberModule struct {
	modules struct {
		actorSystem module.ActorSystemModule
	}
	fiberApp *fiber.Server[*application.Context]
}

func (f *FiberModule) OnInitialize(ctx *application.Context) (err error) {
	f.fiberApp = fiber.New(ctx, gofiber.New())

	return
}

func (f *FiberModule) OnDependencyInitialize(ctx *application.Context) (err error) {
	f.modules.actorSystem = application.LoadModule[module.ActorSystemModule](ctx)
	return
}

func (f *FiberModule) OnDependencySetup() (err error) {
	return
}

func (f *FiberModule) OnSetup() (err error) {
	f.modules.actorSystem.ActorSystem().ActorOfF(func() vivid.Actor {
		return newFiberActor(f)
	}, func(descriptor *vivid.ActorDescriptor) {
		descriptor.WithSupervisionStrategyProvider(supervision.FunctionalStrategyProvider(func() supervision.Strategy {
			return supervision.OneForOne(10, time.Microsecond*100, time.Second*3, supervision.FunctionalDecide(func(record *supervision.AccidentRecord) supervision.Directive {
				return supervision.DirectiveRestart
			}))
		}))
	})
	return
}

func (f *FiberModule) Fiber() *fiber.Server[*application.Context] {
	return f.fiberApp
}
