package ecs

type Status struct {
	ComponentNum       int            // 组件数量
	ArchetypeNum       int            // 原型数量
	ArchetypeEntityNum map[string]int // 原型中实体的数量
	LivingEntityNum    int            // 正在使用的实体数量
}

func NewStatus(world *World) Status {
	s := Status{
		ComponentNum:       len(world.componentIds),
		ArchetypeNum:       len(world.archetypes) + len(world.archetypes64),
		ArchetypeEntityNum: make(map[string]int),
		LivingEntityNum:    world.entityPool.Living(),
	}

	for _, arch := range world.archetypes {
		s.ArchetypeEntityNum[arch.name()] = len(arch.entityIndex)
	}

	for _, arch := range world.archetypes64 {
		s.ArchetypeEntityNum[arch.name()] = len(arch.entityIndex)
	}

	return s
}
