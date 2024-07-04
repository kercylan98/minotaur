package vivid

import (
	"github.com/kercylan98/minotaur/core"
)

// ActorRef 是 Actor 的引用，同时也是进程引用的别名，进程是 Actor 的载体，因此 ActorRef 也是进程引用
type ActorRef = *core.ProcessRef
