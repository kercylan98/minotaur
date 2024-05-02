package network

import (
	"context"
	"github.com/kercylan98/minotaur/server"
	"github.com/xtaci/kcp-go/v5"
	"runtime"
	"sync"
	"sync/atomic"
)

var kcpInitOnce sync.Once

func init() {
	kcp.SystemTimedSched.Close() // 默认禁用 KCP 系统定时器
	kcp.SystemTimedSched = nil
}

func newKcpCore(addr string) server.Network {
	kcpInitOnce.Do(func() {
		if kcp.SystemTimedSched == nil {
			kcp.SystemTimedSched = kcp.NewTimedSched(runtime.NumCPU())
		}
	})
	return &kcpCore{
		addr: addr,
	}
}

type kcpCore struct {
	ctx        context.Context
	controller server.Controller
	addr       string
	closed     atomic.Bool
}

func (k *kcpCore) OnSetup(ctx context.Context, controller server.Controller) error {
	k.ctx = ctx
	k.controller = controller
	return nil
}

func (k *kcpCore) OnRun() error {
	lis, err := kcp.ListenWithOptions(k.addr, nil, 0, 0)
	if err != nil {
		return err
	}
	defer func(lis *kcp.Listener) {
		_ = lis.Close()
	}(lis)
	for !k.closed.Load() {
		var conn *kcp.UDPSession
		var srvConn server.Conn
		if conn, err = lis.AcceptKCP(); err != nil {
			continue
		}

		// 注册连接
		k.controller.RegisterConnection(conn,
			func(packet server.Packet) (err error) {
				if _, err = conn.Write(packet.GetBytes()); err != nil {
					k.controller.OnConnectionAsyncWriteError(srvConn, packet, err)
				}
				return
			}, func(conn server.Conn) {
				srvConn = conn
			},
		)

		// 处理连接数据
		go func(ctx context.Context, conn *kcp.UDPSession, srvConn server.Conn) {
			var buf = make([]byte, 1024)
			var n int
			for {
				select {
				case <-ctx.Done():
					return
				default:
					if n, err = conn.Read(buf); err != nil {
						srvConn.Close()
						return
					}
					k.controller.ReactPacket(conn, server.NewPacket(buf[:n]))
				}
			}
		}(k.ctx, conn, srvConn)
	}

	return nil
}

func (k *kcpCore) OnShutdown() error {
	k.closed.Store(true)
	return nil
}

func (k *kcpCore) Schema() string {
	return "kcp(udp)"
}

func (k *kcpCore) Address() string {
	return k.addr
}
