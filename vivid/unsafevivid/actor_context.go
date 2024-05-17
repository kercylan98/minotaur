package unsafevivid

import (
	"context"
	"github.com/kercylan98/minotaur/toolkit/log"
	vivid "github.com/kercylan98/minotaur/vivid/vivids"
)

// ActorContext 是 Actor 的上下文
type ActorContext struct {
	context.Context                                // 上下文
	System          *ActorSystem                   // Actor 所属的 Actor 系统
	Core            *ActorCore                     // Actor 的核心
	Parent          *ActorCore                     // 父 Actor
	Children        map[vivid.ActorName]*ActorCore // 子 Actor
	Id              vivid.ActorId                  // Actor 的 ID
	IsEnd           bool                           // 是否是末级 Actor
	Started         bool                           // 是否已经启动
}

// GetSystem 获取 Actor 所属的 Actor 系统
func (c *ActorContext) GetSystem() vivid.ActorSystem {
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

func (c *ActorContext) ActorOf(actor vivid.Actor, opts ...*vivid.ActorOptions) (vivid.ActorRef, error) {
	var opt *vivid.ActorOptions
	if len(opts) > 0 {
		opt = opts[0]
		if opt == nil {
			opt = vivid.NewActorOptions()
		}
	} else {
		opt = vivid.NewActorOptions()
	}
	opt = opt.WithParent(c)

	return c.System.generateActor(c, actor, opt, !c.Started)
}

func (c *ActorContext) GetActor() vivid.Query {
	return NewQuery(c.System, c.Core)
}

func (c *ActorContext) NotifyTerminated(v ...vivid.Message) {
	terminatedContext := NewActorTerminatedContext(c.Core, v...)
	c.Parent.OnChildTerminated(c.Parent, terminatedContext)
	if terminatedContext.cancelTerminate {
		return
	}
	c.System.releaseActor(c.Core, !c.Core.Started)
}

func (c *ActorContext) GetParentActor() vivid.ActorRef {
	return c.Parent.ActorRef
}

func (c *ActorContext) GetActorId() vivid.ActorId {
	return c.Id
}

func (c *ActorContext) Future(handler func() vivid.Message) vivid.Future {
	return NewFuture(c, handler)
}
