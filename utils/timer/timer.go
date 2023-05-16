package timer

import (
	"sync"

	"github.com/RussellLuo/timingwheel"
)

var timer = new(Timer)

func GetTicker(size int, options ...Option) *Ticker {
	ticker := timer.NewTicker(size)
	for _, option := range options {
		option(ticker)
	}
	return ticker
}

type Timer struct {
	tickers []*Ticker
	lock    sync.Mutex
}

func (slf *Timer) NewTicker(size int) *Ticker {
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
	ticker.wheel.Start()
	return ticker
}
