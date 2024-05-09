package astar

import "github.com/kercylan98/minotaur/toolkit/constraints"

// Graph 适用于 A* 算法的图数据结构接口定义，表示导航网格，其中包含了节点和连接节点的边。
type Graph[I constraints.Ordered, T any] interface {
	// GetNodeId 返回节点的唯一标识。
	GetNodeId(node T) I
	// GetNeighbours 返回与给定节点相邻的节点列表。
	GetNeighbours(t T) []T
}
