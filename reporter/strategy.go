package reporter

import (
	"fmt"
	"github.com/kercylan98/minotaur/utils/timer"
	"time"
)

type ReportStrategy[Data any] func(buried *Buried[Data])

// ReportStrategyInstantly 命中时将立即上报
func ReportStrategyInstantly[Data any]() ReportStrategy[Data] {
	return func(buried *Buried[Data]) {
		oldHitHandle := buried.hitHandle
		buried.hitHandle = func(actions *HitOperationSet[Data]) {
			oldHitHandle(actions)
			ticker.After(fmt.Sprintf("ReportStrategyInstantly_%s", buried.GetName()), timer.Instantly, func() {
				if disableBuried.Get(buried.GetName()) {
					return
				}
				if err := buried.Reporting(); err != nil {
					if buried.errorHandle != nil {
						buried.errorHandle(buried, err)
					}
				}
			})
		}
	}
}

// ReportStrategyAfter 命中后一段时间后上报
//   - 该模式下如果连续在时间范围内命中，将仅上报最后一次结果
func ReportStrategyAfter[Data any](t time.Duration) ReportStrategy[Data] {
	return func(buried *Buried[Data]) {
		oldHitHandle := buried.hitHandle
		buried.hitHandle = func(actions *HitOperationSet[Data]) {
			oldHitHandle(actions)
			ticker.After(fmt.Sprintf("ReportStrategyAfter_%s_%d", buried.GetName(), t.Milliseconds()), t, func() {
				if disableBuried.Get(buried.GetName()) {
					return
				}
				if err := buried.Reporting(); err != nil {
					if buried.errorHandle != nil {
						buried.errorHandle(buried, err)
					}
				}
			})
		}
	}
}

// ReportStrategyLoop 循环上报
func ReportStrategyLoop[Data any](t time.Duration) ReportStrategy[Data] {
	return func(buried *Buried[Data]) {
		ticker.Loop(fmt.Sprintf("ReportStrategyLoop_%s_%d", buried.GetName(), t.Milliseconds()), timer.Instantly, t, timer.Forever, func() {
			if disableBuried.Get(buried.GetName()) {
				return
			}
			if err := buried.Reporting(); err != nil {
				if buried.errorHandle != nil {
					buried.errorHandle(buried, err)
				}
			}
		})
	}
}

// ReportStrategyFixedTime 将在每天的固定时间上报
func ReportStrategyFixedTime[Data any](hour, min, sec int) ReportStrategy[Data] {
	return func(buried *Buried[Data]) {
		now := time.Now()
		current := now.Unix()
		next := time.Date(now.Year(), now.Month(), now.Day(), hour, min, sec, 0, now.Location())
		target := next.Unix()
		if current >= target {
			next = next.AddDate(0, 0, 1)
			target = next.Unix()
		}
		ticker.Loop(fmt.Sprintf("ReportStrategyFixedTime_%s_%d_%d_%d", buried.GetName(), hour, min, sec), time.Duration(target-current)*time.Second, 24*time.Hour, -1, func() {
			if disableBuried.Get(buried.GetName()) {
				return
			}
			if err := buried.Reporting(); err != nil {
				if buried.errorHandle != nil {
					buried.errorHandle(buried, err)
				}
			}
		})
	}
}
