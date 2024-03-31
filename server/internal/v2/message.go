package server

type Message interface {
	Execute()
}

func HandlerMessage(srv *server, handler func(srv *server)) Message {
	return &handlerMessage{srv: srv, handler: handler}
}

type handlerMessage struct {
	srv     *server
	handler func(srv *server)
}

func (s *handlerMessage) Execute() {
	s.handler(s.srv)
}
