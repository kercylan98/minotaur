package report

import (
	"fmt"
	"github.com/kercylan98/minotaur/utils/timer"
	"time"
)

// ReporterStrategy 上报器策略
type ReporterStrategy func(reporter *Reporter)

// ReportStrategyLoop 循环上报
//   - 将在创建后上报一次，并且在每隔一段时间后继续上报
func ReportStrategyLoop(t time.Duration) ReporterStrategy {
	return func(reporter *Reporter) {
		reporter.ticker.Loop(fmt.Sprintf("ReportStrategyLoop_%d", t.Milliseconds()), timer.Instantly, t, timer.Forever, func() {
			if err := reporter.Report(); err != nil && reporter.errorHandle != nil {
				reporter.errorHandle(reporter, err)
			}
		})
	}
}

// ReportStrategyFixedTime 将在每天的固定时间上报
func ReportStrategyFixedTime(hour, min, sec int) ReporterStrategy {
	return func(reporter *Reporter) {
		now := time.Now()
		current := now.Unix()
		next := time.Date(now.Year(), now.Month(), now.Day(), hour, min, sec, 0, now.Location())
		target := next.Unix()
		if current >= target {
			next = next.AddDate(0, 0, 1)
			target = next.Unix()
		}
		reporter.ticker.Loop(fmt.Sprintf("ReportStrategyFixedTime_%d_%d_%d", hour, min, sec), time.Duration(target-current)*time.Second, 24*time.Hour, -1, func() {
			if err := reporter.Report(); err != nil && reporter.errorHandle != nil {
				reporter.errorHandle(reporter, err)
			}
		})
	}
}
