package effect

type ManagerConfigurator interface {
	Configure(c *ManagerConfiguration)
}

type FunctionalManagerConfigurator func(c *ManagerConfiguration)

func (f FunctionalManagerConfigurator) Configure(c *ManagerConfiguration) {
	f(c)
}
