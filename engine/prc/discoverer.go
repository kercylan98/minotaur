package prc

import (
	"fmt"
	"github.com/hashicorp/memberlist"
	"sync"
	"time"
)

// NewDiscoverer 创建一个新的 Discoverer
func NewDiscoverer(rc *ResourceController, bindAddr string, bindPort int, configurator ...DiscovererConfigurator) *Discoverer {
	d := &Discoverer{
		rc:       rc,
		config:   newDiscovererConfiguration(),
		bindAddr: bindAddr,
		bindPort: bindPort,
		metadata: &DiscovererMetadata{
			RcPhysicalAddress: rc.GetPhysicalAddress(),
			UserMetadata:      make(map[string]string),
		},
	}

	for _, c := range configurator {
		c.Configure(d.config)
	}

	d.metadata.UserMetadata, d.config.userMetadata = d.config.userMetadata, nil

	return d
}

// Discoverer 是用于实现多个 ResourceController 之间的自主发现的数据结构
type Discoverer struct {
	config     *DiscovererConfiguration
	rc         *ResourceController
	state      *DiscovererState
	metadata   *DiscovererMetadata
	memberlist *memberlist.Memberlist
	stateLock  sync.RWMutex
	bindAddr   string
	bindPort   int
}

func (d *Discoverer) Discover() error {
	config := memberlist.DefaultLocalConfig()
	config.Name = d.config.name
	config.BindAddr, config.BindPort = d.bindAddr, d.bindPort
	config.AdvertisePort = config.BindPort
	d.metadata.LaunchAt = time.Now().UnixMilli()
	if d.config.advertiseAddr != "" {
		config.AdvertiseAddr = d.config.advertiseAddr
	}
	if d.config.advertisePort != 0 {
		config.AdvertisePort = d.config.advertisePort
	}

	config.Delegate = newDiscovererDelegate(d)

	list, err := memberlist.Create(config)
	if err != nil {
		return err
	}

	d.memberlist = list
	if len(d.config.joinNodes) == 0 {
		_, err = d.memberlist.Join([]string{fmt.Sprintf("%s:%d", d.bindAddr, d.bindPort)})
	} else {
		_, err = d.memberlist.Join(d.config.joinNodes)
	}

	return err
}

func (d *Discoverer) Leave() {
	var err error
	if err = d.memberlist.Leave(15 * time.Second); err != nil {
		panic(err)
	}
	if err = d.memberlist.Shutdown(); err != nil {
		panic(err)
	}
}

// GetNodes 获取集群中的所有节点
func (d *Discoverer) GetNodes() []*DiscoverNode {
	members := d.memberlist.Members()
	nodes := make([]*DiscoverNode, len(members))
	for i, member := range members {
		nodes[i] = newDiscoverNode(member)
	}
	return nodes
}
