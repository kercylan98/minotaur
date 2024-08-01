package room

import (
	"github.com/kercylan98/minotaur/engine/prc"
	"github.com/kercylan98/minotaur/engine/vivid"
)

type (
	JoinRoomMessage struct {
		UserId string
		User   vivid.ActorRef
	}

	ChatMessage struct {
		UserId string
		Chat   string
	}
)

func New(roomId string) *Room {
	return &Room{
		roomId:  roomId,
		users:   map[string]vivid.ActorRef{},
		userMap: make(map[prc.LogicalAddress]string),
	}
}

type Room struct {
	roomId  string
	users   map[string]vivid.ActorRef
	userMap map[prc.LogicalAddress]string
}

func (r *Room) OnReceive(ctx vivid.ActorContext) {
	switch m := ctx.Message().(type) {
	case *JoinRoomMessage:
		r.onJoinRoom(ctx, m)
	case *ChatMessage:
		r.onChat(ctx, m)
	case *vivid.OnTerminated:
		r.onTerminated(ctx, m)
	}
}

func (r *Room) onJoinRoom(ctx vivid.ActorContext, m *JoinRoomMessage) {
	if _, exist := r.users[m.UserId]; exist {
		return
	}

	r.userMap[m.User.GetLogicalAddress()] = m.UserId
	r.users[m.UserId] = m.User
	message := &ChatMessage{
		UserId: m.UserId,
		Chat:   "Joined the room",
	}

	ctx.Watch(m.User)

	for _, ref := range r.users {
		ctx.Tell(ref, message)
	}
}

func (r *Room) onChat(ctx vivid.ActorContext, m *ChatMessage) {
	if _, exist := r.users[m.UserId]; !exist {
		return
	}
	for _, ref := range r.users {
		ctx.Tell(ref, m)
	}
}

func (r *Room) onTerminated(ctx vivid.ActorContext, m *vivid.OnTerminated) {
	la := m.TerminatedActor.GetLogicalAddress()
	userId, exist := r.userMap[la]
	if !exist {
		return
	}

	delete(r.users, userId)
	delete(r.userMap, la)

	message := &ChatMessage{
		UserId: userId,
		Chat:   "Leaved the room",
	}
	for _, ref := range r.users {
		ctx.Tell(ref, message)
	}

	if len(r.userMap) == 0 {
		ctx.Terminate(ctx.Ref(), true)
	}
}
