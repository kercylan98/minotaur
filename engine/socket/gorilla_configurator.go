package socket

type GorillaConfigurator interface {
	Configure(config *GorillaConfiguration)
}

type FunctionalGorillaConfigurator func(config *GorillaConfiguration)

func (f FunctionalGorillaConfigurator) Configure(config *GorillaConfiguration) {
	f(config)
}
