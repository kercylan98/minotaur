package cluster

import (
	"fmt"
	"github.com/hashicorp/memberlist"
	"github.com/kercylan98/minotaur/toolkit/balancer"
	"github.com/kercylan98/minotaur/toolkit/collection"
	"log"
	"time"
)

func NewNode(options ...Option) (*Node, error) {
	opts := new(Options).apply(options...)
	config := memberlist.DefaultLANConfig()
	var node = &Node{
		config: config,
		delegate: &Delegate{
			nodes:      make(map[string]*NodeInfo),
			clusterTab: make(map[string]*balancer.ConsistentHashWeight[string, *NodeInfo]),
			metadata: &Metadata{
				ClusterName: opts.ClusterName,
				Region:      opts.Region,
				Zone:        opts.Zone,
				ShardId:     opts.ShardId,
				Weight:      opts.Weight,
			},
		},
		defaultJoinAddresses: opts.DefaultJoinAddresses,
	}

	if opts.Logger != nil {
		config.Logger = log.New(&logger{opts.Logger}, "", log.LstdFlags)
	}
	config.BindAddr = opts.BindAddr
	config.BindPort = int(opts.BindPort)
	config.AdvertiseAddr = opts.AdvertiseAddr
	config.AdvertisePort = int(opts.AdvertisePort)
	config.Name = node.delegate.metadata.Name(config.AdvertiseAddr, config.AdvertisePort)

	node.delegate.Node = node
	node.delegate.broadcast = &memberlist.TransmitLimitedQueue{
		NumNodes: func() int {
			return node.delegate.list.NumMembers()
		},
		RetransmitMult: 1,
	}

	config.Delegate = node.delegate
	config.Events = node.delegate

	var list, err = memberlist.Create(config)
	if err != nil {
		return nil, err
	}

	node.delegate.list = list
	node.delegate.messages = make(chan []byte)

	return node, nil
}

// Node 集群节点
type Node struct {
	config               *memberlist.Config // 配置
	delegate             *Delegate          // 代理
	defaultJoinAddresses []string           // 默认加入地址
}

// GetClusterName 获取集群名称
func (n *Node) GetClusterName() string {
	if n == nil {
		return ""
	}
	return n.delegate.metadata.ClusterName
}

// GetHost 获取主机地址
func (n *Node) GetHost() string {
	if n == nil {
		return ""
	}
	return n.config.AdvertiseAddr
}

// GetPort 获取端口
func (n *Node) GetPort() uint16 {
	if n == nil {
		return 0
	}
	return uint16(n.config.AdvertisePort)
}

// Read 读取数据
func (n *Node) Read() <-chan []byte {
	return n.delegate.messages
}

// SendToCluster 发送数据到集群
func (n *Node) SendToCluster(cluster string, data []byte) error {
	n.delegate.tabMutex.RLock()
	defer n.delegate.tabMutex.RUnlock()
	tabs, exist := n.delegate.clusterTab[cluster]
	if !exist {
		return fmt.Errorf("cluster %s not found", cluster)
	}
	info, err := tabs.Select()
	if err != nil {
		return err
	}
	return n.delegate.list.SendReliable(info.Node, data)
}

// SendToNode 发送数据到指定节点
func (n *Node) SendToNode(address string, data []byte) error {
	n.delegate.tabMutex.RLock()
	defer n.delegate.tabMutex.RUnlock()
	info, exist := n.delegate.nodes[address]
	if !exist {
		return fmt.Errorf("node %s not found", address)
	}
	return n.delegate.list.SendReliable(info.Node, data)
}

// Broadcast 广播数据
func (n *Node) Broadcast(data memberlist.Broadcast) {
	n.delegate.broadcast.QueueBroadcast(data)
}

// Join 加入集群，addresses 为集群中的任意多个节点地址，当 addresses 为空时，将会以该节点作为集群的第一个节点
//   - addresses 无需包含所有节点地址，加入集群后会逐渐扩散节点信息，最终所有节点都会知道所有节点的地址
func (n *Node) Join(addresses ...string) error {
	addresses = append(addresses, n.defaultJoinAddresses...)
	addresses = collection.DeduplicateSlice(addresses)
	_, err := n.delegate.list.Join(addresses)
	return err
}

// Leave 离开集群
func (n *Node) Leave() error {
	return n.delegate.list.Leave(time.Second)
}

// Shutdown 关闭节点
func (n *Node) Shutdown() error {
	if err := n.Leave(); err != nil {
		return err
	}
	return n.delegate.list.Shutdown()
}
