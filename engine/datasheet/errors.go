package datasheet

import "errors"

var (
	ErrorIndexStart              = errors.New("index must start at 1")      // 索引必须从 1 开始
	ErrorIndexContinuous         = errors.New("index must be continuous")   // 索引必须连续
	ErrorNotSupportDatasheetType = errors.New("not support datasheet type") // 不支持的数据表类型
)
