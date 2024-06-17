package cluster

import (
	"github.com/hashicorp/memberlist"
	"github.com/kercylan98/minotaur/toolkit"
	"github.com/kercylan98/minotaur/toolkit/balancer"
	"sync"
)

type Delegate struct {
	*Node

	list       *memberlist.Memberlist                                       // 成员列表
	metadata   *Metadata                                                    // 节点元数据
	messages   chan []byte                                                  // 消息通道
	broadcast  *memberlist.TransmitLimitedQueue                             // 广播队列
	clusterTab map[string]*balancer.ConsistentHashWeight[string, *NodeInfo] // 一致性哈希表
	nodes      map[string]*NodeInfo                                         // 节点表
	tabMutex   sync.RWMutex                                                 // 一致性哈希表锁
}

func (d *Delegate) NotifyJoin(node *memberlist.Node) {
	var info = &NodeInfo{
		Node:     node,
		metadata: new(Metadata),
	}
	toolkit.UnmarshalJSON(node.Meta, info.metadata)

	d.tabMutex.Lock()
	defer d.tabMutex.Unlock()

	tabs, exist := d.clusterTab[info.metadata.ClusterName]
	if !exist {
		tabs = balancer.NewConsistentHashWeight[string, *NodeInfo](10)
		d.clusterTab[info.metadata.ClusterName] = tabs
	}
	tabs.Add(info)

	d.nodes[info.GetId()] = info
}

func (d *Delegate) NotifyLeave(node *memberlist.Node) {
	var info = &NodeInfo{
		Node:     node,
		metadata: new(Metadata),
	}
	toolkit.UnmarshalJSON(node.Meta, info.metadata)

	d.tabMutex.Lock()
	defer d.tabMutex.Unlock()

	tabs, exist := d.clusterTab[info.metadata.ClusterName]
	if !exist {
		return
	}

	tabs.Remove(info.GetId())
	if len(tabs.GetInstances()) == 0 {
		delete(d.clusterTab, info.metadata.ClusterName)
	}

	delete(d.nodes, info.GetId())
}

func (d *Delegate) NotifyUpdate(node *memberlist.Node) {
	var info = &NodeInfo{
		Node:     node,
		metadata: new(Metadata),
	}
	toolkit.UnmarshalJSON(node.Meta, info.metadata)

	d.tabMutex.Lock()
	defer d.tabMutex.Unlock()

	tabs, exist := d.clusterTab[info.metadata.ClusterName]
	if !exist {
		tabs = balancer.NewConsistentHashWeight[string, *NodeInfo](10)
		d.clusterTab[info.metadata.ClusterName] = tabs
	} else {
		tabs.Update(info)
		return
	}

	tabs.Add(info)

	d.nodes[info.GetId()] = info
}

func (d *Delegate) NodeMeta(limit int) []byte {
	return d.metadata.NodeMeta(limit)
}

func (d *Delegate) NotifyMsg(bytes []byte) {
	d.messages <- bytes
}

func (d *Delegate) GetBroadcasts(overhead, limit int) [][]byte {
	return d.broadcast.GetBroadcasts(overhead, limit)
}

func (d *Delegate) LocalState(join bool) []byte {
	return []byte{}
}

func (d *Delegate) MergeRemoteState(buf []byte, join bool) {

}
