package astar

// Graph 适用于 A* 算法的图数据结构接口定义
type Graph[Node comparable] interface {
	// Neighbours 返回特定节点的邻居节点
	Neighbours(node Node) []Node
}
