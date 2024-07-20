package prc

// DiscovererConfigurator 是用于配置 DiscovererConfiguration 的接口
type DiscovererConfigurator interface {
	// Configure 用于配置 DiscovererConfiguration
	Configure(configuration *DiscovererConfiguration)
}

// FunctionalDiscovererConfigurator 是用于配置 DiscovererConfiguration 的函数式配置器
type FunctionalDiscovererConfigurator func(configuration *DiscovererConfiguration)

// Configure 用于配置 DiscovererConfiguration
func (f FunctionalDiscovererConfigurator) Configure(configuration *DiscovererConfiguration) {
	f(configuration)
}
