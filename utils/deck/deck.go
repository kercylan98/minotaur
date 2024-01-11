package deck

import (
	"github.com/kercylan98/minotaur/utils/collection"
)

// NewDeck 创建一个新的甲板
func NewDeck[I Item]() *Deck[I] {
	deck := &Deck[I]{
		groups: make(map[int64]*Group[I]),
	}
	return deck
}

// Deck 甲板，用于针对一堆 Group 进行管理的数据结构
type Deck[I Item] struct {
	groups map[int64]*Group[I]
	sort   []int64
}

// AddGroup 将一个组添加到甲板中
func (slf *Deck[I]) AddGroup(group *Group[I]) {
	if !collection.KeyInMap(slf.groups, group.GetGuid()) {
		slf.groups[group.GetGuid()] = group
		slf.sort = append(slf.sort, group.GetGuid())
	}
}

// RemoveGroup 移除甲板中的一个组
func (slf *Deck[I]) RemoveGroup(guid int64) {
	delete(slf.groups, guid)
	index := collection.FindIndexInComparableSlice(slf.sort, guid)
	if index != -1 {
		slf.sort = append(slf.sort[:index], slf.sort[index+1:]...)
	}
}

// GetCount 获取甲板中的组数量
func (slf *Deck[I]) GetCount() int {
	return len(slf.groups)
}

// GetGroups 获取所有组
func (slf *Deck[I]) GetGroups() map[int64]*Group[I] {
	return collection.CloneMap(slf.groups)
}

// GetGroupsSlice 获取所有组
func (slf *Deck[I]) GetGroupsSlice() []*Group[I] {
	var groups = make([]*Group[I], 0, len(slf.groups))
	for _, guid := range slf.sort {
		groups = append(groups, slf.groups[guid])
	}
	return groups
}

// GetNext 获取特定组的下一个组
func (slf *Deck[I]) GetNext(guid int64) *Group[I] {
	index := collection.FindIndexInComparableSlice(slf.sort, guid)
	if index == -1 {
		return nil
	}
	if index == len(slf.sort)-1 {
		return slf.groups[slf.sort[0]]
	}
	return slf.groups[slf.sort[index+1]]
}

// GetPrev 获取特定组的上一个组
func (slf *Deck[I]) GetPrev(guid int64) *Group[I] {
	index := collection.FindIndexInComparableSlice(slf.sort, guid)
	if index == -1 {
		return nil
	}
	if index == 0 {
		return slf.groups[slf.sort[len(slf.sort)-1]]
	}
	return slf.groups[slf.sort[index-1]]
}
