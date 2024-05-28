package vivid

import "reflect"

// Flow 消息流，可用于控制消息的过滤及流向
type Flow interface {
	// forward 消息是否需要转发
	forward(m Message) bool

	// dest 获取目标 Actor
	dest() ActorRef
}

func FlowOf[T Message](system *ActorSystem, source ActorRef, target ActorRef, filter ...func(T) bool) Flow {
	f := &flow[T]{
		source: source,
		target: target,
	}
	if len(filter) > 0 {
		f.filter = filter[0]
	}

	system.flowRW.Lock()
	defer system.flowRW.Unlock()

	flows, exist := system.flows[source.Id()]
	if !exist {
		flows = map[reflect.Type]Flow{}
		system.flows[source.Id()] = flows
	}

	flows[reflect.TypeOf((*T)(nil)).Elem()] = f
	return f
}

type flow[T Message] struct {
	source ActorRef
	target ActorRef
	filter func(T) bool
}

func (f *flow[T]) forward(m Message) bool {
	return f.filter == nil || f.filter(m.(T))
}

func (f *flow[T]) dest() ActorRef {
	return f.target
}
