package vivid

import "reflect"

// ActorCore 是 Actor 的核心接口定义，被用于高级功能的实现
type ActorCore interface {
	Actor
	ActorRef
	ActorContext
	ActorCoreExpose
}

// ActorCoreExpose 额外的暴露项
type ActorCoreExpose interface {
	GetOptions() *ActorOptions
}

func newActorCore(system *ActorSystem, actorId ActorId, actor Actor, opts *ActorOptions) *actorCore {
	core := &actorCore{
		Actor: actor,
		opts:  opts,
	}
	core.ActorRef = newLocalActorRef(system, actorId)
	core.actorContext = &actorContext{
		id:        actorId,
		vof:       reflect.Value{},
		state:     actorContextStatePreStart,
		behaviors: make(map[reflect.Type]reflect.Value),
	}
	core.actorContext.vof = reflect.ValueOf(core.actorContext)

	return core
}

type actorCore struct {
	Actor
	ActorRef
	*actorContext
	opts *ActorOptions
}

// onPreStart 在 Actor 启动之前执行的逻辑
func (a *actorCore) onPreStart() error {
	if err := a.Actor.OnPreStart(a); err != nil {
		return err
	}

	a.state = actorContextStateStarted
	return nil
}

// GetOptions 获取 Actor 的配置项
func (a *actorCore) GetOptions() *ActorOptions {
	return a.opts
}
