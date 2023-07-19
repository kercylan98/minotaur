package storage

import "github.com/kercylan98/minotaur/utils/generic"

// IndexDataStorage 全局数据存储器接口
type IndexDataStorage[I generic.Ordered, T any] interface {
	// Load 加载特定索引数据
	//   - 通常情况下当数据不存在时，应当返回空指针
	Load(name string, index I) T
	// LoadAll 加载所有数据
	LoadAll(name string) map[I]T
	// Save 保存特定索引数据
	Save(name string, index I, data T)
	// SaveAll 保存所有数据
	SaveAll(name string, data map[I]T)
	// Delete 删除特定索引数据
	Delete(name string, index I)
	// DeleteAll 删除所有数据
	DeleteAll(name string)
}

// LoadIndexData 加载索引数据
func LoadIndexData[I generic.Ordered, T any](indexData *IndexData[I, T], index I, storage IndexDataStorage[I, T]) *IndexData[I, T] {
	indexData.data[index] = storage.Load(indexData.GetName(), index)
	return indexData
}

// LoadIndexDataAll 加载所有索引数据
func LoadIndexDataAll[I generic.Ordered, T any](indexData *IndexData[I, T], storage IndexDataStorage[I, T]) *IndexData[I, T] {
	indexData.data = storage.LoadAll(indexData.GetName())
	return indexData
}

// SaveIndexData 保存索引数据
func SaveIndexData[I generic.Ordered, T any](indexData *IndexData[I, T], index I, storage IndexDataStorage[I, T]) *IndexData[I, T] {
	storage.Save(indexData.GetName(), index, indexData.GetData(index))
	return indexData
}

// SaveIndexDataALl 保存所有所索引数据
func SaveIndexDataALl[I generic.Ordered, T any](indexData *IndexData[I, T], storage IndexDataStorage[I, T]) *IndexData[I, T] {
	storage.SaveAll(indexData.GetName(), indexData.GetAllData())
	return indexData
}

// DeleteIndexData 删除索引数据
func DeleteIndexData[I generic.Ordered, T any](indexData *IndexData[I, T], index I, storage IndexDataStorage[I, T]) *IndexData[I, T] {
	storage.Delete(indexData.GetName(), index)
	delete(indexData.data, index)
	return indexData
}

// DeleteIndexDataAll 删除所有索引数据
func DeleteIndexDataAll[I generic.Ordered, T any](indexData *IndexData[I, T], storage IndexDataStorage[I, T]) *IndexData[I, T] {
	storage.DeleteAll(indexData.GetName())
	for k := range indexData.data {
		delete(indexData.data, k)
	}
	return indexData
}
