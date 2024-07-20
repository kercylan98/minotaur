package ecs

import (
	"github.com/kercylan98/minotaur/engine/ecs/storage/column"
	"github.com/kercylan98/minotaur/toolkit"
)

func newArchetypes(world *world) *archetypes {
	arts := &archetypes{
		world:    world,
		arts:     make([]*archetype, 0),
		masks:    make(map[toolkit.DynamicBitSetKey]int),
		entities: make(map[Entity]*archetype),
	}
	arts.root = arts.lockCreateRootArchetype(toolkit.NewDynamicBitSet(), column.New[entityId, ComponentId]())
	return arts
}

type archetypes struct {
	world    *world                           // 世界
	root     *archetype                       // 根原型
	arts     []*archetype                     // 原型列表
	masks    map[toolkit.DynamicBitSetKey]int // 掩码索引
	guid     int                              // 原型 ID（列表索引）
	entities map[Entity]*archetype            // 实体索引
}

func (a *archetypes) lockCreateRootArchetype(mask *toolkit.DynamicBitSet, storage Storage) *archetype {
	return a.lockCreateArchetype(mask, nil, storage)
}

func (a *archetypes) lockCreateArchetype(mask *toolkit.DynamicBitSet, prev *archetype, storage Storage) *archetype {
	guid := a.guid
	a.guid++
	a.masks[mask.Key()] = guid

	art := newArchetype(a, guid, mask, prev, storage)
	a.arts = append(a.arts, art)
	return art
}

func (a *archetypes) noneLockCreateArchetype(mask *toolkit.DynamicBitSet, prev *archetype, storage Storage) *archetype {
	guid := a.guid
	a.guid++
	a.masks[mask.Key()] = guid

	art := newArchetype(a, guid, mask, prev, storage)
	a.arts = append(a.arts, art)
	return art
}

func (a *archetypes) noneLockGetArchetypeWithMask(mask *toolkit.DynamicBitSet) *archetype {
	key := mask.Key()
	idx, exist := a.masks[key]
	if !exist {
		return nil
	}
	return a.arts[idx]
}

func (a *archetypes) get(ids ...ComponentId) *archetype {
	return a.root.mutation(ids, nil, func() Storage {
		return column.New[entityId, ComponentId]()
	})
}

func (a *archetypes) unbind(entity Entity) {
	if _, exists := a.entities[entity]; exists {
		delete(a.entities, entity)
	}
}

func (a *archetypes) unBindMany(entities []Entity) {
	for _, entity := range entities {
		delete(a.entities, entity)
	}
}
