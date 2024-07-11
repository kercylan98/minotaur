package room

import (
	"errors"
	"github.com/kercylan98/minotaur/core/ecs"
	"github.com/kercylan98/minotaur/core/vivid"
	"github.com/kercylan98/minotaur/examples/internal/room_manager/ecs/components"
)

type (
	AskJoinRoom      struct{ UserId string }
	AskJoinRoomReply struct{ PlayerId ecs.Entity }
)

func NewRoom() *Room {
	r := &Room{
		world: ecs.NewWorld(),
	}
	return r
}

type Room struct {
	world ecs.World

	playerComponentId ecs.ComponentId
	seatComponentId   ecs.ComponentId
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
	r.world.Query(ecs.Equal(r.playerComponentId)).Each(func(entity ecs.Entity) bool {
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
