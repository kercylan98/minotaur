package main

import (
	"github.com/kercylan98/minotaur/skeleton/internal/controller"
	"github.com/kercylan98/minotaur/skeleton/internal/module"
	"github.com/kercylan98/minotaur/skeleton/internal/module/moduleimpl"
	"github.com/kercylan98/minotaur/skeleton/pkg/application"
)

func main() {
	ctx := application.New(application.FunctionalConfigurator(func(c *application.Configuration) {
		c.WithConfigDir("./conf")
	}))
	application.RegisterModule[module.ActorSystemModule](ctx, moduleimpl.NewActorSystemModule())
	application.RegisterModule[module.FiberModule](ctx, moduleimpl.NewFiberModule())
	application.RegisterModule[module.RPCModule](ctx, moduleimpl.NewRPCModule())

	application.RegisterController(ctx, controller.NewOAuthController())

	if err := ctx.Run(); err != nil {
		panic(err)
	}
}
