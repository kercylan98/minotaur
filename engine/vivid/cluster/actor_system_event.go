package cluster

import (
	"github.com/hashicorp/memberlist"
	"github.com/kercylan98/minotaur/engine/vivid"
	"github.com/kercylan98/minotaur/engine/vivid/cluster/internal/cm"
	"google.golang.org/protobuf/proto"
)

func newActorSystemEvent(system *ActorSystem) *actorSystemEvent {
	return &actorSystemEvent{
		system: system,
		ch:     make(chan *vivid.SubscriptionContactEvent),
	}
}

type actorSystemEvent struct {
	system *ActorSystem
	ch     chan *vivid.SubscriptionContactEvent
}

func (e *actorSystemEvent) ChangeNotify() <-chan *vivid.SubscriptionContactEvent {
	return e.ch
}

func (e *actorSystemEvent) NotifyJoin(node *memberlist.Node) {
	metadata := &cm.Metadata{}
	if err := proto.Unmarshal(node.Meta, metadata); err != nil {
		panic(err)
	}

	e.ch <- &vivid.SubscriptionContactEvent{Address: metadata.ProcessId.PhysicalAddress}
}

func (e *actorSystemEvent) NotifyLeave(node *memberlist.Node) {
	metadata := &cm.Metadata{}
	if err := proto.Unmarshal(node.Meta, metadata); err != nil {
		panic(err)
	}

	e.ch <- &vivid.SubscriptionContactEvent{Address: metadata.ProcessId.PhysicalAddress, Stop: true}
}

func (e *actorSystemEvent) NotifyUpdate(node *memberlist.Node) {

}
