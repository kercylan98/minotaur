package collection

import "errors"

var ErrCircularDependencyDetected = errors.New("circular dependency detected")

type topologicalSortNode[V any] struct {
	value      V
	dependsOn  []*topologicalSortNode[V]
	dependents []*topologicalSortNode[V]
}

// TopologicalSort 拓扑排序是一种对有向图进行排序的算法，它可以用来解决一些依赖关系的问题，比如计算字段的依赖关系。拓扑排序会将存在依赖关系的元素进行排序，使得依赖关系的元素总是排在被依赖的元素之前。
//   - slice: 需要排序的切片
//   - queryIndexHandler: 用于查询切片中每个元素的索引
//   - queryDependsHandler: 用于查询切片中每个元素的依赖关系，返回的是一个索引切片，如果没有依赖关系，那么返回空切片
//
// 该函数在存在循环依赖的情况下将会返回 ErrCircularDependencyDetected 错误
func TopologicalSort[S ~[]V, Index comparable, V any](slice S, queryIndexHandler func(item V) Index, queryDependsHandler func(item V) []Index) (S, error) {

	var nodes = make(map[Index]*topologicalSortNode[V])

	for _, item := range slice {
		node := &topologicalSortNode[V]{value: item}
		nodes[queryIndexHandler(item)] = node
	}

	for _, item := range slice {
		depends := queryDependsHandler(item)
		for _, depend := range depends {
			if node, exists := nodes[depend]; exists {
				node.dependsOn = append(node.dependsOn, nodes[queryIndexHandler(item)])
				node.dependents = append(node.dependents, nodes[queryIndexHandler(item)])
			}
		}
	}

	var sorted = make([]V, 0, len(slice))
	var visited = make(map[Index]bool)

	var visit func(node *topologicalSortNode[V])
	visit = func(node *topologicalSortNode[V]) {
		index := queryIndexHandler(node.value)
		if node == nil || visited[index] {
			return
		}
		visited[index] = true
		for _, n := range node.dependsOn {
			visit(n)
		}
		sorted = append(sorted, node.value)
	}

	for _, node := range nodes {
		visit(node)
	}

	if len(sorted) != len(slice) {
		return nil, ErrCircularDependencyDetected
	}

	return sorted, nil
}
