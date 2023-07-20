package storage

import (
	"github.com/kercylan98/minotaur/utils/generic"
)

// IndexDataStorage 全局数据存储器接口
type IndexDataStorage[I generic.Ordered, T IndexDataItem[I]] interface {
	// Load 加载特定索引数据
	//   - 通常情况下当数据不存在时，应当返回空指针
	Load(name string, index I) (T, error)
	// LoadAll 加载所有数据
	LoadAll(name string) (map[I]T, error)
	// Save 保存特定索引数据
	Save(name string, index I, data T) error
	// SaveAll 保存所有数据
	SaveAll(name string, data map[I]T) error
	// Delete 删除特定索引数据
	Delete(name string, index I)
	// DeleteAll 删除所有数据
	DeleteAll(name string)
}
