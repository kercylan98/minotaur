package transport

import (
	"github.com/kercylan98/minotaur/minotaur/vivid"
	"net"
	"time"
)

type Conn = ConnActorTyped

type ConnActorTyped interface {
	vivid.ActorTyped

	Init(conn net.Conn, writer ConnWriter)

	React(packet Packet)

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

func (c *ConnActor) Init(conn net.Conn, writer ConnWriter) {
	c.Tell(ConnectionInitMessage{Conn: conn, Writer: writer})
}

func (c *ConnActor) React(packet Packet) {
	c.Tell(ConnectionReactPacketMessage{Packet: packet})
}

func (c *ConnActor) Write(packet Packet) {
	c.Writer.Tell(packet)
}

func (c *ConnActor) Close() {
	c.Stop(true)
}

func (c *ConnActor) SetPacketHandler(handler ConnPacketHandler) {
	c.Tell(ConnectionSetPacketHandlerMessage{Handler: handler})
}

func (c *ConnActor) SetTerminateHandler(handler ConnTerminateHandler) {
	c.Tell(ConnectionSetTerminateHandlerMessage{Handler: handler})
}

func (c *ConnActor) LoadMod(mods ...vivid.ModInfo) {
	c.Tell(ConnectionLoadModMessage{Mods: mods}, vivid.WithInstantly(true))
}

func (c *ConnActor) UnloadMod(mods ...vivid.ModInfo) {
	c.Tell(ConnectionUnloadModMessage{Mods: mods}, vivid.WithInstantly(true))
}

func (c *ConnActor) ApplyMod() {
	c.Tell(ConnectionApplyModMessage{}, vivid.WithInstantly(true))
}

func (c *ConnActor) SetZombieTimeout(timeout time.Duration) {
	c.Tell(ConnectionSetZombieTimeoutMessage{Timeout: timeout})
}
