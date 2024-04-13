package ecs

import (
	"reflect"
)

type archetypeId = string

type archetype struct {
	id archetypeId

	entityIndex map[EntityId]int
	entityData  map[ComponentId][]any
}

func (a *archetype) init(world *World, id archetypeId, componentIds []ComponentId) {
	a.id = id
	a.entityIndex = make(map[EntityId]int)
	a.entityData = make(map[ComponentId][]any)
	for _, componentId := range componentIds {
		a.entityData[componentId] = make([]any, 0)
	}
}

func (a *archetype) addEntity(world *World, entityId EntityId) {
	a.entityIndex[entityId] = len(a.entityIndex)

	for componentId, data := range a.entityData {
		componentType := world.componentTypes[componentId]
		data = append(data, reflect.New(componentType).Interface())
		a.entityData[componentId] = data
	}
}

func (a *archetype) getEntityData(entityId EntityId, id ComponentId) any {
	if idx, ok := a.entityIndex[entityId]; ok {
		return a.entityData[id][idx]
	}
	return nil
}
