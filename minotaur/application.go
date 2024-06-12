package minotaur

import (
	"context"
	"github.com/kercylan98/minotaur/minotaur/transport"
	"github.com/kercylan98/minotaur/minotaur/vivid"
	"os"
	"os/signal"
	"syscall"
)

func NewApplication(options ...Option) *Application {
	opts := new(Options).apply(options...)

	actorSystem := vivid.NewActorSystem(opts.ActorSystemName)
	ctx, cancel := context.WithCancel(actorSystem.Context())
	return &Application{
		options:     opts,
		ctx:         ctx,
		cancel:      cancel,
		closed:      make(chan struct{}),
		actorSystem: &actorSystem,
	}
}

// Application 基于 Minotaur 的应用程序结构
type Application struct {
	options     *Options
	ctx         context.Context
	cancel      context.CancelFunc
	closed      chan struct{}
	actorSystem *vivid.ActorSystem
	server      vivid.TypedActorRef[transport.ServerActorTyped]
}

func (a *Application) GetSystem() *vivid.ActorSystem {
	return a.actorSystem
}

func (a *Application) GetServer() vivid.TypedActorRef[transport.ServerActorTyped] {
	if a.server == nil {
		panic("server actor not initialized or not launched, please with WithNetwork option to initialize it")
	}
	return a.server
}

func (a *Application) GetContext() vivid.ActorContext {
	return a.actorSystem.GetContext()
}

func (a *Application) Launch() {
	defer func(a *Application) {
		close(a.closed)
	}(a)

	if a.options.Network != nil {
		a.server = transport.NewServerActor(a.actorSystem, vivid.NewActorOptions[*transport.ServerActor]().WithName("server"))
		a.server.Protocol().Launch(a.options.Network)
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

func (a *Application) ActorSystem() *vivid.ActorSystem {
	return a.actorSystem
}

func (a *Application) ActorOf(ofo vivid.ActorOfO) vivid.ActorRef {
	return a.actorSystem.ActorOf(ofo)
}
