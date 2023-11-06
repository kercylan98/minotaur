package survey

import (
	"github.com/kercylan98/minotaur/utils/super"
	"strings"
)

func newReport(analyzer *Analyzer) *Report {
	report := &Report{
		analyzer: analyzer,
		Name:     "ROOT",
		Values:   analyzer.v,
		Counter:  analyzer.vc,
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
	Name     string           // 报告名称（默认为 ROOT）
	Values   map[string]any   `json:"Values,omitempty"`
	Counter  map[string]int64 `json:"Count,omitempty"`
	Subs     []*Report        `json:"Reports,omitempty"`
}

// Avg 计算平均值
func (slf *Report) Avg(key string) float64 {
	value, exist := slf.Values[key]
	if !exist {
		return 0
	}
	valF, ok := value.(float64)
	if !ok {
		return 0
	}
	return valF / float64(slf.Counter[key])
}

// Count 获取特定 key 的计数次数
func (slf *Report) Count(key string) int64 {
	return slf.Counter[key]
}

// Sum 获取特定 key 的总和
func (slf *Report) Sum(keys ...string) float64 {
	var sum float64
	for _, key := range keys {
		value, exist := slf.Values[key]
		if !exist {
			continue
		}
		valF, ok := value.(float64)
		if !ok {
			continue
		}
		sum += valF
	}
	return sum
}

// Sub 获取特定名称的子报告
func (slf *Report) Sub(name string) *Report {
	for _, sub := range slf.Subs {
		if sub.Name == name {
			return sub
		}
	}
	return nil
}

// ReserveSubByPrefix 仅保留特定前缀的子报告
func (slf *Report) ReserveSubByPrefix(prefix string) *Report {
	report := newReport(slf.analyzer)
	var newSub []*Report
	for _, sub := range slf.Subs {
		if strings.HasPrefix(sub.Name, prefix) {
			newSub = append(newSub, sub)
		}
	}
	report.Subs = newSub
	return report
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

// FilterSub 将特定名称的子报告过滤掉
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
