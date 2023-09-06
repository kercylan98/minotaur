package survey

import (
	"github.com/kercylan98/minotaur/utils/super"
)

func newReport(analyzer *Analyzer) *Report {
	report := &Report{
		analyzer: analyzer,
		Name:     "ROOT",
		Values:   analyzer.v,
		Subs:     make([]*Report, 0, len(analyzer.subs)),
	}
	for k, v := range analyzer.subs {
		sub := newReport(v)
		sub.Name = k
		report.Subs = append(report.Subs, sub)
	}
	return report
}

// Report 分析报告
type Report struct {
	analyzer *Analyzer
	Name     string             // 报告名称（默认为 ROOT）
	Values   map[string]float64 `json:"Values,omitempty"`
	Subs     []*Report          `json:"Reports,omitempty"`
}

// ReserveSub 仅保留特定名称子报告
func (slf *Report) ReserveSub(names ...string) *Report {
	report := newReport(slf.analyzer)
	var newSub []*Report
	for _, sub := range slf.Subs {
		var exist bool
		for _, name := range names {
			if sub.Name == name {
				exist = true
				break
			}
		}
		if exist {
			newSub = append(newSub, sub)
		}
	}
	report.Subs = newSub
	return report
}

// FilterSub 过滤特定名称的子报告
func (slf *Report) FilterSub(names ...string) *Report {
	report := newReport(slf.analyzer)
	var newSub []*Report
	for _, sub := range slf.Subs {
		var exist bool
		for _, name := range names {
			if sub.Name == name {
				exist = true
				break
			}
		}
		if !exist {
			newSub = append(newSub, sub)
		}
	}
	report.Subs = newSub
	return report
}

func (slf *Report) String() string {
	return string(super.MarshalIndentJSON(slf, "", "  "))
}
