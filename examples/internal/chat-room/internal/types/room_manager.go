package types

import (
	"github.com/kercylan98/minotaur/engine/vivid"
)

type (
	GetRoomMessage struct{ RoomId }
)

func NewRoomManager() *RoomManager {
	return &RoomManager{}
}

type RoomManager struct {
	rooms map[RoomId]vivid.ActorRef
}

func (r *RoomManager) OnReceive(ctx vivid.ActorContext) {
	switch m := ctx.Message().(type) {
	case *vivid.OnLaunch:
		r.onLaunch(ctx)
	case GetRoomMessage:
		r.onFindOrCreateRoom(ctx, m)
	}
}

func (r *RoomManager) onLaunch(ctx vivid.ActorContext) {
	r.rooms = make(map[RoomId]vivid.ActorRef)
}

func (r *RoomManager) onFindOrCreateRoom(ctx vivid.ActorContext, m GetRoomMessage) {
	room, exist := r.rooms[m.RoomId]
	if !exist {
		room = ctx.ActorOfF(func() vivid.Actor {
			return newRoom(m.RoomId)
		})

		r.rooms[m.RoomId] = room
	}

	ctx.Reply(room)
}
