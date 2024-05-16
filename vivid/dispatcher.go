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

	// Stop 用于停止调度器
	Stop()
}

// newDispatcher 创建一个新的消息调度器
func newDispatcher() Dispatcher {
	d := &dispatcher{
		mailboxes:   make(map[ActorId]*Mailbox),
		mailboxWait: map[ActorId]chan struct{}{},
	}

	return d
}

type dispatcher struct {
	mailboxes   map[ActorId]*Mailbox      // ActorId -> Mailbox
	mailboxWait map[ActorId]chan struct{} // ActorId -> chan struct{}
	mailboxesRW sync.RWMutex              // 保护 mailboxes 的读写锁
	wait        sync.WaitGroup            // 等待所有 Actor 关闭
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

	wait := make(chan struct{})
	d.mailboxesRW.Lock()
	d.mailboxes[actor.GetId()] = mailbox
	d.mailboxWait[actor.GetId()] = wait
	d.mailboxesRW.Unlock()

	go mailbox.Start()
	go d.watchReceive(actor, wait, mailbox.Dequeue())
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
	wait := d.mailboxWait[actorId]
	delete(d.mailboxes, actorId)
	delete(d.mailboxWait, actorId)
	d.mailboxesRW.Unlock()

	d.wait.Add(1)
	go func() { // 异步等待邮箱关闭
		defer d.wait.Done()
		<-wait
	}()
	mailbox.Stop()

	return nil
}

func (d *dispatcher) Stop() {
	d.wait.Wait()
}

func (d *dispatcher) watchReceive(actor ActorCore, wait chan struct{}, dequeue <-chan MessageContext) {
	defer close(wait)
	for ctx := range dequeue {
		ctx.(*messageContext).ActorContext = actor
		actor.OnReceived(ctx)
	}
}
