package main

import (
	"github.com/kercylan98/minotaur/skeleton/cmd/user-certification-center/internal/components/user"
	"github.com/kercylan98/minotaur/skeleton/pkg/application"
	"github.com/kercylan98/minotaur/skeleton/pkg/components/fiber"
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
