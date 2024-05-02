package rpc

type Service interface {
	// OnRPCSetup 装载 RPC 服务
	OnRPCSetup(router Router)
}

// ServiceInfo 描述该 RPC 服务进程的信息
type ServiceInfo struct {
	UniqueId string // 在整个 RPC 网络中的唯一标识
	Name     string // 服务名称
}
