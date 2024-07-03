package cluster

import (
	"github.com/kercylan98/minotaur/core/vivid"
	"github.com/kercylan98/minotaur/toolkit/log"
	"github.com/kercylan98/minotaur/toolkit/random"
	"sort"
)

type Cluster interface {
	// System 返回当前集群节点的 vivid.ActorSystem
	System() *vivid.ActorSystem

	// RegKind 注册一个新的 actor.Kind
	RegKind(kind vivid.Kind, producer vivid.ActorProducer, options ...vivid.ActorOptionDefiner)

	// KindOf 通过 Kind 创建一个新的 Actor
	KindOf(kind vivid.Kind, parentPath ...string) vivid.ActorRef
}

// Invoke 获取 vivid.ActorSystem 对应的 Cluster 实例，如果不存在则返回 nil
func Invoke(actorSystem *vivid.ActorSystem) Cluster {
	nodeMu.RLock()
	defer nodeMu.RUnlock()
	return nodes[actorSystem]
}

var _ Cluster = (*cluster)(nil)

type cluster struct {
	*Node
}

func (c *cluster) System() *vivid.ActorSystem {
	return c.Node.support.System()
}

func (c *cluster) RegKind(kind vivid.Kind, producer vivid.ActorProducer, options ...vivid.ActorOptionDefiner) {
	c.Node.support.System().RegKind(kind, producer, options...)
}

func (c *cluster) KindOf(kind vivid.Kind, parentPath ...string) vivid.ActorRef {
	var pp string
	if len(parentPath) > 0 {
		pp = parentPath[0]
	} else {
		pp = "user"
	}
	members := c.memberlist.Members()
	sort.Slice(members, func(i, j int) bool {
		return random.Bool()
	})

	for _, node := range members {
		//if node.Name == c.name {
		//	continue
		//}

		c.state.local.mu.RLock()
		if _, ok := c.state.local.Kinds[node.Name][kind]; !ok {
			c.state.local.mu.RUnlock()
			continue
		}
		c.state.local.mu.RUnlock()

		md, err := parseMetadata(node.Meta)
		if err != nil {
			c.support.Logger().Error("cluster", log.String("status", "parse metadata"), log.Err(err))
			continue
		}

		return c.support.System().KindOf(kind, md.ActorSystemRootAddress.Join(pp).Ref())
	}

	c.support.Logger().Debug("cluster", log.String("status", "no available node"))
	return c.Node.support.GetDeadLetter().Ref()
}
