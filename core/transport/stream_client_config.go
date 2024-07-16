package transport

import (
	"github.com/kercylan98/minotaur/core/vivid"
)

type StreamClientConfig struct {
	ConnectionOpenedHandler func(ctx vivid.ActorContext)
	ConnectionPacketHandler func(ctx vivid.ActorContext, packet Packet)
	ConnectionClosedHandler func(ctx vivid.ActorContext, err error)
}
