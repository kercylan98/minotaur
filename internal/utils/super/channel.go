package super

// TryWriteChannel 尝试写入 channel，如果 channel 无法写入则忽略，返回是否写入成功
//   - 无法写入的情况包括：channel 已满、channel 已关闭
func TryWriteChannel[T any](ch chan<- T, data T) bool {
	select {
	case ch <- data:
		return true
	default:
		return false
	}
}

// TryWriteChannelByHandler 尝试写入 channel，如果 channel 无法写入则执行 handler
//   - 无法写入的情况包括：channel 已满、channel 已关闭
func TryWriteChannelByHandler[T any](ch chan<- T, data T, handler func()) {
	select {
	case ch <- data:
	default:
		handler()
	}
}

// TryReadChannel 尝试读取 channel，如果 channel 无法读取则忽略，返回是否读取成功
//   - 无法读取的情况包括：channel 已空、channel 已关闭
func TryReadChannel[T any](ch <-chan T) (v T, suc bool) {
	select {
	case data := <-ch:
		return data, true
	default:
		return v, false
	}
}

// TryReadChannelByHandler 尝试读取 channel，如果 channel 无法读取则执行 handler
//   - 无法读取的情况包括：channel 已空、channel 已关闭
func TryReadChannelByHandler[T any](ch <-chan T, handler func(ch <-chan T) T) (v T) {
	select {
	case data := <-ch:
		return data
	default:
		return handler(ch)
	}
}
