package game

const (
	// MessageTypePlayer 玩家通过客户端发送到服务端的消息
	MessageTypePlayer byte = iota
	// MessageTypeEvent 服务端内部事件消息
	MessageTypeEvent
)

// EventTypeGuestJoin 服务器内部事件消息类型
const (
	EventTypeGuestJoin  byte = iota // 访客加入事件
	EventTypeGuestLeave             // 访客离开事件
)

// Message 游戏消息数据结构
type Message struct {
	t    byte
	args []any
}

func (slf *Message) Init(t byte, args ...any) *Message {
	slf.t = t
	slf.args = args
	return slf
}

func (slf *Message) Type() byte {
	return slf.t
}

func (slf *Message) Args() []any {
	return slf.args
}
