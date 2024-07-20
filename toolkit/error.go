package toolkit

// ErrorHandler 是一个通用的错误处理器
type ErrorHandler interface {
	// Handle 处理错误，根据不同的使用场景，err 可能会为空值
	Handle(err error)
}

// FunctionalErrorHandler 是一个函数式的错误处理器
type FunctionalErrorHandler func(err error)

// Handle 处理错误，根据不同的使用场景，err 可能会为空值
func (f FunctionalErrorHandler) Handle(err error) {
	f(err)
}

// ErrorPolicyDecisionHandler 是一个包含决策的错误处理器
type ErrorPolicyDecisionHandler[T any] interface {
	// Handle 处理错误，根据不同的使用场景，err 可能会为空值
	Handle(err error) T
}

// FunctionalErrorPolicyDecisionHandler 是一个函数式的包含决策的错误处理器
type FunctionalErrorPolicyDecisionHandler[T any] func(err error) T

// Handle 处理错误，根据不同的使用场景，err 可能会为空值
func (f FunctionalErrorPolicyDecisionHandler[T]) Handle(err error) T {
	return f(err)
}
