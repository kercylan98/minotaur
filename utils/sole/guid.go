package sole

import "sync"

var (
	global    int64
	namespace map[any]int64
	mutex     sync.Mutex
)

func init() {
	namespace = map[any]int64{}
}

func Get() int64 {
	global++
	return global
}

func GetWith(name any) int64 {
	namespace[name]++
	return namespace[name]
}

func GetSync() int64 {
	mutex.Lock()
	defer mutex.Unlock()
	global++
	return global
}

func GetSyncWith(name any) int64 {
	mutex.Lock()
	defer mutex.Unlock()
	namespace[name]++
	return namespace[name]
}
