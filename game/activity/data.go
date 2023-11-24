package activity

import (
	"sync"
	"time"
)

// DataMeta 全局活动数据
type DataMeta[Data any] struct {
	once       sync.Once
	Data       Data      `json:"data,omitempty"`       // 活动数据
	LastNewDay time.Time `json:"lastNewDay,omitempty"` // 上次跨天时间
}

// EntityDataMeta 活动实体数据
type EntityDataMeta[Data any] struct {
	once       sync.Once
	Data       Data      `json:"data,omitempty"`       // 活动数据
	LastNewDay time.Time `json:"lastNewDay,omitempty"` // 上次跨天时间
}
