package report

import "github.com/kercylan98/minotaur/utils/timer"

func NewReporter(reportHandle func() error, options ...ReporterOption) *Reporter {
	reporter := &Reporter{
		reportHandle: reportHandle,
	}
	for _, option := range options {
		option(reporter)
	}
	if reporter.ticker == nil {
		reporter.ticker = timer.GetTicker(50)
	}
	for _, strategy := range reporter.strategies {
		strategy(reporter)
	}
	return reporter
}

// Reporter 数据上报器
type Reporter struct {
	ticker       *timer.Ticker
	strategies   []ReporterStrategy
	reportHandle func() error
	errorHandle  func(reporter *Reporter, err error)
}

// Report 上报
func (slf *Reporter) Report() error {
	return slf.reportHandle()
}
