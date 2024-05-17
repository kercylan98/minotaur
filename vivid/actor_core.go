package vivid

import (
	"fmt"
	"reflect"
)

// ActorCore 是 Actor 的核心接口定义，被用于高级功能的实现
type ActorCore interface {
	Actor
	ActorRef
	ActorContext
	ActorCoreExpose
}

// ActorCoreExpose 额外的暴露项
type ActorCoreExpose interface {
	// GetOptions 获取 Actor 的配置项
	GetOptions() *ActorOptions

	// SetContext 为 Actor 设置一个上下文
	SetContext(ctx any)

	// GetContext 获取 Actor 的上下文
	GetContext() any
}

func newActorCore(system *ActorSystem, actorId ActorId, actor Actor, opts *ActorOptions) *actorCore {
	core := &actorCore{
		Actor: actor,
		opts:  opts,
	}
	core.ActorRef = newLocalActorRef(system, actorId)
	core.actorContext = &actorContext{
		system:    system,
		id:        actorId,
		core:      core,
		state:     actorContextStatePreStart,
		behaviors: make(map[reflect.Type]reflect.Value),
		children:  map[ActorName]*actorCore{},
	}
	return core
}

type actorCore struct {
	Actor                       // 外部 Actor 实现
	ActorRef                    // Actor 的引用
	*actorContext               // Actor 的上下文
	opts          *ActorOptions // Actor 的配置项
	restartNum    int           // 重启次数
	ctx           any           // 外部上下文
}

// onPreStart 在 Actor 启动之前执行的逻辑
func (a *actorCore) onPreStart() (err error) {
	defer func() {
		if r := recover(); r != nil {
			a.NotifyTerminated(r)
			err = fmt.Errorf("%w: %v", ErrActorPreStart, r)
		}
	}()
	if err = a.Actor.OnPreStart(a); err != nil {
		return err
	}

	a.state = actorContextStateStarted
	return
}

// GetOptions 获取 Actor 的配置项
func (a *actorCore) GetOptions() *ActorOptions {
	return a.opts
}

// SetContext 为 Actor 设置一个上下文
func (a *actorCore) SetContext(ctx any) {
	a.ctx = ctx
}

// GetContext 获取 Actor 的上下文
func (a *actorCore) GetContext() any {
	return a.ctx
}
