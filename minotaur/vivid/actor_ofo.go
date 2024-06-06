package vivid

// ActorOfO 该接口用于根据 ActorOptions 创建 Actor 并返回 ActorRef
//   - 接口用于内部实现，外部无法直接实现
type ActorOfO interface {
	generate(of ActorOwner) ActorRef
}

type actorOfO[T Actor] struct {
	options []func(actorOptions *ActorOptions[T])
}

func (o *actorOfO[T]) generate(of ActorOwner) ActorRef {
	var opts = new(ActorOptions[T])
	for _, opt := range o.options {
		opt(opts)
	}
	return ActorOf(of, opts)
}

// OfO 创建一个 ActorOfO 对象
func OfO[T Actor](options ...func(actorOptions *ActorOptions[T])) ActorOfO {
	return &actorOfO[T]{
		options: options,
	}
}
