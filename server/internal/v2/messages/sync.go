package messages

import (
	"github.com/kercylan98/minotaur/server/internal/v2"
	"github.com/kercylan98/minotaur/server/internal/v2/queue"
	"github.com/kercylan98/minotaur/server/internal/v2/reactor"
	"github.com/kercylan98/minotaur/utils/super"
)

func NewSync(handler func()) server.MessageI {
	return &Sync{handler: handler}
}

type Sync struct {
	handler func()
}

func (s *Sync) OnInitialize(srv server.Server, reactor *reactor.Reactor[server.Message], queue *queue.Queue[int, string, server.Message], ident string) {

}

func (s *Sync) OnProcess(finish func(err error)) {
	defer finish(super.RecoverTransform(recover()))

	s.handler()
}
