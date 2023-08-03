package arrangement

import "github.com/kercylan98/minotaur/utils/hash"

// Area 编排区域
type Area[ID comparable, AreaInfo any] struct {
	info        AreaInfo
	items       map[ID]Item[ID]
	constraints []AreaConstraintHandle[ID, AreaInfo]
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
func (slf *Area[ID, AreaInfo]) IsAllow(item Item[ID]) (Item[ID], bool) {
	for _, constraint := range slf.constraints {
		if item, allow := constraint(slf, item); !allow {
			return item, false
		}
	}
	return nil, true
}

// GetScore 获取该编排区域的评估分数
//   - 当 extra 不为空时，将会将 extra 中的内容添加到 items 中进行评估
func (slf *Area[ID, AreaInfo]) GetScore(extra ...Item[ID]) float64 {
	if slf.evaluate == nil {
		return 0
	}
	var items = hash.Copy(slf.items)
	for _, item := range extra {
		items[item.GetID()] = item
	}
	return slf.evaluate(slf.GetAreaInfo(), items)
}
