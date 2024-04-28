package configuration

import (
	"github.com/kercylan98/minotaur/utils/log"
	"github.com/kercylan98/minotaur/utils/timer"
	"time"
)

const (
	tickerLoadRefresh = "_tickerLoadRefresh"
)

var (
	cTicker   *timer.Ticker
	cInterval time.Duration
	cLoader   []Loader
)

// Init 配置初始化
//   - 在初始化后会立即加载配置
func Init(loader ...Loader) {
	cLoader = loader
	Load()
	Refresh()
}

// Load 加载配置
//   - 加载后并不会刷新线上配置，需要执行 Refresh 函数对线上配置进行刷新
func Load() {
	defer func() {
		if err := recover(); err != nil {
			log.Error("Config", log.String("Action", "Load"), log.Err(err.(error)))
		}
	}()
	for _, loader := range cLoader {
		loader.Load()
	}
}

// Refresh 刷新配置
func Refresh() {
	defer func() {
		if err := recover(); err != nil {
			log.Error("Config", log.String("Action", "Refresh"), log.Err(err.(error)))
		}
		OnConfigRefreshEvent()
	}()
	for _, loader := range cLoader {
		loader.Refresh()
	}
}

// WithTickerLoad 通过定时器加载配置
//   - 通过定时器加载配置后，会自动刷新线上配置
//   - 调用该函数后不会立即刷新，而是在 interval 后加载并刷新一次配置，之后每隔 interval 加载并刷新一次配置
func WithTickerLoad(ticker *timer.Ticker, interval time.Duration) {
	if ticker != cTicker && cTicker != nil {
		cTicker.StopTimer(tickerLoadRefresh)
	}
	cTicker = ticker
	cInterval = interval
	cTicker.Loop(tickerLoadRefresh, cInterval, cInterval, timer.Forever, func() {
		Load()
		Refresh()
	})
}

// StopTickerLoad 停止通过定时器加载配置
func StopTickerLoad() {
	if cTicker != nil {
		cTicker.StopTimer(tickerLoadRefresh)
	}
}
