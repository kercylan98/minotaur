package loadbalancer

type RoundRobinItem[Id comparable] interface {
	// Id 返回唯一标识
	Id() Id
}
