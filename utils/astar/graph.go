package astar

// Graph 适用于 A* 算法的图数据结构接口定义，表示导航网格，其中包含了节点和连接节点的边。
type Graph[Node comparable] interface {
	// Neighbours 返回与给定节点相邻的节点列表。
	Neighbours(node Node) []Node
}
