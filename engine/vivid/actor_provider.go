package vivid

import "github.com/kercylan98/minotaur/toolkit/charproc"

// ActorProviderName 是一个字符串类型的 ActorProvider 名称
type ActorProviderName = string

// ActorProvider 是一个 Actor 生成器接口，它定义了生成 Actor 实例的方法。
type ActorProvider interface {
	// GetActorProviderName 返回 ActorProvider 的名称
	GetActorProviderName() ActorProviderName

	// Provide 每次调用都应返回一个新的 Actor 实例，错误的使用可能导致 Actor 状态被污染。
	Provide() Actor
}

// FunctionalActorProvider 是一个函数类型的 Actor 生成器，它定义了生成 Actor 实例的方法。
type FunctionalActorProvider func() Actor

// GetActorProviderName 返回 ActorProvider 的名称
func (f FunctionalActorProvider) GetActorProviderName() ActorProviderName {
	return charproc.None
}

// Provide 每次调用都应返回一个新的 Actor 实例，错误的使用可能导致 Actor 状态被污染。
func (f FunctionalActorProvider) Provide() Actor {
	return f()
}

// NewShortcutActorProvider 创建一个包含名称的 ActorProvider 实例，它使用给定的 Actor 生成器作为提供者。
func NewShortcutActorProvider(name ActorProviderName, provide func() Actor) ActorProvider {
	return &shortcutActorProvider{
		name:     name,
		provider: FunctionalActorProvider(provide),
	}
}

type shortcutActorProvider struct {
	name     ActorProviderName
	provider ActorProvider
}

func (s *shortcutActorProvider) GetActorProviderName() ActorProviderName {
	return s.name
}

func (s *shortcutActorProvider) Provide() Actor {
	return s.provider.Provide()
}
