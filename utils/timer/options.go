package timer

type Option func(ticker *Ticker)

// WithCaller 通过其他的 handle 执行 Caller
func WithCaller(handle func(name string, caller func())) Option {
	return func(ticker *Ticker) {
		ticker.handle = handle
	}
}
