package ecs

type Result struct {
	world      *world
	archetypes []*archetype
	count      int
	entities   []Entity
}

func (r *Result) Iterator() ResultIterator {
	return ResultIterator{result: r, index: -1}
}

func (r *Result) expansion() {
	for _, art := range r.archetypes {
		r.count += len(art.entities)
		r.entities = append(r.entities, art.entities...)
	}
}

func (r *Result) Entities() []Entity {
	return r.entities
}

func (r *Result) Count() int {
	return r.count
}

func (r *Result) Get(entity Entity, componentId ComponentId) any {
	art := r.world.archetypes.entities[entity]
	return art.storage.Get(entity.id(), componentId)
}

func (r *Result) Each(handler func(entity Entity) bool) {
	for _, entity := range r.entities {
		if !handler(entity) {
			break
		}
	}
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
	entity := r.Entity()
	art := r.result.world.archetypes.entities[entity]
	return art.storage.Get(entity.id(), componentId)
}

func (r *ResultIterator) Entity() Entity {
	return r.result.entities[r.index]
}
