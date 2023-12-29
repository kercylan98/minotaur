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
