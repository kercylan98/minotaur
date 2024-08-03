package vivid

// ActorProvider 是一个 Actor 生成器接口，它定义了生成 Actor 实例的方法。
type ActorProvider interface {
	// Provide 每次调用都应返回一个新的 Actor 实例，错误的使用可能导致 Actor 状态被污染。
	Provide() Actor
}

// FunctionalActorProvider 是一个函数类型的 Actor 生成器，它定义了生成 Actor 实例的方法。
type FunctionalActorProvider func() Actor

// Provide 每次调用都应返回一个新的 Actor 实例，错误的使用可能导致 Actor 状态被污染。
func (f FunctionalActorProvider) Provide() Actor {
	return f()
}
