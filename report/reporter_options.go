package report

import "github.com/kercylan98/minotaur/utils/timer"

type ReporterOption func(reporter *Reporter)

// WithReporterTicker 通过特定的定时器创建上报器
func WithReporterTicker(ticker *timer.Ticker) ReporterOption {
	return func(reporter *Reporter) {
		reporter.ticker = ticker
	}
}

// WithReporterStrategies 通过特定上报策略进行创建
func WithReporterStrategies(strategies ...ReporterStrategy) ReporterOption {
	return func(reporter *Reporter) {
		reporter.strategies = append(reporter.strategies, strategies...)
	}
}

func WithReporterErrorHandle(errorHandle func(reporter *Reporter, err error)) ReporterOption {
	return func(reporter *Reporter) {
		reporter.errorHandle = errorHandle
	}
}
