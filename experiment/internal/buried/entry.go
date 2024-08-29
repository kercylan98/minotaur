package buried

import "time"

// Hit 命中条目
func Hit(eventName string) Entries {
	return Entries{m: make(map[string]any)}
}

// Entries 条目集合
type Entries struct {
	m map[string]any
}

func (es Entries) String(key string, value string) Entries {
	es.m[key] = value
	return es
}

func (es Entries) Int(key string, value int) Entries {
	es.m[key] = value
	return es
}

func (es Entries) Int8(key string, value int8) Entries {
	es.m[key] = value
	return es
}

func (es Entries) Int16(key string, value int16) Entries {
	es.m[key] = value
	return es
}

func (es Entries) Int32(key string, value int32) Entries {
	es.m[key] = value
	return es
}

func (es Entries) Int64(key string, value int64) Entries {
	es.m[key] = value
	return es
}

func (es Entries) Uint(key string, value uint) Entries {
	es.m[key] = value
	return es
}

func (es Entries) Uint8(key string, value uint8) Entries {
	es.m[key] = value
	return es
}

func (es Entries) Uint16(key string, value uint16) Entries {
	es.m[key] = value
	return es
}

func (es Entries) Uint32(key string, value uint32) Entries {
	es.m[key] = value
	return es
}

func (es Entries) Uint64(key string, value uint64) Entries {
	es.m[key] = value
	return es
}

func (es Entries) Float32(key string, value float32) Entries {
	es.m[key] = value
	return es
}

func (es Entries) Float64(key string, value float64) Entries {
	es.m[key] = value
	return es
}

func (es Entries) Bool(key string, value bool) Entries {
	es.m[key] = value
	return es
}

func (es Entries) Time(key string, value time.Time) Entries {
	es.m[key] = value
	return es
}

func (es Entries) Duration(key string, value time.Duration) Entries {
	es.m[key] = value
	return es
}
