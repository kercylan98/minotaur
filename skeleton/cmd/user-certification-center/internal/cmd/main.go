package main

import (
	"github.com/kercylan98/minotaur/skeleton/pkg/application"
	"github.com/kercylan98/minotaur/skeleton/pkg/components/fiber"
	"github.com/kercylan98/minotaur/skeleton/pkg/components/user"
)

func main() {
	app := application.New()
	app.SetupComponents(
		fiber.NewFiberComponent(":8080"),
		user.NewUserComponent(),
	)
	if err := app.Run(); err != nil {
		panic(err)
	}
}
