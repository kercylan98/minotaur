package game

import "minotaur/utils/huge"

// Item 物品接口定义
type Item[ID comparable] interface {
	// GetID 获取物品id
	GetID() ID
	// GetGuid 获取物品guid
	GetGuid() int64
	// SetGuid 设置物品guid
	SetGuid(guid int64)
	// ChangeStackCount 改变物品堆叠数量，返回新数量
	ChangeStackCount(count *huge.Int) *huge.Int
	// GetStackCount 获取物品堆叠数量
	GetStackCount() *huge.Int
}
