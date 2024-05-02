package rpc

// Core RPC 核心接口
type Core interface {
	// OnInit 初始化 RPC 核心
	OnInit(info ServiceInfo, matcher RouteMatcher, routes [][]Route) error

	OnCall(routes ...Route) func(request any) error

	// Close 停止 RPC 核心
	Close()
}
