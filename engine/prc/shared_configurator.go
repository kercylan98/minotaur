package prc

// SharedConfigurator 共享配置器
type SharedConfigurator interface {
	// Config 配置
	Config(config *SharedConfiguration)
}

// FunctionalSharedConfigurator 函数式共享配置器
type FunctionalSharedConfigurator func(config *SharedConfiguration)

// Config 配置
func (f FunctionalSharedConfigurator) Config(config *SharedConfiguration) {
	f(config)
}
