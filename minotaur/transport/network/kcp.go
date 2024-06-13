package network

import (
	"context"
	"github.com/kercylan98/minotaur/minotaur/transport"
	"github.com/kercylan98/minotaur/minotaur/vivid"
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
}

func (k *kcpCore) Launch(ctx context.Context, srv vivid.TypedActorRef[transport.ServerActorExpandTyped]) error {
	//TODO implement me
	panic("implement me")
}

func (k *kcpCore) Shutdown() error {
	//TODO implement me
	panic("implement me")
}

func (k *kcpCore) Schema() string {
	//TODO implement me
	panic("implement me")
}

func (k *kcpCore) Address() string {
	//TODO implement me
	panic("implement me")
}

//
//func (k *kcpCore) OnPreStart(ctx vivids.ActorContext) (err error) {
//	k.ctx = ctx
//	ctx.Future(func() vivids.Message {
//		lis, err := kcp.ListenWithOptions(k.addr, nil, 0, 0)
//		if err != nil {
//			return err
//		}
//		defer func(lis *kcp.Listener) {
//			_ = lis.Close()
//		}(lis)
//		for !k.closed.Load() {
//			var conn *kcp.UDPSession
//			var srvConn server.Conn
//			if conn, err = lis.AcceptKCP(); err != nil {
//				continue
//			}
//
//			// 注册连接
//			srvConn, err = k.ctx.GetParentActor().Ask(server.ServerNetworkConnectionOpenedEvent{
//				Conn: conn,
//				ConnectionWriter: func(packet server.Packet) (err error) {
//					if _, err = conn.Write(packet.GetBytes()); err != nil {
//						return k.ctx.GetParentActor().Tell(server.NetworkConnectionAsyncWriteErrorMessage{
//							Conn:   srvConn,
//							Packet: packet,
//							Error:  err,
//						})
//					}
//					return
//				},
//			})
//
//			// 处理连接数据
//			go func(ctx context.Context, conn *kcp.UDPSession, srvConn server.Conn) {
//				var buf = make([]byte, 1024)
//				var n int
//				for {
//					select {
//					case <-ctx.Done():
//						return
//					default:
//						if n, err = conn.Read(buf); err != nil {
//							srvConn.Close()
//							return
//						}
//						k.controller.ReactPacket(conn, server.NewPacket(buf[:n]))
//					}
//				}
//			}(ctx, conn, srvConn)
//		}
//
//		return nil
//	})
//
//	return
//}
//
//func (k *kcpCore) OnReceived(ctx vivids.MessageContext) (err error) {
//	switch v := ctx.GetMessage().(type) {
//	case error:
//		ctx.NotifyTerminated(v)
//	}
//
//	return
//}
//
//func (k *kcpCore) OnDestroy(ctx vivids.ActorContext) (err error) {
//	k.closed.Store(true)
//	return nil
//}
