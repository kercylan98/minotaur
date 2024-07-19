package behavior

// Performance 行为具体表现
type Performance[T any] interface {
	// Perform 履行行为
	Perform(ctx T)
}

// FunctionalPerformance 函数式的行为表现，它实现了 Performance 接口
type FunctionalPerformance[T any] func(ctx T)

// Perform 履行行为
func (f FunctionalPerformance[T]) Perform(ctx T) {
	f(ctx)
}
