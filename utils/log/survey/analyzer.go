package survey

import (
	"strings"
	"sync"
)

// Analyzer 分析器
type Analyzer struct {
	v      map[string]float64
	repeat map[string]struct{}
	subs   map[string]*Analyzer
	m      sync.Mutex
}

// Sub 获取子分析器
func (slf *Analyzer) Sub(key string) *Analyzer {
	slf.m.Lock()
	defer slf.m.Unlock()
	if slf.subs == nil {
		slf.subs = make(map[string]*Analyzer)
	}
	sub, e := slf.subs[key]
	if !e {
		sub = &Analyzer{}
		slf.subs[key] = sub
	}
	return sub
}

// Increase 在指定 key 现有值的基础上增加 recordKey 的值
func (slf *Analyzer) Increase(key string, record R, recordKey string) {
	slf.m.Lock()
	defer slf.m.Unlock()
	if !record.Exist(recordKey) {
		return
	}
	if slf.v == nil {
		slf.v = make(map[string]float64)
	}
	v, e := slf.v[key]
	if !e {
		slf.v[key] = record.GetFloat64(recordKey)
		return
	}
	slf.v[key] = v + record.GetFloat64(recordKey)
}

// IncreaseValue 在指定 key 现有值的基础上增加 value
func (slf *Analyzer) IncreaseValue(key string, value float64) {
	slf.m.Lock()
	defer slf.m.Unlock()
	if slf.v == nil {
		slf.v = make(map[string]float64)
	}
	slf.v[key] += value
}

// IncreaseNonRepeat 在指定 key 现有值的基础上增加 recordKey 的值，但是当去重维度 dimension 相同时，不会增加
func (slf *Analyzer) IncreaseNonRepeat(key string, record R, recordKey string, dimension ...string) {
	slf.m.Lock()
	if !record.Exist(recordKey) {
		slf.m.Unlock()
		return
	}
	if slf.repeat == nil {
		slf.repeat = make(map[string]struct{})
	}
	dvs := make([]string, 0, len(dimension))
	for _, v := range dimension {
		dvs = append(dvs, record.GetString(v))
	}
	dk := strings.Join(dvs, "_")
	if _, e := slf.repeat[dk]; e {
		slf.m.Unlock()
		return
	}
	slf.m.Unlock()
	slf.Increase(key, record, recordKey)
}

// IncreaseValueNonRepeat 在指定 key 现有值的基础上增加 value，但是当去重维度 dimension 相同时，不会增加
func (slf *Analyzer) IncreaseValueNonRepeat(key string, record R, value float64, dimension ...string) {
	slf.m.Lock()
	if slf.repeat == nil {
		slf.repeat = make(map[string]struct{})
	}
	dvs := make([]string, 0, len(dimension))
	for _, v := range dimension {
		dvs = append(dvs, record.GetString(v))
	}
	dk := strings.Join(dvs, "_")
	if _, e := slf.repeat[dk]; e {
		slf.m.Unlock()
		return
	}
	slf.m.Unlock()
	slf.IncreaseValue(key, value)
}
