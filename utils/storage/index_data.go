package storage

import (
	"github.com/kercylan98/minotaur/utils/generic"
)

// NewIndexData 创建索引数据
func NewIndexData[I generic.Ordered, T any](name string, storage IndexDataStorage[I, T]) *IndexData[I, T] {
	data := &IndexData[I, T]{
		name: name,
		data: map[I]T{},
	}
	return data
}

// IndexData 全局数据
type IndexData[I generic.Ordered, T any] struct {
	storage GlobalDataStorage[T] // 存储器
	name    string               // 数据组名称
	data    map[I]T              // 数据
}

// GetName 获取名称
func (slf *IndexData[I, T]) GetName() string {
	return slf.name
}

// GetData 获取数据
func (slf *IndexData[I, T]) GetData(index I) T {
	return slf.data[index]
}

// GetAllData 获取所有数据
func (slf *IndexData[I, T]) GetAllData() map[I]T {
	return slf.data
}

// LoadData 加载数据
func (slf *IndexData[I, T]) LoadData(index I, storage IndexDataStorage[I, T]) *IndexData[I, T] {
	return LoadIndexData(slf, index, storage)
}

// SaveData 保存数据
func (slf *IndexData[I, T]) SaveData(storage IndexDataStorage[I, T], index I) *IndexData[I, T] {
	return SaveIndexData(slf, index, storage)
}

// Handle 处理数据
func (slf *IndexData[I, T]) Handle(index I, handler func(name string, index I, data T)) *IndexData[I, T] {
	handler(slf.GetName(), index, slf.GetData(index))
	return slf
}

// HandleWithCallback 处理数据
func (slf *IndexData[I, T]) HandleWithCallback(index I, handler func(name string, index I, data T) error, callback func(err error)) *IndexData[I, T] {
	callback(handler(slf.GetName(), index, slf.GetData(index)))
	return slf
}
