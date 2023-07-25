package storage

import "github.com/kercylan98/minotaur/utils/generic"

type Storage[PrimaryKey generic.Ordered, Body any] interface {
	// Load 加载数据
	Load(index PrimaryKey) (Body, error)

	// Save 保存数据
	Save(set *Set[PrimaryKey, Body], index PrimaryKey, data any) error
}
