package balancer

import (
	"github.com/kercylan98/minotaur/toolkit/constraints"
	"sync"
)

func NewRoundRobin[I constraints.Ordered, T Item[I]]() *RoundRobin[I, T] {
	return &RoundRobin[I, T]{
		instances: []T{},
		curr:      0,
		idx:       make(map[I]int),
	}
}

// RoundRobin 轮询负载均衡器
type RoundRobin[I constraints.Ordered, T Item[I]] struct {
	instances []T
	curr      int
	idx       map[I]int
	rw        sync.RWMutex
}

func (r *RoundRobin[I, T]) Select(opts ...*SelectOptions) (t T, err error) {
	if len(r.instances) == 0 {
		return t, ErrNoInstance
	}

	r.curr = (r.curr + 1) % len(r.instances)
	t = r.instances[r.curr]
	return
}

func (r *RoundRobin[I, T]) Add(instance T) {
	r.rw.Lock()
	defer r.rw.Unlock()

	if _, ok := r.idx[instance.GetId()]; ok {
		return
	}
	r.idx[instance.GetId()] = len(r.instances)
	r.instances = append(r.instances, instance)

}

func (r *RoundRobin[I, T]) Remove(instance T) {
	r.rw.Lock()
	defer r.rw.Unlock()

	if idx, ok := r.idx[instance.GetId()]; ok {
		delete(r.idx, instance.GetId())
		r.instances = append(r.instances[:idx], r.instances[idx+1:]...)
		for i := idx; i < len(r.instances); i++ {
			r.idx[r.instances[i].GetId()] = i // 更新索引
		}

		if r.curr >= len(r.instances) {
			r.curr = 0
		}
	}
}

func (r *RoundRobin[I, T]) GetInstances() []T {
	r.rw.RLock()
	defer r.rw.RUnlock()

	return r.instances
}
