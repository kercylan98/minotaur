package astar

import "github.com/kercylan98/minotaur/utils/generic"

type hm[T any, V generic.SignedNumber] struct {
	v T
	p V
}

type h[T any, V generic.SignedNumber] []*hm[T, V]

func (slf *h[T, V]) Len() int {
	return len(*slf)
}

func (slf *h[T, V]) Less(i, j int) bool {
	h := *slf
	return h[i].p > h[j].p
}

func (slf *h[T, V]) Swap(i, j int) {
	h := *slf
	h[i], h[j] = h[j], h[i]
}

func (slf *h[T, V]) Push(x any) {
	*slf = append(*slf, x.(*hm[T, V]))
}

func (slf *h[T, V]) Pop() any {
	h := *slf
	size := len(h)
	t := h[size-1]
	*slf = h[0 : size-1]
	return t
}
