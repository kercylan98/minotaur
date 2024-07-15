package transport

import (
	"github.com/kercylan98/minotaur/core/vivid"
	"net/http"
)

type WebSocketClientConfig struct {
	Header                  http.Header
	ConnectionOpenedHandler func(ctx vivid.ActorContext)
	ConnectionPacketHandler func(ctx vivid.ActorContext, packet Packet)
	ConnectionClosedHandler func(ctx vivid.ActorContext, err error)
}
