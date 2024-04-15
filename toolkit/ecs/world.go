package ecs

import (
	"github.com/kercylan98/minotaur/toolkit/ident"
	"github.com/kercylan98/minotaur/utils/collection"
	"reflect"
)

// NewWorld 创建一个新的世界
func NewWorld() World {
	w := World{
		componentIds:   make(map[reflect.Type]ComponentId),
		componentTypes: make(map[ComponentId]reflect.Type),
		archetypes:     make(map[string]archetype),
		archetypes64:   make(map[uint64]archetype),
		entities:       make(map[EntityId]*archetype),
	}

	w.entityPool.init(1024)
	return w
}

type World struct {
	componentIds   map[reflect.Type]ComponentId // 已经注册的组件清单
	componentTypes map[ComponentId]reflect.Type // 已经注册的组件类型清单

	archetypes   map[string]archetype // 原型清单
	archetypes64 map[uint64]archetype // 小于等于64组件数量的原型清单
	entityPool   entityPool           // 实体池
	entities     map[EntityId]*archetype
}

// GetComponentId 获取或注册一个组件
func GetComponentId[C any](world *World) ComponentId {
	return world.registerComponent(reflect.TypeOf((*C)(nil)).Elem())
}

// CreateEntity 创建一个实体
func (w *World) CreateEntity(componentIds ...ComponentId) EntityId {
	arch, _ := w.findOrCreateArchetype(componentIds...)
	eid := w.entityPool.Get()
	arch.addEntity(eid)
	w.entities[eid] = &arch
	return eid
}

// CreateEntities 创建多个实体
func (w *World) CreateEntities(count int, componentIds []ComponentId) []EntityId {
	arch, _ := w.findOrCreateArchetype(componentIds...)
	var ids = make([]EntityId, count)
	for i := 0; i < count; i++ {
		entityId := w.entityPool.Get()
		ids[i] = entityId
		arch.addEntity(entityId)
	}
	arch.addEntities(ids...)
	return nil
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

// AddComponent 为实体添加组件
func (w *World) AddComponent(entityId EntityId, componentIds ...ComponentId) {
	curr, exist := w.entities[entityId]
	if exist {
		var m = make(map[ComponentId]struct{})
		for _, id := range componentIds {
			m[id] = struct{}{}
		}
		for _, id := range curr.components() {
			m[id] = struct{}{}
		}
		componentIds = collection.ConvertMapKeysToSlice(m)
	}
	arch, _ := w.findOrCreateArchetype(componentIds...)
	w.setEntityArchetype(entityId, arch)
}

// RemoveComponent 为实体移除组件
func (w *World) RemoveComponent(entityId EntityId, componentIds ...ComponentId) {
	curr, exist := w.entities[entityId]
	if exist {
		var m = make(map[ComponentId]struct{})
		for _, id := range curr.components() {
			m[id] = struct{}{}
		}
		for _, id := range componentIds {
			delete(m, id)
		}
		componentIds = collection.ConvertMapKeysToSlice(m)
	}
	arch, _ := w.findOrCreateArchetype(componentIds...)
	w.setEntityArchetype(entityId, arch)

}

// setEntityArchetype 设置实体的原型
func (w *World) setEntityArchetype(entityId EntityId, arch archetype) {
	curr, exist := w.entities[entityId]
	if exist {
		arch.addEntity(entityId)
		arch.moveEntityData(entityId, curr)
		curr.removeEntity(entityId)
		w.entities[entityId] = &arch
	}
}

// findOrCreateArchetype 查找或创建一个原型并返回
func (w *World) findOrCreateArchetype(componentIds ...ComponentId) (arch archetype, loaded bool) {
	// 组件数量小于等于 64 时采用位运算
	if len(componentIds) <= 64 {
		mask := uint64(0)
		for _, id := range componentIds {
			mask |= 1 << id
		}
		if arch, ok := w.archetypes64[mask]; ok {
			return arch, true
		} else {
			arch.init(w, componentIds)
			w.archetypes64[mask] = arch
			return arch, false
		}
	}

	// 组件数量大于 64 时采用字符串拼接
	aid := ident.GenerateOrderedUniqueIdentStringWithUInt64(componentIds...)
	if arch, ok := w.archetypes[aid]; ok {
		return arch, true
	} else {
		arch.init(w, componentIds)
		w.archetypes[aid] = arch
		return arch, false
	}
}

// getComponentType 获取组件类型
func (w *World) getComponentType(id ComponentId) reflect.Type {
	return w.componentTypes[id]
}
