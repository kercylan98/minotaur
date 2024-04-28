package modular

import (
	"github.com/samber/do/v2"
)

// NewApplication 创建一个新的模块化应用程序
func NewApplication() *Application {
	a := &Application{
		injector: do.New(),
	}
	return a
}

// Application 模块化应用程序
type Application struct {
	injector *do.RootScope   // 依赖注入器
	services []GlobalService // 服务列表
}

// run 启动应用程序
func (a *Application) run() {
	startLifecycle(a.services).
		next("onInit", func(service GlobalService) bool {
			service.OnInit(a)
			return true
		}).
		next("onPreload", func(service GlobalService) bool {
			service.OnPreload(a)
			return true
		}).
		next("onMount", func(service GlobalService) bool {
			service.OnMount(a)
			return true
		}).
		next("onStart", func(service GlobalService) bool {
			service.OnStart(a)
			return true
		}).
		next("onBlock", func(service GlobalService) bool {
			block, ok := service.(BlockService)
			if ok {
				block.OnBlock(a)
			}
			return ok
		}).
		run()
}
