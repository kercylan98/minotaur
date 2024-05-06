package registry

import (
	"encoding/json"
	"github.com/kercylan98/minotaur/rpc"
	"github.com/kercylan98/minotaur/toolkit"
	"github.com/nats-io/nats.go"
)

type NatsCaller struct {
	msg    *nats.Msg
	Route  []rpc.Route
	Packet json.RawMessage
}

func (n *NatsCaller) Marshal(route []rpc.Route, packet []byte) []byte {
	n.Route = route
	n.Packet = packet
	return toolkit.MarshalJSON(n)
}

func (n *NatsCaller) Unmarshal(msg *nats.Msg) *NatsCaller {
	n.msg = msg
	toolkit.UnmarshalJSON(msg.Data, n)
	return n
}

func (n *NatsCaller) GetRoute() []rpc.Route {
	return n.Route
}

func (n *NatsCaller) GetPacket() []byte {
	return n.Packet
}

func (n *NatsCaller) Respond(packet []byte) error {
	return n.msg.Respond(packet)
}
