package time

import "time"

// Time 带有偏移量的时间
type Time struct {
	offset time.Duration
}

func (slf *Time) SetOffset(offset time.Duration) {
	slf.offset = offset
}

func (slf *Time) Now() time.Time {
	return time.Now().Add(slf.offset)
}

func (slf *Time) Since(t time.Time) time.Duration {
	return slf.Now().Sub(t)
}
