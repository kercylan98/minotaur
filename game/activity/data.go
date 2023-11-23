package activity

import "time"

// DataMeta 全局活动数据
type DataMeta[Data any] struct {
	Start      time.Time `json:"start,omitempty"`      // 活动开始时间
	End        time.Time `json:"end,omitempty"`        // 活动结束时间
	Data       Data      `json:"data,omitempty"`       // 活动数据
	LastNewDay time.Time `json:"lastNewDay,omitempty"` // 上次跨天时间
}

// EntityDataMeta 活动实体数据
type EntityDataMeta[Data any] struct {
	Start      time.Time `json:"start,omitempty"`      // 对象参与活动时间
	End        time.Time `json:"end,omitempty"`        // 对象结束活动时间
	Data       Data      `json:"data,omitempty"`       // 活动数据
	LastNewDay time.Time `json:"lastNewDay,omitempty"` // 上次跨天时间
}
