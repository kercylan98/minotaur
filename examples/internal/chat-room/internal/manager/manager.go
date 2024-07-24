package manager

import (
	"github.com/kercylan98/minotaur/engine/prc"
	"github.com/kercylan98/minotaur/engine/vivid"
	"github.com/kercylan98/minotaur/examples/internal/chat-room/internal/room"
	"github.com/kercylan98/minotaur/toolkit/log"
)

type (
	FindOrCreateRoomMessage struct {
		RoomId string
	}

	FindOrCreateRoomMessageReply struct {
		Room vivid.ActorRef
	}
)

func New() *Manager {
	return &Manager{
		rooms:   map[string]vivid.ActorRef{},
		roomMap: make(map[prc.LogicalAddress]string),
	}
}

type Manager struct {
	rooms   map[string]vivid.ActorRef
	roomMap map[prc.LogicalAddress]string
}

func (m *Manager) OnReceive(ctx vivid.ActorContext) {
	switch msg := ctx.Message().(type) {
	case *FindOrCreateRoomMessage:
		m.onFindOrCreateRoom(ctx, msg)
	case *vivid.OnTerminated:
		m.onTerminated(ctx, msg)
	}
}

func (m *Manager) onFindOrCreateRoom(ctx vivid.ActorContext, msg *FindOrCreateRoomMessage) {
	r, exist := m.rooms[msg.RoomId]
	if !exist {
		r = ctx.ActorOfF(func() vivid.Actor {
			return room.New(msg.RoomId)
		})
		m.roomMap[r.LogicalAddress()] = msg.RoomId
		m.rooms[msg.RoomId] = r
		ctx.System().Logger().Info("RoomManager", log.String("status", "create"), log.String("room_id", msg.RoomId))
	}

	ctx.Reply(&FindOrCreateRoomMessageReply{Room: r})
}

func (m *Manager) onTerminated(ctx vivid.ActorContext, msg *vivid.OnTerminated) {
	roomId, exist := m.roomMap[msg.TerminatedActor.LogicalAddress()]
	if exist {
		delete(m.roomMap, msg.TerminatedActor.LogicalAddress())
		delete(m.rooms, roomId)
		ctx.System().Logger().Info("RoomManager", log.String("status", "destroy"), log.String("room_id", roomId))
	}
}
