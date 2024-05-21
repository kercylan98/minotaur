package vivid

import "sync"

type internalActorContext interface {
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
}

type _internalActorContext struct {
	*_ActorCore // ActorContext
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
