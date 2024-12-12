package application

import (
	"github.com/kercylan98/minotaur/engine/vivid"
	"os"
)

func New(configurator ...Configurator) *Application {
	config := newConfiguration()
	for _, conf := range configurator {
		conf.Configure(config)
	}

	ins := &Application{
		actorSystem: vivid.NewActorSystem(),
	}

	setupServer(ins, config)
	return ins
}

type Application struct {
	config      *Configuration
	actorSystem *vivid.ActorSystem
}

func (i *Application) ActorSystem() *vivid.ActorSystem {
	return i.actorSystem
}

func (i *Application) Run() {
	i.actorSystem.Signal(func(system *vivid.ActorSystem, signal os.Signal) {
		system.Logger().Info("system ready shutdown, signal: %s", signal)
		system.Shutdown(true)
		system.Logger().Info("system shutdown")
	})
}
