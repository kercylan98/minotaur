package vivid

import (
	"fmt"
	"reflect"
)

// TypedActorRef 是一个类型化的 ActorRef，可以通过 Api 方法获取其类型化的协议
type TypedActorRef[Protocol any] interface {
	ActorRef
	Api() Protocol
}

// typedWrapper 用于实现 TypedActorRef 接口的包装器
type typedWrapper[Protocol any] struct {
	ActorRef
	protocol Protocol
}

// Api 获取 ActorRef 的类型化协议
func (m *typedWrapper[Protocol]) Api() Protocol {
	return m.protocol
}

// Typed 创建一个类型化包装的 ActorRef，该 ActorRef 可以通过 Api 方法获取其类型化的协议
func Typed[Protocol any](ref ActorRef, protocol Protocol) TypedActorRef[Protocol] {
	// 检查 Api 是否为接口类型
	typ := reflect.TypeOf((*Protocol)(nil)).Elem()
	if typ.Kind() != reflect.Interface {
		panic(fmt.Errorf("protocol must be an interface type, but got %s", typ.String()))
	}

	return &typedWrapper[Protocol]{
		ActorRef: ref,
		protocol: protocol,
	}
}
