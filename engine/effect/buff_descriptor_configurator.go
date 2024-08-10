package effect

type BuffDescriptorConfigurator interface {
	Configure(c *BuffDescriptor)
}

type FunctionalBuffDescriptorConfigurator func(c *BuffDescriptor)

func (f FunctionalBuffDescriptorConfigurator) Configure(c *BuffDescriptor) {
	f(c)
}
