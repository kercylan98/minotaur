package ecs

import (
	"strconv"
	"strings"
	"unsafe"
)

type archetype struct {
	archetypeSize  uintptr // 表示了该原型内所有组件的大小总和
	archetypeAlign uintptr // 表示了该原型内所有组件的对齐方式

	entityIndex map[EntityId]int
	entityData  map[ComponentId]*component
}

func (a *archetype) init(world *World, componentIds []ComponentId) {
	a.entityIndex = make(map[EntityId]int)
	a.entityData = make(map[ComponentId]*component)

	for _, componentId := range componentIds {
		a.entityData[componentId] = newComponent(world.getComponentType(componentId))
	}
}

// name 返回原型的名称
func (a *archetype) name() string {
	var builder strings.Builder
	builder.WriteString("Archetype[")
	var last = len(a.entityData) - 1
	var index = 0
	for id, cmp := range a.entityData {
		builder.WriteString(cmp.typ.String())
		builder.WriteString("(")
		builder.WriteString(strconv.FormatUint(id, 10))
		builder.WriteString(")")
		if index != last {
			builder.WriteString(", ")
		}
		index++
	}
	builder.WriteString("]")
	return builder.String()
}

func (a *archetype) addEntity(entityId EntityId) {
	index := len(a.entityIndex)
	a.entityIndex[entityId] = index
	for _, cmp := range a.entityData {
		cmp.Append(1)
	}
}

func (a *archetype) moveEntityData(entityId EntityId, from *archetype) {
	fromIdx := from.entityIndex[entityId]
	toIdx := a.entityIndex[entityId]

	for id, cmp := range a.entityData {
		fromCmp, exist := from.entityData[id]
		if !exist {
			continue
		}
		cmp.Set(toIdx, fromCmp.Get(fromIdx))
		fromCmp.Delete(fromIdx)
	}
}

func (a *archetype) removeEntity(entityId EntityId) {
	if index, ok := a.entityIndex[entityId]; ok {
		for _, cmp := range a.entityData {
			cmp.Delete(index)
		}
		delete(a.entityIndex, entityId)
	}
}

func (a *archetype) addEntities(entityIds ...EntityId) {
	if len(entityIds) == 0 {
		return
	}

	index := len(a.entityIndex)
	for _, entityId := range entityIds {
		a.entityIndex[entityId] = index
		index++
	}

	for _, cmp := range a.entityData {
		cmp.Append(len(entityIds))
	}
}

func (a *archetype) getEntityData(entityId EntityId, id ComponentId) unsafe.Pointer {
	if index, ok := a.entityIndex[entityId]; ok {
		if cmp, ok := a.entityData[id]; ok {
			return cmp.Get(index)
		}
	}
	return nil
}

func (a *archetype) hasComponent(id ComponentId) bool {
	_, ok := a.entityData[id]
	return ok
}

func (a *archetype) components() []ComponentId {
	var ids = make([]ComponentId, 0, len(a.entityData))
	for id := range a.entityData {
		ids = append(ids, id)
	}
	return ids
}
