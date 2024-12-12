package application

import (
	"github.com/gofiber/contrib/websocket"
	"github.com/gofiber/fiber/v2"
	"github.com/kercylan98/minotaur/engine/socket"
	"github.com/kercylan98/minotaur/engine/vivid"
	"github.com/kercylan98/minotaur/engine/vivid/supervision"
	"github.com/kercylan98/minotaur/toolkit/log"
)

func setupServer(app *Application, config *Configuration) {
	app.actorSystem.ActorOfF(func() vivid.Actor {
		return &serverActor{app: app, config: config}
	}, func(descriptor *vivid.ActorDescriptor) {
		descriptor.WithSupervisionStrategyProvider(supervision.FunctionalStrategyProvider(func() supervision.Strategy {
			return supervision.RestartStrategy()
		}))
	})
}

type serverActor struct {
	app    *Application
	config *Configuration

	socketFactory socket.Factory
	fiberApp      *fiber.App
}

func (s *serverActor) OnReceive(ctx vivid.ActorContext) {
	switch m := ctx.Message().(type) {
	case *vivid.OnLaunch:
		s.socketFactory = socket.NewFactory(ctx.System())
		s.fiberApp = fiber.New(fiber.Config{
			DisableStartupMessage: true,
		})

		s.fiberApp.Get("/websocket", websocket.New(func(conn *websocket.Conn) {
			socket.ProduceGorillaSocket(s.socketFactory, conn, newConn())
		}))

		if s.config.fiberSettingHandler != nil {
			s.config.fiberSettingHandler(s.fiberApp)
		}

		ctx.AwaitForward(ctx.Ref(), func() vivid.Message {
			return s.fiberApp.Listen(s.config.addr)
		})
	case *vivid.OnRestarted:
		ctx.System().Logger().Warn("server restarted")
	case *vivid.OnTerminated:
		if err := s.fiberApp.Shutdown(); err != nil {
			ctx.System().Logger().Error("server shutdown error", log.Err(err))
		}
	case error:
		ctx.ReportAbnormal(m)
	}
}
