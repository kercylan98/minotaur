package actors

func NewActorTerminatedContext(c *actorCore, v ...Message) *ActorTerminatedContext {
	return &ActorTerminatedContext{
		core:               c,
		terminatedMessages: v,
	}
}

type ActorTerminatedContext struct {
	core               *actorCore // Actor 核心
	terminatedMessages []Message  // 销毁消息
	cancelTerminate    bool       // 是否取消销毁
}

func (c *ActorTerminatedContext) GetActorId() ActorId {
	return c.core.id
}

func (c *ActorTerminatedContext) GetActor() any {
	return c.core.Actor
}

func (c *ActorTerminatedContext) HasTerminatedMessage() bool {
	return len(c.terminatedMessages) > 0
}

func (c *ActorTerminatedContext) GetTerminatedMessage() Message {
	if len(c.terminatedMessages) == 1 {
		return c.terminatedMessages[0]
	} else {
		return c.terminatedMessages
	}
}

func (c *ActorTerminatedContext) Restart() error {
	c.core.restartNum++
	if _, err := c.core.restart(false); err != nil {
		return err
	}
	c.core.restartNum = 0
	c.cancelTerminate = true
	return nil
}

func (c *ActorTerminatedContext) Recover() error {
	c.cancelTerminate = true
	return nil
}

func (c *ActorTerminatedContext) IsPreStart() bool {
	return c.core.state == actorContextStatePreStart
}

func (c *ActorTerminatedContext) GetRestartNum() int {
	return c.core.restartNum
}
