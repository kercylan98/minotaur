package transport

import (
	"github.com/kercylan98/minotaur/core/vivid"
	"github.com/kercylan98/minotaur/toolkit/log"
	"os"
)

var testLogger = log.New(log.NewHandler(os.Stdout, log.NewDevHandlerOptions()))
var benchmarkLogger = log.NewSilentLogger()

func NewTestActorSystem(options ...func(options *vivid.ActorSystemOptions)) *vivid.ActorSystem {
	return vivid.NewActorSystem(append(options, func(options *vivid.ActorSystemOptions) {
		options.WithLoggerProvider(func() *log.Logger {
			return testLogger
		})
	})...)
}

func NewBenchmarkActorSystem(options ...func(options *vivid.ActorSystemOptions)) *vivid.ActorSystem {
	return vivid.NewActorSystem(append(options, func(options *vivid.ActorSystemOptions) {
		options.WithLoggerProvider(func() *log.Logger {
			return benchmarkLogger
		})
	})...)
}
