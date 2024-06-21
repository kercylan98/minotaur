package vivid

type ActorRef interface {
	// Id 获取 Actor ID
	Id() ActorId

	// Tell 向 Actor 发送一条消息，当消息发送失败时将会进入死信队列中，以下列举一些特殊场景
	//  - 当接收人处理消息失败或不存在时，将会进入到接收人所在的 ActorSystem 的死信队列中
	//  - 当发送人已经销毁或不存在时，将会进入到发送人所在的 ActorSystem 的死信队列中
	//  - 当自己给自己发送消息时，消息不会立即执行，而是会进入到自己的邮箱中，等待下一次消息循环时执行
	Tell(msg Message, opts ...MessageOption)

	// Ask 向 Actor 发送消息并等待回复
	//  - 当消息发送失败或等待超时的时候将会返回 nil，并进入死信队列中
	//  - 当接收人处理消息失败、回复失败或不存在时，将会进入到接收人所在的 ActorSystem 的死信队列中
	//  - 当发送人已经销毁或不存在时，将会进入到发送人所在的 ActorSystem 的死信队列中
	//  - 当给自己给自己发送消息时需特别注意，自己的邮箱会收到消息，但是由于不会立即执行，所以回复始终会等待到超时位置，而超时后收到的消息将被执行，回复将无效
	Ask(msg Message, opts ...MessageOption) Message

	// Stop 停止 Actor
	Stop(ctx ...any)

	// GetSystem 获取 Actor 所在的 ActorSystem
	GetSystem() *ActorSystem

	// 内部发送消息实现
	send(ctx MessageContext)
}

// _LocalActorRef 本地 Actor 引用
type _LocalActorRef struct {
	core *_ActorCore // Actor 核心
}

func (r *_LocalActorRef) Id() ActorId {
	return r.core._LocalActorRef.core.GetId()
}

func (r *_LocalActorRef) Tell(msg Message, opts ...MessageOption) {
	r.core.system.sendMessage(r.core._LocalActorRef, msg, opts...)
}

func (r *_LocalActorRef) Ask(msg Message, opts ...MessageOption) Message {
	return r.core.system.sendMessage(r.core._LocalActorRef, msg, append(opts, func(options *MessageOptions) {
		options.reply = true
	})...)
}

func (r *_LocalActorRef) Stop(ctx ...any) {
	switch len(ctx) {
	case 0:
		r.Tell(OnTerminate{})
	case 1:
		r.Tell(OnTerminate{Context: ctx[0]})
	default:
		r.Tell(OnTerminate{Context: ctx})
	}
}

func (r *_LocalActorRef) GetSystem() *ActorSystem {
	return r.core.system
}

func (r *_LocalActorRef) send(ctx MessageContext) {
	r.core.system.sendToDispatcher(r.core.dispatcher, r.core, ctx)
}

func newRemoteActorRef(system *ActorSystem, actorId ActorId) *_RemoteActorRef {
	return &_RemoteActorRef{
		system:  system,
		actorId: actorId,
	}
}

// _RemoteActorRef 远程 Actor 引用
type _RemoteActorRef struct {
	system  *ActorSystem // Actor 系统
	actorId ActorId      // 远程 Actor ID
}

func (r *_RemoteActorRef) Id() ActorId {
	return r.actorId
}

func (r *_RemoteActorRef) Tell(msg Message, opts ...MessageOption) {
	r.system.sendMessage(r, msg, opts...)
}

func (r *_RemoteActorRef) Ask(msg Message, opts ...MessageOption) Message {
	return r.system.sendMessage(r, msg, append(opts, func(options *MessageOptions) {
		options.reply = true
	})...)
}

func (r *_RemoteActorRef) Stop(ctx ...any) {
	switch len(ctx) {
	case 0:
		r.Tell(OnTerminate{})
	case 1:
		r.Tell(OnTerminate{Context: ctx[0]})
	default:
		r.Tell(OnTerminate{Context: ctx})
	}
}

func (r *_RemoteActorRef) GetSystem() *ActorSystem {
	return r.system
}

func (r *_RemoteActorRef) send(ctx MessageContext) {
	data, err := r.system.codec.Encode(ctx)
	if err != nil {
		return
	}
	if err = r.system.cluster.SendToNode(r.actorId.Address(), data); err != nil {
		//panic(err)
	}
}

// newDeadLetterActorRef 创建一个新的死信 Actor 引用
func newDeadLetterActorRef(system *ActorSystem) *_DeadLetterActorRef {
	return &_DeadLetterActorRef{
		system: system,
	}
}

// _DeadLetterActorRef 死信 Actor 引用
type _DeadLetterActorRef struct {
	system *ActorSystem // Actor 系统
}

func (r *_DeadLetterActorRef) Id() ActorId {
	return ""
}

func (r *_DeadLetterActorRef) Tell(msg Message, opts ...MessageOption) {
	r.system.deadLetters.DeadLetter(NewDeadLetterEvent(DeadLetterEventTypeMessage, DeadLetterEventMessage{
		Error:   ErrActorDeadOrNotExist,
		Message: msg,
	}))
}

func (r *_DeadLetterActorRef) Ask(msg Message, opts ...MessageOption) Message {
	r.system.deadLetters.DeadLetter(NewDeadLetterEvent(DeadLetterEventTypeMessage, DeadLetterEventMessage{
		Error:   ErrActorDeadOrNotExist,
		Message: msg,
	}))
	return nil
}

func (r *_DeadLetterActorRef) Stop(ctx ...any) {
	switch len(ctx) {
	case 0:
		r.Tell(OnTerminate{})
	case 1:
		r.Tell(OnTerminate{Context: ctx[0]})
	default:
		r.Tell(OnTerminate{Context: ctx})
	}
}

func (r *_DeadLetterActorRef) GetSystem() *ActorSystem {
	return r.system
}

func (r *_DeadLetterActorRef) send(ctx MessageContext) {
	panic("dead letter actor ref can't send message")
}

// newNoSenderActorRef 创建一个新的无发送者 Actor 引用
func newNoSenderActorRef(system *ActorSystem) ActorRef {
	return &_NoSenderActorRef{
		_DeadLetterActorRef: newDeadLetterActorRef(system),
	}
}

// _NoSenderActorRef 无发送者 Actor 引用
type _NoSenderActorRef struct {
	*_DeadLetterActorRef
}
