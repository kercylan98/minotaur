package timer

import (
	"sync"

	"github.com/RussellLuo/timingwheel"
)

var tickerPoolSize = 96

var timer = new(Timer)

// SetTickerPoolSize 设置定时器池大小
//   - 默认值为 96，当定时器池中的定时器不足时，会自动创建新的定时器，当定时器释放时，会将多余的定时器进行释放，否则将放入定时器池中
func SetTickerPoolSize(size int) {
	if size <= 0 {
		panic("ticker pool size must be greater than 0")
	}
	timer.lock.Lock()
	defer timer.lock.Unlock()
	tickerPoolSize = size
}

func GetTicker(size int, options ...Option) *Ticker {
	return timer.NewTicker(size, options...)
}

type Timer struct {
	tickers []*Ticker
	lock    sync.Mutex
}

func (slf *Timer) NewTicker(size int, options ...Option) *Ticker {
	slf.lock.Lock()
	defer slf.lock.Unlock()

	var ticker *Ticker
	if len(slf.tickers) > 0 {
		ticker = slf.tickers[0]
		slf.tickers = slf.tickers[1:]
		return ticker
	}

	ticker = &Ticker{
		timer:  slf,
		wheel:  timingwheel.NewTimingWheel(timingWheelTick, int64(size)),
		timers: make(map[string]*Scheduler),
	}
	for _, option := range options {
		option(ticker)
	}
	ticker.wheel.Start()
	return ticker
}
