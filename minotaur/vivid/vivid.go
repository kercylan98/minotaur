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

func ActorOfI[T Actor](actorOf ActorOwner, actor T, options ...func(options *ActorOptions[T])) ActorRef {
	var opts = NewActorOptions[T]().WithConstruct(actor)
	for _, opt := range options {
		opt(opts)
	}
	return ActorOf(actorOf, opts)
}

func ActorOfF[T Actor](actorOf ActorOwner, options ...func(options *ActorOptions[T])) ActorRef {
	var opts = NewActorOptions[T]()
	for _, opt := range options {
		opt(opts)
	}
	return ActorOf(actorOf, opts)
}

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

func FreeActorOf[T any](actorOf ActorOwner, options ...*ActorOptions[*FreeActor[T]]) ActorRef {
	var opts = parseActorOptions(options...)
	var ins = opts.Construct.actor
	if reflect.ValueOf(ins).IsNil() {
		tof := reflect.TypeOf((*T)(nil)).Elem().Elem()
		ins = reflect.New(tof).Interface().(T)
	}
	var system = actorOf.GetSystem()

	if opts.Parent == nil {
		opts.Parent = actorOf.GetContext()
	}

	ctx, err := generateActor[*FreeActor[T]](actorOf.GetSystem(), &FreeActor[T]{actor: ins}, opts, false)
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

// GetActor 快捷的将 Actor 上下文转换为对应的 Actor 对象。该函数兼容 FreeActorOf 创建的对象
//   - 需要注意的是，当被转换的对象类型不匹配时将会返回空指针
func GetActor[T any](ctx Context) (t T) {
	switch v := ctx.GetActor().(type) {
	case *FreeActor[T]:
		return v.GetActor()
	default:
		t, _ = v.(T)
		return t
	}
}
