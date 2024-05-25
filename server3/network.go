package server

import "context"

// Network 提供给 server 使用的网络接口
type Network interface {
	// Launch 启动函数，该函数在执行的时候应该完成对资源的初始化及启动
	Launch(ctx context.Context, srv Core) error

	// Shutdown 停止函数，该函数在执行的时候应该完成对资源的停止及释放
	Shutdown() error

	// Schema 获取协议标识
	Schema() string

	// Address 获取监听地址信息
	Address() string
}
