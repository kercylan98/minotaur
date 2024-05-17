package unsafevivid

import (
	"github.com/kercylan98/minotaur/toolkit/log"
	"github.com/kercylan98/minotaur/toolkit/queues"
	vivid "github.com/kercylan98/minotaur/vivid/vivids"
	"reflect"
	"sync"
)

var dispatcherActorContextKey = reflect.TypeOf((*DispatcherActorContext)(nil)).Elem()

type DispatcherActorContext struct {
	mailbox vivid.Mailbox
	wait    chan struct{}
}

// NewDispatcher 创建一个新的消息调度器
func NewDispatcher() *Dispatcher {
	d := &Dispatcher{}
	return d
}

type Dispatcher struct {
	system vivid.ActorSystemExternal
	wait   sync.WaitGroup // 等待所有 Actor 关闭
}

func (d *Dispatcher) OnInit(system vivid.ActorSystemExternal) {
	d.system = system
}

func (d *Dispatcher) Attach(actor vivid.ActorCore) error {
	opts := actor.GetOptions()
	// 为 Actor 创建一个邮箱
	var mailbox vivid.Mailbox
	if opts.Mailbox != nil {
		mailbox = opts.Mailbox()
	} else {
		mailbox = NewMailbox(queues.NewFIFO[vivid.MessageContext](queues.NewFIFOOptions().WithBufferSize(128)))
	}

	ctx := &DispatcherActorContext{
		mailbox: mailbox,
		wait:    make(chan struct{}),
	}
	actor.SetContext(dispatcherActorContextKey, ctx)

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

func (d *Dispatcher) Detach(actor vivid.ActorCore) error {
	ctx, ok := actor.GetContext(dispatcherActorContextKey).(*DispatcherActorContext)
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

func (d *Dispatcher) Send(receiver vivid.ActorCore, msg vivid.MessageContext) error {
	ctx, ok := receiver.GetContext(dispatcherActorContextKey).(*DispatcherActorContext)
	if !ok {
		return vivid.ErrActorTerminated
	}

	ctx.mailbox.Enqueue(msg)
	log.Debug("send message to actor", log.Uint64("seq", msg.GetSeq()),
		log.String("msg", reflect.TypeOf(msg.GetMessage()).String()),
		log.String("sender", msg.GetSenderId().String()),
		log.String("receiver", msg.GetReceiverId().String()),
	)
	return nil
}

func (d *Dispatcher) Stop() {
	d.wait.Wait()
}

func (d *Dispatcher) watchReceive(actor vivid.ActorCore, wait chan struct{}, dequeue <-chan vivid.MessageContext) {
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
		if err := actor.OnReceived(ctx); err != nil {
			log.Error("actor receive message failed", log.Err(err))
		}
	}
}
