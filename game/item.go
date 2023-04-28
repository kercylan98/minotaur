package game

// Item 物品接口定义
type Item[ID comparable] interface {
	// GetID 获取物品id
	GetID() ID
	// GetCount 获取物品数量
	GetCount() int
}
