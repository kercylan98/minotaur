package combination

// Option 组合器选项
type Option[T Item] func(*Combination[T])

// WithEvaluation 设置组合评估函数
//   - 用于对组合进行评估，返回一个分值的评价函数
//   - 通过该选项将设置所有匹配器的默认评估函数为该函数
//   - 通过匹配器选项 WithMatcherEvaluation 可以覆盖该默认评估函数
//   - 默认的评估函数将返回一个随机数
func WithEvaluation[T Item](evaluate func(items []T) float64) Option[T] {
	return func(c *Combination[T]) {
		c.evaluate = evaluate
	}
}
