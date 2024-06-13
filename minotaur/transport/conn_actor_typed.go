package transport

import "github.com/kercylan98/minotaur/minotaur/vivid"

type ConnActorTyped interface {
	// Write 向连接内写入数据包
	Write(packet Packet)

	// Close 关闭连接
	Close()

	// BecomeReactPacketMessage 切换响应数据包消息行为，默认情况下会发布 ConnectionReceivePacketEvent 事件，可以通过 SubscribePacketReceivedEvent 订阅
	BecomeReactPacketMessage(handler func(vivid.MessageContext, ConnectionReactPacketMessage))

	// SubscribePacketReceivedEvent 订阅数据包接收事件
	SubscribePacketReceivedEvent(subscriber vivid.Subscriber)

	// LoadMod 加载模组
	LoadMod(mods ...vivid.ModInfo)

	// UnloadMod 卸载模组
	UnloadMod(mods ...vivid.ModInfo)

	// ApplyMod 应用模组
	ApplyMod()
}

type ConnActorTypedImpl struct {
	ConnActorRef vivid.ActorRef
}

func (c *ConnActorTypedImpl) Write(packet Packet) {
	c.ConnActorRef.Tell(ConnectionWritePacketMessage{Packet: packet})
}

func (c *ConnActorTypedImpl) Close() {
	c.ConnActorRef.Stop()
}

func (c *ConnActorTypedImpl) SubscribePacketReceivedEvent(subscriber vivid.Subscriber) {
	c.ConnActorRef.GetSystem().Subscribe(subscriber, ConnectionReceivePacketEvent{})
}

func (c *ConnActorTypedImpl) BecomeReactPacketMessage(handler func(vivid.MessageContext, ConnectionReactPacketMessage)) {
	c.ConnActorRef.Tell(ConnectionBecomeReactPacketMessage{Behavior: vivid.BehaviorOf[ConnectionReactPacketMessage](handler)})
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
