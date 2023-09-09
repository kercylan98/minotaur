package survey

import (
	"strings"
	"sync"
)

// Analyzer 分析器
type Analyzer struct {
	v      map[string]float64  // 记录了每个 key 的当前值
	vc     map[string]int64    // 记录了每个 key 生效的计数数量
	repeat map[string]struct{} // 去重信息
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
		slf.vc = make(map[string]int64)
	}
	slf.v[key] += record.GetFloat64(recordKey)
	slf.vc[key]++
}

// IncreaseValue 在指定 key 现有值的基础上增加 value
func (slf *Analyzer) IncreaseValue(key string, value float64) {
	slf.m.Lock()
	defer slf.m.Unlock()
	if slf.v == nil {
		slf.v = make(map[string]float64)
		slf.vc = make(map[string]int64)
	}
	slf.v[key] += value
	slf.vc[key]++
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
	dk := strings.Join(append([]string{key}, dvs...), "_")
	if _, e := slf.repeat[dk]; e {
		slf.m.Unlock()
		return
	}
	slf.repeat[dk] = struct{}{}
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
	dk := strings.Join(append([]string{key}, dvs...), "_")
	if _, e := slf.repeat[dk]; e {
		slf.m.Unlock()
		return
	}
	slf.repeat[dk] = struct{}{}
	slf.m.Unlock()
	slf.IncreaseValue(key, value)
}
