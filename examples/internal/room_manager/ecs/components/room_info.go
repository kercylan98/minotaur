package components

import (
	"github.com/kercylan98/minotaur/core/vivid"
	"github.com/kercylan98/minotaur/engine/ecs"
	"github.com/kercylan98/minotaur/toolkit/collection"
)

type RoomInfo struct {
	Id      ecs.Entity              // 房间 ID
	Ref     vivid.ActorRef          // 房间 ActorRef
	Players map[ecs.Entity]struct{} // 房间内的玩家
}

func (r *RoomInfo) Clone() *RoomInfo {
	c := *r
	c.Players = collection.CloneMap(r.Players)
	return &c
}
