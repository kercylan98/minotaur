package timer

import (
	"fmt"
	"sync"

	"github.com/RussellLuo/timingwheel"
)

// NewPool 创建一个定时器池，当 tickerPoolSize 小于等于 0 时，将会引发 panic，可指定为 DefaultTickerPoolSize
func NewPool(tickerPoolSize int) *Pool {
	if tickerPoolSize <= 0 {
		panic(fmt.Errorf("timer tickerPoolSize must greater than 0, got: %d", tickerPoolSize))
	}
	return &Pool{
		tickerPoolSize: tickerPoolSize,
	}
}

// Pool 定时器池
type Pool struct {
	tickers        []*Ticker
	lock           sync.Mutex
	tickerPoolSize int
}

// ChangePoolSize 改变定时器池大小
//   - 当传入的大小小于或等于 0 时，将会返回错误，并且不会发生任何改变
func (slf *Pool) ChangePoolSize(size int) error {
	if size <= 0 {
		return fmt.Errorf("timer tickerPoolSize must greater than 0, got: %d", tickerPoolSize)
	}
	slf.lock.Lock()
	defer slf.lock.Unlock()
	slf.tickerPoolSize = size
	return nil
}

// GetTicker 获取一个新的定时器
func (slf *Pool) GetTicker(size int, options ...Option) *Ticker {
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
