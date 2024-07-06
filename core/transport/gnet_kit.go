package transport

import (
	"github.com/kercylan98/minotaur/core/vivid"
	"github.com/panjf2000/gnet/v2"
)

type (
	GNETConnectionOpenedHook func(kit *GNETKit, conn *Conn) error
	GNETConnectionClosedHook func(kit *GNETKit, conn *Conn, err error)
	GNETConnectionPacketHook func(kit *GNETKit, conn *Conn, packet Packet) error
)

type GNETKit struct {
	ownerRef    vivid.ActorRef
	app         gnet.Engine
	actorSystem *vivid.ActorSystem

	connectionOpenedHook GNETConnectionOpenedHook
	connectionClosedHook GNETConnectionClosedHook
	connectionPacketHook GNETConnectionPacketHook
}

func (k *GNETKit) App() gnet.Engine {
	return k.app
}

func (k *GNETKit) System() *vivid.ActorSystem {
	return k.actorSystem
}

// ConnectionOpenedHook 绑定连接打开钩子函数，该函数将在连接打开时调用，返回错误则关闭连接
//   - 该函数的执行会在独立的 FiberActor 中执行，所以是并发安全的
func (k *GNETKit) ConnectionOpenedHook(hook GNETConnectionOpenedHook) *GNETKit {
	k.connectionOpenedHook = hook
	return k
}

// ConnectionClosedHook 绑定连接关闭钩子函数，该函数将在连接关闭时调用
//   - 该函数的执行会在独立的 FiberActor 中执行，所以是并发安全的
func (k *GNETKit) ConnectionClosedHook(hook GNETConnectionClosedHook) *GNETKit {
	k.connectionClosedHook = hook
	return k
}

// ConnectionPacketHook 绑定连接数据包钩子函数，该函数将在接收到数据包时调用，返回错误则关闭连接
//   - 该函数的执行是在维护连接的 Actor 中进行的，连接与连接之间是相互隔离的
func (k *GNETKit) ConnectionPacketHook(hook GNETConnectionPacketHook) *GNETKit {
	k.connectionPacketHook = hook
	return k
}
