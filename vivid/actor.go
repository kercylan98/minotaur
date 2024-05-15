package vivid

import (
	"reflect"
)

var actorType = reflect.TypeOf((*Actor)(nil)).Elem()

// Actor 是 Actor 模型的接口，该接口用于定义一个 Actor
type Actor interface {
	// OnPreStart 在 Actor 启动之前执行的逻辑，适用于对 Actor 状态的初始化
	OnPreStart(ctx ActorContext) error

	// OnReceived 当 Actor 接收到消息时执行的逻辑
	OnReceived(ctx MessageContext) error

	// OnDestroy 当 Actor 被要求销毁时将会调用该函数，需要在该函数中释放 Actor 的资源
	OnDestroy(ctx ActorContext) error
}
