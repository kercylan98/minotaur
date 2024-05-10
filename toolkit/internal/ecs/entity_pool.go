package ecs

import "github.com/kercylan98/minotaur/toolkit/collection/listings"

type entityPool struct {
	entities  *listings.PagedSlice[EntityId] // 所有的实体
	next      int                            // 下一个实体的索引
	available int                            // 可用的实体数量
	pageSize  int                            // 每页的大小
}

func (p *entityPool) init(pageSize int) {
	p.pageSize = pageSize
	p.Reset()
	return
}

// GetN 返回多个新的实体
func (p *entityPool) GetN(n int) []EntityId {
	var entities = make([]EntityId, n)
	var length = uint32(p.entities.Len())
	for i := uint32(0); i < uint32(n); i++ {
		var entity EntityId
		entity.initFromId(length + i)
		p.entities.Add(entity)
		entities[i] = entity
	}
	return entities
}

// Get 返回一个新的或者已经回收的实体
func (p *entityPool) Get() EntityId {
	if p.available == 0 {
		// 如果没有可用的实体，则创建一个新的实体
		var entity EntityId
		entity.initFromId(uint32(p.entities.Len()))
		p.entities.Add(entity)
		return entity
	}

	// 从实体池中获取一个实体
	curr := p.next
	next := p.next
	p.next = int(p.entities.Get(p.next).getId())
	p.entities.Get(next).setId(uint32(next))
	p.available--
	return *p.entities.Get(curr)
}

// Recycle 回收一个实体
func (p *entityPool) Recycle(entity EntityId) {
	eid := entity.getId()
	if eid == 0 {
		panic("can't recycle reserved zero entity")
	}
	iEid := int(eid)
	p.entities.Get(iEid).addGeneration()
	next := p.next
	p.next = iEid
	p.entities.Get(iEid).setId(uint32(next))
	p.available++
}

// Reset 重置实体池
func (p *entityPool) Reset() {
	var zero EntityId
	if p.entities == nil {
		p.entities = listings.NewPagedSlice[EntityId](p.pageSize)
	}
	p.entities.Clear()
	p.entities.Add(zero)
	p.next = 0
	p.available = 0
}

// Alive 返回一个实体是否还活着，基于实体的生成
func (p *entityPool) Alive(entity EntityId) bool {
	return entity.getGeneration() == p.entities.Get(int(entity.getId())).getGeneration()
}

// Living 返回当前正在使用的实体数量
func (p *entityPool) Living() int {
	return p.entities.Len() - 1 - p.available
}

// Capacity 返回当前的容量（使用和回收的实体）
func (p *entityPool) Capacity() int {
	return p.entities.Len() - 1
}
