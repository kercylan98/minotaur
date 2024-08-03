package cluster

import (
	"fmt"
	"google.golang.org/protobuf/proto"
)

func newActorSystemDelegate(system *ActorSystem) *actorSystemDelegate {
	return &actorSystemDelegate{
		system: system,
	}
}

type actorSystemDelegate struct {
	system *ActorSystem
}

func (d *actorSystemDelegate) NodeMeta(limit int) []byte {
	mdBytes, err := proto.Marshal(d.system.metadata)
	if err != nil {
		panic(err)
	}
	if len(mdBytes) > limit {
		panic(fmt.Errorf("metadata size %d exceeds limit %d", len(mdBytes), limit))
	}

	return mdBytes
}

func (d *actorSystemDelegate) NotifyMsg(bytes []byte) {

}

func (d *actorSystemDelegate) GetBroadcasts(overhead, limit int) [][]byte {
	return nil
}

func (d *actorSystemDelegate) LocalState(join bool) []byte {
	d.system.state.rw.RLock()
	defer d.system.state.rw.RUnlock()

	stateBytes, err := proto.Marshal(d.system.state.data)
	if err != nil {
		panic(err)
	}
	return stateBytes
}

func (d *actorSystemDelegate) MergeRemoteState(buf []byte, join bool) {
	// 应同步集群内 Actor 状态
}
