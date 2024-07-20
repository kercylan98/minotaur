package room

import (
	"github.com/kercylan98/minotaur/core/vivid"
	ecs2 "github.com/kercylan98/minotaur/engine/ecs"
	"github.com/kercylan98/minotaur/examples/internal/room_manager/ecs/components"
)

type (
	AskRoomInfos       struct{}
	AskRoomInfosReply  []*components.RoomInfo
	AskCreateRoom      struct{}
	AskCreateRoomReply vivid.ActorRef
)

func NewManager() *Manager {
	m := &Manager{
		world: ecs2.NewWorld(),
	}
	return m
}

type Manager struct {
	world ecs2.World

	roomInfoComponentId ecs2.ComponentId
}

func (r *Manager) OnReceive(ctx vivid.ActorContext) {
	switch ctx.Message().(type) {
	case vivid.OnLaunch:
		r.onLaunch(ctx)
	case AskRoomInfos:
		r.onAskRoomInfos(ctx)
	case AskCreateRoom:
		r.onAskCreateRoom(ctx)
	}
}

func (r *Manager) onLaunch(ctx vivid.ActorContext) {
	r.roomInfoComponentId = r.world.RegComponent(new(components.RoomInfo))
}

func (r *Manager) onAskRoomInfos(ctx vivid.ActorContext) {
	rooms := make([]*components.RoomInfo, 0)
	r.world.QueryF(ecs2.Equal(r.roomInfoComponentId), func(result *ecs2.Result) {
		result.Each(func(entity ecs2.Entity) bool {
			room := result.Get(entity, r.roomInfoComponentId).(*components.RoomInfo)
			rooms = append(rooms, room.Clone())
			return true

		})
	})
	ctx.Reply(AskRoomInfosReply(rooms))
}

func (r *Manager) onAskCreateRoom(ctx vivid.ActorContext) {
	ref := ctx.ActorOf(func() vivid.Actor {
		return NewRoom()
	}, func(options *vivid.ActorOptions) {
		options.WithNamePrefix("room")
	})

	ctx.Reply(AskCreateRoomReply(ref))
}
