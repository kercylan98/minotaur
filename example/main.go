package main

import (
	"github.com/gofiber/fiber/v2"
	"github.com/kercylan98/minotaur/example/application"
)

func main() {
	app := application.New(application.FunctionalConfigurator(func(configuration *application.Configuration) {
		configuration.WithAddr(":8888")
		configuration.WithFiberSettings(func(fiberApp *fiber.App) {
			fiberApp.Get("/ping", func(ctx *fiber.Ctx) error {
				if _, err := ctx.WriteString("pong"); err != nil {
					return err
				}
				return nil
			})
		})
	}))
	app.Run()
}
