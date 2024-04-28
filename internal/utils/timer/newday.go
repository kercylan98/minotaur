package timer

import (
	"github.com/kercylan98/minotaur/utils/offset"
	"github.com/kercylan98/minotaur/utils/times"
	"time"
)

type (
	SystemNewDayEventHandle     func()
	OffsetTimeNewDayEventHandle func()
)

var (
	systemNewDayEventHandles     = make(map[string][]SystemNewDayEventHandle)
	offsetTimeNewDayEventHandles = make(map[string][]OffsetTimeNewDayEventHandle)
)

// RegSystemNewDayEvent 注册系统新的一天事件
//   - 建议全局注册一个事件后再另行拓展
//   - 将特定 name 的定时任务注册到 ticker 中，在系统时间到达每天的 00:00:00 时触发，如果 trigger 为 true，则立即触发一次
func RegSystemNewDayEvent(ticker *Ticker, name string, trigger bool, handle SystemNewDayEventHandle) {
	ticker.Loop(name, times.GetNextDayInterval(time.Now()), times.Day, Forever, OnSystemNewDayEvent, name)
	systemNewDayEventHandles[name] = append(systemNewDayEventHandles[name], handle)
	if trigger {
		OnSystemNewDayEvent(name)
	}
}

// OnSystemNewDayEvent 系统新的一天事件
func OnSystemNewDayEvent(name string) {
	for _, handle := range systemNewDayEventHandles[name] {
		handle()
	}
}

// RegOffsetTimeNewDayEvent 注册偏移时间新的一天事件
//   - 建议全局注册一个事件后再另行拓展
//   - 与 RegSystemNewDayEvent 类似，但是触发时间为 offset 时间到达每天的 00:00:00
func RegOffsetTimeNewDayEvent(ticker *Ticker, name string, offset *offset.Time, trigger bool, handle OffsetTimeNewDayEventHandle) {
	ticker.Loop(name, times.GetNextDayInterval(offset.Now()), times.Day, Forever, OnOffsetTimeNewDayEvent, name)
	offsetTimeNewDayEventHandles[name] = append(offsetTimeNewDayEventHandles[name], handle)
	if trigger {
		OnOffsetTimeNewDayEvent(name)
	}
}

// OnOffsetTimeNewDayEvent 偏移时间新的一天事件
func OnOffsetTimeNewDayEvent(name string) {
	for _, handle := range offsetTimeNewDayEventHandles[name] {
		handle()
	}
}
