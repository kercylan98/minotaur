package loadbalancer

import "sync"

func NewRoundRobin[Id comparable, T RoundRobinItem[Id]]() *RoundRobin[Id, T] {
	return &RoundRobin[Id, T]{
		head: nil,
		curr: nil,
		size: 0,
	}
}

type roundRobinNode[Id comparable, T RoundRobinItem[Id]] struct {
	Value T
	Next  *roundRobinNode[Id, T]
}

type RoundRobin[Id comparable, T RoundRobinItem[Id]] struct {
	head *roundRobinNode[Id, T]
	curr *roundRobinNode[Id, T]
	size int
	rw   sync.RWMutex
}

func (r *RoundRobin[Id, T]) Add(t T) {
	r.rw.Lock()
	defer r.rw.Unlock()

	newNode := &roundRobinNode[Id, T]{Value: t}

	if r.head == nil {
		r.head = newNode
		r.curr = newNode
		newNode.Next = newNode
	} else {
		newNode.Next = r.head.Next
		r.head.Next = newNode
	}
	r.size++
}

func (r *RoundRobin[Id, T]) Remove(t T) {
	r.rw.Lock()
	defer r.rw.Unlock()

	if r.head == nil {
		return
	}

	prev := r.head
	for i := 0; i < r.size; i++ {
		if prev.Next.Value.GetId() == t.GetId() {
			if prev.Next == r.curr {
				r.curr = prev
			}
			prev.Next = prev.Next.Next
			r.size--
			if r.size == 0 {
				r.head = nil
				r.curr = nil
			}
			return
		}
		prev = prev.Next
	}
}

func (r *RoundRobin[Id, T]) Next() (t T) {
	r.rw.Lock()
	defer r.rw.Unlock()

	if r.curr == nil {
		return
	}

	r.curr = r.curr.Next
	return r.curr.Value
}

func (r *RoundRobin[Id, T]) Refresh() {
	r.rw.Lock()
	defer r.rw.Unlock()

	if r.head == nil {
		return
	}

	curr := r.head
	for i := 0; i < r.size; i++ {
		if curr.Value.GetId() == r.curr.Value.GetId() {
			r.curr = curr
			return
		}
		curr = curr.Next
	}

	r.curr = r.head
}
