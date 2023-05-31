package reporter

import (
	"github.com/kercylan98/minotaur/utils/synchronization"
	"github.com/kercylan98/minotaur/utils/timer"
)

var (
	ticker          *timer.Ticker // 定时器
	tickerIsDefault bool
	registerBuried  *synchronization.Map[string, bool] // 已注册的数据埋点
	disableBuried   *synchronization.Map[string, bool] // 已排除的数据埋点
)

func init() {
	ticker = timer.GetTicker(50)
	tickerIsDefault = true
	registerBuried = synchronization.NewMap[string, bool]()
	disableBuried = synchronization.NewMap[string, bool]()
}

// UseTicker 使用特定定时器取代默认上报定时器
func UseTicker(t *timer.Ticker) {
	if t == nil {
		return
	}
	if tickerIsDefault {
		tickerIsDefault = false
		ticker.Release()
		ticker = t
	}
}

// DisableBuried 禁用特定数据埋点
func DisableBuried[Data any](buried *Buried[Data]) {
	DisableBuriedWithName(buried.GetName())
}

// DisableBuriedWithName 禁用特定名称的数据埋点
func DisableBuriedWithName(name string) {
	disableBuried.Set(name, true)
}

// EnableBuried 启用特定数据埋点
func EnableBuried[Data any](buried *Buried[Data]) {
	EnableBuriedWithName(buried.GetName())
}

// EnableBuriedWithName 启用特定名称的数据埋点
func EnableBuriedWithName(name string) {
	disableBuried.Delete(name)
}
