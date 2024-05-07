package balancer

import "github.com/kercylan98/minotaur/toolkit/constraints"

// Item 负载均衡器的实例
type Item[I constraints.Ordered] interface {
	// GetId 返回唯一标识
	GetId() I

	// GetWeight 返回权重
	GetWeight() int
}

// MetadataItem 元数据负载均衡器的实例
type MetadataItem[I constraints.Ordered, M any] interface {
	Item[I]

	// GetMetadata 返回元数据
	GetMetadata() M
}
