package timer

import (
	"fmt"
	"sync"

	"github.com/RussellLuo/timingwheel"
)

var (
	tickerPoolSize = DefaultTickerPoolSize
	standardTimer  = NewTimer(tickerPoolSize)
)

// SetTickerPoolSize 设置定时器池大小
//   - 默认值为 DefaultTickerPoolSize，当定时器池中的定时器不足时，会自动创建新的定时器，当定时器释放时，会将多余的定时器进行释放，否则将放入定时器池中
func SetTickerPoolSize(size int) {
	_ = standardTimer.ChangeTickerPoolSize(size)
}

func GetTicker(size int, options ...Option) *Ticker {
	return standardTimer.NewTicker(size, options...)
}

func NewTimer(tickerPoolSize int) *Timer {
	if tickerPoolSize <= 0 {
		panic(fmt.Errorf("timer tickerPoolSize must greater than 0, got: %d", tickerPoolSize))
	}
	return &Timer{
		tickerPoolSize: tickerPoolSize,
	}
}

type Timer struct {
	tickers        []*Ticker
	lock           sync.Mutex
	tickerPoolSize int
}

// ChangeTickerPoolSize 改变定时器池大小
//   - 当传入的大小小于或等于 0 时，将会返回错误，并且不会发生任何改变
func (slf *Timer) ChangeTickerPoolSize(size int) error {
	if size <= 0 {
		return fmt.Errorf("timer tickerPoolSize must greater than 0, got: %d", tickerPoolSize)
	}
	slf.lock.Lock()
	defer slf.lock.Unlock()
	slf.tickerPoolSize = size
	return nil
}

// NewTicker 获取一个新的定时器
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
