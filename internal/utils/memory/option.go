package memory

import (
	"github.com/kercylan98/minotaur/utils/timer"
	"time"
)

func NewOption() *Option {
	return new(Option)
}

type Option struct {
	ticker     *timer.Ticker
	firstDelay time.Duration // 首次持久化延迟
	interval   time.Duration // 持久化间隔
	delay      time.Duration // 每条数据持久化间隔
}

// WithPeriodicity 设置持久化周期
//   - ticker 定时器，通常建议使用服务器的定时器，这样可以降低多线程的程序复杂性
//   - firstDelay 首次持久化延迟，当首次持久化为 0 时，将会在下一个持久化周期开始时持久化
//   - interval 持久化间隔
//   - delay 每条数据持久化间隔，适当的设置该值可以使持久化期间尽量降低对用户体验的影响，如果为0，将会一次性持久化所有数据
func (slf *Option) WithPeriodicity(ticker *timer.Ticker, firstDelay, interval, delay time.Duration) *Option {
	if interval <= 0 {
		panic("interval must be greater than 0")
	}
	slf.ticker = ticker
	slf.firstDelay = firstDelay
	slf.interval = interval
	slf.delay = delay
	return slf
}
