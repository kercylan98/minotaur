package transport

import (
	"github.com/gofiber/fiber/v2"
	"reflect"
)

func newFiberApp(app *fiber.App) *FiberApp {
	f := &FiberApp{
		App:     app,
		exposes: map[reflect.Type]FiberService{},
	}
	f.hooks = newFiberHooks(f)
	return f
}

type FiberApp struct {
	*fiber.App
	hooks   *FiberHooks
	exposes map[reflect.Type]FiberService
}

// ExposeService 暴露服务使其透过接口对外公开，暴露其可通过 transport.FiberExpose 函数生成
func (f *FiberApp) ExposeService(exposer *FiberExposer) {
	f.exposes[exposer.k] = exposer.v
}

// ProvideService 提供特定服务注入到 service 中
func (f *FiberApp) ProvideService(provider FiberProvider) FiberService {
	return f.exposes[reflect.Type(provider)]
}

func (f *FiberApp) Hooks() *FiberHooks {
	return f.hooks
}
