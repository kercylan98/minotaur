package buried

type Logger interface {
	// Write 写入埋点条目到日志记录器中，通常来说该函数应当简单存储到内存即可
	Write(name string, entries ...Entries)

	// Flush 将内存中的日志记录器数据刷入磁盘或网络
	Flush() error
}
