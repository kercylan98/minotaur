package ecs

import (
	"github.com/alphadose/haxmap"
	"github.com/kercylan98/minotaur/toolkit"
	"github.com/kercylan98/minotaur/toolkit/charproc"
)

func newArchetype(archetypes *archetypes, id int, mask *toolkit.DynamicBitSet, prev *archetype, storage Storage) *archetype {
	art := &archetype{
		archetypes: archetypes,
		mask:       mask,
		delEdges:   make(map[toolkit.DynamicBitSetKey]*archetype),
		addEdges:   make(map[toolkit.DynamicBitSetKey]*archetype),
		id:         id,
		cache:      haxmap.New[string, *archetype](),
		storage:    storage,
	}

	if prev != nil {
		art.delEdges[prev.mask.Key()] = prev
		prev.addEdges[art.mask.Key()] = art
	}

	for _, componentId := range mask.Bits() {
		componentInfo := archetypes.world.components.getInfo(componentId)
		art.storage.SetColumn(componentId, func() any {
			return componentInfo.instantiate()
		})
	}
	return art
}

type archetype struct {
	archetypes *archetypes                     // 原型列表
	mask       *toolkit.DynamicBitSet          // 原型掩码
	storage    Storage                         // 存储
	id         int                             // 原型 ID
	cache      *haxmap.Map[string, *archetype] // 缓存
	entities   []Entity                        // 查询用途实体列表

	delEdges map[toolkit.DynamicBitSetKey]*archetype // 删除边
	addEdges map[toolkit.DynamicBitSetKey]*archetype // 添加边
}

// mutation 变异原型，返回新的原型
func (a *archetype) mutation(add, del []ComponentId, storageFactory func() Storage) *archetype {
	curr := a
	mask := a.mask.Copy()
	cacheKey := charproc.NumberJoin(add, "") + charproc.NumberJoin(del, "")
	if cache, exists := a.cache.Get(cacheKey); exists {
		return cache
	}
	defer func() { // 建立缓存
		a.cache.Set(cacheKey, curr)
	}()

	for _, id := range del {
		mask.Clear(id)
		next, exists := curr.delEdges[mask.Key()]
		if !exists {
			next = a.archetypes.noneLockGetArchetypeWithMask(mask)
			if next == nil {
				next = a.archetypes.noneLockCreateArchetype(mask.Copy(), curr, storageFactory())
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
			next = a.archetypes.noneLockGetArchetypeWithMask(mask)
			if next == nil {
				next = a.archetypes.noneLockCreateArchetype(mask.Copy(), curr, storageFactory())
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

func (a *archetype) migrate(target *archetype, entityIds []entityId) {
	a.storage.Migrate(target.storage, entityIds...)
}

func (a *archetype) bind(entity Entity) {
	a.storage.AddRow(entity.id())
	a.archetypes.entities[entity] = a
	a.entities = append(a.entities, entity)
}

func (a *archetype) bindMany(e []Entity) {
	ids := make([]entityId, len(e))
	for i, entity := range e {
		ids[i] = entity.id()
		a.archetypes.entities[entity] = a
	}

	a.storage.AddRows(ids)
	a.entities = append(a.entities, e...)
}
