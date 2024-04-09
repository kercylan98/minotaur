package loadbalancer

type RoundRobinItem[Id comparable] interface {
	// GetId 返回唯一标识
	GetId() Id
}
