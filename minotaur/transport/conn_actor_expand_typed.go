package transport

import "github.com/kercylan98/minotaur/minotaur/vivid"

type ConnActorExpandTyped interface {
	React(packet Packet)
}

type ConnActorExpandTypedImpl struct {
	ConnActorRef vivid.ActorRef
}

func (c *ConnActorExpandTypedImpl) React(packet Packet) {
	c.ConnActorRef.Tell(ConnectionReactPacketMessage{Packet: packet})
}
