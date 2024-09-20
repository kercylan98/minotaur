package socket

import "github.com/kercylan98/minotaur/engine/vivid"

// Actor 是为 vivid.Actor 扩展网络功能的必选接口，该接口将实现 Socket 的数据包响应功能，以确保来自网络连接的数据得到承接
type Actor interface {
	vivid.Actor

	// OnPacket 当收到数据包时将会触发该函数，该函数不会阻止 vivid.Actor.OnReceive 的处理，这也就意味着可以在 OnReceive 中断言 *Packet 类型的消息进行处理
	OnPacket(ctx vivid.ActorContext, socket Socket, packet *Packet)
}

// OpenedActor 是一个 Actor 的可选扩展接口，在实现该接口后，该接口将会在 Socket 连接成功时触发 OnOpened 函数
type OpenedActor interface {
	Actor

	// OnOpened 当 Socket 连接成功时将会触发该函数
	OnOpened(ctx vivid.ActorContext, socket Socket)
}

// CloseActor 是一个 Actor 的可选扩展接口，在实现该接口后，该接口将会在 Socket 连接断开时触发 OnClose 函数
type CloseActor interface {
	Actor

	// OnClose 当 Socket 连接断开时将会触发该函数，根据不同的场景，此刻 Socket 可能存在无法写入消息的情况（例如网络原因已断开等）
	OnClose(ctx vivid.ActorContext, socket Socket, err error)
}
