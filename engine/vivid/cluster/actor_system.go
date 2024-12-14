package cluster

import (
	"context"
	"errors"
	"github.com/kercylan98/minotaur/engine/prc"
	"github.com/kercylan98/minotaur/engine/vivid"
	"github.com/kercylan98/minotaur/toolkit/log"
)

func init() {
	vivid.RegisterFutureAskType(func(ctx any) *vivid.ActorSystem {
		as, _ := ctx.(*ActorSystem)
		return as.ActorSystem
	})
}

func NewActorSystem(address prc.PhysicalAddress, seedNodes []prc.PhysicalAddress, configurator ...ActorSystemConfigurator) *ActorSystem {
	config := newActorSystemConfiguration()
	for _, c := range configurator {
		c.Configure(config)
	}

	system := &ActorSystem{
		config: config,
	}

	config.WithShared(address)
	config.WithComponents(&component{cluster: system})

	system.ActorSystem = vivid.NewActorSystemWithConfiguration(config.ActorSystemConfiguration)

	system.systemRef = system.ActorOfF(func() vivid.Actor {
		return newActorSystemActor(system, seedNodes)
	}, func(descriptor *vivid.ActorDescriptor) {
		descriptor.WithName("cluster")
	})

	return system
}

type ActorSystem struct {
	*vivid.ActorSystem                           // 如果单独使用，那么一切行为将越过集群
	config             *ActorSystemConfiguration // 集群配置
	systemRef          vivid.ActorRef            // 集群 ActorSystem 的 Actor 引用
}

func (sys *ActorSystem) onShutdown() {
	sys.Logger().Info("cluster", log.String("status", "shutdown in progress"))

	ctx, cancel := context.WithTimeout(context.Background(), sys.config.shutdownTimeout)
	sys.Tell(sys.systemRef, &actorSystemActorExitMessage{cancel: cancel})
	<-ctx.Done()

	if errors.Is(ctx.Err(), context.DeadlineExceeded) {
		sys.Logger().Error("cluster", log.String("status", "shutdown timeout"))
	} else {
		sys.Logger().Info("cluster", log.String("status", "shutdown success"))
	}
}
