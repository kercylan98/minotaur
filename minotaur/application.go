package minotaur

import (
	"context"
	"github.com/kercylan98/minotaur/minotaur/pulse"
	"github.com/kercylan98/minotaur/minotaur/transport"
	"github.com/kercylan98/minotaur/minotaur/vivid"
	"os"
	"os/signal"
	"syscall"
)

func NewApplication(options ...Option) *Application {
	opts := new(Options).apply(options...)

	actorSystem := vivid.NewActorSystem(opts.ActorSystemName)
	eventBus := pulse.NewPulse(&actorSystem, vivid.NewActorOptions[*pulse.EventBusActor]().WithName(opts.EventBusActorName))
	ctx, cancel := context.WithCancel(actorSystem.Context())
	return &Application{
		options:     opts,
		ctx:         ctx,
		cancel:      cancel,
		closed:      make(chan struct{}),
		actorSystem: &actorSystem,
		eventBus:    &eventBus,
	}
}

// Application 基于 Minotaur 的应用程序结构
type Application struct {
	options     *Options
	ctx         context.Context
	cancel      context.CancelFunc
	closed      chan struct{}
	actorSystem *vivid.ActorSystem
	eventBus    *pulse.Pulse
	server      *transport.Server
}

func (a *Application) Launch() {
	defer func(a *Application) {
		close(a.closed)
	}(a)

	if a.options.Network != nil {
		srv := transport.NewServer(a.actorSystem, vivid.NewActorOptions[*transport.ServerActor]().WithName("server"))
		a.server = &srv
		a.server.Launch(a.eventBus, a.options.Network)
	}

	var systemSignal = make(chan os.Signal, 1)
	signal.Notify(systemSignal, syscall.SIGHUP, syscall.SIGQUIT, syscall.SIGTERM, syscall.SIGINT)
	select {
	case <-systemSignal:
	case <-a.ctx.Done():
	}
	a.actorSystem.Shutdown()
}

func (a *Application) Shutdown() {
	a.cancel()
	<-a.closed
}

func (a *Application) EventBus() *pulse.Pulse {
	return a.eventBus
}

func (a *Application) ActorSystem() *vivid.ActorSystem {
	return a.actorSystem
}
