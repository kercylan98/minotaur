package vivid

import (
	"reflect"
)

var actorType = reflect.TypeOf((*Actor)(nil)).Elem()

// Actor 是 Actor 模型的接口，该接口用于定义一个 Actor
type Actor interface {
	// OnPreStart 在 Actor 启动之前执行的逻辑，适用于对 Actor 状态的初始化
	OnPreStart(ctx ActorContext) (err error)

	// OnReceived 当 Actor 接收到消息时执行的逻辑
	OnReceived(ctx MessageContext) (err error)

	// OnDestroy 当 Actor 被要求销毁时将会调用该函数，需要在该函数中释放 Actor 的资源
	//  - 该函数可能会在重启前被调用，被用于重置 Actor 的状态
	OnDestroy(ctx ActorContext) (err error)

	// OnSaveSnapshot 当 Actor 被要求保存快照时将会调用该函数
	OnSaveSnapshot(ctx ActorContext) (snapshot []byte, err error)

	// OnRecoverSnapshot 当 Actor 被要求恢复快照时将会调用该函数
	OnRecoverSnapshot(ctx ActorContext, snapshot []byte) (err error)

	// OnChildTerminated 当 Actor 的子 Actor 被销毁时将会调用该函数
	OnChildTerminated(ctx ActorContext, child ActorTerminatedContext)
}

// ActorTerminatedContext 是 Actor 销毁时的上下文
type ActorTerminatedContext interface {
	// GetActorId 获取 Actor 的 ID
	GetActorId() ActorId

	// GetActor 获取 Actor 原始对象
	//  - 常被用于类型断言进行不同 Actor 类型的处理
	GetActor() any

	// HasTerminatedMessage 判断是否有销毁消息
	HasTerminatedMessage() bool

	// GetTerminatedMessage 获取销毁消息
	GetTerminatedMessage() Message

	// Restart 以全新的状态重启 Actor，包括所有的子 Actor
	//  - 该函数将会一次执行 Actor.OnSaveSnapshot、 Actor.OnDestroy、 Actor.OnPreStart 三个函数来完成重启
	//  - 当重启过程中发生错误时将会通过 Actor.OnRecoverSnapshot 函数来恢复 Actor 的状态
	Restart() error

	// Recover 保留当前的状态恢复 Actor
	Recover() error

	// IsPreStart 是否在启动之前发生的销毁
	IsPreStart() bool

	// GetRestartNum 获取重启次数，该值将在重启成功后清零
	GetRestartNum() int
}

func newActorTerminatedContext(c *actorCore, v ...Message) *actorTerminatedContext {
	return &actorTerminatedContext{
		core:               c,
		terminatedMessages: v,
	}
}

type actorTerminatedContext struct {
	core               *actorCore // Actor 核心
	terminatedMessages []Message  // 销毁消息
	cancelTerminate    bool       // 是否取消销毁
}

func (c *actorTerminatedContext) GetActorId() ActorId {
	return c.core.id
}

func (c *actorTerminatedContext) GetActor() any {
	return c.core.Actor
}

func (c *actorTerminatedContext) HasTerminatedMessage() bool {
	return len(c.terminatedMessages) > 0
}

func (c *actorTerminatedContext) GetTerminatedMessage() Message {
	if len(c.terminatedMessages) == 1 {
		return c.terminatedMessages[0]
	} else {
		return c.terminatedMessages
	}
}

func (c *actorTerminatedContext) Restart() error {
	c.core.restartNum++
	if _, err := c.core.restart(false); err != nil {
		return err
	}
	c.core.restartNum = 0
	c.cancelTerminate = true
	return nil
}

func (c *actorTerminatedContext) Recover() error {
	c.cancelTerminate = true
	return nil
}

func (c *actorTerminatedContext) IsPreStart() bool {
	return c.core.state == actorContextStatePreStart
}

func (c *actorTerminatedContext) GetRestartNum() int {
	return c.core.restartNum
}
