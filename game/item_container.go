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
	GetItem(id ItemID) (ItemContainerMember[ItemID], error)
	// GetItemWithGuid 根据guid获取物品
	GetItemWithGuid(id ItemID, guid int64) (ItemContainerMember[ItemID], error)
	// GetItems 获取所有物品
	GetItems() []ItemContainerMember[ItemID]

	// AddItem 添加物品
	//  - 当物品guid相同时，如果相同物品id及guid的堆叠数量未达到上限，将增加数量，否则增加非堆叠数量
	//  - 当物品guid不同时，堆叠将不可用，每次都将增加非堆叠数量
	AddItem(item I, count *huge.Int) error
	// DeductItem 扣除特定物品数量，当数量为0将被移除，数量不足时将不进行任何改变
	//  - 将查找特定id的物品，无论guid是否相同，都有可能被扣除直到达到扣除数量
	//  - 当count为负数时，由于负负得正，无论guid是否相同，都有可能被增加物品数量直到达到扣除数量
	DeductItem(id ItemID, count *huge.Int) error
	// DeductItemWithGuid 更为精准的扣除特定物品数量，可参考 DeductItem
	DeductItemWithGuid(id ItemID, guid int64, count *huge.Int) error
}
