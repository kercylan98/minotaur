package survey

import (
	"github.com/tidwall/gjson"
	"strings"
	"sync"
)

// Analyzer 分析器
type Analyzer struct {
	v      map[string]any      // 记录了每个 key 的当前值
	vc     map[string]int64    // 记录了每个 key 生效的计数数量
	repeat map[string]struct{} // 去重信息
	subs   map[string]*Analyzer
	format map[string]func(v any) any // 格式化函数
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

// SetFormat 设置格式化函数
func (slf *Analyzer) SetFormat(key string, format func(v any) any) {
	slf.m.Lock()
	defer slf.m.Unlock()
	if slf.format == nil {
		slf.format = make(map[string]func(v any) any)
	}
	slf.format[key] = format
}

// SetValueIfGreaterThan 设置指定 key 的值，当新值大于旧值时
//   - 当已有值不为 float64 时，将会被忽略
func (slf *Analyzer) SetValueIfGreaterThan(key string, value float64) {
	slf.m.Lock()
	defer slf.m.Unlock()
	if slf.v == nil {
		slf.v = make(map[string]any)
		slf.vc = make(map[string]int64)
	}
	v, exist := slf.v[key]
	if !exist {
		slf.v[key] = value
		slf.vc[key]++
		return
	}
	switch v := v.(type) {
	case float64:
		if v < value {
			slf.v[key] = value
			slf.vc[key]++
		}
	}
}

// SetValueIfLessThan 设置指定 key 的值，当新值小于旧值时
//   - 当已有值不为 float64 时，将会被忽略
func (slf *Analyzer) SetValueIfLessThan(key string, value float64) {
	slf.m.Lock()
	defer slf.m.Unlock()
	if slf.v == nil {
		slf.v = make(map[string]any)
		slf.vc = make(map[string]int64)
	}
	v, exist := slf.v[key]
	if !exist {
		slf.v[key] = value
		slf.vc[key]++
		return
	}
	switch v := v.(type) {
	case float64:
		if v > value {
			slf.v[key] = value
			slf.vc[key]++
		}
	}
}

// SetValueIf 当表达式满足的时候将设置指定 key 的值为 value
func (slf *Analyzer) SetValueIf(key string, expression bool, value float64) {
	if !expression {
		return
	}
	slf.m.Lock()
	defer slf.m.Unlock()
	slf.v[key] = value
	slf.vc[key]++
}

// SetValueStringIf 当表达式满足的时候将设置指定 key 的值为 value
func (slf *Analyzer) SetValueStringIf(key string, expression bool, value string) {
	if !expression {
		return
	}
	slf.m.Lock()
	defer slf.m.Unlock()
	slf.v[key] = value
	slf.vc[key]++
}

// SetValue 设置指定 key 的值
func (slf *Analyzer) SetValue(key string, value float64) {
	slf.m.Lock()
	defer slf.m.Unlock()
	if slf.v == nil {
		slf.v = make(map[string]any)
		slf.vc = make(map[string]int64)
	}
	slf.v[key] = value
	slf.vc[key]++
}

// SetValueString 设置指定 key 的值
func (slf *Analyzer) SetValueString(key string, value string) {
	slf.m.Lock()
	defer slf.m.Unlock()
	if slf.v == nil {
		slf.v = make(map[string]any)
		slf.vc = make(map[string]int64)
	}
	slf.v[key] = value
	slf.vc[key]++
}

// Increase 在指定 key 现有值的基础上增加 recordKey 的值
//   - 当分析器已经记录过相同 key 的值时，会根据已有的值类型进行不同处理
//
// 处理方式：
//   - 当已有值类型为 string 时，将会使用新的值的 string 类型进行覆盖
//   - 当已有值类型为 float64 时，当新的值类型不为 float64 时，将会被忽略
func (slf *Analyzer) Increase(key string, record R, recordKey string) {
	slf.m.Lock()
	defer slf.m.Unlock()
	if !record.Exist(recordKey) {
		return
	}
	if slf.v == nil {
		slf.v = make(map[string]any)
		slf.vc = make(map[string]int64)
	}
	value, exist := slf.v[key]
	if !exist {
		result := gjson.Get(string(record), recordKey)
		switch result.Type {
		case gjson.String:
			slf.v[key] = result.String()
		case gjson.Number:
			slf.v[key] = result.Float()
		default:
			return
		}
		slf.vc[key]++
		return
	}
	switch v := value.(type) {
	case string:
		slf.v[key] = record.GetString(recordKey)
	case float64:
		slf.v[key] = v + record.GetFloat64(recordKey)
	default:
		return
	}
	slf.vc[key]++
}

// IncreaseValue 在指定 key 现有值的基础上增加 value
func (slf *Analyzer) IncreaseValue(key string, value float64) {
	slf.m.Lock()
	defer slf.m.Unlock()
	if slf.v == nil {
		slf.v = make(map[string]any)
		slf.vc = make(map[string]int64)
	}
	v, exist := slf.v[key]
	if !exist {
		slf.v[key] = value
		slf.vc[key]++
		return
	}
	switch v := v.(type) {
	case float64:
		slf.v[key] = v + value
		slf.vc[key]++
	}
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

// GetValue 获取当前记录的值
func (slf *Analyzer) GetValue(key string) float64 {
	slf.m.Lock()
	defer slf.m.Unlock()
	value, exist := slf.v[key]
	if !exist {
		return 0
	}
	switch v := value.(type) {
	case float64:
		return v
	}
	return 0
}

// GetValueString 获取当前记录的值
func (slf *Analyzer) GetValueString(key string) string {
	slf.m.Lock()
	defer slf.m.Unlock()
	value, exist := slf.v[key]
	if !exist {
		return ""
	}
	switch v := value.(type) {
	case string:
		return v
	}
	return ""
}
