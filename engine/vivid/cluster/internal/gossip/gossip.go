package gossip

import gossipv1 "github.com/kercylan98/minotaur/engine/vivid/cluster/internal/gossip/v1"

const (
	NodeStatusJoining     = gossipv1.GossipNodeStatus_GOSSIP_NODE_STATUS_JOINING
	NodeStatusAlive       = gossipv1.GossipNodeStatus_GOSSIP_NODE_STATUS_ALIVE
	NodeStatusLeaving     = gossipv1.GossipNodeStatus_GOSSIP_NODE_STATUS_LEAVING
	NodeStatusExiting     = gossipv1.GossipNodeStatus_GOSSIP_NODE_STATUS_EXITING
	NodeStatusExited      = gossipv1.GossipNodeStatus_GOSSIP_NODE_STATUS_EXITED
	NodeStatusRemoved     = gossipv1.GossipNodeStatus_GOSSIP_NODE_STATUS_REMOVED
	NodeStatusUnreachable = gossipv1.GossipNodeStatus_GOSSIP_NODE_STATUS_UNREACHABLE
	NodeStatusReachable   = gossipv1.GossipNodeStatus_GOSSIP_NODE_STATUS_REACHABLE
	NodeStatusDown        = gossipv1.GossipNodeStatus_GOSSIP_NODE_STATUS_DOWN
)

const (
	VectorClockOrderingConcurrent = gossipv1.VectorClockOrdering_VECTOR_CLOCK_ORDERING_CONCURRENT
	VectorClockOrderingAfter      = gossipv1.VectorClockOrdering_VECTOR_CLOCK_ORDERING_AFTER
	VectorClockOrderingBefore     = gossipv1.VectorClockOrdering_VECTOR_CLOCK_ORDERING_BEFORE
	VectorClockOrderingSame       = gossipv1.VectorClockOrdering_VECTOR_CLOCK_ORDERING_SAME
)

type (
	Node                          = gossipv1.Node
	NodeId                        = gossipv1.NodeId
	Gossip                        = gossipv1.Gossip
	VectorClock                   = gossipv1.VectorClock
	NodeStatus                    = gossipv1.GossipNodeStatus
	ActorLeaveClusterMessage      = gossipv1.GossipActorLeaveClusterMessage
	ActorInitClusterMessage       = gossipv1.GossipActorInitClusterMessage
	ActorTryJoinClusterMessage    = gossipv1.GossipActorTryJoinClusterMessage
	ActorCreateClusterMessage     = gossipv1.GossipActorCreateClusterMessage
	ActorTryJoinClusterAckMessage = gossipv1.GossipActorTryJoinClusterAckMessage
	ActorClusterConvergedMessage  = gossipv1.GossipActorClusterConvergedMessage
	Gossiped                      = gossipv1.Gossiped
	GossipedAckMessage            = gossipv1.GossipedAckMessage
	ActorPingPongMessage          = gossipv1.GossipActorPingPongMessage
	ActorClusterExitingMessage    = gossipv1.GossipActorClusterExitingMessage
	ActorClusterExitedMessage     = gossipv1.GossipActorClusterExitedMessage
)
