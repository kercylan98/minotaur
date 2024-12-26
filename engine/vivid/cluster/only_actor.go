package cluster

import (
	"github.com/kercylan98/minotaur/engine/vivid"
)

type OnlyActorProvider interface {
	ProvideActor() (actor vivid.Actor)
	ProvideConfigurator() (configurator vivid.ActorDescriptorConfigurator)
}
