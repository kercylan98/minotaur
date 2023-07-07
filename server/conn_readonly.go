package server

import (
	"net"
)

type ConnReadonly interface {
	RemoteAddr() net.Addr
	GetID() string
	GetIP() string
	GetData(key any) any
	IsWebsocket() bool
}
