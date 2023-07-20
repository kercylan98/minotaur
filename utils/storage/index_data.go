package storage

import (
	"github.com/kercylan98/minotaur/utils/generic"
	"time"
)

// NewIndexData 创建索引数据
func NewIndexData[I generic.Ordered, T IndexDataItem[I]](name string, storage IndexDataStorage[I, T]) (*IndexData[I, T], error) {
	data := &IndexData[I, T]{
		storage: storage,
		name:    name,
	}
	var err error
	data.data, err = storage.LoadAll(name)
	if err != nil {
		return nil, err
	}
	return data, nil
}

// IndexData 全局数据
type IndexData[I generic.Ordered, T IndexDataItem[I]] struct {
	storage IndexDataStorage[I, T] // 存储器
	name    string                 // 数据组名称
	data    map[I]T                // 数据
}

// GetName 获取名称
func (slf *IndexData[I, T]) GetName() string {
	return slf.name
}

// GetData 获取数据
func (slf *IndexData[I, T]) GetData(index I) T {
	data, _ := slf.data[index]
	return data
}

// GetAllData 获取所有数据
func (slf *IndexData[I, T]) GetAllData() map[I]T {
	return slf.data
}

// LoadData 加载数据
func (slf *IndexData[I, T]) LoadData(index I) error {
	data, err := slf.storage.Load(slf.GetName(), index)
	if err != nil {
		return err
	}
	slf.data[index] = data
	return nil
}

// LoadAllData 加载所有数据
func (slf *IndexData[I, T]) LoadAllData() error {
	data, err := slf.storage.LoadAll(slf.GetName())
	if err != nil {
		return err
	}
	slf.data = data
	return nil
}

// SaveData 保存数据
func (slf *IndexData[I, T]) SaveData(index I) error {
	return slf.storage.Save(slf.GetName(), index, slf.GetData(index))
}

// SaveAllData 保存所有数据
//   - errHandle 错误处理中如果返回 false 将重试，否则跳过当前保存下一个
func (slf *IndexData[I, T]) SaveAllData(errHandle func(err error) bool, retryInterval time.Duration) {
	slf.storage.SaveAll(slf.GetName(), slf.GetAllData(), errHandle, retryInterval)
}

// DeleteData 删除数据
func (slf *IndexData[I, T]) DeleteData(index I) *IndexData[I, T] {
	slf.storage.Delete(slf.GetName(), index)
	delete(slf.data, index)
	return slf
}

// DeleteAllData 删除所有数据
func (slf *IndexData[I, T]) DeleteAllData() *IndexData[I, T] {
	slf.storage.DeleteAll(slf.GetName())
	for k := range slf.data {
		delete(slf.data, k)
	}
	return slf
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
