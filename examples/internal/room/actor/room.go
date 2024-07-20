package actor

import (
	"github.com/kercylan98/minotaur/core/vivid"
	ecs2 "github.com/kercylan98/minotaur/engine/ecs"
	"github.com/kercylan98/minotaur/examples/internal/room/components"
	"github.com/kercylan98/minotaur/examples/internal/room/messages"
	"github.com/kercylan98/minotaur/toolkit/log"
)

func NewRoom() *Room {
	r := &Room{
		world: ecs2.NewWorld(),
	}
	return r
}

type Room struct {
	world ecs2.World

	roomEntity ecs2.ComponentId
}

func (r *Room) OnReceive(ctx vivid.ActorContext) {
	switch m := ctx.Message().(type) {
	case vivid.OnLaunch:
		r.onLaunch(ctx)
	case *messages.JoinRoomAsk:
		r.onJoinRoom(ctx, m)
	}
}

func (r *Room) onLaunch(ctx vivid.ActorContext) {
	r.roomEntity = r.world.RegComponent(new(components.Room))
}

func (r *Room) onJoinRoom(ctx vivid.ActorContext, m *messages.JoinRoomAsk) {
	eid := r.world.Spawn(r.roomEntity)
	roomEntity := r.world.Get(eid, r.roomEntity).(*components.Room)
	roomEntity.Id = m.EntityId
	ctx.Reply(&messages.JoinRoomReply{Eid: uint64(eid)})

	log.Info("joinRoom", log.String("entityId", m.EntityId), log.String("ecsId", eid.String()))

	r.world.Query(ecs2.Equal(r.roomEntity)).Each(func(entity ecs2.Entity) bool {
		roomEntity := r.world.Get(entity, r.roomEntity).(*components.Room)
		log.Info("roomEntity", log.String("entityId", roomEntity.Id), log.String("ecsId", entity.String()))
		return true
	})
}
