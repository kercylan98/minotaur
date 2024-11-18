package socket

import "github.com/kercylan98/minotaur/engine/vivid"

func NewFactoryConfiguration() *FactoryConfiguration {
	return &FactoryConfiguration{}
}

type FactoryConfiguration struct {
	socketActorDescriptor vivid.ActorDescriptorConfigurator
}

func (fc *FactoryConfiguration) WithSocketActorDescriptorFunc(descriptor vivid.ActorDescriptorConfigurator) {
	fc.socketActorDescriptor = descriptor
}
