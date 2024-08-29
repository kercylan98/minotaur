package stream

// Deprecated: 该设计加大了理解成本，且不易于使用，将考虑新的方案用于处理网络连接。至 v0.7.0 版本及以后，stream 包将被移除。
type Configurator interface {
	Configure(c *Configuration)
}

// Deprecated: 该设计加大了理解成本，且不易于使用，将考虑新的方案用于处理网络连接。至 v0.7.0 版本及以后，stream 包将被移除。
type FunctionalConfigurator func(c *Configuration)

func (f FunctionalConfigurator) Configure(c *Configuration) {
	f(c)
}
