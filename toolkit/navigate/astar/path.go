package astar

// path 表示一条路径，是一系列节点的有序序列。
type path[T any] []T

// Last 获取路径中的最后一个节点
func (p path[T]) Last() T {
	if len(p) == 0 {
		panic("empty path")
	}
	return p[len(p)-1]
}

// Extend 通过追加一个节点创建一个新的路径
// 这个方法返回一个新的路径，而不会改变原始路径
func (p path[T]) Extend(n T) path[T] {
	newPath := append(make(path[T], 0, len(p)+1), p...)
	return append(newPath, n)
}

// Cost 计算路径的总成本
// 参数 `costFunc` 是一个函数，计算两个节点之间的成本
func (p path[T]) Cost(f func(a, b T) float64) float64 {
	totalCost := 0.0
	for i := 1; i < len(p); i++ {
		totalCost += f(p[i-1], p[i])
	}
	return totalCost
}
