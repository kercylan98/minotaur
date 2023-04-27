package notify

func NewManager(senders ...Sender) *Manager {
	manager := &Manager{
		senders:       senders,
		closeChannel:  make(chan struct{}),
		notifyChannel: make(chan Notify, 48),
	}

	go func() {
		for {
			select {
			case <-manager.closeChannel:
				close(manager.closeChannel)
				close(manager.notifyChannel)
				return
			case notify := <-manager.notifyChannel:
				for _, sender := range manager.senders {
					sender.Push(notify)
				}
			}
		}
	}()

	return manager
}

type Manager struct {
	senders       []Sender
	notifyChannel chan Notify
	closeChannel  chan struct{}
}

// PushNotify 推送通知
func (slf *Manager) PushNotify(notify Notify) {
	slf.notifyChannel <- notify
}

func (slf *Manager) Release() {
	slf.closeChannel <- struct{}{}
}
