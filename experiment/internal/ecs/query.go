package ecs

import "sync"

type QueryCondition interface {
	Evaluate(mask *DynamicBitSet) bool
	String() string
}

type Result struct {
	world      *World
	archetypes map[DynamicBitSetKey]*archetype
	once       sync.Once
	count      int
	entities   []EntityId
}

func (r *Result) Iterator() *ResultIterator {
	return &ResultIterator{result: r, index: -1}
}

func (r *Result) expansion() {
	for _, art := range r.archetypes {
		r.count += len(art.entities)
		r.entities = append(r.entities, art.entityList...)
	}
}

func (r *Result) Entities() []EntityId {
	return r.entities
}

func (r *Result) Count() int {
	return r.count
}

func (r *Result) Get(entityId EntityId, componentId ComponentId) any {
	return r.world.getEntityComponentData(entityId, componentId)
}

type ResultIterator struct {
	result *Result
	index  int
}

func (r *ResultIterator) Next() bool {
	r.index++
	return r.index < r.result.count
}

func (r *ResultIterator) Get(componentId ComponentId) any {
	return r.result.world.getEntityComponentData(r.result.entities[r.index], componentId)
}

func (r *ResultIterator) Entity() EntityId {
	return r.result.entities[r.index]
}
