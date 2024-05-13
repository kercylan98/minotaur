package vivid

import "encoding/json"

type Context interface {
	// GetCommand 获取消息的命令
	GetCommand() string

	// GetSender 获取消息的发送者
	GetSender() ActorId

	// Reply 用于回复消息
	Reply(v any)

	// ReadTo 用于将消息的数据解码到指定的结构体
	ReadTo(v any) error

	// MustReadTo 用于将消息的数据解码到指定的结构体，如果解码失败则 panic
	MustReadTo(v any)
}

type Result func(dst any) error

func (r Result) ReadTo(v any) error {
	return r(v)
}

func (r Result) MustReadTo(v any) {
	if err := r(v); err != nil {
		panic(err)
	}
}

type context struct {
	Sender   ActorId
	Receiver ActorId
	Command  string
	Data     json.RawMessage
	reply    any
	reader   func(dst any) error
	done     chan struct{}
	err      error
}

func (m *context) Reply(v any) {
	m.reply = v
}

func (m *context) ReadTo(v any) error {
	return m.reader(v)
}

func (m *context) MustReadTo(v any) {
	if err := m.reader(v); err != nil {
		panic(err)
	}
}

func (m *context) GetCommand() string {
	return m.Command
}

func (m *context) GetSender() ActorId {
	return m.Sender
}

func (m *context) GetReceiver() ActorId {
	return m.Receiver
}
