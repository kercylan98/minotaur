package astar

// heapItem 表示一个带有优先级的项目
type heapItem[T any] struct {
	value    T
	priority float64
}

// heapQueue 是一个基于堆实现的优先级队列
type heapQueue[T any] []*heapItem[T]

// Len 返回优先级队列的长度
func (pq *heapQueue[T]) Len() int {
	return len(*pq)
}

// Less 比较两个项目的优先级
// 如果返回 true，表示第一个项目的优先级高于第二个项目
func (pq *heapQueue[T]) Less(i, j int) bool {
	return (*pq)[i].priority > (*pq)[j].priority
}

// Swap 交换两个项目的位置
func (pq *heapQueue[T]) Swap(i, j int) {
	(*pq)[i], (*pq)[j] = (*pq)[j], (*pq)[i]
}

// Push 向优先级队列添加一个新项目
func (pq *heapQueue[T]) Push(x any) {
	item := x.(*heapItem[T])
	*pq = append(*pq, item)
}

// Pop 从优先级队列中移除并返回优先级最高的项目
func (pq *heapQueue[T]) Pop() any {
	old := *pq
	n := len(old)
	item := old[n-1]
	*pq = old[0 : n-1]
	return item
}
