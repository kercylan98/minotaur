package server

import (
	"net"
)

// ConnReadonly 连接只读接口
type ConnReadonly interface {
	// RemoteAddr 获取远程地址
	RemoteAddr() net.Addr
	// GetID 获取连接 ID
	GetID() string
	// GetIP 获取连接 IP
	GetIP() string
	// GetData 获取连接数据
	GetData(key any) any
	// IsWebsocket 是否是 websocket 连接
	IsWebsocket() bool
}
