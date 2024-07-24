package user

import (
	"fmt"
	"github.com/gofiber/contrib/websocket"
	"github.com/kercylan98/minotaur/engine/stream"
	"github.com/kercylan98/minotaur/engine/vivid"
	"github.com/kercylan98/minotaur/examples/internal/chat-room/internal/manager"
	"github.com/kercylan98/minotaur/examples/internal/chat-room/internal/manager/room"
	"github.com/kercylan98/minotaur/toolkit/charproc"
)

func New(mgr vivid.ActorRef) *User {
	return &User{
		manager: mgr,
	}
}

type User struct {
	manager vivid.ActorRef
	writer  stream.Writer
	conn    *websocket.Conn
	userId  string
	room    vivid.ActorRef
}

func (u *User) OnReceive(ctx vivid.ActorContext) {
	switch m := ctx.Message().(type) {
	case stream.Writer:
		u.writer = m
	case *websocket.Conn:
		u.onConnOpened(ctx, m)
	case *stream.Packet:
		u.onPacket(ctx, m)
	case *room.ChatMessage:
		ctx.Tell(u.writer, stream.NewPacketSC(fmt.Sprintf("%s: %s", m.UserId, m.Chat), websocket.TextMessage))
	}
}

func (u *User) onConnOpened(ctx vivid.ActorContext, m *websocket.Conn) {
	userId := m.Query("userId")
	roomId := m.Query("roomId")
	if userId == charproc.None {
		ctx.Tell(u.writer, stream.NewPacketSC("please input userId query param, example for: ws://127.0.0.1:8080/ws?userId=kercylan&roomId=123", websocket.TextMessage))
		ctx.Terminate(ctx.Ref(), true)
		return
	}
	if roomId == charproc.None {
		ctx.Tell(u.writer, stream.NewPacketSC("please input roomId query param, example for: ws://127.0.0.1:8080/ws?userId=kercylan&roomId=123", websocket.TextMessage))
		ctx.Terminate(ctx.Ref(), true)
		return
	}

	// find or create
	reply, err := ctx.FutureAsk(u.manager, &manager.FindOrCreateRoomMessage{
		RoomId: roomId,
	}).Result()
	if err != nil {
		ctx.Tell(u.writer, stream.NewPacketSC("get room info failed, please retry!", websocket.TextMessage))
		ctx.Terminate(ctx.Ref(), true)
		return
	}

	r := reply.(*manager.FindOrCreateRoomMessageReply)
	ctx.Tell(r.Room, &room.JoinRoomMessage{
		UserId: userId,
		User:   ctx.Ref(),
	})

	u.userId = userId
	u.room = r.Room
}

func (u *User) onPacket(ctx vivid.ActorContext, m *stream.Packet) {
	if u.room == nil {
		return
	}

	ctx.Tell(u.room, &room.ChatMessage{
		UserId: u.userId,
		Chat:   string(m.Data()),
	})
}
