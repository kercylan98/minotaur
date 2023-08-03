package arrangement

import (
	"github.com/kercylan98/minotaur/utils/hash"
	"github.com/kercylan98/minotaur/utils/slice"
)

// Editor 编排器
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
	if hash.Exist(slf.falls, item.GetID()) {
		return
	}
	area.items[item.GetID()] = item
	slf.falls[item.GetID()] = struct{}{}
}

// GetAreas 获取所有的编排区域
func (slf *Editor[ID, AreaInfo]) GetAreas() []*Area[ID, AreaInfo] {
	return slice.Copy(slf.a.areas)
}

// GetRetryCount 获取重试次数
func (slf *Editor[ID, AreaInfo]) GetRetryCount() int {
	return slf.retryCount
}

// GetThresholdProgressRate 获取重试次数阈值进度
func (slf *Editor[ID, AreaInfo]) GetThresholdProgressRate() float64 {
	return float64(slf.retryCount) / float64(slf.a.threshold)
}
