package moduleimpl

import (
	"github.com/kercylan98/minotaur/engine/vivid"
	"github.com/kercylan98/minotaur/engine/vivid/cluster"
	"github.com/kercylan98/minotaur/skeleton/internal/module"
	"github.com/kercylan98/minotaur/skeleton/pkg/application"
	"github.com/kercylan98/minotaur/toolkit"
)

var _ module.ActorSystemModule = (*ActorSystemModule)(nil)

func NewActorSystemModule() *ActorSystemModule {
	return &ActorSystemModule{}
}

type ActorSystemModule struct {
	cluster     *cluster.ActorSystem
	actorSystem *vivid.ActorSystem
}

func (a *ActorSystemModule) OnInitialize(ctx *application.Context) (err error) {
	name := ctx.GetBootstrapConfig().GetString(application.ConfigurationVividNameKey)
	addr := ctx.GetBootstrapConfig().GetString(application.ConfigurationVividAddrKey)
	seedNodes := ctx.GetBootstrapConfig().GetStringSlice(application.ConfigurationVividClusterSeedNodesKey)

	if len(seedNodes) > 0 {
		// Cluster mode
		a.cluster = cluster.NewActorSystem(addr, cluster.FunctionalActorSystemConfigurator(func(config *cluster.ActorSystemConfiguration) {
			toolkit.IfThen(name != "", func() {
				config.WithName(name)
			})
			config.WithSeeds(seedNodes...)
		}))
		a.actorSystem = a.cluster.ActorSystem
	} else {
		// Standalone mode
		a.actorSystem = vivid.NewActorSystem(vivid.FunctionalActorSystemConfigurator(func(config *vivid.ActorSystemConfiguration) {
			toolkit.IfThen(name != "", func() {
				config.WithName(name)
			})
			toolkit.IfThen(addr != "", func() {
				config.WithShared(addr)
			})
		}))
	}

	return
}

func (a *ActorSystemModule) OnDependencyInitialize(ctx *application.Context) (err error) {
	return
}

func (a *ActorSystemModule) OnDependencySetup() (err error) {
	return
}

func (a *ActorSystemModule) OnSetup() (err error) {
	return
}

func (a *ActorSystemModule) ActorSystem() *vivid.ActorSystem {
	return a.actorSystem
}

func (a *ActorSystemModule) Cluster() *cluster.ActorSystem {
	return a.cluster
}
