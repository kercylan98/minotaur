package vivid

import (
	"context"
	"github.com/kercylan98/minotaur/toolkit/log"
	"github.com/samber/do/v2"
	"reflect"
	"time"
)

// ActorContext 针对 Actor 的上下文，该上下文暴露给 Actor 自身使用，但不提供外部自行实现
//   - 上下文代表了 Actor 完整的生命周期，该上下文将在 Actor 的生命周期中一直存在
type ActorContext interface {
	context.Context
	internalActorContext
	ActorRef
	// GetId 获取当前 ActorContext 的 ID
	GetId() ActorId

	// GetSystem 获取当前 ActorContext 所属的 ActorSystem
	GetSystem() *ActorSystem

	// GetActor 获取当前 ActorContext 的 Actor 对象
	GetActor() Actor

	// GetParent 获取当前上下文的父 Actor 的引用
	GetParent() ActorRef

	// Become 切换 Actor 在面对特定消息的行为。通过调用 BehaviorOf 函数可以创建一个特定消息类型的行为，将 Actor 的消息处理逻辑切换为指定的新行为。新的行为会覆盖当前的行为，直到下次调用 Become 或 UnBecome 为止
	Become(behavior Behavior, discardOld ...bool)

	// UnBecome 恢复 Actor 在面对特性消息的行为为之前的行为，多次调用 UnBecome 会依次恢复之前的行为，直到没有行为为止
	UnBecome(message Message)

	// ActorOf 创建一个 Actor 并返回 ActorRef
	//  - ActorOfO 对象可通过 OfO 函数快速创建
	ActorOf(ofo ActorOfO) ActorRef

	// Terminate 销毁当前 Actor，该操作将会触发 Actor 的 OnTerminate 生命周期
	Terminate()

	// Subscribe 订阅事件
	Subscribe(event Message, options ...SubscribeOption)

	// Unsubscribe 取消订阅事件
	Unsubscribe(event Message)

	// Publish 发布事件
	Publish(event Message)

	// LoadMod 加载模组
	LoadMod(mods ...ModInfo)

	// UnloadMod 卸载模组
	UnloadMod(mods ...ModInfo)

	// ApplyMod 应用模组
	ApplyMod()

	// SetIdleTimeout 设置 Actor 的空闲超时时间，当 Actor 被重启时，该值不会被使用
	SetIdleTimeout(timeout time.Duration)
}

const (
	actorStatusRunning actorStatus = iota
	actorStatusTerminated
)

type actorStatus = int32

type _ActorContext struct {
	*_internalActorContext
	*_ActorCore
	behaviors   map[reflect.Type][]Behavior // 行为栈，用于存储 Actor 在面对特定消息时的行为
	mods        []ModInfo                   // 声明的模组
	currentMods []ModInfo                   // 当前生命周期的模组
	runtimeMods do.Injector                 // 运行时模组
	status      actorStatus                 // 是否已经终止
}

func (c *_ActorContext) GetId() ActorId {
	return c.id
}

func (c *_ActorContext) GetSystem() *ActorSystem {
	return c.system
}

func (c *_ActorContext) GetActor() Actor {
	return c.Actor
}

func (c *_ActorContext) GetParent() ActorRef {
	return c.parent
}

func (c *_ActorContext) Become(behavior Behavior, discardOld ...bool) {
	if c.behaviors == nil {
		c.behaviors = make(map[reflect.Type][]Behavior)
	}

	messageType := behavior.getMessageType()
	if len(discardOld) > 0 && discardOld[0] {
		c.behaviors[messageType] = c.behaviors[messageType][:0]
	}
	c.behaviors[messageType] = append([]Behavior{behavior}, c.behaviors[messageType]...)
}

func (c *_ActorContext) UnBecome(message Message) {
	if c.behaviors == nil {
		return
	}

	messageType := reflect.TypeOf(message)
	if behaviors, ok := c.behaviors[messageType]; ok {
		if len(behaviors) > 0 {
			behaviors = behaviors[1:]
			c.behaviors[messageType] = behaviors
		}
		if len(behaviors) == 0 {
			delete(c.behaviors, messageType)
		}
	}
}

func (c *_ActorContext) ActorOf(ofo ActorOfO) ActorRef {
	return ofo.generate(c)
}

func (c *_ActorContext) Terminate() {
	c.Tell(OnTerminate{})
}

func (c *_ActorContext) Subscribe(event Message, options ...SubscribeOption) {
	c.system.Subscribe(c, event, options...)
}

func (c *_ActorContext) Unsubscribe(event Message) {
	c.system.Unsubscribe(c, event)
}

func (c *_ActorContext) Publish(event Message) {
	c.system.Publish(c, event)
}

func (c *_ActorContext) LoadMod(mods ...ModInfo) {
	c.UnloadMod(mods...)
	c.mods = append(c.mods, mods...)
}

func (c *_ActorContext) UnloadMod(mods ...ModInfo) {
	// 标记相同类型的模组为卸载状态
	for _, mod := range mods {
		modType := mod.getModType()
		for _, m := range c.mods {
			if m.getModType() == modType {
				m.setStatus(modStatusUnload)
				break
			}
		}
	}
}

func (c *_ActorContext) ApplyMod() {
	if c.runtimeMods == nil {
		c.runtimeMods = do.New()
	}
	var currentMods []ModInfo
	for _, mod := range c.mods {
		switch mod.getStatus() {
		case modStatusWaiting:
			c.getSystem().GetLogger().Debug("LoadMod", log.String("actor", c.GetId().String()), log.String("mod", mod.getModType().String()))
			mod.provide(c, c.runtimeMods)
			mod.setStatus(modStatusLoaded)
			currentMods = append(currentMods, mod)
		case modStatusLoaded:
			currentMods = append(currentMods, mod)
		case modStatusUnload:
			c.getSystem().GetLogger().Debug("UnloadMod", log.String("actor", c.GetId().String()), log.String("mod", mod.getModType().String()))
			mod.shutdown()
		}
	}
	c.currentMods = currentMods

	for i := ModLifeCycleOnInit; i <= ModLifeCycleOnStart; i++ {
		for _, mod := range c.currentMods {
			mod.onLifeCycle(c, i)
		}
	}
	c.getSystem().GetLogger().Debug("ApplyMod", log.String("actor", c.GetId().String()), log.Int("count", len(c.currentMods)))

	c.mods = currentMods
}

func (c *_ActorContext) SetIdleTimeout(timeout time.Duration) {
	c.idleTimeout = timeout
}
