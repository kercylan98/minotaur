package main

import (
	"github.com/kercylan98/minotaur/skeleton/internal/controller"
	"github.com/kercylan98/minotaur/skeleton/internal/module"
	"github.com/kercylan98/minotaur/skeleton/internal/module/moduleimpl"
	"github.com/kercylan98/minotaur/skeleton/pkg/application"
)

func main() {
	ctx := application.New()

	application.RegisterModule[module.ActorSystemModule](ctx, moduleimpl.NewActorSystemModule())
	application.RegisterModule[module.FiberModule](ctx, moduleimpl.NewFiberModule())

	application.RegisterController(ctx, controller.NewWebSocketController())
}
