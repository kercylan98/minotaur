package arrangement

// Item 编排成员
type Item[ID comparable] interface {
	// GetID 获取成员的唯一标识
	GetID() ID
	// Equal 比较两个成员是否相等
	Equal(item Item[ID]) bool
}
