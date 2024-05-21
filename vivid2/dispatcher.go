package vivid

import "sync"

const (
	DefaultDispatcherId DispatcherId = 1 // 默认调度器 ID
)

// DispatcherId 调度器 ID
type DispatcherId = uint64

// Dispatcher Actor 调度器
type Dispatcher interface {
	// Attach 将一个 Actor 添加到调度器中
	Attach(actor ActorCore)

	// Detach 将一个 Actor 从调度器中移除
	Detach(actor ActorCore)

	// Send 向一个 Actor 发送消息
	Send(actor ActorCore, message MessageContext)
}

type _Dispatcher struct {
	group sync.WaitGroup
}

func (d *_Dispatcher) Attach(actor ActorCore) {
	factory := actor.GetMailboxFactory()
	mailbox := factory.Get()
	actor.BindMailbox(mailbox)
	d.group.Add(1)
	go func(group *sync.WaitGroup, mailbox Mailbox) {
		defer group.Done()
		mailbox.Start()
	}(&d.group, mailbox)
}

func (d *_Dispatcher) Detach(actor ActorCore) {
	factory := actor.GetMailboxFactory()
	mailbox := actor.GetMailbox()
	mailbox.Stop()
	factory.Put(mailbox)
}

func (d *_Dispatcher) Send(actor ActorCore, message MessageContext) {
	mailbox := actor.GetMailbox()
	mailbox.Enqueue(message)
}
