package vivid

import (
	"context"
	"time"
)

type MessageContext interface {
	context.Context
	ActorOwner
	ActorRef

	// GetContext 获取 Actor 上下文
	GetContext() ActorContext

	// GetSeq 获取消息的序号
	GetSeq() uint64

	// GetSender 获取消息的发送者
	GetSender() ActorRef

	// GetRef 获取上下文的 ActorRef
	GetRef() ActorRef

	// GetMessage 获取消息内容
	GetMessage() Message

	// Reply 回复消息
	Reply(Message)

	// GetActor 获取 Actor 对象，该函数是 ActorContext.GetActor 的快捷方式
	GetActor() Actor

	// GetPriority 获取消息的优先级
	GetPriority() int64

	// Instantly 是否是立即执行的消息
	Instantly() bool

	// Become 该函数是 ActorContext.Become 的快捷方式
	Become(behavior Behavior)

	// UnBecome 该函数是 ActorContext.UnBecome 的快捷方式
	UnBecome(message Message)

	// ActorOf 创建一个 Actor 对象，该函数是 ActorContext.ActorOf 的快捷方式
	ActorOf(ofo ActorOfO) ActorRef

	// Subscribe 订阅事件，该函数是 ActorContext.Subscribe 的快捷方式
	Subscribe(event Message, options ...SubscribeOption)

	// Unsubscribe 取消订阅事件，该函数是 ActorContext.Unsubscribe 的快捷方式
	Unsubscribe(event Message)

	// Publish 发布事件，该函数是 ActorContext.Publish 的快捷方式
	Publish(event Message)
}

func newMessageContext(system *ActorSystem, message Message, priority int64, instantly, hasReply bool) *_MessageContext {
	return &_MessageContext{
		system:        system,
		Seq:           system.messageSeq.Add(1),
		Network:       system.network,
		Host:          system.host,
		Port:          system.port,
		Message:       message,
		Priority:      priority,
		InstantlyExec: instantly,
		HasReply:      hasReply,
	}
}

// _MessageContext 消息上下文，消息上下文实现了兼容本地及远程消息的上下文
//   - 该结构体中，除开公共信息外，内部字段被用于本地消息，公开字段被用于远程消息，需要保证公共及公开字段的可序列化
type _MessageContext struct {
	system   *ActorSystem // 创建上下文的 Actor 系统
	Seq      uint64       // 消息序号
	Network  string       // 产生消息的网络
	Host     string       // 产生消息的主机
	Port     uint16       // 产生消息的端口
	Message  Message      // 消息内容
	Priority int64        // 消息优先级
	HasReply bool         // 是否需要回复

	// 本地消息是直接根据实现了 ActorRef 的 _ActorCore 来投递的，所以可以直接将消息投递到 ActorCore 绑定的 Dispatcher 中
	actorContext ActorContext // 本地接收者的上下文
	sender       ActorRef     // 本地发送者，通过 WithSender 设置后，由于是本地消息，所以可以直接使用 ActorRef

	// 远程消息是通过网络传输的，所以需要将接收者的 ActorId 传递到远程消息的上下文中，以便在远程消息到达后，能够根据 ActorId 获取到 ActorRef
	ReceiverId     ActorId // 远程接收者的 ActorId
	SenderId       ActorId // 远程发送者的 ActorId，通过 WithSender 设置后，由于是远程消息，所以需要将发送者的 ActorId 传递到远程消息的上下文中
	RemoteReplySeq uint64  // 用于远程回复的消息序号
	InstantlyExec  bool    // 是否立即执行
}

func (c *_MessageContext) GetSystem() *ActorSystem {
	return c.system
}

func (c *_MessageContext) Id() ActorId {
	return c.GetRef().Id()
}

func (c *_MessageContext) Tell(msg Message, opts ...MessageOption) {
	c.GetRef().Tell(msg, opts...)
}

func (c *_MessageContext) Ask(msg Message, opts ...MessageOption) Message {
	return c.GetRef().Ask(msg, opts...)
}

func (c *_MessageContext) Stop() {
	c.GetRef().Stop()
}

func (c *_MessageContext) send(ctx MessageContext) {
	c.GetRef().send(ctx)
}

func (c *_MessageContext) Deadline() (deadline time.Time, ok bool) {
	return c.actorContext.Deadline()
}

func (c *_MessageContext) Done() <-chan struct{} {
	return c.actorContext.Done()
}

func (c *_MessageContext) Err() error {
	return c.actorContext.Err()
}

func (c *_MessageContext) Value(key any) any {
	return c.actorContext.Value(key)
}

func (c *_MessageContext) getSystem() *ActorSystem {
	return c.system
}

func (c *_MessageContext) getContext() *_ActorCore {
	return c.GetContext().(*_ActorCore)
}

func (c *_MessageContext) withLocal(receiver *_ActorCore, sender ActorRef) *_MessageContext {
	c.actorContext = receiver
	c.sender = sender
	return c
}

func (c *_MessageContext) withRemote(receiverId ActorId, senderId ActorId) *_MessageContext {
	c.ReceiverId = receiverId
	c.SenderId = senderId
	return c
}

func (c *_MessageContext) GetContext() ActorContext {
	if c.actorContext != nil {
		return c.actorContext
	}

	if c.ReceiverId.IsLocal(c.system) {
		if receiver := c.system.getActor(c.ReceiverId); receiver != nil {
			c.actorContext = receiver
			return receiver
		}
		return nil
	}

	return nil
}

func (c *_MessageContext) GetSeq() uint64 {
	return c.Seq
}

func (c *_MessageContext) GetSender() ActorRef {
	if c.sender != nil {
		return c.sender
	}

	if c.SenderId.Invalid() {
		return newNoSenderActorRef(c.system)
	}

	if c.SenderId.IsLocal(c.system) {
		if c.sender = c.system.getActor(c.SenderId); c.sender == nil {
			c.sender = newDeadLetterActorRef(c.system)
		}
		return c.sender
	}

	c.sender = newRemoteActorRef(c.system, c.SenderId)
	return c.sender
}

func (c *_MessageContext) GetRef() ActorRef {
	if c.actorContext != nil {
		return c.actorContext.(*_ActorCore)._LocalActorRef
	}

	if c.ReceiverId.IsLocal(c.system) {
		if receiver := c.system.getActor(c.ReceiverId); receiver != nil {
			c.actorContext = receiver
			return receiver._LocalActorRef
		}
		return newDeadLetterActorRef(c.system)
	}

	return newRemoteActorRef(c.system, c.ReceiverId)
}

func (c *_MessageContext) GetMessage() Message {
	return c.Message
}

func (c *_MessageContext) Reply(message Message) {
	if !c.HasReply {
		return
	}
	// 本地消息回复直接投递到对应 waiter 中，远程消息则通过网络发送
	sender := c.GetSender()
	localSender, isLocal := sender.(*_LocalActorRef)
	noSender, isNoSender := sender.(*_NoSenderActorRef)
	if isLocal || (isNoSender && c.Network == c.system.network && c.Host == c.system.host && c.Port == c.system.port) {
		var system *ActorSystem
		if isLocal {
			system = localSender.core.system
		} else {
			system = noSender.system
		}
		system.askWaitsLock.RLock()
		waiter, exist := system.askWaits[c.Seq]
		system.askWaitsLock.RUnlock()
		if exist {
			waiter <- message
		}
		// TODO 死信
		return
	}

	// 远程回复
	sender.send(&_MessageContext{
		Message:        message,
		RemoteReplySeq: c.Seq,
	})
}

func (c *_MessageContext) GetActor() Actor {
	return c.GetContext().GetActor()
}

func (c *_MessageContext) GetPriority() int64 {
	return c.Priority
}

func (c *_MessageContext) Instantly() bool {
	return c.InstantlyExec
}

func (c *_MessageContext) Become(behavior Behavior) {
	c.GetContext().Become(behavior)
}

func (c *_MessageContext) UnBecome(message Message) {
	c.GetContext().UnBecome(message)
}

func (c *_MessageContext) ActorOf(ofo ActorOfO) ActorRef {
	return c.GetContext().ActorOf(ofo)
}

func (c *_MessageContext) Subscribe(event Message, options ...SubscribeOption) {
	c.GetContext().Subscribe(event, options...)
}

func (c *_MessageContext) Unsubscribe(event Message) {
	c.GetContext().Unsubscribe(event)
}

func (c *_MessageContext) Publish(event Message) {
	c.GetContext().Publish(event)
}
