package transport

import (
	"github.com/kercylan98/minotaur/minotaur/vivid"
	"time"
)

type Conn = vivid.TypedActorRef[ConnActorTyped]

type ConnActorTyped interface {
	// Write 向连接内写入数据包
	Write(packet Packet)

	// Close 关闭连接
	Close()

	// LoadMod 加载模组
	LoadMod(mods ...vivid.ModInfo)

	// UnloadMod 卸载模组
	UnloadMod(mods ...vivid.ModInfo)

	// ApplyMod 应用模组
	ApplyMod()

	// SetPacketHandler 设置数据包处理器
	//   - 该函数是通过 become 实现，并且会丢弃旧的行为
	SetPacketHandler(handler ConnPacketHandler)

	// SetTerminateHandler 设置连接终止处理器
	SetTerminateHandler(handler ConnTerminateHandler)

	// SetZombieTimeout 设置僵尸连接超时时间
	SetZombieTimeout(timeout time.Duration)
}

type ConnActorTypedImpl struct {
	ConnActorRef       vivid.ActorRef
	ConnWriterActorRef vivid.ActorRef
}

func (c *ConnActorTypedImpl) Write(packet Packet) {
	c.ConnWriterActorRef.Tell(packet)
}

func (c *ConnActorTypedImpl) Close() {
	c.ConnActorRef.Stop()
}

func (c *ConnActorTypedImpl) SetPacketHandler(handler ConnPacketHandler) {
	c.ConnActorRef.Tell(ConnectionSetPacketHandlerMessage{Handler: handler})
}

func (c *ConnActorTypedImpl) SetTerminateHandler(handler ConnTerminateHandler) {
	c.ConnActorRef.Tell(ConnectionSetTerminateHandlerMessage{Handler: handler})
}

func (c *ConnActorTypedImpl) LoadMod(mods ...vivid.ModInfo) {
	c.ConnActorRef.Tell(ConnectionLoadModMessage{Mods: mods}, vivid.WithInstantly(true))
}

func (c *ConnActorTypedImpl) UnloadMod(mods ...vivid.ModInfo) {
	c.ConnActorRef.Tell(ConnectionUnloadModMessage{Mods: mods}, vivid.WithInstantly(true))
}

func (c *ConnActorTypedImpl) ApplyMod() {
	c.ConnActorRef.Tell(ConnectionApplyModMessage{}, vivid.WithInstantly(true))
}

func (c *ConnActorTypedImpl) SetZombieTimeout(timeout time.Duration) {
	c.ConnActorRef.Tell(ConnectionSetZombieTimeoutMessage{Timeout: timeout})
}
