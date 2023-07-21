package concurrent

// BalanceMapOption BalanceMap 的选项
type BalanceMapOption[Key comparable, Value any] func(m *BalanceMap[Key, Value])

// WithBalanceMapSource 通过传入的 map 初始化
func WithBalanceMapSource[Key comparable, Value any](source map[Key]Value) BalanceMapOption[Key, Value] {
	return func(m *BalanceMap[Key, Value]) {
		m.data = source
	}
}
