package vivid

import (
	"github.com/samber/do/v2"
	"reflect"
)

// ModLifeCycle 模组生命周期
type ModLifeCycle uint8

const (
	// ModLifeCycleOnInit 模组初始化阶段
	ModLifeCycleOnInit ModLifeCycle = iota + 1

	// ModLifeCycleOnPreload 模组预加载阶段
	ModLifeCycleOnPreload

	// ModLifeCycleOnMount 模组挂载阶段
	ModLifeCycleOnMount

	// ModLifeCycleOnStart 模组启动阶段
	ModLifeCycleOnStart

	// ModLifeCycleOnStop 模组停止阶段
	ModLifeCycleOnStop
)

// Mod 模组是用于对 Actor 进行功能扩展的接口
type Mod interface {
	OnLifeCycle(ctx ActorContext, lifeCycle ModLifeCycle)
}

type ModInfo interface {
	onLifeCycle(ctx ActorContext, lifeCycle ModLifeCycle)
	provide(injector do.Injector)
	getModType() reflect.Type
	setUnload()
	isUnload() bool
	shutdown()
}

type modInfo[T Mod] struct {
	mod             Mod
	modType         reflect.Type
	unload          bool
	shutdownHandler func()
}

func (m *modInfo[T]) onLifeCycle(ctx ActorContext, lifeCycle ModLifeCycle) {
	m.mod.OnLifeCycle(ctx, lifeCycle)
}

func (m *modInfo[T]) provide(injector do.Injector) {
	do.ProvideValue[T](injector, m.mod.(T))
	m.shutdownHandler = func() {
		do.MustShutdown[T](injector)
	}
}

func (m *modInfo[T]) getModType() reflect.Type {
	return m.modType
}

func (m *modInfo[T]) setUnload() {
	m.unload = true
}

func (m *modInfo[T]) isUnload() bool {
	return m.unload
}

func (m *modInfo[T]) shutdown() {
	m.shutdownHandler()
}

func ModOf[Interface Mod](mod Interface) ModInfo {
	return &modInfo[Interface]{
		mod:     mod,
		modType: reflect.TypeOf(mod),
	}
}

func InvokeMod[Interface Mod](ctx ActorContext) Interface {
	return do.MustInvoke[Interface](ctx.getCore().runtimeMods)
}
