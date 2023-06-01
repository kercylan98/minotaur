package report

import (
	"github.com/kercylan98/minotaur/utils/asynchronization"
	"github.com/kercylan98/minotaur/utils/hash"
)

// NewDataBuried 创建一个数据埋点
func NewDataBuried[DataID comparable, Data any](hitLogic HitLogic[Data], options ...DataBuriedOption[DataID, Data]) *DataBuried[DataID, Data] {
	buried := &DataBuried[DataID, Data]{
		data:     asynchronization.NewMap[DataID, Data](),
		hitLogic: hitLogic,
	}
	buried.setData = func(id DataID, data Data) {
		buried.data.Set(id, data)
	}
	buried.getData = func(id DataID) Data {
		return buried.data.Get(id)
	}
	for _, option := range options {
		option(buried)
	}
	return buried
}

// DataBuried 数据埋点
//   - 数据埋点通常用于统计不同类型的数据，例如用户数据、商城数据等
type DataBuried[DataID comparable, Data any] struct {
	data     hash.Map[DataID, Data]
	hitLogic HitLogic[Data]
	getData  func(DataID) Data
	setData  func(id DataID, data Data)
}

// Hit 命中数据埋点
func (slf *DataBuried[DataID, Data]) Hit(id DataID, data Data) {
	slf.setData(id, slf.hitLogic(slf.getData(id), data))
}

// GetData 获取数据
func (slf *DataBuried[DataID, Data]) GetData(id DataID) Data {
	return slf.getData(id)
}

// GetSize 获取已触发该埋点的id数量
func (slf *DataBuried[DataID, Data]) GetSize() int {
	return slf.data.Size()
}
