package game

import "github.com/kercylan98/minotaur/utils/huge"

// Item 物品
type Item[ID comparable] interface {
	// GetID 获取物品id
	GetID() ID
	// GetCount 获取物品数量
	GetCount() *huge.Int
	// GetStackLimit 获取堆叠上限
	GetStackLimit() *huge.Int
	// SetCount 设置物品数量
	SetCount(count *huge.Int)
	// ChangeCount 改变物品数量
	ChangeCount(count *huge.Int) error
}
