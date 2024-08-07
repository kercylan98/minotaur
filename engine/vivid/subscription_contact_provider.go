package vivid

import "github.com/kercylan98/minotaur/engine/prc"

type SubscriptionContactEvent struct {
	Address prc.PhysicalAddress
	Stop    bool
}

// SubscriptionContactProvider 订阅联络提供者
type SubscriptionContactProvider interface {
	ChangeNotify() <-chan *SubscriptionContactEvent
}
