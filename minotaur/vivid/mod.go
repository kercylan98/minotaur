package vivid

import (
	"github.com/samber/do/v2"
	"reflect"
)

// ModLifecycle 模组生命周期，模组在生命周期中的不同阶段将会通过 Mod.OnLifeCycle 函数进行回调
//
// 生命周期阶段包含如下几个阶段：
//   - ModLifecycleOnInit 初始化阶段
//   - ModLifecycleOnPreload 预加载阶段
//   - ModLifecycleOnMount 挂载阶段
//   - ModLifecycleOnStart 启动阶段
//   - ModLifecycleOnStop 停止阶段
type ModLifecycle uint8

const (
	// ModLifecycleOnInit 模组在被应用后加载的首个阶段，在该阶段模组应该完成对自身非依赖性的初始化工作
	//  - 该阶段不应该依赖其他任何模组
	ModLifecycleOnInit ModLifecycle = iota + 1

	// ModLifecycleOnPreload 模组自身初始化完成后，进入预加载阶段，该阶段可以完成对依赖模组的依赖注入
	//  - 需要注意的是，在该阶段，被依赖的模组可能还未初始化完成，因此在该阶段不应该使用依赖模组的功能
	ModLifecycleOnPreload

	// ModLifecycleOnMount 模组在注入依赖模组后，进入挂载阶段，该阶段可以完成需要依赖其他模组的初始化工作，即二次初始化
	ModLifecycleOnMount

	// ModLifecycleOnStart 模组启动阶段，该阶段所有模组均已初始化完成，可以在该阶段完成的功能定义
	ModLifecycleOnStart

	// ModLifecycleOnStop 模组在被释放时，将会进入停止阶段，该阶段可以完成模组的资源释放工作
	ModLifecycleOnStop
)

const (
	modStatusWaiting modStatus = iota // 等待加载
	modStatusLoaded                   // 已加载
	modStatusUnload                   // 已卸载
)

type modStatus uint8

type ModUnloadedHandler func(mod ModInfo)

// Mod 模组是用于对 Actor 在生命周期中的功能扩展
type Mod interface {
	// OnLifecycle 模组生命周期回调，当模组生命周期发生变化时，将会调用该函数
	//  - 关于生命周期的详细信息可参考 ModLifecycle
	OnLifecycle(ctx ActorContext, lifeCycle ModLifecycle)
}

// ModOf 定义一个模组，返回该模组可被用于 Actor 加载、卸载模块的信息
//   - Exposer 应提供该模组的接口，用于暴露模组的功能给其他模组进行使用，当 Exposer 不为接口类型时，可能会发生意想不到的情况
func ModOf[Exposer Mod](mod Exposer) ModInfo {
	return &modInfo[Exposer]{
		mod:     mod,
		modType: reflect.TypeOf(mod),
	}
}

// InvokeMod 调用模组，返回模组的实例，如果模组未被加载，将会 panic
//   - Exposer 模组的暴露接口
func InvokeMod[Exposer Mod](ctx ActorContext) Exposer {
	return do.MustInvoke[Exposer](ctx.getCore().runtimeMods)
}

// ModInfo 描述了模组的信息，该接口不对外暴露，仅供内部使用
type ModInfo interface {
	onLifecycle(ctx ActorContext, lifeCycle ModLifecycle)
	provide(ctx ActorContext, injector do.Injector)
	getModType() reflect.Type
	setStatus(status modStatus)
	getStatus() modStatus
	shutdown()

	// TypeEqual 判断两个模组是否为同一类型
	TypeEqual(other Mod) bool
}

type modInfo[T Mod] struct {
	mod             Mod
	modType         reflect.Type
	status          modStatus
	shutdownHandler func()
}

func (m *modInfo[T]) TypeEqual(other Mod) bool {
	return reflect.TypeOf(other) == m.modType
}

func (m *modInfo[T]) onLifecycle(ctx ActorContext, lifeCycle ModLifecycle) {
	m.mod.OnLifecycle(ctx, lifeCycle)
}

func (m *modInfo[T]) provide(ctx ActorContext, injector do.Injector) {
	do.ProvideValue[T](injector, m.mod.(T))
	m.shutdownHandler = func() {
		m.onLifecycle(ctx, ModLifecycleOnStop)
		do.MustShutdown[T](injector)
	}
}

func (m *modInfo[T]) getModType() reflect.Type {
	return m.modType
}

func (m *modInfo[T]) setStatus(status modStatus) {
	m.status = status
}

func (m *modInfo[T]) getStatus() modStatus {
	return m.status
}

func (m *modInfo[T]) shutdown() {
	m.shutdownHandler()
}
