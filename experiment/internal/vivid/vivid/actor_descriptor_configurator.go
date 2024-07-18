package vivid

// ActorDescriptorConfigurator 是用于配置 ActorDescriptor 的配置器
type ActorDescriptorConfigurator interface {
	// Configure 配置 ActorDescriptor
	Configure(descriptor *ActorDescriptor)
}

type FunctionalActorDescriptorConfigurator func(descriptor *ActorDescriptor)

// Configure 配置 ActorDescriptor
func (f FunctionalActorDescriptorConfigurator) Configure(descriptor *ActorDescriptor) {
	f(descriptor)
}
