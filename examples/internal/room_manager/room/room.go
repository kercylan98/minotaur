package room

import (
	"errors"
	"github.com/kercylan98/minotaur/core/vivid"
	ecs2 "github.com/kercylan98/minotaur/engine/ecs"
	"github.com/kercylan98/minotaur/examples/internal/room_manager/ecs/components"
)

type (
	AskJoinRoom      struct{ UserId string }
	AskJoinRoomReply struct{ PlayerId ecs2.Entity }
)

func NewRoom() *Room {
	r := &Room{
		world: ecs2.NewWorld(),
	}
	return r
}

type Room struct {
	world ecs2.World

	playerComponentId ecs2.ComponentId
	seatComponentId   ecs2.ComponentId
}

func (r *Room) OnReceive(ctx vivid.ActorContext) {
	switch m := ctx.Message().(type) {
	case vivid.OnLaunch:
		r.onLaunch(ctx)
	case AskJoinRoom:
		r.onAskJoinRoom(ctx, m)
	}
}

func (r *Room) onLaunch(ctx vivid.ActorContext) {
	r.playerComponentId = r.world.RegComponent(new(components.Player))
	r.seatComponentId = r.world.RegComponent(new(components.Seat))
}

func (r *Room) onAskJoinRoom(ctx vivid.ActorContext, m AskJoinRoom) {
	var exists bool
	r.world.Query(ecs2.Equal(r.playerComponentId)).Each(func(entity ecs2.Entity) bool {
		player := r.world.Get(entity, r.playerComponentId).(components.Player)
		if exists = player.Id == m.UserId; exists {
			return false
		}
		return true
	})

	if exists {
		ctx.Reply(errors.New("already in room"))
		return
	}

	player := r.world.Spawn(r.playerComponentId)
	ctx.Reply(AskJoinRoomReply{PlayerId: player})
}
