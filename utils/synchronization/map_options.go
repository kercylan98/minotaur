package synchronization

type MapOption[Key comparable, Value any] func(m *Map[Key, Value])

// WithMapSource 通过传入的 map 初始化
func WithMapSource[Key comparable, Value any](source map[Key]Value) MapOption[Key, Value] {
	return func(m *Map[Key, Value]) {
		m.data = source
	}
}