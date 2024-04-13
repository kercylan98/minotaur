package ecs

const (
	entityIdMask          = 0xffffffff
	entityGenerationMask  = 0xffffffff00000000
	entityGenerationShift = 32
)

// EntityId 实体的唯一标识，其中还包含了实体的其他信息
//
// 实体分区信息
//   - 0 ~ 31 位：实体的索引，可容纳 2^32 个实体，即最多支持 4,294,967,296 个实体
//   - 32 ~ 63 位：实体的迭代次数，可容纳 2^32 次迭代，即一个实体最多支持 4,294,967,296 次迭代
type EntityId uint64

func (e *EntityId) init(id uint32, generation uint32) {
	*e = EntityId(id) | (EntityId(generation) << entityGenerationShift)
}

func (e *EntityId) getId() uint32 {
	return uint32(*e & entityIdMask)
}

func (e *EntityId) setId(id uint32) {
	*e = (*e & entityGenerationMask) | EntityId(id)
}

func (e *EntityId) getGeneration() uint32 {
	return uint32(*e >> entityGenerationShift)
}

func (e *EntityId) addGeneration() {
	*e += 1 << entityGenerationShift
}
