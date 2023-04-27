package notify

import (
	"go.uber.org/zap"
	"minotaur/utils/log"
	"reflect"
)

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
				log.Info("Manager", zap.String("state", "release"))
				return
			case notify := <-manager.notifyChannel:
				for _, sender := range manager.senders {
					if err := sender.Push(notify); err != nil {
						log.Error("Manager", zap.String("sender", reflect.TypeOf(sender).String()), zap.Error(err))
					}
				}
			}
		}
	}()

	log.Info("Manager", zap.String("state", "running"))
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
