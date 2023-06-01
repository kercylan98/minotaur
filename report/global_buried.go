package report

import (
	"sync"
)

// NewGlobalBuried 创建一个全局埋点
func NewGlobalBuried[Data any](hitLogic HitLogic[Data]) *GlobalBuried[Data] {
	return &GlobalBuried[Data]{
		hitLogic: hitLogic,
	}
}

// GlobalBuried 全局埋点
//   - 天然并发安全
//   - 全局埋点适用于活跃用户数等统计
type GlobalBuried[Data any] struct {
	rw       sync.RWMutex
	data     Data
	hitLogic HitLogic[Data]
}

// Hit 命中数据埋点
func (slf *GlobalBuried[Data]) Hit(data Data) {
	slf.rw.Lock()
	defer slf.rw.Unlock()
	slf.data = slf.hitLogic(slf.data, data)
}

// GetData 获取数据
func (slf *GlobalBuried[Data]) GetData() Data {
	slf.rw.RLock()
	defer slf.rw.RUnlock()
	return slf.data
}
