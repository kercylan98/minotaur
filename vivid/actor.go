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
	OnReceive(message *Message) error

	// OnDestroy 销毁 Actor，该函数将在 ActorSystem 主动销毁 Actor 时调用
	//  - 销毁过程应该保持同步阻塞， ActorSystem 会等待 Actor 销毁
	OnDestroy()
}

// BasicActor 基础 Actor 结构体，实现了 Actor 接口
type BasicActor struct {
	system     *ActorSystem
	terminated ActorTerminatedNotifier
	router     map[string]reflect.Value
}

// RegisterTell 注册消息处理函数
func (b *BasicActor) RegisterTell(name string, handler any) error {
	if b.router == nil {
		b.router = make(map[string]reflect.Value)
	}

	tof := reflect.TypeOf(handler)
	if tof.Kind() != reflect.Func {
		return fmt.Errorf("%w: %s", ErrActorMessageHandlerNotFunc, tof.String())
	}

	b.router[name] = reflect.ValueOf(handler)
	return nil
}

func (b *BasicActor) OnSpawn(system *ActorSystem, terminated ActorTerminatedNotifier) error {
	b.system = system
	b.terminated = terminated
	return nil
}

func (b *BasicActor) OnReceive(message *Message) error {
	if b.router == nil {
		return ErrActorNotHasAnyHandler
	}
	handler, exist := b.router[message.Command]
	if !exist {
		return fmt.Errorf("%w: %s", ErrActorMessageHandlerNotFound, message.Command)
	}
	if handler.Type().NumIn() != len(message.Params) {
		return fmt.Errorf("%w: %s, except: %d, got: %d", ErrActorHandlerParamsNotMatch, message.Command, handler.Type(), len(message.Params))
	}

	params := make([]reflect.Value, len(message.Params))
	for i := 0; i < len(message.Params); i++ {
		params[i] = reflect.ValueOf(message.Params[i])
	}

	result := handler.Call(params)
	message.Results = make([]any, len(result))
	for i := 0; i < len(result); i++ {
		message.Results[i] = result[i].Interface()
	}
	return nil
}

func (b *BasicActor) OnDestroy() {
	b.terminated()
}
