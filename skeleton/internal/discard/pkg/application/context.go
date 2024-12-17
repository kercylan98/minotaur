package application

import (
	"github.com/kercylan98/minotaur/engine/vivid"
	"github.com/kercylan98/minotaur/engine/vivid/cluster"
	"github.com/kercylan98/minotaur/skeleton/internal/discard/pkg/module"
	"os"
	"os/signal"
	"syscall"
)

type ActorSystemBinder func(actorSystem *vivid.ActorSystem, actorSystemCluster *cluster.ActorSystem)

func New() (*Context, ActorSystemBinder) {
	ctx := &Context{}
	ctx.manager = module.New(ctx)

	return ctx, func(actorSystem *vivid.ActorSystem, actorSystemCluster *cluster.ActorSystem) {
		ctx.actorSystem = actorSystem
		ctx.actorSystemCluster = actorSystemCluster
	}
}

type Context struct {
	manager            *module.Manager[*Context]
	actorSystem        *vivid.ActorSystem
	actorSystemCluster *cluster.ActorSystem
}

func (c *Context) Register(modules ...module.Module[*Context]) *Context {
	c.manager.Register(modules...)
	return c
}

func (c *Context) ActorSystem() *vivid.ActorSystem {
	return c.actorSystem
}

func (c *Context) Cluster() *cluster.ActorSystem {
	return c.actorSystemCluster
}

func (c *Context) Run() {
	c.manager.Run()
	if c.actorSystem != nil {
		c.actorSystem.Signal(func(system *vivid.ActorSystem, signal os.Signal) {
			system.Shutdown(true)
		})
	} else {
		s := make(chan os.Signal, 1)
		signal.Notify(s, syscall.SIGHUP, syscall.SIGQUIT, syscall.SIGTERM, syscall.SIGINT)
		<-s
	}
}
