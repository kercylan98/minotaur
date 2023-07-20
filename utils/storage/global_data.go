package storage

// NewGlobalData 创建全局数据
func NewGlobalData[T any](name string, storage GlobalDataStorage[T]) (*GlobalData[T], error) {
	data := &GlobalData[T]{
		storage: storage,
		name:    name,
	}
	var err error
	data.data, err = storage.Load(name)
	if err != nil {
		return nil, err
	}
	return data, nil
}

// GlobalData 全局数据
type GlobalData[T any] struct {
	storage GlobalDataStorage[T] // 全局数据存储器
	name    string               // 全局数据名称
	data    T                    // 数据
}

// GetName 获取名称
func (slf *GlobalData[T]) GetName() string {
	return slf.name
}

// GetData 获取数据
func (slf *GlobalData[T]) GetData() T {
	return slf.data
}

// LoadData 加载数据
func (slf *GlobalData[T]) LoadData() error {
	data, err := slf.storage.Load(slf.GetName())
	if err != nil {
		return err
	}
	slf.data = data
	return nil
}

// SaveData 保存数据
func (slf *GlobalData[T]) SaveData() error {
	return slf.storage.Save(slf.GetName(), slf.GetData())
}

// Handle 处理数据
func (slf *GlobalData[T]) Handle(handler func(name string, data T)) *GlobalData[T] {
	handler(slf.GetName(), slf.GetData())
	return slf
}

// HandleWithCallback 处理数据
func (slf *GlobalData[T]) HandleWithCallback(handler func(name string, data T) error, callback func(err error)) *GlobalData[T] {
	callback(handler(slf.GetName(), slf.GetData()))
	return slf
}
