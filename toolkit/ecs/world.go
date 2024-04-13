package ecs

import (
	"github.com/kercylan98/minotaur/toolkit/ident"
	"reflect"
)

// NewWorld 创建一个新的世界
func NewWorld() World {
	w := World{
		componentIds:   make(map[reflect.Type]ComponentId),
		componentTypes: make(map[ComponentId]reflect.Type),
		archetypes:     make(map[string]archetype),
	}

	w.entityPool.init(1024)
	return w
}

type World struct {
	componentIds   map[reflect.Type]ComponentId // 已经注册的组件清单
	componentTypes map[ComponentId]reflect.Type // 已经注册的组件类型清单

	archetypes map[string]archetype // 原型清单
	entityPool entityPool           // 实体池
}

// Component 获取或注册一个组件
func Component[C any](world *World) ComponentId {
	return world.registerComponent(reflect.TypeOf((*C)(nil)).Elem())
}

// CreateEntity 创建一个实体
func (w *World) CreateEntity(componentIds ...ComponentId) EntityId {
	arch := w.findOrCreateArchetype(componentIds...)
	eid := w.entityPool.Get()
	arch.addEntity(w, eid)
	return eid
}

// registerComponent 注册一个组件
func (w *World) registerComponent(componentType reflect.Type) ComponentId {
	if id, ok := w.componentIds[componentType]; ok {
		return id
	}

	id := ComponentId(len(w.componentIds))
	w.componentIds[componentType] = id
	w.componentTypes[id] = componentType
	return id
}

func (w *World) findOrCreateArchetype(componentIds ...ComponentId) archetype {
	aid := ident.GenerateOrderedUniqueIdentStringWithUInt64(componentIds...)
	if arch, ok := w.archetypes[aid]; ok {
		return arch
	} else {
		arch.init(w, aid, componentIds)
		w.archetypes[aid] = arch
		return arch
	}
}
