package config

// RefreshEventHandle 配置刷新事件处理函数
type RefreshEventHandle func()

var configRefreshEventHandles []func()

// RegConfigRefreshEvent 当配置刷新时将立即执行被注册的事件处理函数
func RegConfigRefreshEvent(handle RefreshEventHandle) {
	configRefreshEventHandles = append(configRefreshEventHandles, handle)
}

func OnConfigRefreshEvent() {
	for _, handle := range configRefreshEventHandles {
		handle()
	}
}
