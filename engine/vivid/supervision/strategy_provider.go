package supervision

// StrategyProvider 是监督策略提供者
type StrategyProvider interface {
	// Provide 提供一个监督策略
	Provide() Strategy
}

// FunctionalStrategyProvider 是一个函数类型的监督策略提供者
type FunctionalStrategyProvider func() Strategy

// Provide 提供一个监督策略
func (f FunctionalStrategyProvider) Provide() Strategy {
	return f()
}
