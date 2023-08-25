package survey

// Flusher 用于刷新缓冲区的接口
type Flusher interface {
	// Flush 将缓冲区的数据持久化
	Flush(records []string)
	// Info 返回当前刷新器的信息
	Info() string
}
