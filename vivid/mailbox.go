package vivid

import "github.com/kercylan98/minotaur/toolkit/queues"

// NewMailbox 创建一个新的邮箱
func NewMailbox(queue queues.Queue[MessageContext]) *Mailbox {
	return &Mailbox{
		Queue: queue,
	}
}

// Mailbox 邮箱是 Actor 中用于接收消息的队列，该邮箱接受任意实现了 queues.Queue 接口的队列作为其实现
type Mailbox struct {
	queues.Queue[MessageContext]
}
