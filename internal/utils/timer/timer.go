package timer

var (
	tickerPoolSize = DefaultTickerPoolSize
	standardPool   = NewPool(tickerPoolSize)
)

// SetPoolSize 设置标准池定时器池大小
//   - 默认值为 DefaultTickerPoolSize，当定时器池中的定时器不足时，会自动创建新的定时器，当定时器释放时，会将多余的定时器进行释放，否则将放入定时器池中
func SetPoolSize(size int) {
	_ = standardPool.ChangePoolSize(size)
}

// GetTicker 获取标准池中的一个定时器
func GetTicker(size int, options ...Option) *Ticker {
	return standardPool.GetTicker(size, options...)
}
