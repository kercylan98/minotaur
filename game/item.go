package game

type Item[ID comparable] interface {
	// GetID 获取物品ID
	GetID() ID
	// IsSame 与另一个物品比较是否相同
	IsSame(item Item[ID]) bool
}
