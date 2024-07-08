package ecs

import "github.com/kercylan98/minotaur/toolkit/convert"

// EntityId 高位存储实体 ID，低位存储生代信息的实体 ID
type EntityId uint64

func newEntityId(index, generation uint32) EntityId {
	return EntityId(uint64(generation)<<32 | uint64(index))
}

func (id EntityId) Generation() uint32 {
	return uint32(id >> 32)
}

func (id EntityId) Id() uint32 {
	return uint32(id)
}

func (id EntityId) addGeneration() EntityId {
	return newEntityId(id.Id(), id.Generation()+1)
}

func (id EntityId) changeId(newId uint32) EntityId {
	return newEntityId(newId, id.Generation())

}

func (id EntityId) String() string {
	return "EntityId(generation=" + convert.Uint32ToString(id.Generation()) + ", id=" + convert.Uint32ToString(id.Id()) + ")"
}
