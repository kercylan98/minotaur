package server

import (
	"github.com/kercylan98/minotaur/vivid/vivids"
)

type Network interface {
	vivids.Actor
}

type (
	NetworkConnectionOpenedEvent struct {
		vivids.ActorRef
	}

	NetworkConnectionClosedEvent struct {
		vivids.ActorRef
	}

	NetworkConnectionReceivedMessage struct {
		Packet Packet
	}

	ConnectionWriteMessage struct {
		Packet Packet
	}

	ConnectionAsyncWriteErrorEvent struct {
		Error error
	}
)
