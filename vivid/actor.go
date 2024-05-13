package vivid

import (
	"fmt"
	"reflect"
)

// Actor 外部 Actor 接口，实现该接口的结构体可以被 ActorSystem 管理
type Actor interface {
	// OnSpawn 该函数将在 Actor 被创建时调用，传入 ActorSystem、 ActorTerminatedNotifier，用于 Actor 的初始化
	OnSpawn(system *ActorSystem, terminated ActorTerminatedNotifier) error

	// OnReceive 接收消息，该函数将在 Actor 接收到消息时调用
	OnReceive(message Context) error

	// OnDestroy 销毁 Actor，该函数将在 ActorSystem 主动销毁 Actor 时调用
	//  - 销毁过程应该保持同步阻塞， ActorSystem 会等待 Actor 销毁
	OnDestroy()
}

type ActorHandler func(message Context) error

// BasicActor 基础 Actor 结构体，实现了 Actor 接口
type BasicActor struct {
	system     *ActorSystem
	terminated ActorTerminatedNotifier
	router     map[string]ActorHandler
}

// RegisterTell 注册消息处理函数
func (b *BasicActor) RegisterTell(name string, handler ActorHandler) error {
	if b.router == nil {
		b.router = make(map[string]ActorHandler)
	}

	tof := reflect.TypeOf(handler)
	if tof.Kind() != reflect.Func {
		return fmt.Errorf("%w: %s", ErrActorMessageHandlerNotFunc, tof.String())
	}

	b.router[name] = handler
	return nil
}

func (b *BasicActor) OnSpawn(system *ActorSystem, terminated ActorTerminatedNotifier) error {
	b.system = system
	b.terminated = terminated
	return nil
}

func (b *BasicActor) OnReceive(message Context) error {
	if b.router == nil {
		return ErrActorNotHasAnyHandler
	}
	command := message.GetCommand()
	handler, exist := b.router[command]
	if !exist {
		return fmt.Errorf("%w: %s", ErrActorMessageHandlerNotFound, command)
	}

	return handler(message)
}

func (b *BasicActor) OnDestroy() {
	b.terminated()
}
