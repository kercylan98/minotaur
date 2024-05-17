package vivid

import (
	"fmt"
	"reflect"
)

// ActorCore 是 Actor 的核心接口定义，被用于高级功能的实现
type ActorCore interface {
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

	// IsPause 判断 Actor 当前是否处于暂停处理消息的状态，返回一个只读的通道
	//  - 当 Actor 未处于暂停状态时，返回 nil
	IsPause() <-chan struct{}

	// BindMessageActorContext 绑定消息的 Actor 上下文为当前 Actor
	BindMessageActorContext(ctx MessageContext)

	// OnReceived 用于处理接收到的消息
	OnReceived(ctx MessageContext) (err error)
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
	pause         chan struct{} // 暂停处理消息
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

func (a *actorCore) GetOptions() *ActorOptions {
	return a.opts
}

func (a *actorCore) SetContext(ctx any) {
	a.ctx = ctx
}

func (a *actorCore) GetContext() any {
	return a.ctx
}

func (a *actorCore) IsPause() <-chan struct{} {
	if a.pause == nil {
		return nil
	}
	return a.pause
}

func (a *actorCore) BindMessageActorContext(ctx MessageContext) {
	ctx.(*messageContext).ActorContext = a
}

func (a *actorCore) OnReceived(ctx MessageContext) error {
	return a.Actor.OnReceived(ctx)
}
