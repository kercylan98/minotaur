package minotaur

import (
	"context"
	"github.com/kercylan98/minotaur/minotaur/transport"
	"github.com/kercylan98/minotaur/minotaur/vivid"
	"github.com/kercylan98/minotaur/toolkit"
	"github.com/kercylan98/minotaur/toolkit/log"
	"os"
	"os/signal"
	"syscall"
	"time"
)

// NewApplication 创建一个应用程序
func NewApplication(options ...Option) *Application {
	opts := defaultOptions().apply(options...)

	actorSystem := vivid.NewActorSystem(opts.ActorSystemName, opts.ActorSystemOptions...)
	ctx, cancel := context.WithCancel(actorSystem.Context())
	app := &Application{
		options:     opts,
		ctx:         ctx,
		cancel:      cancel,
		closed:      make(chan struct{}),
		actorSystem: &actorSystem,
	}
	return app
}

// Application 基于 Minotaur 的应用程序结构
type Application struct {
	vivid.ActorRef
	options          *Options
	ctx              context.Context
	cancel           context.CancelFunc
	closed           chan struct{}
	actorSystem      *vivid.ActorSystem
	server           transport.ServerActorTyped
	onReceiveHandler []func(app *Application, ctx vivid.MessageContext)
}

func (a *Application) onReceive(ctx vivid.MessageContext) {
	for _, handler := range a.onReceiveHandler {
		handler(a, ctx)
	}
}

// EnablePProf 启用 PProf
func (a *Application) EnablePProf(addr, prefix string, errorHandler func(err error)) {
	toolkit.EnableHttpPProf(addr, prefix, errorHandler)
}

// DisablePProf 禁用 PProf
func (a *Application) DisablePProf(addr string) {
	toolkit.DisableHttpPProf(addr)
}

// GetSystem 获取 ActorSystem
func (a *Application) GetSystem() *vivid.ActorSystem {
	return a.actorSystem
}

// GetServer 获取 ServerActor
func (a *Application) GetServer() transport.ServerActorTyped {
	if a.server == nil {
		panic("server actor not initialized or not launched, please with WithNetwork option to initialize it")
	}
	return a.server
}

// GetContext 获取 ActorContext
func (a *Application) GetContext() vivid.ActorContext {
	return a.actorSystem.GetContext()
}

// Launch 启动应用程序
func (a *Application) Launch(onReceive ...func(app *Application, ctx vivid.MessageContext)) {
	defer func(a *Application) {
		close(a.closed)
	}(a)
	a.onReceiveHandler = onReceive

	if a.options.Network != nil {
		a.server = transport.NewServerActor(a.actorSystem, vivid.NewActorOptions[*transport.ServerActor]().WithName("server"))
		a.server.Tell(transport.ServerLaunchMessage{Network: a.options.Network})
	}

	a.ActorRef = vivid.ActorOfI(a.actorSystem, &applicationActor{a}, func(options *vivid.ActorOptions[*applicationActor]) {
		options.WithName("app")
	})

	if a.actorSystem.ClusterEnabled() {
		if err := a.actorSystem.JoinCluster(); err != nil {
			panic(err)
		}
	}

	var systemSignal = make(chan os.Signal, 1)
	signal.Notify(systemSignal, syscall.SIGHUP, syscall.SIGQUIT, syscall.SIGTERM, syscall.SIGINT)
	select {
	case <-systemSignal:
	case <-a.ctx.Done():
	}

	a.tryGracefulShutdown()

	a.actorSystem.Shutdown()
}

// Shutdown 关闭应用程序
func (a *Application) Shutdown() {
	a.cancel()
	<-a.closed
}

// ActorSystem 获取 ActorSystem
func (a *Application) ActorSystem() *vivid.ActorSystem {
	return a.actorSystem
}

// ActorOf 创建一个 Actor
func (a *Application) ActorOf(ofo vivid.ActorOfO) vivid.ActorRef {
	return a.actorSystem.ActorOf(ofo)
}

func (a *Application) tryGracefulShutdown() {
	if a.server != nil && a.options.NetworkShutdownMode == NetworkShutdownModeGraceful {
		var ctx context.Context
		var cancel context.CancelFunc
		if a.options.NetworkGracefulShutdownTimeout > 0 {
			ctx, cancel = context.WithTimeout(context.Background(), a.options.NetworkGracefulShutdownTimeout)
		} else {
			ctx, cancel = context.WithCancel(context.Background())
		}
		defer cancel()

		var lastHit = time.Now()
		var logger = func(attrs ...log.Attr) {
			if a.options.NetworkGracefulShutdownPromptHandler == nil {
				return
			}
			now := time.Now()
			if time.Since(lastHit) > a.options.NetworkGracefulShutdownPromptInterval {
				a.options.NetworkGracefulShutdownPromptHandler(attrs...)
				lastHit = now
			}
		}

		for {
			now := time.Now()
			status, err := a.server.Status()
			switch {
			case err != nil:
				logger(log.Err(err))
			case status.ConnectionNum > 0:
				logger(log.String("message", "waiting for all connections to close"), log.Int("connection_num", status.ConnectionNum))
			case status.LastConnectionOpenedTime.After(now.Add(-a.options.NetworkGracefulShutdownWaitTime)):
				logger(log.String("message", "wait for no more new connections to come in"), log.Time("expected_shutdown_time", status.LastConnectionOpenedTime.Add(a.options.NetworkGracefulShutdownWaitTime)))
			default:
				logger(log.String("message", "all connections are closed, shutdown server"))
				return
			}

			select {
			case <-ctx.Done():
				logger(log.String("message", "timeout, shutdown server"))
				return
			default:
				time.Sleep(time.Second)
			}
		}
	}
}
