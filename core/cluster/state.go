package cluster

import (
	"bytes"
	"encoding/gob"
	"github.com/hashicorp/memberlist"
	"sync"
)

type state struct {
	mu         *sync.RWMutex
	node       *Node
	metadata   *metadata
	local      *local
	broadcast  []byte
	broadcasts *memberlist.TransmitLimitedQueue
}

func (c *state) Invalidates(b memberlist.Broadcast) bool {
	return false
}

func (c *state) Message() []byte {
	return c.broadcast
}

func (c *state) Finished() {}

func newState(node *Node, metadata *metadata) *state {
	mu := new(sync.RWMutex)
	metadata.mu = mu
	return &state{
		mu:       mu,
		node:     node,
		metadata: metadata,
		local:    &local{mu: mu},
		broadcasts: &memberlist.TransmitLimitedQueue{
			NumNodes: func() int {
				return node.memberlist.NumMembers()
			},
			RetransmitMult: 3,
		},
	}
}

func (c *state) NodeMeta(limit int) []byte {
	data, err := c.metadata.bytes()
	if err != nil {
		panic(err)
	}
	return data
}

func (c *state) NotifyMsg(buf []byte) {
	var other = new(local)
	b := bytes.NewBuffer(buf)
	decoder := gob.NewDecoder(b)
	err := decoder.Decode(&other)
	if err != nil {
		panic(err)
	}

	c.local.merge(other)
}

func (c *state) GetBroadcasts(overhead, limit int) [][]byte {
	return c.broadcasts.GetBroadcasts(overhead, limit)
}

func (c *state) LocalState(join bool) []byte {
	data, err := c.local.bytes()
	if err != nil {
		panic(err)
	}
	return data
}

func (c *state) MergeRemoteState(buf []byte, join bool) {
	var other = new(local)
	b := bytes.NewBuffer(buf)
	decoder := gob.NewDecoder(b)
	err := decoder.Decode(&other)
	if err != nil {
		panic(err)
	}

	c.local.merge(other)
}

// 关键的状态变更，立即广播
func (c *state) pivotalStateChanged() {
	data, err := c.local.bytes(false)
	if err != nil {
		panic(err)
	}
	c.broadcast = data
	c.broadcasts.QueueBroadcast(c)
}
