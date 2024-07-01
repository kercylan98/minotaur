package transport

import (
	"github.com/kercylan98/minotaur/core"
)

var _ core.Process = &remoteActor{}

func newRemoteActor(network *Network, address core.Address) *remoteActor {
	return &remoteActor{
		network: network,
		address: address,
	}
}

type remoteActor struct {
	network *Network
	address core.Address
}

func (r *remoteActor) GetAddress() core.Address {
	return r.address
}

func (r *remoteActor) SendUserMessage(sender *core.ProcessRef, message core.Message) {
	r.network.send(sender, r.address, message, false)
}

func (r *remoteActor) SendSystemMessage(sender *core.ProcessRef, message core.Message) {
	r.network.send(sender, r.address, message, true)
}

func (r *remoteActor) Terminate(ref *core.ProcessRef) {
	if ref.Address() == r.address {
		return // 远程解析不需要销毁
	}

	r.network.support.System().Terminate(ref)
}
