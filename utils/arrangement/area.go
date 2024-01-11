package arrangement

import (
	"github.com/kercylan98/minotaur/utils/collection"
)

// Area 编排区域
type Area[ID comparable, AreaInfo any] struct {
	info        AreaInfo
	items       map[ID]Item[ID]
	constraints []AreaConstraintHandle[ID, AreaInfo]
	conflicts   []AreaConflictHandle[ID, AreaInfo]
	evaluate    AreaEvaluateHandle[ID, AreaInfo]
}

// GetAreaInfo 获取编排区域的信息
func (slf *Area[ID, AreaInfo]) GetAreaInfo() AreaInfo {
	return slf.info
}

// GetItems 获取编排区域中的所有成员
func (slf *Area[ID, AreaInfo]) GetItems() map[ID]Item[ID] {
	return slf.items
}

// IsAllow 检测一个成员是否可以被添加到该编排区域中
func (slf *Area[ID, AreaInfo]) IsAllow(item Item[ID]) (constraintErr error, conflictItems map[ID]Item[ID], allow bool) {
	for _, constraint := range slf.constraints {
		if err := constraint(slf, item); err != nil {
			return err, nil, false
		}
	}
	for _, conflict := range slf.conflicts {
		if items := conflict(slf, item); len(items) > 0 {
			if conflictItems == nil {
				conflictItems = make(map[ID]Item[ID])
			}
			for id, item := range items {
				conflictItems[id] = item
			}
		}
	}
	return nil, conflictItems, len(conflictItems) == 0
}

// IsConflict 检测一个成员是否会造成冲突
func (slf *Area[ID, AreaInfo]) IsConflict(item Item[ID]) bool {
	if collection.FindInMapKey(slf.items, item.GetID()) {
		return false
	}
	for _, conflict := range slf.conflicts {
		if items := conflict(slf, item); len(items) > 0 {
			return true
		}
	}
	return false
}

// GetConflictItems 获取与一个成员产生冲突的所有其他成员
func (slf *Area[ID, AreaInfo]) GetConflictItems(item Item[ID]) map[ID]Item[ID] {
	if collection.FindInMapKey(slf.items, item.GetID()) {
		return nil
	}
	var conflictItems map[ID]Item[ID]
	for _, conflict := range slf.conflicts {
		if items := conflict(slf, item); len(items) > 0 {
			if conflictItems == nil {
				conflictItems = make(map[ID]Item[ID])
			}
			for id, item := range items {
				conflictItems[id] = item
			}
		}
	}
	return conflictItems
}

// GetScore 获取该编排区域的评估分数
//   - 当 extra 不为空时，将会将 extra 中的内容添加到 items 中进行评估
func (slf *Area[ID, AreaInfo]) GetScore(extra ...Item[ID]) float64 {
	if slf.evaluate == nil {
		return 0
	}
	var items = collection.CloneMap(slf.items)
	for _, item := range extra {
		items[item.GetID()] = item
	}
	return slf.evaluate(slf.GetAreaInfo(), items)
}
