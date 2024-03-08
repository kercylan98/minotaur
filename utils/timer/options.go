package timer

type Option func(ticker *Ticker)

// WithCaller 通过其他的 handler 执行 Caller
func WithCaller(handle func(name string, caller func())) Option {
	return func(ticker *Ticker) {
		ticker.handler = handle
	}
}

// WithMark 通过特定的标记创建定时器
func WithMark(mark string) Option {
	return func(ticker *Ticker) {
		ticker.mark = mark
	}
}
