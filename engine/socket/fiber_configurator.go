package socket

type FiberV2Configurator interface {
	Configure(config *FiberV2Configuration)
}

type FunctionalFiberV2Configurator func(config *FiberV2Configuration)

func (f FunctionalFiberV2Configurator) Configure(config *FiberV2Configuration) {
	f(config)
}
