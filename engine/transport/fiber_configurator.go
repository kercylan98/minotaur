package transport

type FiberConfigurator interface {
	Configure(configuration *FiberConfiguration)
}

type FunctionalFiberConfigurator func(configuration *FiberConfiguration)

func (f FunctionalFiberConfigurator) Configure(configuration *FiberConfiguration) {
	f(configuration)
}
