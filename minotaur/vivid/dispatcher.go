package vivid

const (
	DefaultDispatcherId DispatcherId = 1 // 默认调度器 ID
)

// DispatcherId 调度器 ID
type DispatcherId = uint64

// Dispatcher Actor 调度器
type Dispatcher interface {
	// Attach 将一个 Actor 添加到调度器中
	Attach(system ActorSystemCore, actor ActorCore)

	// Detach 将一个 Actor 从调度器中移除
	Detach(system ActorSystemCore, actor ActorCore)

	// Send 向一个 Actor 发送消息。如果消息发送成功，返回 true；否则返回 false
	Send(system ActorSystemCore, actor ActorCore, message MessageContext) bool
}

type _Dispatcher struct {
}

func (d *_Dispatcher) Attach(system ActorSystemCore, actor ActorCore) {
	factory := actor.GetMailboxFactory()
	mailbox := factory.Get()
	actor.BindMailbox(mailbox)
	mailbox.Start()
}

func (d *_Dispatcher) Detach(system ActorSystemCore, actor ActorCore) {
	factory := actor.GetMailboxFactory()
	mailbox := actor.GetMailbox()
	mailbox.Stop()
	factory.Put(mailbox)
}

func (d *_Dispatcher) Send(system ActorSystemCore, actor ActorCore, message MessageContext) bool {
	mailbox := actor.GetMailbox()
	return mailbox.Enqueue(message, message.Instantly())
}
