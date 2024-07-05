package transport

import "github.com/gofiber/fiber/v2"

type (
	FiberWebSocketUpgradeHook func(kit *FiberKit, ctx *fiber.Ctx) error
	ConnectionOpenedHook      func(kit *FiberKit, ctx *FiberContext, conn *Conn) error
	ConnectionClosedHook      func(kit *FiberKit, ctx *FiberContext, conn *Conn, err error)
	ConnectionPacketHook      func(kit *FiberKit, ctx *FiberContext, conn *Conn, packet Packet) error
)

type FiberWebSocket struct {
	kit *FiberKit

	upgradeHook          FiberWebSocketUpgradeHook
	connectionOpenedHook ConnectionOpenedHook
	connectionClosedHook ConnectionClosedHook
	connectionPacketHook ConnectionPacketHook
}

func (fw *FiberWebSocket) init() {
	fw.upgradeHook = func(kit *FiberKit, ctx *fiber.Ctx) error { return nil }
	fw.connectionOpenedHook = func(kit *FiberKit, ctx *FiberContext, conn *Conn) error { return nil }
	fw.connectionClosedHook = func(kit *FiberKit, ctx *FiberContext, conn *Conn, err error) {}
	fw.connectionPacketHook = func(kit *FiberKit, ctx *FiberContext, conn *Conn, packet Packet) error { return nil }
}

func (fw *FiberWebSocket) UpgradeHook(hook FiberWebSocketUpgradeHook) *FiberWebSocket {
	fw.upgradeHook = hook
	return fw
}

func (fw *FiberWebSocket) ConnectionOpenedHook(hook ConnectionOpenedHook) *FiberWebSocket {
	fw.connectionOpenedHook = hook
	return fw
}

func (fw *FiberWebSocket) ConnectionClosedHook(hook ConnectionClosedHook) *FiberWebSocket {
	fw.connectionClosedHook = hook
	return fw
}

func (fw *FiberWebSocket) ConnectionPacketHook(hook ConnectionPacketHook) *FiberWebSocket {
	fw.connectionPacketHook = hook
	return fw
}
