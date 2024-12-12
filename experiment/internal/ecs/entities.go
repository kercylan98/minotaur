package ecs

import (
	"github.com/kercylan98/minotaur/toolkit/collection/listings"
)

func newEntities(pageSize int) *entities {
	var p = new(entities)
	p.pageSize = pageSize
	p.reset()
	return p
}

type entities struct {
	entities  *listings.PagedSlice[Entity] // 所有的实体
	next      int                          // 下一个实体的索引
	available int                          // 可用的实体数量
	pageSize  int                          // 每页的大小
}

// getMany 返回多个新的实体
func (p *entities) getMany(n int) []Entity {
	var entities = make([]Entity, n)
	var length = uint32(p.entities.Len())
	for i := uint32(0); i < uint32(n); i++ {
		var entity = newEntity(length+i, 0)
		p.entities.Add(entity)
		entities[i] = entity
	}
	return entities
}

// get 返回一个新的或者已经回收的实体
func (p *entities) get() Entity {
	if p.available == 0 {
		// 如果没有可用的实体，则创建一个新的实体
		var entity = newEntity(uint32(p.entities.Len()), 0)
		p.entities.Add(entity)
		return entity
	}

	// 从实体池中获取一个实体
	curr := p.next
	next := p.next
	p.next = int(p.entities.Get(p.next).id())
	p.entities.Set(next, p.entities.Get(next).changeId(uint32(next)))
	p.available--
	return *p.entities.Get(curr)
}

// recycle 回收一个实体
func (p *entities) recycle(entity Entity) {
	eid := entity.id()
	if eid == 0 {
		panic("can't recycle reserved zero entity")
	}
	iEid := int(eid)
	p.entities.Set(iEid, p.entities.Get(iEid).addGeneration())
	next := p.next
	p.next = iEid
	p.entities.Set(iEid, p.entities.Get(iEid).changeId(uint32(next)))
	p.available++
}

// reset 重置实体池
func (p *entities) reset() {
	var zero Entity
	if p.entities == nil {
		p.entities = listings.NewPagedSlice[Entity](p.pageSize)
	}
	p.entities.Add(zero)
	p.next = 0
	p.available = 0
}

// alive 返回一个实体是否还活着，基于实体的生成
func (p *entities) alive(entity Entity) bool {
	return entity.generation() == p.entities.Get(int(entity.id())).generation()
}

// living 返回当前正在使用的实体数量
func (p *entities) living() int {
	return p.entities.Len() - 1 - p.available
}

// capacity 返回当前的容量（使用和回收的实体）
func (p *entities) capacity() int {
	return p.entities.Len() - 1
}
