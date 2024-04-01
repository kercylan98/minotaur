package server

type Message interface {
	Execute()
}

func NativeMessage(srv *server, handler func(srv *server)) Message {
	return &nativeMessage{srv: srv, handler: handler}
}

type nativeMessage struct {
	srv     *server
	handler func(srv *server)
}

func (s *nativeMessage) Execute() {
	s.handler(s.srv)
}
