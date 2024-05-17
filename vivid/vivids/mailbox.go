package vivids

import "github.com/kercylan98/minotaur/toolkit/queues"

type Mailbox interface {
	queues.Queue[MessageContext]
}
