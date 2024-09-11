package socket

import "github.com/kercylan98/minotaur/engine/vivid"

type Actor interface {
	vivid.Actor

	// OnPacket 当收到数据包时将会触发该函数，该函数不会阻止 vivid.Actor.OnReceive 的处理，这也就意味着可以在 OnReceive 中断言 *Packet 类型的消息进行处理
	OnPacket(ctx vivid.ActorContext, socket Socket, packet *Packet)
}

type OpenedActor interface {
	Actor

	// OnOpened 当 Socket 连接成功时将会触发该函数
	OnOpened(ctx vivid.ActorContext, socket Socket)
}

type CloseActor interface {
	Actor

	// OnClose 当 Socket 连接断开时将会触发该函数
	OnClose(ctx vivid.ActorContext, socket Socket, err error)
}
