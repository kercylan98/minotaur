package astar

import (
	"container/heap"
	"github.com/kercylan98/minotaur/utils/generic"
)

// Find 使用成本函数和成本启发式函数在图中找到起点和终点之间的最低成本路径。
func Find[Node comparable, V generic.SignedNumber](graph Graph[Node], start, end Node, cost, heuristic func(a, b Node) V) []Node {
	closed := make(map[Node]bool)

	h := &h[path[Node, V], V]{}
	heap.Init(h)
	heap.Push(h, &hm[path[Node, V], V]{v: path[Node, V]{start}})

	for h.Len() > 0 {
		p := heap.Pop(h).(*hm[path[Node, V], V]).v
		n := p.last()
		if closed[n] {
			continue
		}
		if n == end {
			return p
		}
		closed[n] = true

		for _, nb := range graph.Neighbours(n) {
			cp := p.cont(nb)
			heap.Push(h, &hm[path[Node, V], V]{
				v: cp,
				p: -(cp.cost(cost) + heuristic(nb, end)),
			})
		}
	}

	return nil
}
