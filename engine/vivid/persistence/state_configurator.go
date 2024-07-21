package persistence

type StateConfigurator interface {
	Configure(configuration *StateConfiguration)
}

type FunctionalStateConfigurator func(configuration *StateConfiguration)

func (f FunctionalStateConfigurator) Configure(configuration *StateConfiguration) {
	f(configuration)
}
