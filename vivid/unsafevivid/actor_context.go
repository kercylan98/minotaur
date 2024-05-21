package unsafevivid

import (
	"context"
	"github.com/kercylan98/minotaur/toolkit/log"
	"github.com/kercylan98/minotaur/vivid/vivids"
	"reflect"
)

// ActorContext 是 Actor 的上下文
type ActorContext struct {
	context.Context                                 // 上下文
	System          *ActorSystem                    // Actor 所属的 Actor 系统
	Core            *ActorCore                      // Actor 的核心
	Parent          *ActorCore                      // 父 Actor
	Children        map[vivids.ActorName]*ActorCore // 子 Actor
	Id              vivids.ActorId                  // Actor 的 ID
	IsEnd           bool                            // 是否是末级 Actor
	Started         bool                            // 是否已经启动
	Behaviors       map[reflect.Type]reflect.Value  // 消息处理器
}

func (c *ActorContext) RegisterBehavior(message vivids.Message, behavior any) {
	if c.Started {
		panic("register behavior after actor started")
	}
	msgType := reflect.TypeOf(message)
	behaviorValue := reflect.ValueOf(behavior)
	if behaviorValue.Kind() != reflect.Func {
		panic("behavior must be a function")
	}
	if _, exist := c.Behaviors[msgType]; exist {
		panic("behavior already registered")
	}
	if behaviorValue.Type().NumIn() != 2 {
		panic("behavior must have two parameters")
	}
	if behaviorValue.Type().In(0) != reflect.TypeOf(c) {
		panic("behavior first parameter must be ActorContext")
	}
	if behaviorValue.Type().In(1) != msgType {
		panic("behavior second parameter must be message")
	}
	c.Behaviors[msgType] = behaviorValue
}

// GetSystem 获取 Actor 所属的 Actor 系统
func (c *ActorContext) GetSystem() vivids.ActorSystem {
	return c.System
}

// restart 重启上下文
func (c *ActorContext) restart(reentry bool) (recovery func(), err error) {
	// 尝试获取快照，获取失败则算作重启失败
	actorSnapshot, err := c.Core.Actor.OnSaveSnapshot(c.Core)
	if err != nil {
		return nil, err
	}

	// 重启所有子 Actor
	if !reentry {
		c.System.actorsRW.RLock()
		defer c.System.actorsRW.RUnlock()
	}
	var recoveries []func()
	for _, child := range c.Children {
		if recovery, err = child.restart(true); err != nil {
			for _, f := range recoveries {
				f()
			}
			return nil, err
		}
		recoveries = append(recoveries, recovery)
	}

	var recoveryFunc = func() {
		// TODO: 恢复快照失败如何处理？
		if err := c.Core.OnRecoverSnapshot(c, actorSnapshot); err != nil {
			log.Error("actor recover snapshot failed", log.Err(err))
		}
	}
	if err = c.Core.Actor.OnDestroy(c.Core); err != nil {
		recoveryFunc()
		return nil, err
	}
	if err = c.Core.onPreStart(); err != nil {
		recoveryFunc()
		return nil, err
	}
	return recoveryFunc, nil
}

func (c *ActorContext) bindChildren(core *ActorCore) {
	c.Children[core.GetOptions().Name] = core
}

func (c *ActorContext) ActorOf(actor vivids.Actor, opts ...*vivids.ActorOptions) (vivids.ActorRef, error) {
	var opt *vivids.ActorOptions
	if len(opts) > 0 {
		opt = opts[0]
		if opt == nil {
			opt = vivids.NewActorOptions()
		}
	} else {
		opt = vivids.NewActorOptions()
	}
	opt = opt.WithParent(c)

	return c.System.generateActor(c, actor, opt, !c.Started)
}

func (c *ActorContext) GetActor() vivids.Query {
	return NewQuery(c.System, c.Core)
}

func (c *ActorContext) NotifyTerminated(v ...vivids.Message) {
	terminatedContext := NewActorTerminatedContext(c.Core, v...)
	c.Parent.OnChildTerminated(c.Parent, terminatedContext)
	if terminatedContext.cancelTerminate {
		return
	}
	c.System.releaseActor(c.Core, !c.Core.Started)
}

func (c *ActorContext) GetParentActor() vivids.ActorRef {
	return c.Parent.ActorRef
}

func (c *ActorContext) GetActorId() vivids.ActorId {
	return c.Id
}

func (c *ActorContext) Future(handler func() vivids.Message) vivids.Future {
	return NewFuture(c, handler)
}

func (c *ActorContext) PublishEvent(event vivids.Message) {
	c.Core.EventRW.RLock()
	defer c.Core.EventRW.RUnlock()

	eventType := reflect.TypeOf(event)
	if _, exist := c.Core.Events[eventType]; !exist {
		return
	}

	var err error
	for actorId := range c.Core.Events[eventType] {
		if err = c.System.Tell(actorId, event); err != nil {
			log.Error("publish event failed", log.Err(err))
		}
	}
}

func (c *ActorContext) GetProps() any {
	return c.Core.GetOptions().Props
}
