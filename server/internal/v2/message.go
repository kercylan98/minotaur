package server

import (
	"github.com/kercylan98/minotaur/server/internal/v2/queue"
	"github.com/kercylan98/minotaur/server/internal/v2/reactor"
)

type Message interface {
	// OnInitialize 消息初始化阶段将会被告知消息所在服务器、反应器、队列及标识信息
	OnInitialize(controller Controller, reactor *reactor.Reactor[Message], message queue.MessageWrapper[int, string, Message])

	// OnProcess 消息处理阶段需要完成对消息的处理，并返回处理结果
	OnProcess()
}

// GenerateSystemSyncMessage 生成系统同步消息
func GenerateSystemSyncMessage(handler func(srv Server)) Message {
	return &systemSyncMessage{handler: handler}
}

type systemSyncMessage struct {
	controller Controller
	handler    func(srv Server)
}

func (m *systemSyncMessage) OnInitialize(controller Controller, reactor *reactor.Reactor[Message], message queue.MessageWrapper[int, string, Message]) {
	m.controller = controller
}

func (m *systemSyncMessage) OnProcess() {
	m.handler(m.controller.GetServer())
}

// GenerateSystemAsyncMessage 生成系统异步消息
func GenerateSystemAsyncMessage(handler func(srv Server) error, callback func(srv Server, err error)) Message {
	return &systemAsyncMessage{
		handler:  handler,
		callback: callback,
	}
}

type systemAsyncMessage struct {
	controller Controller
	queue      *queue.Queue[int, string, Message]
	handler    func(srv Server) error
	callback   func(srv Server, err error)
	hasIdent   bool
	ident      string
}

func (m *systemAsyncMessage) OnInitialize(controller Controller, reactor *reactor.Reactor[Message], message queue.MessageWrapper[int, string, Message]) {
	m.controller = controller
	m.queue = message.Queue()
	m.ident = message.Ident()
	m.hasIdent = message.HasIdent()
}

func (m *systemAsyncMessage) OnProcess() {
	var ident = m.ident

	m.queue.WaitAdd(ident, 1)
	err := m.controller.GetAnts().Submit(func() {
		err := m.handler(m.controller.GetServer())
		if !m.hasIdent {
			m.controller.PushSystemMessage(GenerateSystemAsyncCallbackMessage(m.callback, err), func(err error) {
				m.queue.WaitAdd(ident, -1)
			})
		} else {
			m.controller.PushIdentMessage(ident, GenerateSystemAsyncCallbackMessage(m.callback, err), func(err error) {
				m.queue.WaitAdd(ident, -1)
			})
		}
		if err != nil {
			m.queue.WaitAdd(ident, -1)
		}
	})
	if err != nil {
		m.controller.MessageErrProcess(m, err)
		m.queue.WaitAdd(ident, -1)
	}
}

// GenerateSystemAsyncCallbackMessage 生成系统异步回调消息
func GenerateSystemAsyncCallbackMessage(handler func(srv Server, err error), err error) Message {
	return &systemAsyncCallbackMessage{
		err:     err,
		handler: handler,
	}
}

type systemAsyncCallbackMessage struct {
	controller Controller
	err        error
	handler    func(srv Server, err error)
	queue      *queue.Queue[int, string, Message]
	ident      string
}

func (m *systemAsyncCallbackMessage) OnInitialize(controller Controller, reactor *reactor.Reactor[Message], message queue.MessageWrapper[int, string, Message]) {
	m.controller = controller
	m.queue = message.Queue()
	m.ident = message.Ident()
}

func (m *systemAsyncCallbackMessage) OnProcess() {
	defer func(m *systemAsyncCallbackMessage) {
		m.queue.WaitAdd(m.ident, -1)
	}(m)

	if m.handler != nil {
		m.handler(m.controller.GetServer(), m.err)
	}
}

// GenerateConnSyncMessage 生成连接同步消息
func GenerateConnSyncMessage(conn Conn, handler func(srv Server, conn Conn)) Message {
	return &connSyncMessage{handler: handler, conn: conn}
}

type connSyncMessage struct {
	controller Controller
	conn       Conn
	handler    func(srv Server, conn Conn)
}

func (m *connSyncMessage) OnInitialize(controller Controller, reactor *reactor.Reactor[Message], message queue.MessageWrapper[int, string, Message]) {
	m.controller = controller
}

func (m *connSyncMessage) OnProcess() {
	m.handler(m.controller.GetServer(), m.conn)
}

// GenerateConnAsyncMessage 生成连接异步消息
func GenerateConnAsyncMessage(conn Conn, handler func(srv Server, conn Conn) error, callback func(srv Server, conn Conn, err error)) Message {
	return &connAsyncMessage{
		conn:     conn,
		handler:  handler,
		callback: callback,
	}
}

type connAsyncMessage struct {
	controller Controller
	conn       Conn
	queue      *queue.Queue[int, string, Message]
	handler    func(srv Server, conn Conn) error
	callback   func(srv Server, conn Conn, err error)
	ident      string
	hasIdent   bool
}

func (m *connAsyncMessage) OnInitialize(controller Controller, reactor *reactor.Reactor[Message], message queue.MessageWrapper[int, string, Message]) {
	m.controller = controller
	m.queue = message.Queue()
	m.ident = message.Ident()
	m.hasIdent = message.HasIdent()
}

func (m *connAsyncMessage) OnProcess() {
	m.queue.WaitAdd(m.ident, 1)
	err := m.controller.GetAnts().Submit(func() {
		err := m.handler(m.controller.GetServer(), m.conn)
		if !m.hasIdent {
			m.controller.PushSystemMessage(GenerateConnAsyncCallbackMessage(m.conn, m.callback, err), func(err error) {
				m.queue.WaitAdd(m.ident, -1)
			})
		} else {
			m.controller.PushIdentMessage(m.ident, GenerateConnAsyncCallbackMessage(m.conn, m.callback, err), func(err error) {
				m.queue.WaitAdd(m.ident, -1)
			})
		}
		if err != nil {
			m.queue.WaitAdd(m.ident, -1)
		}
	})
	if err != nil {
		m.controller.MessageErrProcess(m, err)
		m.queue.WaitAdd(m.ident, -1)
	}
}

// GenerateConnAsyncCallbackMessage 生成连接异步回调消息
func GenerateConnAsyncCallbackMessage(conn Conn, handler func(srv Server, conn Conn, err error), err error) Message {
	return &connAsyncCallbackMessage{
		conn:    conn,
		err:     err,
		handler: handler,
	}
}

type connAsyncCallbackMessage struct {
	controller Controller
	conn       Conn
	err        error
	handler    func(srv Server, conn Conn, err error)
	queue      *queue.Queue[int, string, Message]
	ident      string
}

func (m *connAsyncCallbackMessage) OnInitialize(controller Controller, reactor *reactor.Reactor[Message], message queue.MessageWrapper[int, string, Message]) {
	m.controller = controller
	m.queue = message.Queue()
	m.ident = message.Ident()
}

func (m *connAsyncCallbackMessage) OnProcess() {
	defer func(m *connAsyncCallbackMessage) {
		m.queue.WaitAdd(m.ident, -1)
	}(m)

	if m.handler != nil {
		m.handler(m.controller.GetServer(), m.conn, m.err)
	}
}
