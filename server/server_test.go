package server_test

import (
	"github.com/kercylan98/minotaur/server"
	"github.com/kercylan98/minotaur/utils/super"
	"runtime/debug"
	"testing"
	"time"
)

// 该单元测试用于测试以不同的基本参数创建服务器是否存在异常
func TestNew(t *testing.T) {
	var cases = []struct {
		name        string
		network     server.Network
		addr        string
		shouldPanic bool
	}{
		{name: "TestNew_Unknown", addr: "", network: "Unknown", shouldPanic: true},
		{name: "TestNew_None", addr: "", network: server.NetworkNone, shouldPanic: false},
		{name: "TestNew_None_Addr", addr: "addr", network: server.NetworkNone, shouldPanic: false},
		{name: "TestNew_Tcp_AddrEmpty", addr: "", network: server.NetworkTcp, shouldPanic: true},
		{name: "TestNew_Tcp_AddrIllegal", addr: "addr", network: server.NetworkTcp, shouldPanic: true},
		{name: "TestNew_Tcp_Addr", addr: ":9999", network: server.NetworkTcp, shouldPanic: false},
		{name: "TestNew_Tcp4_AddrEmpty", addr: "", network: server.NetworkTcp4, shouldPanic: true},
		{name: "TestNew_Tcp4_AddrIllegal", addr: "addr", network: server.NetworkTcp4, shouldPanic: true},
		{name: "TestNew_Tcp4_Addr", addr: ":9999", network: server.NetworkTcp4, shouldPanic: false},
		{name: "TestNew_Tcp6_AddrEmpty", addr: "", network: server.NetworkTcp6, shouldPanic: true},
		{name: "TestNew_Tcp6_AddrIllegal", addr: "addr", network: server.NetworkTcp6, shouldPanic: true},
		{name: "TestNew_Tcp6_Addr", addr: ":9999", network: server.NetworkTcp6, shouldPanic: false},
		{name: "TestNew_Udp_AddrEmpty", addr: "", network: server.NetworkUdp, shouldPanic: true},
		{name: "TestNew_Udp_AddrIllegal", addr: "addr", network: server.NetworkUdp, shouldPanic: true},
		{name: "TestNew_Udp_Addr", addr: ":9999", network: server.NetworkUdp, shouldPanic: false},
		{name: "TestNew_Udp4_AddrEmpty", addr: "", network: server.NetworkUdp4, shouldPanic: true},
		{name: "TestNew_Udp4_AddrIllegal", addr: "addr", network: server.NetworkUdp4, shouldPanic: true},
		{name: "TestNew_Udp4_Addr", addr: ":9999", network: server.NetworkUdp4, shouldPanic: false},
		{name: "TestNew_Udp6_AddrEmpty", addr: "", network: server.NetworkUdp6, shouldPanic: true},
		{name: "TestNew_Udp6_AddrIllegal", addr: "addr", network: server.NetworkUdp6, shouldPanic: true},
		{name: "TestNew_Udp6_Addr", addr: ":9999", network: server.NetworkUdp6, shouldPanic: false},
		{name: "TestNew_Unix_AddrEmpty", addr: "", network: server.NetworkUnix, shouldPanic: true},
		{name: "TestNew_Unix_AddrIllegal", addr: "addr", network: server.NetworkUnix, shouldPanic: true},
		{name: "TestNew_Unix_Addr", addr: "addr", network: server.NetworkUnix, shouldPanic: false},
		{name: "TestNew_Websocket_AddrEmpty", addr: "", network: server.NetworkWebsocket, shouldPanic: true},
		{name: "TestNew_Websocket_AddrIllegal", addr: "addr", network: server.NetworkWebsocket, shouldPanic: true},
		{name: "TestNew_Websocket_Addr", addr: ":9999/ws", network: server.NetworkWebsocket, shouldPanic: false},
		{name: "TestNew_Http_AddrEmpty", addr: "", network: server.NetworkHttp, shouldPanic: true},
		{name: "TestNew_Http_AddrIllegal", addr: "addr", network: server.NetworkHttp, shouldPanic: true},
		{name: "TestNew_Http_Addr", addr: ":9999", network: server.NetworkHttp, shouldPanic: false},
		{name: "TestNew_Kcp_AddrEmpty", addr: "", network: server.NetworkKcp, shouldPanic: true},
		{name: "TestNew_Kcp_AddrIllegal", addr: "addr", network: server.NetworkKcp, shouldPanic: true},
		{name: "TestNew_Kcp_Addr", addr: ":9999", network: server.NetworkKcp, shouldPanic: false},
		{name: "TestNew_GRPC_AddrEmpty", addr: "", network: server.NetworkGRPC, shouldPanic: true},
		{name: "TestNew_GRPC_AddrIllegal", addr: "addr", network: server.NetworkGRPC, shouldPanic: true},
		{name: "TestNew_GRPC_Addr", addr: ":9999", network: server.NetworkGRPC, shouldPanic: false},
	}

	for _, c := range cases {
		t.Run(c.name, func(t *testing.T) {
			defer func() {
				if err := super.RecoverTransform(recover()); err != nil && !c.shouldPanic {
					debug.PrintStack()
					t.Fatal("not should panic, err:", err)
				}
			}()
			if err := server.New(c.network, server.WithLimitLife(time.Millisecond*10)).Run(""); err != nil {
				panic(err)
			}
		})
	}
}

// 这个测试检查了各个类型的服务器是否为 Socket 模式。如需查看为 Socket 模式的网络类型，请参考 [` Network.IsSocket` ](#struct_Network_IsSocket)
func TestServer_IsSocket(t *testing.T) {
	var cases = []struct {
		name    string
		network server.Network
		expect  bool
	}{
		{name: "TestServer_IsSocket_None", network: server.NetworkNone, expect: false},
		{name: "TestServer_IsSocket_Tcp", network: server.NetworkTcp, expect: true},
		{name: "TestServer_IsSocket_Tcp4", network: server.NetworkTcp4, expect: true},
		{name: "TestServer_IsSocket_Tcp6", network: server.NetworkTcp6, expect: true},
		{name: "TestServer_IsSocket_Udp", network: server.NetworkUdp, expect: true},
		{name: "TestServer_IsSocket_Udp4", network: server.NetworkUdp4, expect: true},
		{name: "TestServer_IsSocket_Udp6", network: server.NetworkUdp6, expect: true},
		{name: "TestServer_IsSocket_Unix", network: server.NetworkUnix, expect: true},
		{name: "TestServer_IsSocket_Http", network: server.NetworkHttp, expect: false},
		{name: "TestServer_IsSocket_Websocket", network: server.NetworkWebsocket, expect: true},
		{name: "TestServer_IsSocket_Kcp", network: server.NetworkKcp, expect: true},
		{name: "TestServer_IsSocket_GRPC", network: server.NetworkGRPC, expect: false},
	}

	for _, c := range cases {
		t.Run(c.name, func(t *testing.T) {
			s := server.New(c.network)
			if s.IsSocket() != c.expect {
				t.Fatalf("expect: %v, got: %v", c.expect, s.IsSocket())
			}
		})
	}
}
