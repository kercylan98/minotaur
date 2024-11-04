package phi

import (
	"fmt"
	"math"
	"sync/atomic"
	"time"
)

// NewAccrualFailureDetector 返回一个新的 AccrualFailureDetector 实例
func NewAccrualFailureDetector(threshold float64, maxSampleSize uint, minStdDeviation, acceptableHeartbeatPause, firstHeartbeatEstimate time.Duration, eventStream chan<- time.Duration) (*AccrualFailureDetector, error) {

	if threshold <= 0.0 {
		return nil, fmt.Errorf("threshold must be > 0, got %f", threshold)
	}

	if maxSampleSize <= 0 {
		return nil, fmt.Errorf("maxSampleSize must be > 0, got %d", maxSampleSize)
	}

	if minStdDeviation <= 0 {
		return nil, fmt.Errorf("minStdDeviation must be > 0, got %d", minStdDeviation)
	}

	if acceptableHeartbeatPause < 0 {
		return nil, fmt.Errorf("acceptableHeartbeatPause must be >= 0, got %d", acceptableHeartbeatPause)
	}

	if firstHeartbeatEstimate <= 0 {
		return nil, fmt.Errorf("heartbeatInterval must be > 0, got %d", firstHeartbeatEstimate)
	}

	firstHeartbeat := newAccrualFailureDetectorHistory(maxSampleSize, firstHeartbeatEstimate)

	afd := &AccrualFailureDetector{
		threshold:                  threshold,
		maxSampleSize:              maxSampleSize,
		minStdDeviation:            minStdDeviation,
		acceptableHeartbeatPause:   acceptableHeartbeatPause,
		firstHeartbeatEstimate:     firstHeartbeatEstimate,
		eventStream:                eventStream,
		firstHeartbeat:             firstHeartbeat,
		acceptableHeartbeatPauseMS: uint64(acceptableHeartbeatPause.Milliseconds()),
		minStdDeviationMS:          uint64(minStdDeviation.Milliseconds()),
	}
	afd.state.Store(&accrualFailureDetectorState{history: firstHeartbeat})
	return afd, nil
}

type AccrualFailureDetector struct {
	threshold                  float64       // 较低的阈值很容易产生许多错误的怀疑，但可以确保在真正发生碰撞时能够快速检测到。相反，高阈值产生的错误较少，但需要更多时间来检测实际崩溃
	maxSampleSize              uint          // 用于计算到达间隔时间的平均值和标准差的样本数
	minStdDeviation            time.Duration // 计算 phi 时使用的正态分布的最小标准差。标准偏差太低可能会导致对心跳到达时间的突然但正常的偏差过于敏感。
	acceptableHeartbeatPause   time.Duration // 与在将其视为异常之前将接受的潜在丢失或延迟心跳数量相对应的持续时间。这个余量对于能够承受由于垃圾收集或网络丢失等原因导致的心跳到达的突然、偶尔的暂停非常重要。
	firstHeartbeatEstimate     time.Duration // 使用与该持续时间相对应的心跳引导统计数据，具有相当高的标准差（因为一开始环境是未知的）
	eventStream                chan<- time.Duration
	firstHeartbeat             accrualFailureDetectorHistory
	acceptableHeartbeatPauseMS uint64
	minStdDeviationMS          uint64
	state                      atomic.Pointer[accrualFailureDetectorState]
}

// IsAvailable 返回资源正常运行
func (fd *AccrualFailureDetector) IsAvailable() bool {
	return fd.isAvailableAt(time.Now())
}

// isAvailableAt 返回资源在指定的时间是否可用
func (fd *AccrualFailureDetector) isAvailableAt(time time.Time) bool {
	return fd.phiAt(time) < fd.threshold
}

// IsMonitoring 如果故障检测器已收到任何心跳并开始监视资源，则 IsMonitoring 返回 true
func (fd *AccrualFailureDetector) IsMonitoring() bool {
	return fd.state.Load().timestamp != nil
}

// Heartbeat 通知检测器有来自受监控资源的心跳，更新其状态
func (fd *AccrualFailureDetector) Heartbeat() {
	for {
		timestamp := time.Now()
		oldState := fd.state.Load()

		var newHistory accrualFailureDetectorHistory

		if latestTimestamp := oldState.timestamp; latestTimestamp == nil {
			newHistory = fd.firstHeartbeat
		} else {
			interval := timestamp.Sub(*latestTimestamp)
			if fd.isAvailableAt(timestamp) {
				intervalMS := uint64(interval.Milliseconds())
				if intervalMS >= (fd.acceptableHeartbeatPauseMS/2) && fd.eventStream != nil {
					fd.eventStream <- interval
				}
				newHistory = oldState.history.append(intervalMS)
			} else {
				newHistory = oldState.history
			}
		}

		newState := &accrualFailureDetectorState{history: newHistory, timestamp: &timestamp} // record new timestamp

		// 如果更新失败，那么继续重试
		if fd.state.CompareAndSwap(oldState, newState) {
			break
		}
	}
}

// Phi 返回故障检测器的怀疑级别
func (fd *AccrualFailureDetector) Phi() float64 {
	return fd.phiAt(time.Now())
}

func (fd *AccrualFailureDetector) phiAt(timestamp time.Time) float64 {
	oldState := fd.state.Load()
	oldTimestamp := oldState.timestamp

	if oldTimestamp == nil {
		return 0.0 // treat unmanaged connections, e.g. with zero heartbeats, as healthy connections
	}

	timeDiff := timestamp.Sub(*oldTimestamp)

	history := oldState.history
	mean := history.mean()
	stdDeviation := fd.ensureValidStdDeviation(history.stdDeviation())

	return fd.phi(float64(timeDiff.Milliseconds()), mean+float64(fd.acceptableHeartbeatPauseMS), stdDeviation)
}

func (fd *AccrualFailureDetector) ensureValidStdDeviation(stdDeviation float64) float64 {
	return math.Max(stdDeviation, float64(fd.minStdDeviationMS))
}

func (fd *AccrualFailureDetector) phi(timeDiff, mean, stdDeviation float64) float64 {
	y := (timeDiff - mean) / stdDeviation
	e := math.Exp(-y * (1.5976 + 0.070566*y*y))

	if timeDiff > mean {
		return -math.Log10(e / (1.0 + e))
	}

	return -math.Log10(1.0 - 1.0/(1.0+e))
}
