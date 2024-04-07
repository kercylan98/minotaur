package messages

import (
	"github.com/kercylan98/minotaur/server/internal/v2"
	"github.com/kercylan98/minotaur/server/internal/v2/queue"
	"github.com/kercylan98/minotaur/server/internal/v2/reactor"
	"github.com/kercylan98/minotaur/utils/super"
)

func NewAsync(handler func() error, callback func(err error)) server.MessageI {
	return &Async{handler: handler}
}

type Async struct {
	handler  func() error
	callback func(err error)
}

func (s *Async) OnInitialize(srv server.Server, reactor *reactor.Reactor[server.Message], queue *queue.Queue[int, string, server.Message], ident string) {

}

func (s *Async) OnProcess(finish func(err error)) {
	defer finish(super.RecoverTransform(recover()))

	s.handler()
}
