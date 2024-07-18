package benchmark

import (
	"github.com/kercylan98/minotaur/experiment/internal/vivid/vivid"
	"github.com/kercylan98/minotaur/toolkit/log"
)

func NewBenchmarkActorSystem() *vivid.ActorSystem {
	logger := log.NewSilentLogger()
	return vivid.NewActorSystem(vivid.FunctionalActorSystemConfigurator(func(config *vivid.ActorSystemConfiguration) {
		config.WithLoggerProvider(vivid.FunctionalLoggerProvider(func() *log.Logger {
			return logger
		}))
	}))
}
