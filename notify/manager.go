package notify

import (
	"github.com/kercylan98/minotaur/utils/log"
	"go.uber.org/zap"
	"reflect"
)

// NewManager 通过指定的 Sender 创建一个通知管理器， senders 包中提供了一些内置的 Sender
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

// Manager 通知管理器，可用于将通知同时发送至多个渠道
type Manager struct {
	senders       []Sender
	notifyChannel chan Notify
	closeChannel  chan struct{}
}

// PushNotify 推送通知
func (slf *Manager) PushNotify(notify Notify) {
	slf.notifyChannel <- notify
}

// Release 释放通知管理器
func (slf *Manager) Release() {
	slf.closeChannel <- struct{}{}
}
