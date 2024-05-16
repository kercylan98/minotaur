package vivid

import (
	"sync"
)

// Dispatcher 消息调度器接口
type Dispatcher interface {
	// Send 用于向一个 Actor 发送消息
	Send(receiver ActorCore, msg MessageContext) error

	// Attach 用于将一个 Actor 添加到调度器中
	Attach(actor ActorCore) error

	// Detach 用于将一个 Actor 从调度器中移除
	Detach(actor ActorCore) error
}

// newDispatcher 创建一个新的消息调度器
func newDispatcher() Dispatcher {
	d := &dispatcher{
		mailboxes: make(map[ActorId]*Mailbox),
	}

	return d
}

type dispatcher struct {
	mailboxes   map[ActorId]*Mailbox // ActorId -> Mailbox
	mailboxesRW sync.RWMutex         // 保护 mailboxes 的读写锁
}

func (d *dispatcher) Send(receiver ActorCore, msg MessageContext) error {
	d.mailboxesRW.RLock()
	mailbox, exist := d.mailboxes[receiver.GetId()]
	d.mailboxesRW.RUnlock()
	if !exist {
		return ErrActorTerminated
	}

	mailbox.Enqueue(msg)
	return nil
}

func (d *dispatcher) Attach(actor ActorCore) error {
	opts := actor.GetOptions()
	// 为 Actor 创建一个邮箱
	mailbox := opts.Mailbox()

	d.mailboxesRW.Lock()
	d.mailboxes[actor.GetId()] = mailbox
	d.mailboxesRW.Unlock()

	go mailbox.Start()
	go d.watchReceive(actor, mailbox.Dequeue())
	return nil
}

func (d *dispatcher) Detach(actor ActorCore) error {
	actorId := actor.GetId()
	d.mailboxesRW.Lock()
	mailbox, exist := d.mailboxes[actorId]
	if !exist {
		d.mailboxesRW.Unlock()
		return nil
	}
	delete(d.mailboxes, actorId)
	d.mailboxesRW.Unlock()

	mailbox.Stop()

	return nil
}

func (d *dispatcher) watchReceive(actor ActorCore, dequeue <-chan MessageContext) {
	for ctx := range dequeue {
		ctx.(*messageContext).ActorContext = actor
		actor.OnReceived(ctx)
	}
}
