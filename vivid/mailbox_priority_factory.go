package vivid

import (
	"github.com/kercylan98/minotaur/toolkit/pools"
)

func NewPriorityFactory(handler func(message MessageContext), opts ...*PriorityOptions) MailboxFactory {
	var pool = pools.NewObjectPool[Priority](func() *Priority {
		return NewPriority(handler, opts...)
	}, func(data *Priority) {
		data.reset()
	})

	return &PriorityFactory{pool: pool}
}

type PriorityFactory struct {
	pool *pools.ObjectPool[*Priority]
}

func (P *PriorityFactory) Get() Mailbox {
	return P.pool.Get()
}

func (P *PriorityFactory) Put(mailbox Mailbox) {
	priority, ok := mailbox.(*Priority)
	if !ok {
		return
	}
	P.pool.Release(priority)
}
