package vivid

// ActorRef 是 Actor 的引用
type ActorRef interface {
	// GetId 用于获取 Actor 的 ActorId
	GetId() ActorId

	// Tell 用于向 Actor 发送消息
	Tell(msg Message, opts ...MessageOption) error

	// Ask 用于向 Actor 发送消息并等待返回结果
	Ask(msg Message, opts ...MessageOption) (Message, error)
}

func newLocalActorRef(system *ActorSystem, actorId ActorId) *localActorRef {
	return &localActorRef{
		actorRef: actorRef{
			system:  system,
			actorId: actorId,
		},
	}
}

func newRemoteActorRef(system *ActorSystem, actorId ActorId) *remoteActorRef {
	return &remoteActorRef{
		actorRef: actorRef{
			system:  system,
			actorId: actorId,
		},
	}
}

type actorRef struct {
	system  *ActorSystem
	actorId ActorId
}

func (a *actorRef) GetId() ActorId {
	return a.actorId
}

func (a *actorRef) Tell(msg Message, opts ...MessageOption) error {
	return a.system.tell(a.GetId(), msg, opts...)
}

func (a *actorRef) Ask(msg Message, opts ...MessageOption) (Message, error) {
	return a.system.ask(a.GetId(), msg, opts...)
}

// localActorRef 实现 Actor 模型的核心逻辑
type localActorRef struct {
	actorRef
}

// remoteActorRef 实现 Actor 模型的远程调用逻辑
type remoteActorRef struct {
	actorRef
}
