package prc

import (
	"github.com/hashicorp/memberlist"
	"google.golang.org/protobuf/proto"
	"sync"
)

func newDiscoverNode(node *memberlist.Node) *DiscoverNode {
	return &DiscoverNode{node: node}
}

type DiscoverNode struct {
	node         *memberlist.Node
	metadataOnce sync.Once
	metadata     *DiscovererMetadata
}

func (dn *DiscoverNode) Metadata() *DiscovererMetadata {
	dn.metadataOnce.Do(func() {
		dn.metadata = new(DiscovererMetadata)
		_ = proto.Unmarshal(dn.node.Meta, dn.metadata)
	})
	return dn.metadata
}

func (dn *DiscoverNode) GetUserMetadata(key string) string {
	return dn.Metadata().UserMetadata[key]
}

func (dn *DiscoverNode) HasUserMetadata(key string) bool {
	_, exist := dn.Metadata().UserMetadata[key]
	return exist
}

func (dn *DiscoverNode) GetPhysicalAddress() PhysicalAddress {
	return dn.Metadata().RcPhysicalAddress
}
