package column

import (
	"github.com/kercylan98/minotaur/toolkit/collection/listings"
)

func newColumn(pageSize int, defaultGetter func() any) *column {
	return &column{
		data:          listings.NewPagedSlice[any](pageSize),
		defaultGetter: defaultGetter,
	}
}

type column struct {
	data          *listings.PagedSlice[any] // 列数据
	defaultGetter func() any                // 默认值获取器
}
