package ecs

import "github.com/kercylan98/minotaur/toolkit/collection/listings"

func newRootArchetype(world *World) *archetype {
	return newArchetype(world, newDynamicBitSet(), nil)
}

func newArchetype(world *World, mask *DynamicBitSet, prev *archetype) *archetype {
	art := &archetype{
		world:    world,
		mask:     mask,
		delEdges: make(map[DynamicBitSetKey]*archetype),
		addEdges: make(map[DynamicBitSetKey]*archetype),

		entities:      make(map[uint32]EntityId),
		entityData:    make(map[EntityId]int),
		componentData: make(map[ComponentId]*listings.PagedSlice[any]),
	}

	if prev != nil {
		art.delEdges[prev.mask.Key()] = prev
		prev.addEdges[art.mask.Key()] = art
	}

	for _, id := range mask.Bits() {
		art.componentData[id] = listings.NewPagedSlice[any](32)
	}
	world.archetypes[mask.Key()] = art
	return art
}

type archetype struct {
	world    *World                          // 世界
	mask     *DynamicBitSet                  // 原型掩码
	delEdges map[DynamicBitSetKey]*archetype // 删除边
	addEdges map[DynamicBitSetKey]*archetype // 添加边

	entities      map[uint32]EntityId                       // 所有实体
	entityList    []EntityId                                // 实体列表
	entityData    map[EntityId]int                          // 实体到数据索引的映射
	componentData map[ComponentId]*listings.PagedSlice[any] // 组件数据
}

// mutation 变异原型，返回新的原型
func (a *archetype) mutation(add, del []ComponentId) *archetype {
	curr := a
	mask := a.mask.Copy()

	for _, id := range del {
		mask.Clear(id)
		next, exists := curr.delEdges[mask.Key()]
		if !exists {
			next, exists = a.world.archetypes[mask.Key()]
			if !exists {
				next = newArchetype(a.world, mask.Copy(), curr)
			} else {
				// 建立边
				curr.delEdges[mask.Key()] = next
				next.addEdges[a.mask.Key()] = a
			}
		}
		curr = next
	}

	for _, id := range add {
		mask.Set(id)
		next, exists := curr.addEdges[mask.Key()]
		if !exists {
			next, exists = a.world.archetypes[mask.Key()]
			if !exists {
				next = newArchetype(a.world, mask.Copy(), curr)
			} else {
				// 建立边
				curr.addEdges[mask.Key()] = next
				next.delEdges[a.mask.Key()] = a
			}
		}
		curr = next
	}

	return curr
}

// addEntity 添加实体到原型
func (a *archetype) addEntity(entityId EntityId, specificationInstantiate func(componentInfo *componentInfo) any) {

	curr, exist := a.entities[entityId.Id()]
	if !exist {
		idx := len(a.entities)
		a.entityData[entityId] = idx
		a.entities[entityId.Id()] = entityId
		a.entityList = append(a.entityList, entityId)

		for id, data := range a.componentData {
			cmp := a.world.getComponentInfoById(id)
			var ins any
			if specificationInstantiate != nil {
				ins = specificationInstantiate(cmp)
			} else {
				ins = cmp.instantiate()
			}
			data.Add(ins)
		}
		return
	}

	if curr.Generation() != entityId.Generation() {
		idx := a.entityData[curr]
		a.entityData[entityId] = idx
		delete(a.entityData, curr)
		a.entities[entityId.Id()] = entityId
		a.entityList[idx] = entityId

		for id, data := range a.componentData {
			cmp := a.world.getComponentInfoById(id)
			var ins any
			if specificationInstantiate != nil {
				ins = specificationInstantiate(cmp)
			} else {
				ins = cmp.instantiate()
			}
			data.Set(idx, ins)
		}
	}
}

// getEntityData 获取实体数据
func (a *archetype) getEntityComponentData(entityId EntityId, componentId ComponentId) any {
	idx, exists := a.entityData[entityId]
	if !exists {
		return nil
	}

	data := a.componentData[componentId].Get(idx)
	if data == nil {
		return nil
	}
	return *data
}

func (a *archetype) migrate(target *archetype, entityIds ...EntityId) {
	for _, entityId := range entityIds {
		idx := a.entityData[entityId]
		target.addEntity(entityId, func(componentInfo *componentInfo) any {
			ins := a.getEntityComponentData(entityId, componentInfo.id)
			if ins == nil {
				ins = componentInfo.instantiate()
			}

			return ins
		})

		// 迁移需要删除，避免数据重复
		delete(a.entities, entityId.Id())
		delete(a.entityData, entityId)
		a.entityList = a.entityList[:idx+copy(a.entityList[idx:], a.entityList[idx+1:])]

		a.world.entityArchetype[entityId] = target

		for _, data := range a.componentData {
			data.Del(idx)
		}
	}
}
