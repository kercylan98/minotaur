package vivid

import (
	"sync"
)

// Dispatcher 消息调度器接口
type Dispatcher interface {
	// OnInit 用于初始化调度器
	OnInit(system ActorSystemExternal)

	// Attach 用于将一个 Actor 添加到调度器中
	Attach(actor ActorCore) error

	// Detach 用于将一个 Actor 从调度器中移除
	Detach(actor ActorCore) error

	// Send 用于向一个 Actor 发送消息
	Send(receiver ActorCore, msg MessageContext) error

	// Stop 用于停止调度器
	Stop()
}

type dispatcherActorContext struct {
	mailbox *Mailbox
	wait    chan struct{}
}

// newDispatcher 创建一个新的消息调度器
func newDispatcher() Dispatcher {
	d := &dispatcher{}
	return d
}

type dispatcher struct {
	system ActorSystemExternal
	wait   sync.WaitGroup // 等待所有 Actor 关闭
}

func (d *dispatcher) OnInit(system ActorSystemExternal) {
	d.system = system
}

func (d *dispatcher) Attach(actor ActorCore) error {
	opts := actor.GetOptions()
	// 为 Actor 创建一个邮箱
	mailbox := opts.Mailbox()

	ctx := &dispatcherActorContext{
		mailbox: mailbox,
		wait:    make(chan struct{}),
	}
	actor.SetContext(ctx)

	pool := d.system.GetGoroutinePool()
	if err := pool.Submit(func() {
		mailbox.Start()
	}); err != nil {
		return err
	}

	if err := pool.Submit(func() {
		d.watchReceive(actor, ctx.wait, mailbox.Dequeue())
	}); err != nil {
		mailbox.Stop()
		return err
	}
	return nil
}

func (d *dispatcher) Detach(actor ActorCore) error {
	ctx, ok := actor.GetContext().(*dispatcherActorContext)
	if !ok {
		return nil
	}

	d.wait.Add(1)

	// 异步等待邮箱关闭，当池无法使用时降级使用内置 goroutine
	pool := d.system.GetGoroutinePool()
	if err := pool.Submit(func() {
		defer d.wait.Done()
		<-ctx.wait
	}); err != nil {
		go func() {
			defer d.wait.Done()
			<-ctx.wait
		}()
	}

	ctx.mailbox.Stop()
	return nil
}

func (d *dispatcher) Send(receiver ActorCore, msg MessageContext) error {
	ctx, ok := receiver.GetContext().(*dispatcherActorContext)
	if !ok {
		return ErrActorTerminated
	}

	ctx.mailbox.Enqueue(msg)
	return nil
}

func (d *dispatcher) Stop() {
	d.wait.Wait()
}

func (d *dispatcher) watchReceive(actor ActorCore, wait chan struct{}, dequeue <-chan MessageContext) {
	defer close(wait)
	for {
		if pause := actor.IsPause(); pause != nil {
			<-pause
		}

		ctx, ok := <-dequeue
		if !ok {
			break
		}
		actor.BindMessageActorContext(ctx)
		actor.OnReceived(ctx)
	}
}
