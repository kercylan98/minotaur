package vivid

import (
	"reflect"
)

// Context 上下文对象接口
type Context interface {
	GetActor() Actor
}

// BehaviorOf 创建一个行为
func BehaviorOf[T Message](handler func(ctx MessageContext, message T)) Behavior {
	adp := &behaviorAdapter[T]{
		typ:     reflect.TypeOf((*T)(nil)).Elem(),
		handler: handler,
	}
	return adp
}

// ActorOfI 使用一个已有的 Actor 实例创建 Actor 并返回 ActorRef
func ActorOfI[T Actor](actorOf ActorOwner, actor T, options ...func(options *ActorOptions[T])) ActorRef {
	var opts = NewActorOptions[T]().WithConstruct(actor)
	for _, opt := range options {
		opt(opts)
	}
	return ActorOf(actorOf, opts)
}

// ActorOfIT 使用一个已有的 Actor 实例创建 Actor 并返回 ActorRef 的类型化引用
func ActorOfIT[A Actor, T ActorTyped](actorOf ActorOwner, actor A, options ...func(options *ActorOptions[A])) T {
	ref := ActorOfI(actorOf, actor, options...)
	ask := ref.Ask(onActorRefTyped{
		ref: ref,
	}, WithInstantly(true))
	typed := ask.(T)
	ref.Tell(OnActorTyped[T]{Typed: typed}, WithInstantly(true))
	return typed
}

// ActorOfF 使用传入的可选项创建 Actor 并返回 ActorRef，该函数简化了对于泛型的声明
func ActorOfF[T Actor](actorOf ActorOwner, options ...func(options *ActorOptions[T])) ActorRef {
	var opts = NewActorOptions[T]()
	for _, opt := range options {
		opt(opts)
	}
	return ActorOf(actorOf, opts)
}

// ActorOfFT 使用传入的可选项创建 Actor 并返回 ActorRef 的类型化引用，该函数简化了对于泛型的声明
func ActorOfFT[A Actor, T ActorTyped](actorOf ActorOwner, options ...func(options *ActorOptions[A])) T {
	ref := ActorOfF(actorOf, options...)
	ask := ref.Ask(onActorRefTyped{
		ref: ref,
	}, WithInstantly(true))
	typed := ask.(T)
	ref.Tell(OnActorTyped[T]{Typed: typed}, WithInstantly(true))
	return typed
}

// ActorOfT 使用传入的可选项创建 Actor 并返回 ActorRef 类型化引用
func ActorOfT[A Actor, T ActorTyped](actorOf ActorOwner, options ...*ActorOptions[A]) T {
	ref := ActorOf(actorOf, options...)
	ask := ref.Ask(onActorRefTyped{
		ref: ref,
	}, WithInstantly(true))
	typed := ask.(T)
	ref.Tell(OnActorTyped[T]{Typed: typed}, WithInstantly(true))
	return typed
}

// ActorOf 使用传入的可选项创建 Actor 并返回 ActorRef
func ActorOf[T Actor](actorOf ActorOwner, options ...*ActorOptions[T]) ActorRef {
	var opts = parseActorOptions(options...)
	var ins = opts.Construct
	if reflect.ValueOf(ins).IsNil() {
		tof := reflect.TypeOf((*T)(nil)).Elem().Elem()
		ins = reflect.New(tof).Interface().(T)
	}
	var system = actorOf.GetSystem()

	if opts.Parent == nil {
		opts.Parent = actorOf.GetContext()
	}

	ctx, err := generateActor[T](actorOf.GetSystem(), ins, opts, false)
	if err != nil {
		system.deadLetters.DeadLetter(NewDeadLetterEvent(DeadLetterEventTypeActorOf, DeadLetterEventActorOf{
			Error:  err,
			Parent: opts.Parent,
			Name:   opts.Name,
		}))
	}
	if ctx == nil {
		return &_DeadLetterActorRef{
			system: system,
		}
	}
	return ctx
}
