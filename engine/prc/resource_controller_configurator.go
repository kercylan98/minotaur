package prc

// ResourceControllerConfigurator 是用于配置 ResourceController 的配置器
type ResourceControllerConfigurator interface {
	// Configure 配置 ActorSystem
	Configure(config *ResourceControllerConfiguration)
}

// FunctionalResourceControllerConfigurator 是用于配置 ResourceController 的配置器
type FunctionalResourceControllerConfigurator func(config *ResourceControllerConfiguration)

// Configure 配置 ActorSystem
func (f FunctionalResourceControllerConfigurator) Configure(config *ResourceControllerConfiguration) {
	f(config)
}
