package vivid

import (
	"github.com/kercylan98/minotaur/engine/vivid/dispatcher"
	"github.com/kercylan98/minotaur/toolkit"
	"github.com/panjf2000/ants/v2"
	"time"
)

// DispatcherProvider 是一个提供 dispatcher.Dispatcher 实例的接口
type DispatcherProvider interface {
	// Provide 返回一个可用的 dispatcher.Dispatcher 实例
	Provide() dispatcher.Dispatcher
}

// FunctionalDispatcherProvider 是一个函数类型的 DispatcherProvider，它定义了生成 dispatcher.Dispatcher 实例的方法。
type FunctionalDispatcherProvider func() dispatcher.Dispatcher

// Provide 返回一个可用的 dispatcher.Dispatcher 实例
func (f FunctionalDispatcherProvider) Provide() dispatcher.Dispatcher {
	return f()
}

var defaultDispatcherProviderInstance = new(defaultDispatcherProvider)
var defaultDispatcherSingleton = toolkit.NewInertiaSingleton[dispatcher.Dispatcher](func() dispatcher.Dispatcher {
	ad, err := dispatcher.NewAnts(ants.DefaultAntsPoolSize,
		ants.WithNonblocking(true),
		ants.WithExpiryDuration(time.Second*3),
	)
	if err != nil {
		return dispatcher.NewGoroutine()
	}
	return ad
})

// GetDefaultDispatcherProvider 返回默认的 DispatcherProvider 实例
func GetDefaultDispatcherProvider() DispatcherProvider {
	return defaultDispatcherProviderInstance
}

type defaultDispatcherProvider struct{}

func (d *defaultDispatcherProvider) Provide() dispatcher.Dispatcher {
	return defaultDispatcherSingleton.Get()
}
