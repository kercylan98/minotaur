package moduleimpl

import (
	"github.com/kercylan98/minotaur/engine/vivid"
	"github.com/kercylan98/minotaur/toolkit/log"
	"time"
)

func newFiberActor(module *FiberModule) *fiberActor {
	return &fiberActor{
		fiberModule: module,
	}
}

type fiberActor struct {
	fiberModule *FiberModule
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
		return f.fiberModule.fiberApp.Listen(":8080")
	})
}

func (f *fiberActor) onTerminated(ctx vivid.ActorContext, m *vivid.OnTerminated) {
	if !m.TerminatedActor.Equal(ctx.Ref()) {
		return
	}

	if err := f.fiberModule.fiberApp.ShutdownWithTimeout(time.Minute); err != nil {
		ctx.System().Logger().Error("fiber", log.String("event", "shutdown failed"), log.Err(err))
	} else {
		ctx.System().Logger().Info("fiber", log.String("event", "shutdown success"))
	}
}

func (f *fiberActor) onRestarted(ctx vivid.ActorContext, m *vivid.OnRestarted) {
	ctx.System().Logger().Warn("fiber", log.String("event", "restarted"))
}
