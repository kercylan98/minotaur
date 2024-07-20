package prc

import (
	"fmt"
	"google.golang.org/protobuf/proto"
)

func newDiscovererDelegate(discoverer *Discoverer) *discovererDelegate {
	return &discovererDelegate{
		discoverer: discoverer,
	}
}

type discovererDelegate struct {
	discoverer *Discoverer
}

func (d *discovererDelegate) NodeMeta(limit int) []byte {
	mdBytes, err := proto.Marshal(d.discoverer.metadata)
	if err != nil {
		panic(err)
	}
	if len(mdBytes) > limit {
		panic(fmt.Errorf("metadata size %d exceeds limit %d", len(mdBytes), limit))
	}

	return mdBytes
}

func (d *discovererDelegate) NotifyMsg(bytes []byte) {

}

func (d *discovererDelegate) GetBroadcasts(overhead, limit int) [][]byte {
	return nil
}

func (d *discovererDelegate) LocalState(join bool) []byte {
	d.discoverer.stateLock.RLock()
	defer d.discoverer.stateLock.RUnlock()

	stateBytes, err := proto.Marshal(d.discoverer.state)
	if err != nil {
		panic(err)
	}
	return stateBytes
}

func (d *discovererDelegate) MergeRemoteState(buf []byte, join bool) {

}
