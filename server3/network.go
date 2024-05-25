package server

import (
	"context"
)

// Network 提供给 Server 使用的网络接口
type Network interface {
	// OnSetup 装载函数，该函数将传入 Server 的上下文及一个可用于控制 Server 行为的控制器，在该函数中应该完成 Network 实现的初始化工作
	OnSetup(ctx context.Context) error

	// OnRun 运行函数，在该函数中通常是进行网络监听启动的阻塞行为
	OnRun() error

	// OnShutdown 停止函数，该函数在执行的时候应该完成对资源的停止及释放
	OnShutdown() error

	// Schema 获取协议标识
	Schema() string

	// Address 获取监听地址信息
	Address() string
}
