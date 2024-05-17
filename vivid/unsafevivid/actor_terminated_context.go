package unsafevivid

import vivid "github.com/kercylan98/minotaur/vivid/vivids"

func NewActorTerminatedContext(c *ActorCore, v ...vivid.Message) *ActorTerminatedContext {
	return &ActorTerminatedContext{
		core:               c,
		terminatedMessages: v,
	}
}

type ActorTerminatedContext struct {
	core               *ActorCore      // Actor 核心
	terminatedMessages []vivid.Message // 销毁消息
	cancelTerminate    bool            // 是否取消销毁
}

func (c *ActorTerminatedContext) GetActorId() vivid.ActorId {
	return c.core.GetId()
}

func (c *ActorTerminatedContext) GetActor() any {
	return c.core.Actor
}

func (c *ActorTerminatedContext) HasTerminatedMessage() bool {
	return len(c.terminatedMessages) > 0
}

func (c *ActorTerminatedContext) GetTerminatedMessage() vivid.Message {
	if len(c.terminatedMessages) == 1 {
		return c.terminatedMessages[0]
	} else {
		return c.terminatedMessages
	}
}

func (c *ActorTerminatedContext) Restart() error {
	c.core.RestartNum++
	if _, err := c.core.restart(false); err != nil {
		return err
	}
	c.core.RestartNum = 0
	c.cancelTerminate = true
	return nil
}

func (c *ActorTerminatedContext) Recover() error {
	c.cancelTerminate = true
	return nil
}

func (c *ActorTerminatedContext) GetRestartNum() int {
	return c.core.RestartNum
}
