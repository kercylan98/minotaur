package prc

// SharedConfigurator 共享配置器
type SharedConfigurator interface {
	// Configure 配置
	Configure(config *SharedConfiguration)
}

// FunctionalSharedConfigurator 函数式共享配置器
type FunctionalSharedConfigurator func(config *SharedConfiguration)

// Configure 配置
func (f FunctionalSharedConfigurator) Configure(config *SharedConfiguration) {
	f(config)
}
