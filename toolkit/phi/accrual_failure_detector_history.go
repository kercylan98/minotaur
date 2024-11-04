package phi

import (
	"math"
	"time"
)

func newAccrualFailureDetectorHistory(maxSampleSize uint, firstHeartbeatEstimate time.Duration) accrualFailureDetectorHistory {
	firstHeartbeatEstimateMS := uint64(firstHeartbeatEstimate.Milliseconds())
	stdDeviationMS := firstHeartbeatEstimateMS / 4
	return accrualFailureDetectorHistory{
		maxSampleSize: maxSampleSize,
		intervals:     make([]uint64, 0, maxSampleSize),
	}.
		append(firstHeartbeatEstimateMS - stdDeviationMS).
		append(firstHeartbeatEstimateMS + stdDeviationMS)
}

type accrualFailureDetectorHistory struct {
	maxSampleSize      uint
	intervals          []uint64
	intervalSum        uint64
	squaredIntervalSum uint64
}

// mean of this heartbeat history's intervals
func (h accrualFailureDetectorHistory) mean() float64 {
	return float64(h.intervalSum) / float64(len(h.intervals))
}

// variance of this heartbeat history's intervals
func (h accrualFailureDetectorHistory) variance() float64 {
	mean := h.mean()
	return (float64(h.squaredIntervalSum) / float64(len(h.intervals))) - (mean * mean)
}

// stdDeviation of this heartbeat history's intervals
func (h accrualFailureDetectorHistory) stdDeviation() float64 {
	return math.Sqrt(h.variance())
}

// append 将一个时间间隔添加到历史记录中并返回一个新对象
func (h accrualFailureDetectorHistory) append(interval uint64) accrualFailureDetectorHistory {
	var src []uint64
	var droppedInterval uint64

	if uint(len(h.intervals)) < h.maxSampleSize {
		src = h.intervals
	} else {
		droppedInterval = h.intervals[0]
		src = h.intervals[1:]
	}

	// 避免更改 h.intervals
	dst := make([]uint64, len(src)+1, h.maxSampleSize)
	copy(dst, src)
	dst[len(src)] = interval

	return accrualFailureDetectorHistory{
		maxSampleSize:      h.maxSampleSize,
		intervals:          dst,
		intervalSum:        h.intervalSum - droppedInterval + interval,
		squaredIntervalSum: h.squaredIntervalSum - (droppedInterval * droppedInterval) + (interval * interval)}
}
