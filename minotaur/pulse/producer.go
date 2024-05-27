package pulse

import "github.com/kercylan98/minotaur/minotaur/vivid"

// Producer 事件生产者
type Producer interface {
	vivid.ActorRef
}
