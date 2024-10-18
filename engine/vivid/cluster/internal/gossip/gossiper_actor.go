package gossip

import (
	"github.com/kercylan98/minotaur/engine/future"
	"github.com/kercylan98/minotaur/engine/prc"
	"github.com/kercylan98/minotaur/engine/vivid"
	"github.com/kercylan98/minotaur/toolkit/collection"
	"github.com/kercylan98/minotaur/toolkit/log"
	"log/slog"
	"time"
)

func NewGossiperActor(seedNodes []prc.PhysicalAddress) *GossiperActor {
	return &GossiperActor{
		seedNodes: seedNodes,
	}
}

type GossiperActor struct {
	seedNodes    []prc.PhysicalAddress // 种子节点
	seedNodeRefs []vivid.ActorRef      // 种子节点 Actor 引用
	logger       *slog.Logger          // 日志记录器
	state        *State                // Gossip 状态
	leader       *Node                 // 集群当前确定的领导者
}

func (g *GossiperActor) OnReceive(ctx vivid.ActorContext) {
	switch m := ctx.Message().(type) {
	case *vivid.OnLaunch:
		g.onLaunch(ctx)
	case *GossipActorInitClusterMessage:
		g.onGossipActorInitClusterMessage(ctx, m)
	case *GossipActorTryJoinClusterMessage:
		g.onGossipActorTryJoinClusterMessage(ctx, m)
	case *GossipActorCreateClusterMessage:
		g.onGossipActorCreateClusterMessage(ctx)
	case *GossipActorTryJoinClusterAckMessage:
		g.onGossipActorTryJoinClusterAckMessage(ctx, m)
	case *GossipActorClusterConvergedMessage:
		g.onGossipActorClusterConvergedMessage(ctx)
	case *Gossiped:
		g.onGossiped(ctx, m)
	case *GossipedAckMessage:
		g.onGossipAckMessage(ctx, m)
	case *GossipActorLeaveClusterMessage:
		g.onGossipActorLeaveClusterMessage(ctx)
	}
}

func (g *GossiperActor) onLaunch(ctx vivid.ActorContext) {
	// 日志初始化
	g.logger = ctx.System().Logger().With(log.String("system", ctx.System().PhysicalAddress()))

	// 初始化自身状态
	g.state = newState(ctx, g)
	g.seedNodeRefs = make([]vivid.ActorRef, len(g.seedNodes))
	for i, seedNode := range g.seedNodes {
		g.seedNodeRefs[i] = vivid.NewActorRef(seedNode, ctx.LogicalAddress())
	}

	// 尝试加入集群
	//ctx.AfterTask("onLaunch.initCluster", time.Second*3, func(ctx vivid.ActorContext) {
	//	ctx.Tell(ctx.Ref(), &GossipActorInitClusterMessage{RetryIntervalDuration: int64(3 * time.Second)})
	//})
	ctx.Tell(ctx.Ref(), &GossipActorInitClusterMessage{RetryIntervalDuration: int64(3 * time.Second)})
}

func (g *GossiperActor) onGossipActorInitClusterMessage(ctx vivid.ActorContext, m *GossipActorInitClusterMessage) {
	type Entry struct {
		Future future.Future[vivid.Message]
		Ref    vivid.ActorRef
	}

	var futures = make([]Entry, 0, len(g.seedNodeRefs)-1)
	for _, ref := range g.seedNodeRefs {
		if ref.PhysicalAddress == ctx.PhysicalAddress() {
			continue // 排除自身对自身尝试的加入
		}

		futures = append(futures, Entry{Future: ctx.FutureAsk(ref, &GossipActorTryJoinClusterMessage{Node: g.state.node}, time.Second*99999), Ref: ref})
	}

	// 避免互相等待对方直到超时，协程内需要严格保证竞态问题
	go func() {
		var fail bool
		defer func() {
			if fail {
				ctx.AfterTask("onGossipActorInitClusterMessage.retry", time.Duration(m.RetryIntervalDuration), func(ctx vivid.ActorContext) {
					ctx.Tell(ctx.Ref(), m)
				})
			}
		}()

		var ackList []*GossipActorTryJoinClusterAckMessage
		for _, entry := range futures {
			g.logger.Debug("cluster", log.String("event", "try join cluster"), log.String("ref", entry.Ref.URL().String()))
			ack, err := entry.Future.Result()
			if err != nil {
				g.logger.Error("cluster", log.String("event", "try join cluster"), log.String("ref", entry.Ref.URL().String()), log.Err(err))
				fail = true
				break
			}
			ackList = append(ackList, ack.(*GossipActorTryJoinClusterAckMessage))
		}

		for _, ack := range ackList {
			if ack.Refuse {
				continue
			}
			ctx.Tell(ctx.Ref(), ack)
			return
		}

		if !fail && collection.IsFirst(g.seedNodes, ctx.PhysicalAddress()) {
			ctx.Tell(ctx.Ref(), &GossipActorCreateClusterMessage{})
			return
		}

		// 如果所有响应都是拒绝，重新尝试
		fail = true
	}()
}

func (g *GossiperActor) onGossipActorCreateClusterMessage(ctx vivid.ActorContext) {
	g.logger.Info("cluster", log.String("status", "create"))
	g.state.AddMember(g.state.node)
	ctx.Tell(ctx.Ref(), &GossipActorClusterConvergedMessage{})
}

func (g *GossiperActor) onGossipActorTryJoinClusterMessage(ctx vivid.ActorContext, m *GossipActorTryJoinClusterMessage) {
	g.logger.Debug("cluster", log.String("event", "received join"), log.String("node", m.Node.Id.Ref.URL().String()))
	var ack = &GossipActorTryJoinClusterAckMessage{Handler: ctx.Ref()}

	switch g.state.node.Status {
	case GossipNodeStatus_GNS_Joining:
		ack.Refuse = true
	default:
		g.state.AddMember(m.Node)
		ack.Gossiped = &Gossiped{
			Gossip:          g.state.gossip,
			GossiperVersion: g.state.node.Vc,
		}
	}

	ctx.Reply(ack)
}

func (g *GossiperActor) onGossipActorTryJoinClusterAckMessage(ctx vivid.ActorContext, m *GossipActorTryJoinClusterAckMessage) {
	g.logger.Debug("cluster", log.String("event", "received join ack"), log.String("node", m.Handler.URL().String()))
	g.state.MergeGossip(m.Gossiped)
	g.state.MarkSeen(g.state.node.Id)
	g.state.GossipUpdate()
}

func (g *GossiperActor) onGossiped(ctx vivid.ActorContext, m *Gossiped) {
	g.state.MergeGossip(m)
	g.state.MarkSeen(g.state.node.Id)

	ctx.Reply(&GossipedAckMessage{Gossiped: &Gossiped{Gossip: g.state.gossip, GossiperVersion: g.state.node.Vc}})

	// 继续传播
	g.state.GossipUpdate()
}

func (g *GossiperActor) onGossipAckMessage(ctx vivid.ActorContext, m *GossipedAckMessage) {
	g.state.MergeGossip(m.Gossiped)
	g.state.MarkSeen(g.state.node.Id)

	if g.state.node.Status == GossipNodeStatus_GNS_Exit {
		g.logger.Info("cluster", log.String("status", "exit"))
		ctx.Terminate(ctx.Ref(), true)
		return
	}

	// 检查是否已收敛
	var seenNum int
	for _, member := range g.state.gossip.Members {
		switch member.Status {
		case GossipNodeStatus_GNS_Removed:
			continue
		}
		seenNum++
	}
	if seenNum == len(g.state.gossip.Seen) && g.state.node.Vc.CompareTo(m.Gossiped.GossiperVersion) == VectorClockOrdering_VCO_Same {
		ctx.Tell(ctx.Ref(), &GossipActorClusterConvergedMessage{})
	}
}

func (g *GossiperActor) onGossipActorClusterConvergedMessage(ctx vivid.ActorContext) {
	g.logger.Info("cluster", log.String("status", "converged"))
	for _, member := range g.state.gossip.Members {
		g.logger.Debug("cluster", log.String("member", member.Id.Ref.URL().String()), log.String("status", member.Status.String()))
	}

	// 确定领导者
	before := g.leader
	g.leader = g.state.CalcLeaderNode()
	if before != nil {
		if g.leader.Id.Ref.PhysicalAddress != before.Id.Ref.PhysicalAddress {
			g.logger.Info("cluster", log.String("info", "LeaderChange"), log.String("leader", g.leader.Id.Ref.URL().String()), log.String("before", before.Id.Ref.URL().String()))
		}
	} else {
		g.logger.Info("cluster", log.String("info", "LeaderInit"), log.String("leader", g.leader.Id.Ref.URL().String()))
	}

	// 将 Joining 节点设置为活跃
	if g.leader.Id.Ref.PhysicalAddress == ctx.PhysicalAddress() {
		var changed bool
		for i, member := range g.state.gossip.Members {
			switch member.Status {
			case GossipNodeStatus_GNS_Joining:
				changed = true
				member.Status = GossipNodeStatus_GNS_Alive
				g.logger.Info("cluster", log.String("node", member.Id.Ref.URL().String()), log.String("status", "joining -> alive"))
			case GossipNodeStatus_GNS_Leaving:
				// 确认节点离开行为已经执行完毕后标记为 Exit
				changed = true
				member.Status = GossipNodeStatus_GNS_Exit
				g.logger.Info("cluster", log.String("node", member.Id.Ref.URL().String()), log.String("status", "leaving -> exit"))
			case GossipNodeStatus_GNS_Exit:
				changed = true
				g.state.gossip.Members = append(g.state.gossip.Members[:i], g.state.gossip.Members[i+1:]...)
				g.logger.Info("cluster", log.String("node", member.Id.Ref.URL().String()), log.String("status", "exit, remove from gossip"))
			}
		}
		if changed {
			// 清空 Seen 列表，并标记自己为已看到
			g.state.gossip.Seen = []*NodeId{g.state.node.Id}

			// 自增领导者节点的 VectorClock，表示集群状态更新
			g.state.Increment()

			// 传播新的 Gossip 状态
			g.state.GossipUpdate()
		}
	}
}

func (g *GossiperActor) onGossipActorLeaveClusterMessage(ctx vivid.ActorContext) {
	g.logger.Info("cluster", log.String("status", "leave"))

	g.state.node.Status = GossipNodeStatus_GNS_Leaving
	g.state.gossip.Seen = []*NodeId{g.state.node.Id}
	g.state.Increment()
	g.state.GossipUpdate()
}
