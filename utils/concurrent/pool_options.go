package concurrent

// PoolOption 线程安全的对象缓冲池选项
type PoolOption[T any] func(pool *Pool[T])

// WithPoolSilent 静默模式
//   - 静默模式下，当缓冲区大小不足时，将不再输出警告日志
func WithPoolSilent[T any]() PoolOption[T] {
	return func(pool *Pool[T]) {
		pool.silent = true
	}
}
