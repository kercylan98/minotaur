package cluster

import (
	"fmt"
	"github.com/hashicorp/memberlist"
	"github.com/kercylan98/minotaur/core/vivid"
	"github.com/kercylan98/minotaur/toolkit/log"
	"math"
	"sync"
	"time"
)

var nodes = make(map[*vivid.ActorSystem]Cluster)
var nodeMu sync.RWMutex

var _ vivid.Module = (*Node)(nil)
var _ vivid.PriorityModule = (*Node)(nil)
var _ vivid.ShutdownModule = (*Node)(nil)
var _ vivid.Module = (*cluster)(nil)

func NewNode(name, bindAddr string, bindPort int, nodes ...string) *Node {
	n := &Node{
		name:     name,
		bindAddr: bindAddr,
		bindPort: bindPort,
		nodes:    nodes,
	}
	n.cluster = &cluster{n}
	return n
}

type Node struct {
	*cluster
	support    *vivid.ModuleSupport
	state      *state
	memberlist *memberlist.Memberlist
	nodes      []string
	name       string
	bindAddr   string
	bindPort   int
}

func (n *Node) Priority() int {
	return math.MinInt + 1 // 优先于内置网络模块
}

func (n *Node) OnRegKind(kind vivid.Kind) {
	n.state.mu.Lock()
	defer n.state.mu.Unlock()
	kinds, exists := n.state.local.Kinds[n.name]
	if !exists {
		kinds = make(map[vivid.Kind]struct{})
		if n.state.local.Kinds == nil {
			n.state.local.Kinds = make(map[string]map[vivid.Kind]struct{})
		}
		n.state.local.Kinds[n.name] = kinds
	}
	kinds[kind] = struct{}{}
	n.state.pivotalStateChanged()
}

func (n *Node) OnLoad(support *vivid.ModuleSupport, hasTransportModule bool) {
	if !hasTransportModule {
		panic(fmt.Errorf("cluster.Node module requires transport module, can use default transport: transport.NewNetwork"))
	}
	n.support = support
	n.state = newState(n, &metadata{
		ActorSystemAddr:        n.support.Address().Host(),
		ActorSystemPort:        n.support.Address().Port(),
		ActorSystemRootAddress: n.support.Address(),
	})

	n.support.Logger().Info("cluster", log.String("status", "metadata initialized"), log.Any("metadata", n.state.metadata))

	config := memberlist.DefaultLocalConfig()
	config.Name = n.name
	config.BindAddr, config.BindPort = n.bindAddr, n.bindPort
	config.AdvertisePort = config.BindPort
	config.Delegate = n.state
	config.Logger = newMemberlistLogger(n.support.Logger)

	// join cluster
	list, err := memberlist.Create(config)
	if err != nil {
		panic(err)
	}

	n.memberlist = list
	if len(n.nodes) == 0 {
		_, err = n.memberlist.Join([]string{fmt.Sprintf("%s:%d", n.bindAddr, n.bindPort)})
	} else {
		_, err = n.memberlist.Join(n.nodes)
	}

	if err != nil {
		panic(err)
	}

	n.support.Logger().Info("cluster", log.String("status", "cluster connected"))
	for _, node := range n.memberlist.Members() {
		n.support.Logger().Info("cluster", log.String("status", "member"), log.String("name", node.Name), log.String("addr", node.Addr.String()), log.Int("port", node.Port))
	}

	nodeMu.Lock()
	defer nodeMu.Unlock()
	nodes[n.support.System()] = n
}

func (n *Node) OnShutdown() {
	var err error
	if err = n.memberlist.Leave(15 * time.Second); err != nil {
		n.support.Logger().Error("cluster", log.String("status", "leave failed"), log.Err(err))
	}
	if err = n.memberlist.Shutdown(); err != nil {
		n.support.Logger().Error("cluster", log.String("status", "shutdown failed"), log.Err(err))
	}
	n.support.Logger().Info("cluster", log.String("status", "shutdown"))
}
