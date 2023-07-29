package combination

// CombinationOption 组合器选项
type CombinationOption[T Item] func(*Combination[T])

// WithCombinationEvaluation 设置组合评估函数
//   - 用于对组合进行评估，返回一个分值的评价函数
//   - 通过该选项将设置所有匹配器的默认评估函数为该函数
//   - 通过匹配器选项 WithMatcherEvaluation 可以覆盖该默认评估函数
//   - 默认的评估函数将返回一个随机数
func WithCombinationEvaluation[T Item](evaluate func(items []T) float64) CombinationOption[T] {
	return func(c *Combination[T]) {
		c.evaluate = evaluate
	}
}
