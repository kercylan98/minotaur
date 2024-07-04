package vivid

type PersistenceStatus interface {
	// GetSnapshot 获取当前快照
	GetSnapshot() Message

	// GetEvents 获取当前事件
	GetEvents() []Message
}

var _ PersistenceStatus = &persistenceStatus{}

type persistenceStatus struct {
	ctx            *actorContext
	snapshot       Message   // 当前最新的快照
	events         []Message // 当前最新的事件
	eventLimit     int       // 事件数量限制，超过限制则需要生成快照
	recovery       bool      // 是否正在恢复
	persistentDone bool      // 持久化完成（需要确保发送持久化快照消息被处理，否则可能导致事件丢失）

	persistenceName       string  // Actor 持久化名称
	persistenceStorage    Storage // Actor 持久化存储器
	persistenceEventLimit int     // Actor 持久化事件数量限制，达到限制时将会触发快照的生成
}

func (m *persistenceStatus) GetSnapshot() Message {
	return m.snapshot
}

func (m *persistenceStatus) GetEvents() []Message {
	return m.events
}

func (m *persistenceStatus) PersistSnapshot(snapshot Message) {
	if m.recovery {
		return
	}
	m.snapshot = snapshot
	m.events = m.events[:0]
}

func (m *persistenceStatus) StatusChanged(event Message) {
	if m.recovery {
		return
	}
	if len(m.events) >= m.eventLimit {
		m.persistentDone = false
		m.ctx.ProcessSystemMessage(onPersistenceSnapshot)
		if !m.persistentDone {
			m.events = append(m.events, event)
		}
	} else {
		m.events = append(m.events, event)
	}
}
