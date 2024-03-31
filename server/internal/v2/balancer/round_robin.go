package balancer

import "sync"

func NewRoundRobin[Id comparable, T Item[Id]]() *RoundRobin[Id, T] {

}

type RoundRobin[Id comparable, T Item[Id]] struct {
	ref   map[Id]int
	items []T
	rw    sync.RWMutex
	curr  int
}

func (r *RoundRobin[Id, T]) Add(t T) {
	r.rw.Lock()
	defer r.rw.Unlock()
	_, exist := r.ref[t.Id()]
	if exist {
		return
	}
	r.ref[t.Id()] = len(r.items)
	r.items = append(r.items, t)
}

func (r *RoundRobin[Id, T]) Remove(t T) {
	r.rw.Lock()
	defer r.rw.Unlock()
	index, exist := r.ref[t.Id()]
	if !exist {
		return
	}
	r.items = append(r.items[:index], r.items[index+1:]...)
	delete(r.ref, t.Id())
}

func (r *RoundRobin[Id, T]) Next() T {
	r.rw.RLock()
	defer r.rw.RUnlock()
	if r.curr >= len(r.items) {
		r.curr = 0
	}
	t := r.items[r.curr]
	r.curr++
}
