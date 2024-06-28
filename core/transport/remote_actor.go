package transport

import (
	core2 "github.com/kercylan98/minotaur/core"
)

var _ core2.Process = &remoteActor{}

func newRemoteActor(network *Network, address core2.Address) *remoteActor {
	return &remoteActor{
		network: network,
		address: address,
	}
}

type remoteActor struct {
	network *Network
	address core2.Address
}

func (r *remoteActor) GetAddress() core2.Address {
	return r.address
}

func (r *remoteActor) SendUserMessage(sender *core2.ProcessRef, message core2.Message) {
	r.network.send(sender, r.address, message)
}

func (r *remoteActor) SendSystemMessage(sender *core2.ProcessRef, message core2.Message) {
	r.network.send(sender, r.address, message)
}

func (r *remoteActor) Terminate(ref *core2.ProcessRef) {
	if ref.Address() == r.address {
		return // 远程解析不需要销毁
	}

	r.network.support.System().Terminate(ref)
}
