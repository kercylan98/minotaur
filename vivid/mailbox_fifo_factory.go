package vivid

import (
	"github.com/kercylan98/minotaur/toolkit/pools"
)

func NewFIFOFactory(handler func(message MessageContext), opts ...*FIFOOptions) MailboxFactory {
	var pool = pools.NewObjectPool[FIFO](func() *FIFO {
		return NewFIFO(handler, opts...)
	}, func(data *FIFO) {
		data.reset()
	})

	return &FIFOFactory{pool: pool}
}

type FIFOFactory struct {
	pool *pools.ObjectPool[*FIFO]
}

func (F *FIFOFactory) Get() Mailbox {
	return F.pool.Get()
}

func (F *FIFOFactory) Put(mailbox Mailbox) {
	fifo, ok := mailbox.(*FIFO)
	if !ok {
		return
	}
	F.pool.Release(fifo)
}
