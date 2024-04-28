package chrono

import "time"

// NewAdjuster 创建一个时间调节器
func NewAdjuster(adjust time.Duration) *Adjuster {
	return &Adjuster{adjust: adjust}
}

// Adjuster 时间调节器是一个用于获取偏移时间的工具
type Adjuster struct {
	adjust time.Duration
}

// Adjust 获取偏移调整的时间量
func (a *Adjuster) Adjust() time.Duration {
	return a.adjust
}

// SetAdjust 设置偏移调整的时间量
func (a *Adjuster) SetAdjust(adjust time.Duration) {
	a.adjust = adjust
}

// AddAdjust 增加偏移调整的时间量
func (a *Adjuster) AddAdjust(adjust time.Duration) {
	a.adjust += adjust
}

// Now 获取经过偏移调整的当前时间
func (a *Adjuster) Now() time.Time {
	return time.Now().Add(a.adjust)
}

// Since 获取经过偏移调整的时间间隔
func (a *Adjuster) Since(t time.Time) time.Duration {
	return a.Now().Sub(t)
}
