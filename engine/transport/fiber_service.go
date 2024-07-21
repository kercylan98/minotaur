package transport

import (
	"github.com/kercylan98/minotaur/engine/vivid"
)

type FiberService interface {
	// OnInit 该函数将在 Actor 启动后，Fiber 运行前进行调用，用于注册路由等初始化操作。
	//  - ctx 为 Fiber Actor 的上下文。
	OnInit(ctx vivid.ActorContext, app *FiberApp)
}

type FunctionalFiberService func(ctx vivid.ActorContext, app *FiberApp)

func (f FunctionalFiberService) OnInit(ctx vivid.ActorContext, app *FiberApp) {
	f(ctx, app)
}
