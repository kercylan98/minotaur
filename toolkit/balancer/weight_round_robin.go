package balancer

import (
	"github.com/kercylan98/minotaur/toolkit/constraints"
	"sync"
)

// NewWeightRoundRobin 创建一个加权轮询负载均衡器
func NewWeightRoundRobin[I constraints.Ordered, T Item[I]]() *WeightRoundRobin[I, T] {
	return &WeightRoundRobin[I, T]{
		instances: []T{},
		curr:      0,
		idx:       make(map[I]int),
	}
}

// WeightRoundRobin 加权轮询负载均衡器
type WeightRoundRobin[I constraints.Ordered, T Item[I]] struct {
	instances []T
	curr      int
	idx       map[I]int
	rw        sync.RWMutex
}

func (r *WeightRoundRobin[I, T]) Select(opts ...*SelectOptions) (t T, err error) {
	r.rw.Lock()
	defer r.rw.Unlock()

	if len(r.instances) == 0 {
		return t, ErrNoInstance
	}

	// 计算总权重
	totalWeight := 0
	for _, instance := range r.instances {
		totalWeight += instance.GetWeight()
	}

	// 选择实例
	for i := 0; i < len(r.instances); i++ {
		r.curr = (r.curr + 1) % len(r.instances)
		t = r.instances[r.curr]
		if t.GetWeight() >= totalWeight {
			break
		}
	}
	return
}

func (r *WeightRoundRobin[I, T]) Add(instance T) {
	r.rw.Lock()
	defer r.rw.Unlock()

	if _, ok := r.idx[instance.GetId()]; ok {
		return
	}
	r.idx[instance.GetId()] = len(r.instances)
	r.instances = append(r.instances, instance)
}

func (r *WeightRoundRobin[I, T]) Remove(instance T) {
	r.rw.Lock()
	defer r.rw.Unlock()

	instanceId := instance.GetId()
	if idx, ok := r.idx[instanceId]; ok {
		delete(r.idx, instanceId)
		r.instances = append(r.instances[:idx], r.instances[idx+1:]...)
		for i := idx; i < len(r.instances); i++ {
			r.idx[r.instances[i].GetId()] = i // 更新索引
		}

		if r.curr >= len(r.instances) {
			r.curr = 0
		}
	}
}

func (r *WeightRoundRobin[I, T]) GetInstances() []T {
	r.rw.RLock()
	defer r.rw.RUnlock()

	return r.instances
}
