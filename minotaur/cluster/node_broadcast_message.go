package cluster

import (
	"github.com/hashicorp/memberlist"
	"github.com/kercylan98/minotaur/toolkit"
)

type BroadcastMessage struct {
	Key   string `json:"key"`
	Value uint64 `json:"value"`
}

func (m BroadcastMessage) Invalidates(other memberlist.Broadcast) bool {
	return false
}

func (m BroadcastMessage) Finished() {
	// nop
}

func (m BroadcastMessage) Message() []byte {
	return toolkit.MarshalJSON(m)
}
