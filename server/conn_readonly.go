package server

import (
	"net"
)

type ConnReadonly interface {
	RemoteAddr() net.Addr
	GetID() string
	GetIP() string // GetData 获取连接数据
	GetData(key any) any
	IsWebsocket() bool
}
