package astar

import (
	"container/heap"
	"github.com/kercylan98/minotaur/utils/generic"
)

// Find 使用 A* 算法在导航网格上查找从起点到终点的最短路径，并返回路径上的节点序列。
//
// 参数：
//   - graph: 图对象，类型为 Graph[Node]，表示导航网格。
//   - start: 起点节点，类型为 Node，表示路径的起点。
//   - end: 终点节点，类型为 Node，表示路径的终点。
//   - cost: 路径代价函数，类型为 func(a, b Node) V，用于计算两个节点之间的代价。
//   - heuristic: 启发函数，类型为 func(a, b Node) V，用于估计从当前节点到目标节点的启发式代价。
//
// 返回值：
//   - []Node: 节点序列，表示从起点到终点的最短路径。如果找不到路径，则返回空序列。
//
// 注意事项：
//   - graph 对象表示导航网格，其中包含了节点和连接节点的边。
//   - start 和 end 分别表示路径的起点和终点。
//   - cost 函数用于计算两个节点之间的代价，可以根据实际情况自定义实现。
//   - heuristic 函数用于估计从当前节点到目标节点的启发式代价，可以根据实际情况自定义实现。
//   - 函数使用了 A* 算法来搜索最短路径。
//   - 函数内部使用了堆数据结构来管理待处理的节点。
//   - 函数返回一个节点序列，表示从起点到终点的最短路径。如果找不到路径，则返回空序列。
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
