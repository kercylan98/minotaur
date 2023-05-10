package game

import "github.com/kercylan98/minotaur/utils/huge"

// ItemContainer 物品容器
type ItemContainer[ItemID comparable, I Item[ItemID]] interface {
	// GetSize 获取容器物品非堆叠数量
	GetSize() int
	// GetSizeLimit 获取容器物品非堆叠数量上限
	GetSizeLimit() int
	// SetExpandSize 设置拓展非堆叠数量上限
	SetExpandSize(size int)

	// GetItem 获取物品
	GetItem(guid int64) (ItemContainerMember[ItemID, I], error)
	// GetItems 获取所有物品
	GetItems() []ItemContainerMember[ItemID, I]
	// GetItemsMap 获取所有物品
	GetItemsMap() map[int64]ItemContainerMember[ItemID, I]
	// ExistItem 物品是否存在
	ExistItem(guid int64) bool
	// ExistItemWithID 是否存在特定ID的物品
	ExistItemWithID(id ItemID) bool

	// AddItem 添加物品
	AddItem(item I, count *huge.Int) error
	// DeductItem 扣除特定物品数量，当数量为0将被移除，数量不足时将不进行任何改变
	DeductItem(guid int64, count *huge.Int) error
}
