package application

type Configurator interface {
	Configure(configuration *Configuration)
}

type FunctionalConfigurator func(configuration *Configuration)

func (f FunctionalConfigurator) Configure(configuration *Configuration) {
	f(configuration)
}
