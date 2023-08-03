package arrangement

import (
	"github.com/kercylan98/minotaur/utils/hash"
	"sort"
)

// NewArrangement 创建一个新的编排
func NewArrangement[ID comparable, AreaInfo any](options ...Option[ID, AreaInfo]) *Arrangement[ID, AreaInfo] {
	arrangement := &Arrangement[ID, AreaInfo]{
		items:    map[ID]Item[ID]{},
		fixed:    map[ID]ItemFixedAreaHandle[AreaInfo]{},
		priority: map[ID][]ItemPriorityHandle[ID, AreaInfo]{},
	}
	for _, option := range options {
		option(arrangement)
	}
	return arrangement
}

// Arrangement 用于针对多条数据进行合理编排的数据结构
//   - 我不知道这个数据结构的具体用途，但是我觉得这个数据结构应该是有用的
//   - 目前我能想到的用途只有我的过往经历：排课
//   - 如果是在游戏领域，或许适用于多人小队匹配编排等类似情况
type Arrangement[ID comparable, AreaInfo any] struct {
	areas    []*Area[ID, AreaInfo]                     // 所有的编排区域
	items    map[ID]Item[ID]                           // 所有的成员
	fixed    map[ID]ItemFixedAreaHandle[AreaInfo]      // 固定编排区域的成员
	priority map[ID][]ItemPriorityHandle[ID, AreaInfo] // 成员的优先级函数
}

// AddArea 添加一个编排区域
func (slf *Arrangement[ID, AreaInfo]) AddArea(areaInfo AreaInfo, options ...AreaOption[ID, AreaInfo]) {
	area := &Area[ID, AreaInfo]{
		info:  areaInfo,
		items: map[ID]Item[ID]{},
	}
	for _, option := range options {
		option(area)
	}
	slf.areas = append(slf.areas, area)
}

// AddItem 添加一个成员
func (slf *Arrangement[ID, AreaInfo]) AddItem(item Item[ID]) {
	slf.items[item.GetID()] = item
}

// Arrange 编排
func (slf *Arrangement[ID, AreaInfo]) Arrange(threshold int) (areas []*Area[ID, AreaInfo], noSolution map[ID]Item[ID]) {
	if len(slf.areas) == 0 {
		return slf.areas, slf.items
	}
	if threshold <= 0 {
		threshold = 10
	}
	var items = hash.Copy(slf.items)
	var fixed = hash.Copy(slf.fixed)

	// 将固定编排的成员添加到对应的编排区域中，当成员无法添加到对应的编排区域中时，将会被转移至未编排区域
	for id, isFixed := range fixed {
		var notFoundArea = true
		for _, area := range slf.areas {
			if isFixed(area.GetAreaInfo()) {
				area.items[id] = items[id]
				delete(items, id)
				notFoundArea = false
				break
			}
		}
		if notFoundArea {
			delete(fixed, id)
			items[id] = slf.items[id]
		}
	}

	// 优先级处理
	var priorityInfo = map[ID]float64{}             // itemID -> priority
	var itemAreaPriority = map[ID]map[int]float64{} // itemID -> areaIndex -> priority
	for id, item := range items {
		itemAreaPriority[id] = map[int]float64{}
		for i, area := range slf.areas {
			for _, getPriority := range slf.priority[id] {
				priority := getPriority(area.GetAreaInfo(), item)
				priorityInfo[id] += priority
				itemAreaPriority[id][i] = priority
			}
		}
	}
	var pending = hash.ToSlice(items)
	sort.Slice(pending, func(i, j int) bool {
		return priorityInfo[pending[i].GetID()] > priorityInfo[pending[j].GetID()]
	})

	var current Item[ID]
	var fails []Item[ID]
	var retryCount = 0
	for len(pending) > 0 {
		current = pending[0]
		pending = pending[1:]

		var maxPriority = float64(0)
		var area *Area[ID, AreaInfo]
		for areaIndex, priority := range itemAreaPriority[current.GetID()] {
			if priority > maxPriority {
				a := slf.areas[areaIndex]
				if _, allow := a.IsAllow(current); allow {
					maxPriority = priority
					area = a
				}
			}
		}

		if area == nil { // 无法通过优先级找到合适的编排区域
			for i, a := range slf.areas {
				if _, exist := itemAreaPriority[current.GetID()][i]; exist {
					continue
				}
				if _, allow := a.IsAllow(current); allow {
					area = a
					break
				}
			}
			if area == nil {
				fails = append(fails, current)
				goto end
			}
		}

		area.items[current.GetID()] = current

	end:
		{
			if len(fails) > 0 {
				noSolution = map[ID]Item[ID]{}
				for _, item := range fails {
					noSolution[item.GetID()] = item
				}
			}

			if len(pending) == 0 && len(fails) > 0 {
				pending = fails
				fails = fails[:0]
				retryCount++
				if retryCount > threshold {
					break
				}
			}
		}
	}

	return slf.areas, noSolution
}
