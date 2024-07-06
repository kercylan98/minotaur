package ecs

import (
	"github.com/kercylan98/minotaur/core/vivid"
	"reflect"
)

type (
	// WorldRegisterComponentMessage 注册组件消息，该消息将回复一个 ComponentId
	WorldRegisterComponentMessage ComponentType
	// WorldCreateEntityMessage 创建实体消息，该消息将回复一个 EntityId
	WorldCreateEntityMessage []ComponentId
	// WorldLoadComponentMessage 加载组件消息
	WorldLoadComponentMessage struct {
		EntityId      EntityId
		ComponentType ComponentType
	}
)

// SendRegisterComponentMessage 发送注册组件消息并返回 ComponentId
func SendRegisterComponentMessage[T any](system *vivid.ActorSystem, world vivid.ActorRef) ComponentId {
	tof := reflect.TypeOf((*T)(nil)).Elem()
	return system.FutureAsk(world, WorldRegisterComponentMessage(tof)).AssertResult().(ComponentId)
}

// SendCreateEntityMessage 发送创建实体消息并返回 EntityId
func SendCreateEntityMessage(system *vivid.ActorSystem, world vivid.ActorRef, components ...ComponentId) EntityId {
	return system.FutureAsk(world, WorldCreateEntityMessage(components)).AssertResult().(EntityId)
}

// SendLoadComponentMessage 发送加载组件消息
func SendLoadComponentMessage[T any](system *vivid.ActorSystem, world vivid.ActorRef, eid EntityId) T {
	tof := reflect.TypeOf((*T)(nil)).Elem()
	result, _ := system.FutureAsk(world, WorldLoadComponentMessage{EntityId: eid, ComponentType: tof}).AssertResult().(T)
	return result
}
