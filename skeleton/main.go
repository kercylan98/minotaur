package main

import (
	"github.com/kercylan98/minotaur/skeleton/pkg/application"
	"github.com/kercylan98/minotaur/skeleton/pkg/modules/module/fiber"
	"time"
)

func main() {
	app := application.New()
	app.SetupModule(fiber.NewFiber(":8888"))
	app.SetupModule(fiber.NewWebSocket())
	app.Run()
	time.Sleep(time.Minute)
}
