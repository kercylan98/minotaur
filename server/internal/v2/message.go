package server

import (
	"github.com/kercylan98/minotaur/server/internal/v2/queue"
	"github.com/kercylan98/minotaur/server/internal/v2/reactor"
	"github.com/kercylan98/minotaur/utils/log/v2"
	"github.com/kercylan98/minotaur/utils/super"
	"runtime/debug"
)

type MessageI interface {
	// OnInitialize 消息初始化阶段将会被告知消息所在服务器、反应器、队列及标识信息
	OnInitialize(srv Server, reactor *reactor.Reactor[Message], queue *queue.Queue[int, string, Message], ident string)

	// OnProcess 消息处理阶段需要完成对消息的处理，并返回处理结果
	OnProcess(finish func(err error))
}

type Message interface {
	OnExecute()
}

func SyncMessage(srv *server, handler func(srv *server)) Message {
	return &syncMessage{srv: srv, handler: handler}
}

type syncMessage struct {
	srv     *server
	handler func(srv *server)
}

func (s *syncMessage) OnExecute() {
	s.handler(s.srv)
}

func AsyncMessage(srv *server, ident string, handler func(srv *server) error, callback func(srv *server, err error)) Message {
	return &asyncMessage{
		ident:    ident,
		srv:      srv,
		handler:  handler,
		callback: callback,
	}
}

type asyncMessage struct {
	ident    string
	srv      *server
	handler  func(srv *server) error
	callback func(srv *server, err error)
}

func (s *asyncMessage) OnExecute() {
	var q *queue.Queue[int, string, Message]
	var dispatch = func(ident string, message Message, beforeHandler ...func(queue *queue.Queue[int, string, Message], msg Message)) {
		_ = s.srv.reactor.AutoDispatch(ident, message, beforeHandler...)
	}

	dispatch(
		s.ident,
		SyncMessage(s.srv, func(srv *server) {
			_ = srv.ants.Submit(func() {
				defer func(srv *server, msg *asyncMessage) {
					if err := super.RecoverTransform(recover()); err != nil {
						if errHandler := srv.GetMessageErrorHandler(); errHandler != nil {
							errHandler(srv, msg, err)
						} else {
							srv.GetLogger().Error("Message", log.Err(err))
							debug.PrintStack()
						}
					}
				}(s.srv, s)

				err := s.handler(srv)
				var msg Message
				msg = SyncMessage(srv, func(srv *server) {
					defer func() {
						q.WaitAdd(s.ident, -1)
						if err := super.RecoverTransform(recover()); err != nil {
							if errHandler := srv.GetMessageErrorHandler(); errHandler != nil {
								errHandler(srv, msg, err)
							} else {
								srv.GetLogger().Error("Message", log.Err(err))
								debug.PrintStack()
							}
						}
					}()
					if s.callback != nil {
						s.callback(srv, err)
					}
				})
				dispatch(s.ident, msg)

			})
		}),
		func(queue *queue.Queue[int, string, Message], msg Message) {
			queue.WaitAdd(s.ident, 1)
			q = queue
		},
	)
}
