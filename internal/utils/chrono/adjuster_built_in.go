package chrono

import "time"

var (
	builtInAdjuster *Adjuster
)

func init() {
	builtInAdjuster = NewAdjuster(0)
}

// BuiltInAdjuster 获取内置的由 NewAdjuster(0) 函数创建的时间调节器
func BuiltInAdjuster() *Adjuster {
	return builtInAdjuster
}

// Now 调用内置时间调节器 BuiltInAdjuster 的 Adjuster.Now 函数
func Now() time.Time {
	return BuiltInAdjuster().Now()
}

// Since 调用内置时间调节器 BuiltInAdjuster 的 Adjuster.Since 函数
func Since(t time.Time) time.Duration {
	return BuiltInAdjuster().Since(t)
}
