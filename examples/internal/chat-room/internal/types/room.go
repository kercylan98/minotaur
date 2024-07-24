package types

import (
	"fmt"
	"github.com/gofiber/contrib/websocket"
	"github.com/kercylan98/minotaur/engine/stream"
	"github.com/kercylan98/minotaur/engine/vivid"
)

type (
	JoinRoomMessage struct {
		UserId
		User vivid.ActorRef
	}
	Chat struct {
		UserId
		*stream.Packet
	}
)

type RoomId string
type UserId string

func newRoom(roomId RoomId) *Room {
	return &Room{
		roomId: roomId,
	}
}

type Room struct {
	roomId RoomId
	users  map[UserId]vivid.ActorRef
}

func (r *Room) OnReceive(ctx vivid.ActorContext) {
	switch m := ctx.Message().(type) {
	case *vivid.OnLaunch:
		r.onLaunch(ctx)
	case JoinRoomMessage:
		r.onJoinRoom(ctx, m)
	case *Chat:
		r.onPacket(ctx, m)
	}
}

func (r *Room) onLaunch(ctx vivid.ActorContext) {
	r.users = make(map[UserId]vivid.ActorRef)
}

func (r *Room) onJoinRoom(ctx vivid.ActorContext, m JoinRoomMessage) {
	r.users[m.UserId] = m.User

	r.onPacket(ctx, &Chat{
		UserId: m.UserId,
		Packet: stream.NewPacketDC([]byte("joined"), websocket.TextMessage),
	})
}

func (r *Room) onPacket(ctx vivid.ActorContext, m *Chat) {
	for _, ref := range r.users {
		ctx.Tell(ref, m.Packet.Derivation([]byte(fmt.Sprintf("[%s]: %s", m.UserId, m.Data()))))
	}
}
