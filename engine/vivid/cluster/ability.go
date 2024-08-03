package cluster

import "github.com/kercylan98/minotaur/engine/vivid"

func newAbility(provider ActorProvider, configurator ...vivid.ActorDescriptorConfigurator) *ability {
	return &ability{
		provider:     provider,
		configurator: configurator,
	}
}

type ability struct {
	provider     ActorProvider
	configurator []vivid.ActorDescriptorConfigurator
}
