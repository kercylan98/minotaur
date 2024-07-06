package transport

import "github.com/gofiber/fiber/v2"

type (
	FiberWebSocketUpgradeHook func(kit *FiberKit, ctx *fiber.Ctx) error
	FiberConnectionOpenedHook func(kit *FiberKit, ctx *FiberContext, conn *Conn) error
	FiberConnectionClosedHook func(kit *FiberKit, ctx *FiberContext, conn *Conn, err error)
	FiberConnectionPacketHook func(kit *FiberKit, ctx *FiberContext, conn *Conn, packet Packet) error
)

type FiberWebSocket struct {
	kit *FiberKit

	upgradeHook          FiberWebSocketUpgradeHook
	connectionOpenedHook FiberConnectionOpenedHook
	connectionClosedHook FiberConnectionClosedHook
	connectionPacketHook FiberConnectionPacketHook
}

func (fw *FiberWebSocket) init() {
	fw.upgradeHook = func(kit *FiberKit, ctx *fiber.Ctx) error { return nil }
	fw.connectionOpenedHook = func(kit *FiberKit, ctx *FiberContext, conn *Conn) error { return nil }
	fw.connectionClosedHook = func(kit *FiberKit, ctx *FiberContext, conn *Conn, err error) {}
	fw.connectionPacketHook = func(kit *FiberKit, ctx *FiberContext, conn *Conn, packet Packet) error { return nil }
}

// UpgradeHook 绑定升级钩子函数，该函数将在 WebSocket 升级时调用，返回错误则终止升级
func (fw *FiberWebSocket) UpgradeHook(hook FiberWebSocketUpgradeHook) *FiberWebSocket {
	fw.upgradeHook = hook
	return fw
}

// ConnectionOpenedHook 绑定连接打开钩子函数，该函数将在连接打开时调用，返回错误则关闭连接
//   - 该函数的执行会在独立的 FiberActor 中执行，所以是并发安全的
func (fw *FiberWebSocket) ConnectionOpenedHook(hook FiberConnectionOpenedHook) *FiberWebSocket {
	fw.connectionOpenedHook = hook
	return fw
}

// ConnectionClosedHook 绑定连接关闭钩子函数，该函数将在连接关闭时调用
//   - 该函数的执行会在独立的 FiberActor 中执行，所以是并发安全的
func (fw *FiberWebSocket) ConnectionClosedHook(hook FiberConnectionClosedHook) *FiberWebSocket {
	fw.connectionClosedHook = hook
	return fw
}

// ConnectionPacketHook 绑定连接数据包钩子函数，该函数将在接收到数据包时调用，返回错误则关闭连接
//   - 该函数的执行是在维护连接的 Actor 中进行的，连接与连接之间是相互隔离的
func (fw *FiberWebSocket) ConnectionPacketHook(hook FiberConnectionPacketHook) *FiberWebSocket {
	fw.connectionPacketHook = hook
	return fw
}
