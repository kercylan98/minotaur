package storage

// GlobalDataStorage 全局数据存储器接口
type GlobalDataStorage[T any] interface {
	// Load 加载全局数据
	//   - 当全局数据不存在时，应当返回新的全局数据实例
	Load(name string) T
	// Save 保存全局数据
	Save(name string, data T)
}

// LoadGlobalData 加载全局数据
func LoadGlobalData[T any](globalData *GlobalData[T], storage GlobalDataStorage[T]) *GlobalData[T] {
	globalData.data = storage.Load(globalData.GetName())
	return globalData
}

// SaveGlobalData 保存全局数据
func SaveGlobalData[T any](globalData *GlobalData[T], storage GlobalDataStorage[T]) *GlobalData[T] {
	storage.Save(globalData.GetName(), globalData.GetData())
	return globalData
}
