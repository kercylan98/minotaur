package room

import (
	"github.com/kercylan98/minotaur/core/ecs"
	"github.com/kercylan98/minotaur/core/vivid"
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
		world: ecs.NewWorld(),
	}
	return m
}

type Manager struct {
	world ecs.World

	roomInfoComponentId ecs.ComponentId
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
	r.world.QueryF(ecs.Equal(r.roomInfoComponentId), func(result *ecs.Result) {
		result.Each(func(entity ecs.Entity) bool {
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
