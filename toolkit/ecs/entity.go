package ecs

type EntityId = int

// Entity 仅包含一个实体的唯一标识
type Entity struct {
	id      EntityId // 实体的唯一标识
	archIdx int      // 实体所在的原型索引
}

func (e Entity) GetId() EntityId {
	return e.id
}

func (e Entity) GetArchetypeIndex() int {
	return e.archIdx
}

func (e Entity) setArchetypeIndex(idx int) {
	e.archIdx = idx
}
