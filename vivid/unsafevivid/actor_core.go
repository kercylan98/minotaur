package unsafevivid

import (
	"context"
	"fmt"
	vivid "github.com/kercylan98/minotaur/vivid/vivids"
	"reflect"
	"sync"
)

func NewActorCore(ctx context.Context, system *ActorSystem, actorId vivid.ActorId, actor vivid.Actor, opts *vivid.ActorOptions) *ActorCore {
	core := &ActorCore{
		Actor:    actor,
		Options:  opts,
		ActorRef: NewActorRef(system, actorId),
		Events:   map[reflect.Type]map[vivid.ActorId]struct{}{},
	}
	core.ActorContext = &ActorContext{
		Context:   ctx,
		System:    system,
		Id:        actorId,
		Core:      core,
		Children:  map[vivid.ActorName]*ActorCore{},
		Behaviors: map[reflect.Type]reflect.Value{},
	}
	return core
}

type ActorCore struct {
	vivid.Actor                                                // 外部 Actor 实现
	*ActorRef                                                  // Actor 的引用
	*ActorContext                                              // Actor 的上下文
	Options        *vivid.ActorOptions                         // Actor 的配置项
	RestartNum     int                                         // 重启次数
	OutsideContext sync.Map                                    // 外部上下文
	Pause          chan struct{}                               // 暂停处理消息
	Events         map[reflect.Type]map[vivid.ActorId]struct{} // 事件订阅
	EventRW        sync.RWMutex                                // 事件订阅读写锁
}

// onPreStart 在 Actor 启动之前执行的逻辑
func (a *ActorCore) onPreStart() (err error) {
	defer func() {
		if r := recover(); r != nil {
			a.NotifyTerminated(r)
			err = fmt.Errorf("%w: %v", vivid.ErrActorPreStart, r)
		}
	}()
	if err = a.Actor.OnPreStart(a); err != nil {
		return err
	}

	a.Started = true
	return
}

func (a *ActorCore) GetOptions() *vivid.ActorOptions {
	return a.Options
}

func (a *ActorCore) SetContext(key, value any) {
	a.OutsideContext.Store(key, value)
}

func (a *ActorCore) GetContext(key any) any {
	v, _ := a.OutsideContext.Load(key)
	return v
}

func (a *ActorCore) IsPause() <-chan struct{} {
	if a.Pause == nil {
		return nil
	}
	return a.Pause
}

func (a *ActorCore) BindMessageActorContext(ctx vivid.MessageContext) {
	ctx.(*MessageContext).ActorContext = a
}

func (a *ActorCore) OnReceived(ctx vivid.MessageContext) error {
	switch v := ctx.GetMessage().(type) {
	case SubscribeEventMessage:
		eventType := reflect.TypeOf(v.Event)
		a.EventRW.Lock()
		defer a.EventRW.Unlock()
		subscribers, exist := a.Events[eventType]
		if !exist {
			subscribers = map[vivid.ActorId]struct{}{}
			a.Events[eventType] = subscribers
		}
		subscribers[v.Subscriber] = struct{}{}
		return nil
	default:
		return a.Actor.OnReceived(ctx)
	}
}
