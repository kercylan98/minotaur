package cluster

import (
	"context"
	"github.com/kercylan98/minotaur/engine/prc"
	"github.com/kercylan98/minotaur/engine/vivid"
	"github.com/kercylan98/minotaur/engine/vivid/cluster/internal/gossip"
)

type (
	actorSystemActorExitMessage struct {
		cancel context.CancelFunc
	}
)

func newActorSystemActor(system *ActorSystem, seedNodes []prc.PhysicalAddress) *actorSystemActor {
	return &actorSystemActor{
		system:    system,
		seedNodes: seedNodes,
	}
}

type actorSystemActor struct {
	system    *ActorSystem
	seedNodes []prc.PhysicalAddress // 种子节点
	gossipRef vivid.ActorRef        // gossip actor ref
	exited    context.CancelFunc    // 退出信号
	leaderRef vivid.ActorRef        // 集群当前的领导者引用
}

func (a *actorSystemActor) OnReceive(ctx vivid.ActorContext) {
	switch m := ctx.Message().(type) {
	case *vivid.OnLaunch:
		a.onLaunch(ctx)
	case vivid.ActorRef:
		a.onLeaderChanged(ctx, m)
	case *actorSystemActorExitMessage:
		a.onActorSystemActorExitMessage(ctx, m)
	case *gossip.ActorClusterExitingMessage:
		a.onGossipActorClusterExitingMessage(ctx, m)
	case *gossip.ActorClusterExitedMessage:
		a.onGossipActorClusterExitedMessage(ctx, m)
	}
}

func (a *actorSystemActor) onLaunch(ctx vivid.ActorContext) {
	a.gossipRef = ctx.ActorOfF(func() vivid.Actor {
		return gossip.NewGossiperActor(a.seedNodes)
	}, func(descriptor *vivid.ActorDescriptor) {
		descriptor.WithName("gossip")
	})
}

func (a *actorSystemActor) onActorSystemActorExitMessage(ctx vivid.ActorContext, m *actorSystemActorExitMessage) {
	a.exited = m.cancel
	ctx.Tell(a.gossipRef, &gossip.ActorLeaveClusterMessage{})
}

func (a *actorSystemActor) onGossipActorClusterExitingMessage(ctx vivid.ActorContext, m *gossip.ActorClusterExitingMessage) {
	// 善后工作
}

func (a *actorSystemActor) onGossipActorClusterExitedMessage(ctx vivid.ActorContext, m *gossip.ActorClusterExitedMessage) {
	a.exited()
}

func (a *actorSystemActor) onLeaderChanged(ctx vivid.ActorContext, m vivid.ActorRef) {
	a.leaderRef = m
}
