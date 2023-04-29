package game

import "minotaur/utils/huge"

// ItemContainer 物品容器
type ItemContainer[ItemID comparable, I Item[ItemID]] interface {
	// GetItem 根据guid获取物品
	GetItem(guid int64) I
	// GetItems 获取容器中的所有物品
	GetItems() map[int64]I
	// GetItemsWithId 根据id获取容器中所有物品
	GetItemsWithId(id ItemID) map[int64]I
	// AddItem 添加物品到容器中
	AddItem(item I) error
	// ChangeItemCount 改变容器中特定数量的物品（扣除时当数量不足时会尝试扣除相同ID的物品）
	ChangeItemCount(guid int64, count *huge.Int) error
	// DeleteItem 删除容器中的物品
	DeleteItem(guid int64)
	// DeleteItemsWithId 删除容器中所有特定id的物品
	DeleteItemsWithId(id ItemID)
}
