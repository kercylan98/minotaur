package dispatcher

import (
	"sync"
)

const SystemName = "*system"

// NewManager 生成消息分发器管理器
func NewManager[P Producer, M Message[P]](bufferSize int, handler Handler[P, M]) *Manager[P, M] {
	mgr := &Manager[P, M]{
		handler:     handler,
		dispatchers: make(map[string]*Dispatcher[P, M]),
		member:      make(map[string]map[P]struct{}),
		sys:         NewDispatcher(bufferSize, SystemName, handler).Start(),
		curr:        make(map[P]*Dispatcher[P, M]),
		size:        bufferSize,
	}

	return mgr
}

// Manager 消息分发器管理器
type Manager[P Producer, M Message[P]] struct {
	handler     Handler[P, M]                // 消息处理器
	sys         *Dispatcher[P, M]            // 系统消息分发器
	dispatchers map[string]*Dispatcher[P, M] // 当前所有正在工作的消息分发器
	member      map[string]map[P]struct{}    // 当前正在工作的消息分发器对应的生产者
	curr        map[P]*Dispatcher[P, M]      // 当前特定生产者正在使用的消息分发器
	lock        sync.RWMutex                 // 消息分发器锁
	w           sync.WaitGroup               // 消息分发器等待组
	size        int                          // 消息分发器缓冲区大小

	closedHandler  func(name string)
	createdHandler func(name string)
}

// SetDispatcherClosedHandler 设置消息分发器关闭时的回调函数
func (m *Manager[P, M]) SetDispatcherClosedHandler(handler func(name string)) *Manager[P, M] {
	m.closedHandler = handler
	return m
}

// SetDispatcherCreatedHandler 设置消息分发器创建时的回调函数
func (m *Manager[P, M]) SetDispatcherCreatedHandler(handler func(name string)) *Manager[P, M] {
	m.createdHandler = handler
	return m
}

// HasDispatcher 检查是否存在指定名称的消息分发器
func (m *Manager[P, M]) HasDispatcher(name string) bool {
	m.lock.RLock()
	defer m.lock.RUnlock()
	_, exist := m.dispatchers[name]
	return exist
}

// GetDispatcherNum 获取当前正在工作的消息分发器数量
func (m *Manager[P, M]) GetDispatcherNum() int {
	m.lock.RLock()
	defer m.lock.RUnlock()
	return len(m.dispatchers) + 1 // +1 系统消息分发器
}

// GetSystemDispatcher 获取系统消息分发器
func (m *Manager[P, M]) GetSystemDispatcher() *Dispatcher[P, M] {
	return m.sys
}

// GetDispatcher 获取生产者正在使用的消息分发器，如果生产者没有绑定消息分发器，则会返回系统消息分发器
func (m *Manager[P, M]) GetDispatcher(p P) *Dispatcher[P, M] {
	m.lock.Lock()
	defer m.lock.Unlock()

	curr, exist := m.curr[p]
	if exist {
		return curr
	}

	return m.sys
}

// BindProducer 绑定生产者使用特定的消息分发器，如果生产者已经绑定了消息分发器，则会先解绑
func (m *Manager[P, M]) BindProducer(p P, name string) {
	if name == SystemName {
		return
	}
	m.lock.Lock()
	defer m.lock.Unlock()
	member, exist := m.member[name]
	if !exist {
		member = make(map[P]struct{})
		m.member[name] = member
	}

	if _, exist = member[p]; exist {
		d := m.dispatchers[name]
		d.SetProducerDoneHandler(p, nil)
		d.UnExpel()

		return
	}

	curr, exist := m.curr[p]
	if exist {
		delete(m.member[curr.name], p)
		if len(m.member[curr.name]) == 0 {
			curr.Expel()
		}
	}

	dispatcher, exist := m.dispatchers[name]
	if !exist {
		dispatcher = NewDispatcher(m.size, name, m.handler).SetClosedHandler(func(dispatcher *Dispatcher[P, M]) {
			// 消息分发器关闭时，将会将其从管理器中移除
			m.lock.Lock()
			delete(m.dispatchers, dispatcher.name)
			delete(m.member, dispatcher.name)
			m.lock.Unlock()
			if m.closedHandler != nil {
				m.closedHandler(dispatcher.name)
			}
		}).Start()
		m.dispatchers[name] = dispatcher
		defer func(m *Manager[P, M], name string) {
			if m.createdHandler != nil {
				m.createdHandler(name)
			}
		}(m, dispatcher.Name())
	}
	m.curr[p] = dispatcher
	member[p] = struct{}{}
}

// UnBindProducer 解绑生产者使用特定的消息分发器
func (m *Manager[P, M]) UnBindProducer(p P) {
	m.lock.Lock()
	defer m.lock.Unlock()
	curr, exist := m.curr[p]
	if !exist {
		return
	}

	curr.SetProducerDoneHandler(p, func(p P, dispatcher *Dispatcher[P, M]) {
		m.lock.Lock()
		defer m.lock.Unlock()
		delete(m.member[dispatcher.name], p)
		delete(m.curr, p)
		if len(m.member[dispatcher.name]) == 0 {
			dispatcher.Expel()
		}
	})
}
