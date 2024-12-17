package main

import (
	"github.com/kercylan98/minotaur/skeleton/discard/internal/modules/configurator"
	"github.com/kercylan98/minotaur/skeleton/internal/discard/pkg/application"
)

func main() {
	app, binder := application.New()
	app.Register(
		configurator.NewConfiguratorModule(binder),
	)
	app.Run()
}
