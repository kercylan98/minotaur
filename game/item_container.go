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
	// GetItems 获取所有非空物品
	//  - 物品顺序为容器内顺序
	//  - 空的容器空间将被忽略
	GetItems() []ItemContainerMember[ItemID, I]
	// GetItemsFull 获取所有物品
	//  - 物品顺序为容器内顺序
	//  - 空的容器空间将被设置为nil
	//  - 当容器非堆叠物品上限为0时，最后一个非空物品之后的所有空物品都将被忽略
	//  - 当容器非堆叠物品未达到上限时，其余空间将使用nil填充
	GetItemsFull() []ItemContainerMember[ItemID, I]
	// GetItemsMap 获取所有物品
	GetItemsMap() map[int64]ItemContainerMember[ItemID, I]
	// ExistItem 物品是否存在
	ExistItem(guid int64) bool
	// ExistItemWithID 是否存在特定ID的物品
	ExistItemWithID(id ItemID) bool
	// AddItem 添加物品
	AddItem(item I, count *huge.Int) (guid int64, err error)
	// DeductItem 扣除特定物品数量，当数量为0将被移除，数量不足时将不进行任何改变
	DeductItem(guid int64, count *huge.Int) error
	// TransferTo 转移特定物品到另一个容器中
	TransferTo(guid int64, count *huge.Int, target ItemContainer[ItemID, I]) error
	// CheckAllowAdd 检查是否允许添加特定物品
	CheckAllowAdd(item I, count *huge.Int) error
	// CheckDeductItem 检查是否允许扣除特定物品
	CheckDeductItem(guid int64, count *huge.Int) error
	// Remove 移除特定guid的物品
	Remove(guid int64)
	// RemoveWithID 移除所有物品ID匹配的物品
	RemoveWithID(id ItemID)
	// Clear 清空物品容器
	Clear()
}
