package main

import (
	"github.com/kercylan98/minotaur/modular"
	"github.com/kercylan98/minotaur/modular/example/internal/service/services/attack"
	"github.com/kercylan98/minotaur/modular/example/internal/service/services/login"
	"github.com/kercylan98/minotaur/modular/example/internal/service/services/server"
)

func main() {
	modular.RegisterServices(
		new(attack.Service),
		new(server.Service),
		new(login.Service),
	)
	modular.Run()
}
