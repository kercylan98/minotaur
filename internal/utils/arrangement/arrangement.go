package arrangement

import (
"github.com/kercylan98/minotaur/toolkit/collection"
"sort"
)

// NewArrangement 创建一个新的编排
func NewArrangement[ID comparable, AreaInfo any](options ...Option[ID, AreaInfo]) *Arrangement[ID, AreaInfo] {
	arrangement := &Arrangement[ID, AreaInfo]{
		items:        map[ID]Item[ID]{},
		fixed:        map[ID]ItemFixedAreaHandle[AreaInfo]{},
		priority:     map[ID][]ItemPriorityHandle[ID, AreaInfo]{},
		itemNotAllow: map[ID][]ItemNotAllowVerifyHandle[ID, AreaInfo]{},
		threshold:    10,
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
	areas        []*Area[ID, AreaInfo]                           // 所有的编排区域
	items        map[ID]Item[ID]                                 // 所有的成员
	fixed        map[ID]ItemFixedAreaHandle[AreaInfo]            // 固定编排区域的成员
	priority     map[ID][]ItemPriorityHandle[ID, AreaInfo]       // 成员的优先级函数
	itemNotAllow map[ID][]ItemNotAllowVerifyHandle[ID, AreaInfo] // 成员的不允的编排区域检测函数
	threshold    int                                             // 重试次数阈值

	constraintHandles []ConstraintHandle[ID, AreaInfo]
	conflictHandles   []ConflictHandle[ID, AreaInfo]
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
func (slf *Arrangement[ID, AreaInfo]) Arrange() (areas []*Area[ID, AreaInfo], noSolution map[ID]Item[ID]) {
	if len(slf.areas) == 0 {
		return slf.areas, slf.items
	}

	var items = collection.CloneMap(slf.items)
	var fixed = collection.CloneMap(slf.fixed)

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
	var editor = &Editor[ID, AreaInfo]{
		a:       slf,
		pending: collection.ConvertMapValuesToSlice(items),
		falls:   map[ID]struct{}{},
	}
	sort.Slice(editor.pending, func(i, j int) bool {
		return priorityInfo[editor.pending[i].GetID()] > priorityInfo[editor.pending[j].GetID()]
	})

	var current Item[ID]
	for editor.GetPendingCount() > 0 {
		current = editor.pending[0]
		editor.pending = editor.pending[1:]

		var maxPriority = float64(0)
		var area *Area[ID, AreaInfo]
		for areaIndex, priority := range itemAreaPriority[current.GetID()] {
			if priority > maxPriority {
				a := slf.areas[areaIndex]
				if slf.try(editor, a, current) {
					maxPriority = priority
					area = a
				}
			}
		}

		if area == nil { // 无法通过优先级找到合适的编排区域
			for i, a := range editor.GetAreasWithScoreDesc(current) {
				if _, exist := itemAreaPriority[current.GetID()][i]; exist {
					continue
				}
				if slf.try(editor, a, current) {
					area = a
				}
			}
			if area == nil {
				editor.fails = append(editor.fails, current)
				goto end
			}
		}

		area.items[current.GetID()] = current
		editor.falls[current.GetID()] = struct{}{}

	end:
		{
			if len(editor.fails) > 0 {
				noSolution = map[ID]Item[ID]{}
				for _, item := range editor.fails {
					noSolution[item.GetID()] = item
				}
			}

			if len(editor.pending) == 0 && len(editor.fails) > 0 {
				editor.pending = editor.fails
				editor.fails = editor.fails[:0]
				editor.retryCount++
				if editor.retryCount > slf.threshold {
					break
				}
			}
		}
	}

	return slf.areas, noSolution
}

// try 尝试将 current 编排到 a 中
func (slf *Arrangement[ID, AreaInfo]) try(editor *Editor[ID, AreaInfo], a *Area[ID, AreaInfo], current Item[ID]) bool {
	allow := true
	for _, verify := range slf.itemNotAllow[current.GetID()] {
		if verify(a.GetAreaInfo(), current) {
			allow = false
			break
		}
	}
	if !allow {
		return false
	}

	err, conflictItems, allow := a.IsAllow(current)
	if !allow {
		if err != nil {
			var solve = err
			for _, handle := range slf.constraintHandles {
				if solve = handle(editor, a, current, solve); solve == nil {
					err, conflictItems, allow = a.IsAllow(current)
					if allow {
						break
					} else {
						// 当 err 依旧不为 nil 时，发生约束处理函数欺骗行为，不做任何处理
						if len(conflictItems) > 0 {
							goto conflictHandle
						}
					}
					break
				}
			}
		}
	conflictHandle:
		{
			if err == nil && len(conflictItems) > 0 { // 硬性约束未解决时，不考虑冲突解决
				var solve = conflictItems
				for _, handle := range slf.conflictHandles {
					if solve = handle(editor, a, current, solve); len(solve) == 0 {
						if a.IsConflict(current) {
							allow = true
							break
						}
						// 依旧存在冲突时，表明冲突处理函数存在欺骗行为，不做任何处理
					}
				}
			}
		}
	}
	return allow
}
