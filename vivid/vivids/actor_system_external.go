package vivids

import "github.com/panjf2000/ants/v2"

type ActorSystemExternal interface {
	// GetGoroutinePool 用于获取 goroutine 池
	GetGoroutinePool() *ants.Pool
}
