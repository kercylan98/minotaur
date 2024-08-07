package prc

import "google.golang.org/grpc"

// SharedStartHook 共享启动时的钩子
type SharedStartHook interface {
	// OnSharedStart 共享启动
	OnSharedStart()
}

// FunctionalSharedStartHook 是一个函数式的共享启动时的钩子
type FunctionalSharedStartHook func()

// OnSharedStart 共享启动
func (f FunctionalSharedStartHook) OnSharedStart() {
	f()
}

// ShareOpenedHook 共享打开后的钩子
type ShareOpenedHook interface {
	// OnShareOpened 共享打开后
	OnShareOpened(target PhysicalAddress)
}

// FunctionalShareOpenedHook 是一个函数式的共享打开后的钩子
type FunctionalShareOpenedHook func(target PhysicalAddress)

// OnShareOpened 共享打开后
func (f FunctionalShareOpenedHook) OnShareOpened(target PhysicalAddress) {
	f(target)
}

// SharedClosedHook 共享关闭后的钩子
type SharedClosedHook interface {
	// OnShareClosed 共享关闭后
	OnShareClosed(target PhysicalAddress)
}

// FunctionalSharedClosedHook 是一个函数式的共享关闭后的钩子
type FunctionalSharedClosedHook func(target PhysicalAddress)

// OnShareClosed 共享关闭后
func (f FunctionalSharedClosedHook) OnShareClosed(target PhysicalAddress) {
	f(target)
}

// GRPCLaunchBeforeHook GRPC启动前的钩子
type GRPCLaunchBeforeHook interface {
	// OnGRPCLaunchBefore GRPC启动前
	OnGRPCLaunchBefore(server *grpc.Server)
}

// FunctionalGRPCLaunchBeforeHook 是一个函数式的GRPC启动前的钩子
type FunctionalGRPCLaunchBeforeHook func(server *grpc.Server)

// OnGRPCLaunchBefore GRPC启动前
func (f FunctionalGRPCLaunchBeforeHook) OnGRPCLaunchBefore(server *grpc.Server) {
	f(server)
}
