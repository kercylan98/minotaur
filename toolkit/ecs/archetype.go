package ecs

import (
	"github.com/kercylan98/minotaur/utils/super"
	"reflect"
)

func newArchetype(world *World, mask *super.BitSet[ComponentId]) *archetype {
	arch := &archetype{
		world:      world,
		mask:       mask,
		entityData: make(map[ComponentId][]reflect.Value),
	}

	return arch
}

// archetype 原型是一种实体的集合，它们都包含了相同的组件
type archetype struct {
	world      *World
	mask       *super.BitSet[ComponentId]
	entities   []Entity
	entityData map[ComponentId][]reflect.Value
}

func (a *archetype) addEntity(entity Entity) Entity {
	entity.setArchetypeIndex(len(a.entities))
	a.entities = append(a.entities, entity)
	for _, componentId := range a.mask.Bits() {
		t := a.world.getComponentTypeById(componentId)
		if t == nil {
			continue
		}

		v := reflect.New(t)
		a.entityData[componentId] = append(a.entityData[componentId], v)
	}
	return entity
}

func (a *archetype) removeEntity(entity Entity) {
	idx := entity.GetArchetypeIndex()
	for componentId, values := range a.entityData {
		a.entityData[componentId] = append(values[:idx], values[idx+1:]...)
	}
	a.entities = append(a.entities[:idx], a.entities[idx+1:]...)
}

func (a *archetype) getEntityComponentData(entity Entity, componentId ComponentId) reflect.Value {
	return a.entityData[componentId][entity.GetArchetypeIndex()]
}

func (a *archetype) getEntityData(entity Entity) []reflect.Value {
	var idx = entity.GetArchetypeIndex()
	var data []reflect.Value
	for _, componentId := range a.mask.Bits() {
		data = append(data, a.entityData[componentId][idx])
	}
	return data
}
