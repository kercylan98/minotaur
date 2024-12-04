package typemarshalers

type GoConfigurator interface {
	Configure(config *GoConfiguration)
}

type FunctionalGoConfigurator func(config *GoConfiguration)

func (f FunctionalGoConfigurator) Configure(config *GoConfiguration) {
	f(config)
}
