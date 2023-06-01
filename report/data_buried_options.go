package report

import "github.com/kercylan98/minotaur/utils/synchronization"

type DataBuriedOption[DataID comparable, Data any] func(buried *DataBuried[DataID, Data])

// WithDataBuriedSync 通过并发安全的方式创建数据埋点
func WithDataBuriedSync[DataId comparable, Data any]() DataBuriedOption[DataId, Data] {
	return func(buried *DataBuried[DataId, Data]) {
		buried.data = synchronization.NewMap[DataId, Data]()
	}
}
