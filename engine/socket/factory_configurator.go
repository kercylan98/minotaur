package socket

type FactoryConfigurator interface {
	Configure(config *FactoryConfiguration)
}

type FunctionalFactoryConfigurator func(config *FactoryConfiguration)

func (f FunctionalFactoryConfigurator) Configure(config *FactoryConfiguration) {
	f(config)
}
