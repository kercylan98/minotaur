package network

import (
	"context"
	"github.com/kercylan98/minotaur/minotaur/transport"
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

func newKcpCore(addr string) transport.Network {
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
	addr   string
	closed atomic.Bool
	ctx    context.Context
	cancel context.CancelFunc
}

func (k *kcpCore) Launch(ctx context.Context, srv transport.ServerActorTyped) error {
	k.ctx, k.cancel = context.WithCancel(ctx)
	lis, err := kcp.ListenWithOptions(k.addr, nil, 0, 0)
	if err != nil {
		return err
	}
	defer func(lis *kcp.Listener) {
		_ = lis.Close()
	}(lis)
	for !k.closed.Load() {
		var session *kcp.UDPSession
		if session, err = lis.AcceptKCP(); err != nil {
			continue
		}

		// 注册连接
		conn := srv.Attach(session, func(packet transport.Packet) error {
			_, err = session.Write(packet.GetBytes())
			return err
		})

		// 处理连接数据
		go func(ctx context.Context, session *kcp.UDPSession, conn transport.Conn) {
			var buf = make([]byte, 1024)
			var n int
			for {
				select {
				case <-ctx.Done():
					return
				default:
					if n, err = session.Read(buf); err != nil {
						conn.Stop()
						return
					}
					conn.React(transport.NewPacket(buf[:n]))
				}
			}
		}(k.ctx, session, conn)
	}
	return nil
}

func (k *kcpCore) Shutdown() error {
	k.closed.Store(true)
	k.cancel()
	return nil
}

func (k *kcpCore) Schema() string {
	return "udp(kcp)"
}

func (k *kcpCore) Address() string {
	return k.addr
}
