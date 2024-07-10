package ecs

import "github.com/kercylan98/minotaur/toolkit/convert"

type (
	Entity           uint64
	entityId         = uint32
	entityGeneration = uint32
)

func newEntity(id entityId, generation entityGeneration) Entity {
	return Entity(uint64(generation)<<32 | uint64(id))
}

func (id Entity) generation() entityGeneration {
	return entityGeneration(id >> 32)
}

func (id Entity) id() entityId {
	return uint32(id)
}

func (id Entity) String() string {
	return "EntityId(generation=" + convert.Uint32ToString(id.generation()) + ", id=" + convert.Uint32ToString(id.id()) + ")"
}

func (id Entity) addGeneration() Entity {
	return newEntity(id.id(), id.generation()+1)
}

func (id Entity) changeId(newId entityId) Entity {
	return newEntity(newId, id.generation())

}
