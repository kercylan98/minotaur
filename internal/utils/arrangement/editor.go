package arrangement

import (
"github.com/kercylan98/minotaur/toolkit/collection"
"sort"
)

// Editor 提供了大量辅助函数的编辑器
type Editor[ID comparable, AreaInfo any] struct {
	a          *Arrangement[ID, AreaInfo]
	pending    []Item[ID]
	fails      []Item[ID]
	falls      map[ID]struct{}
	retryCount int
}

// GetPendingCount 获取待编排的成员数量
func (slf *Editor[ID, AreaInfo]) GetPendingCount() int {
	return len(slf.pending)
}

// RemoveAreaItem 从编排区域中移除一个成员到待编排队列中，如果该成员不存在于编排区域中，则不进行任何操作
func (slf *Editor[ID, AreaInfo]) RemoveAreaItem(area *Area[ID, AreaInfo], item Item[ID]) {
	target := area.items[item.GetID()]
	if target == nil {
		return
	}
	delete(area.items, item.GetID())
	delete(slf.falls, item.GetID())
	slf.pending = append(slf.pending, target)
}

// AddAreaItem 将一个成员添加到编排区域中，如果该成员已经存在于编排区域中，则不进行任何操作
func (slf *Editor[ID, AreaInfo]) AddAreaItem(area *Area[ID, AreaInfo], item Item[ID]) {
	if collection.KeyInMap(slf.falls, item.GetID()) {
		return
	}
	area.items[item.GetID()] = item
	slf.falls[item.GetID()] = struct{}{}
}

// GetAreas 获取所有的编排区域
func (slf *Editor[ID, AreaInfo]) GetAreas() []*Area[ID, AreaInfo] {
	return collection.CloneSlice(slf.a.areas)
}

// GetAreasWithScoreAsc 获取所有的编排区域，并按照分数升序排序
func (slf *Editor[ID, AreaInfo]) GetAreasWithScoreAsc(extra ...Item[ID]) []*Area[ID, AreaInfo] {
	areas := collection.CloneSlice(slf.a.areas)
	sort.Slice(areas, func(i, j int) bool {
		return areas[i].GetScore(extra...) < areas[j].GetScore(extra...)
	})
	return areas
}

// GetAreasWithScoreDesc 获取所有的编排区域，并按照分数降序排序
func (slf *Editor[ID, AreaInfo]) GetAreasWithScoreDesc(extra ...Item[ID]) []*Area[ID, AreaInfo] {
	areas := collection.CloneSlice(slf.a.areas)
	sort.Slice(areas, func(i, j int) bool {
		return areas[i].GetScore(extra...) > areas[j].GetScore(extra...)
	})
	return areas
}

// GetRetryCount 获取重试次数
func (slf *Editor[ID, AreaInfo]) GetRetryCount() int {
	return slf.retryCount
}

// GetThresholdProgressRate 获取重试次数阈值进度
func (slf *Editor[ID, AreaInfo]) GetThresholdProgressRate() float64 {
	return float64(slf.retryCount) / float64(slf.a.threshold)
}

// GetAllowAreas 获取允许的编排区域
func (slf *Editor[ID, AreaInfo]) GetAllowAreas(item Item[ID]) []*Area[ID, AreaInfo] {
	var areas []*Area[ID, AreaInfo]
	for _, area := range slf.a.areas {
		if _, _, allow := area.IsAllow(item); allow {
			areas = append(areas, area)
		}
	}
	return areas
}

// GetNoAllowAreas 获取不允许的编排区域
func (slf *Editor[ID, AreaInfo]) GetNoAllowAreas(item Item[ID]) []*Area[ID, AreaInfo] {
	var areas []*Area[ID, AreaInfo]
	for _, area := range slf.a.areas {
		if _, _, allow := area.IsAllow(item); !allow {
			areas = append(areas, area)
		}
	}
	return areas
}

// GetBestAllowArea 获取最佳的允许的编排区域，如果不存在，则返回 nil
func (slf *Editor[ID, AreaInfo]) GetBestAllowArea(item Item[ID]) *Area[ID, AreaInfo] {
	var areas = slf.GetAllowAreas(item)
	if len(areas) == 0 {
		return nil
	}
	var bestArea = areas[0]
	var score = bestArea.GetScore(item)
	for _, area := range areas {
		if area.GetScore(item) > score {
			bestArea = area
			score = area.GetScore(item)
		}
	}
	return bestArea
}

// GetBestNoAllowArea 获取最佳的不允许的编排区域，如果不存在，则返回 nil
func (slf *Editor[ID, AreaInfo]) GetBestNoAllowArea(item Item[ID]) *Area[ID, AreaInfo] {
	var areas = slf.GetNoAllowAreas(item)
	if len(areas) == 0 {
		return nil
	}
	var bestArea = areas[0]
	var score = bestArea.GetScore(item)
	for _, area := range areas {
		if area.GetScore(item) > score {
			bestArea = area
			score = area.GetScore(item)
		}
	}
	return bestArea
}

// GetWorstAllowArea 获取最差的允许的编排区域，如果不存在，则返回 nil
func (slf *Editor[ID, AreaInfo]) GetWorstAllowArea(item Item[ID]) *Area[ID, AreaInfo] {
	var areas = slf.GetAllowAreas(item)
	if len(areas) == 0 {
		return nil
	}
	var worstArea = areas[0]
	var score = worstArea.GetScore(item)
	for _, area := range areas {
		if area.GetScore(item) < score {
			worstArea = area
			score = area.GetScore(item)
		}
	}
	return worstArea
}

// GetWorstNoAllowArea 获取最差的不允许的编排区域，如果不存在，则返回 nil
func (slf *Editor[ID, AreaInfo]) GetWorstNoAllowArea(item Item[ID]) *Area[ID, AreaInfo] {
	var areas = slf.GetNoAllowAreas(item)
	if len(areas) == 0 {
		return nil
	}
	var worstArea = areas[0]
	var score = worstArea.GetScore(item)
	for _, area := range areas {
		if area.GetScore(item) < score {
			worstArea = area
			score = area.GetScore(item)
		}
	}
	return worstArea
}
