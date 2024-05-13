package rpc

import "net"

// Transporter 是一个 RPC 传输器的接口，该接口用于定义一个 RPC 传输器
type Transporter interface {
	// OnInit 用于初始化一个传输器
	OnInit(srv Server)

	// Serve 用于启动一个 RPC 传输器
	Serve(l net.Listener) error
}
