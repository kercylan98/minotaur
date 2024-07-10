package main

import (
	"github.com/gofiber/fiber/v2"
	"github.com/kercylan98/minotaur/core/transport"
	"github.com/kercylan98/minotaur/core/vivid"
)

type PingService struct {
}

func (f *PingService) OnInit(kit *transport.FiberKit) {
	kit.App().Get("/ping", f.onPing)
}

func (f *PingService) onPing(ctx *fiber.Ctx) error {
	_, err := ctx.WriteString("pong")
	return err
}

func main() {
	system := vivid.NewActorSystem(func(options *vivid.ActorSystemOptions) {
		options.WithModule(transport.NewFiber(":8877").BindService(new(PingService)))
	})

	system.ShutdownGracefully()
}
