package balancer

type Item[Id comparable] interface {
	// Id 返回唯一标识
	Id() Id

	// Weight 返回权重
	Weight() int
}

type Balancer[Id comparable, T Item[Id]] interface {
	// Add 添加一个负载均衡目标
	Add(t T)

	// Remove 移除一个负载均衡目标
	Remove(t T)

	// Next 根据负载均衡策略选择下一个目标
	Next() T
}
