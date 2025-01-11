package balancer

type InstanceID = string

// Instance 负载均衡器实例
type Instance interface {
	// GetID 获取实例 ID，该 ID 在整个负载均衡器中应该是唯一的
	GetID() InstanceID
}
