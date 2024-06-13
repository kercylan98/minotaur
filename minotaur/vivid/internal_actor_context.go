package vivid

import (
	"fmt"
	"reflect"
	"sync"
)

type internalActorContext interface {
	ActorOwner

	// getLock 获取当前 ActorContext 的锁
	//   - 所有函数均不操作锁，应由外部调用者自行操作
	getLockable() *sync.RWMutex

	// hasChild 判断当前 ActorContext 是否存在指定名字的子 Actor
	hasChild(name ActorName) bool

	// getChildren 获取当前 ActorContext 的所有子 Actor
	getChildren() map[ActorName]*_ActorCore

	// bindChild 绑定子 Actor
	bindChild(name ActorName, child *_ActorCore)

	// unbindChild 解绑子 Actor
	//   - 该函数不会删除子 Actor，只会解除绑定关系
	unbindChild(name ActorName)

	// matchBehavior 匹配行为
	matchBehavior(message Message) Behavior

	// getParent 获取父 Actor
	getParent() ActorContext

	// supervisorExec 监管策略执行
	supervisorExec(core *_ActorCore, message, reason Message)

	// getCore 获取 ActorCore
	getCore() *_ActorCore
}

type _internalActorContext struct {
	*_ActorCore // ActorContext
}

func (c *_internalActorContext) getSystem() *ActorSystem {
	return c.system
}

func (c *_internalActorContext) getContext() *_ActorCore {
	return c._ActorCore
}

func (c *_internalActorContext) getLockable() *sync.RWMutex {
	return c.childrenRW
}

func (c *_internalActorContext) hasChild(name ActorName) bool {
	_, ok := c.children[name]
	return ok
}

func (c *_internalActorContext) bindChild(name ActorName, child *_ActorCore) {
	c.children[name] = child
	child.parent = c
}

func (c *_internalActorContext) unbindChild(name ActorName) {
	delete(c.children, name)
}

func (c *_internalActorContext) getChildren() map[ActorName]*_ActorCore {
	return c.children
}

func (c *_internalActorContext) matchBehavior(message Message) Behavior {
	if c.behaviors == nil {
		return nil
	}

	behaviors, ok := c.behaviors[reflect.TypeOf(message)]
	if !ok {
		return nil
	}

	return behaviors[len(behaviors)-1]
}

func (c *_internalActorContext) getParent() ActorContext {
	return c.parent
}

func (c *_internalActorContext) supervisorExec(core *_ActorCore, message, reason Message) {
	if c._ActorCore.supervisor != nil {
		directive := c._ActorCore.supervisor(message, reason)
		switch directive {
		case DirectiveStop:
			core.Terminate()
		case DirectiveRestart:
			core.Tell(OnTerminate{
				restart: true,
			})
		case DirectiveResume:
			// ignore
		case DirectiveEscalate:
			c.getParent().supervisorExec(core, message, reason)
		default:
			panic(fmt.Errorf("unknown directive: %d", directive))
		}
	} else {
		c.getParent().supervisorExec(core, message, reason)
	}
}

func (c *_internalActorContext) getCore() *_ActorCore {
	return c._ActorCore
}
