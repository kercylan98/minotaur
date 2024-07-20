package supervision

import "github.com/kercylan98/minotaur/toolkit/charproc"

// StrategyProvider 是监督策略提供者
type StrategyProvider interface {
	// GetStrategyProviderName 返回监督策略提供者的名称
	GetStrategyProviderName() StrategyName
	// Provide 提供一个监督策略
	Provide() Strategy
}

// FunctionalStrategyProvider 是一个函数类型的监督策略提供者
type FunctionalStrategyProvider func() Strategy

// GetStrategyProviderName 返回监督策略提供者的名称
func (f FunctionalStrategyProvider) GetStrategyProviderName() StrategyName {
	return charproc.None
}

// Provide 提供一个监督策略
func (f FunctionalStrategyProvider) Provide() Strategy {
	return f()
}

// NewShortcutStrategyProvider 创建一个包含名称的监督策略提供者实例，它使用给定的监督策略生成器作为提供者。
func NewShortcutStrategyProvider(name StrategyName, provide func() Strategy) StrategyProvider {
	return &shortcutStrategyProvider{
		name:     name,
		provider: FunctionalStrategyProvider(provide),
	}
}

type shortcutStrategyProvider struct {
	name     StrategyName
	provider StrategyProvider
}

func (s *shortcutStrategyProvider) GetStrategyProviderName() StrategyName {
	return s.name
}

func (s *shortcutStrategyProvider) Provide() Strategy {
	return s.provider.Provide()
}
