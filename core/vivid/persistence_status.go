package vivid

type PersistenceStatus interface {
	// GetSnapshot 获取当前快照
	GetSnapshot() Message

	// GetEvents 获取当前事件
	GetEvents() []Message
}

var _ PersistenceStatus = &persistenceStatus{}

type persistenceStatus struct {
	ctx        *actorContext
	snapshot   Message   // 当前最新的快照
	events     []Message // 当前最新的事件
	eventLimit int       // 事件数量限制，超过限制则需要生成快照
	recovery   bool      // 是否正在恢复
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
		m.ctx.ProcessSystemMessage(onPersistenceSnapshot)
	} else {
		m.events = append(m.events, event)
	}
}
