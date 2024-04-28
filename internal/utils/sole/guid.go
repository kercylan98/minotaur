package sole

import (
	"sync/atomic"
)

var (
	global    atomic.Int64              // 全局唯一标识符
	namespace = map[any]*atomic.Int64{} // 唯一标识符命名空间
)

// RegNameSpace 注册特定命名空间的唯一标识符
func RegNameSpace(name any) {
	if namespace == nil {
		namespace = map[any]*atomic.Int64{}
	}
	namespace[name] = new(atomic.Int64)
}

// UnRegNameSpace 解除注销特定命名空间的唯一标识符
func UnRegNameSpace(name any) {
	delete(namespace, name)
}

// Get 获取全局唯一标识符
func Get() int64 {
	return global.Add(1)
}

// Reset 重置全局唯一标识符
func Reset() {
	global.Store(0)
}

// GetWith 获取特定命名空间的唯一标识符
func GetWith(name any) int64 {
	return namespace[name].Add(1)
}

// ResetWith 重置特定命名空间的唯一标识符
func ResetWith(name any) {
	namespace[name].Store(0)
}
