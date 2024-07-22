package stream

type Configurator interface {
	Configure(c *Configuration)
}

type FunctionalConfigurator func(c *Configuration)

func (f FunctionalConfigurator) Configure(c *Configuration) {
	f(c)
}
