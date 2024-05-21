package vivid

type ActorRef interface {
	Tell(msg Message, opts ...MessageOption)
	Ask(msg Message, opts ...MessageOption) Message
}

// _LocalActorRef 本地 Actor 引用
type _LocalActorRef struct {
	core *_ActorCore // Actor 核心
}

func (r *_LocalActorRef) Tell(msg Message, opts ...MessageOption) {
	r.core.system.sendMessage(r.core._LocalActorRef, msg, opts...)
}

func (r *_LocalActorRef) Ask(msg Message, opts ...MessageOption) Message {
	return r.core.system.sendMessage(r.core._LocalActorRef, msg, append(opts, func(options *MessageOptions) {
		options.reply = true
	})...)
}

// _RemoteActorRef 远程 Actor 引用
type _RemoteActorRef struct {
	system  *ActorSystem // Actor 系统
	actorId ActorId      // 远程 Actor ID
}

func (r *_RemoteActorRef) Tell(msg Message, opts ...MessageOption) {
	r.system.sendMessage(r, msg, opts...)
}

func (r *_RemoteActorRef) Ask(msg Message, opts ...MessageOption) Message {
	return r.system.sendMessage(r, msg, append(opts, func(options *MessageOptions) {
		options.reply = true
	})...)
}

// _DeadLetterActorRef 死信 Actor 引用
type _DeadLetterActorRef struct {
	system *ActorSystem // Actor 系统
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
