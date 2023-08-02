package deck

// NewGroup 创建一个新的组
func NewGroup[I Item](guid int64, fillHandle func(guid int64) []I) *Group[I] {
	if fillHandle == nil {
		panic("deck.NewGroup: fillHandle is nil")
	}
	group := &Group[I]{
		guid:       guid,
		fillHandle: fillHandle,
	}
	return group
}

// Group 甲板中的组，用于针对一堆内容进行管理的数据结构
type Group[I Item] struct {
	guid       int64                // 组的 guid
	fillHandle func(guid int64) []I // 组的填充函数
	items      []I                  // 组中的所有内容
}

// GetGuid 获取组的 guid
func (slf *Group[I]) GetGuid() int64 {
	return slf.guid
}

// Fill 将该组的数据填充为 WithGroupFillHandle 中设置的内容
func (slf *Group[I]) Fill() {
	slf.items = slf.fillHandle(slf.guid)
}

// Pop 从顶部获取一个内容
func (slf *Group[I]) Pop() (item I) {
	if len(slf.items) == 0 {
		return item
	}
	item = slf.items[0]
	slf.items = slf.items[1:]
	return item
}

// PopN 从顶部获取指定数量的内容
func (slf *Group[I]) PopN(n int) (items []I) {
	if len(slf.items) == 0 {
		return items
	}
	if len(slf.items) < n {
		n = len(slf.items)
	}
	items = slf.items[:n]
	slf.items = slf.items[n:]
	return items
}

// PressOut 从底部压出一个内容
func (slf *Group[I]) PressOut() (item I) {
	if len(slf.items) == 0 {
		return item
	}
	item = slf.items[len(slf.items)-1]
	slf.items = slf.items[:len(slf.items)-1]
	return item
}

// PressOutN 从底部压出指定数量的内容
func (slf *Group[I]) PressOutN(n int) (items []I) {
	if len(slf.items) == 0 {
		return items
	}
	if len(slf.items) < n {
		n = len(slf.items)
	}
	items = slf.items[len(slf.items)-n:]
	slf.items = slf.items[:len(slf.items)-n]
	return items
}

// Push 向顶部压入一个内容
func (slf *Group[I]) Push(item I) {
	slf.items = append([]I{item}, slf.items...)
}

// PushN 向顶部压入指定数量的内容
func (slf *Group[I]) PushN(items []I) {
	slf.items = append(items, slf.items...)
}

// Insert 向底部插入一个内容
func (slf *Group[I]) Insert(item I) {
	slf.items = append(slf.items, item)
}

// InsertN 向底部插入指定数量的内容
func (slf *Group[I]) InsertN(items []I) {
	slf.items = append(slf.items, items...)
}

// Pull 从特定位置拔出一个内容
func (slf *Group[I]) Pull(index int) (item I) {
	if len(slf.items) == 0 {
		return item
	}
	item = slf.items[index]
	slf.items = append(slf.items[:index], slf.items[index+1:]...)
	return item
}

// Thrust 向特定位置插入一个内容
func (slf *Group[I]) Thrust(index int, item I) {
	if len(slf.items) == 0 {
		return
	}
	slf.items = append(slf.items[:index], append([]I{item}, slf.items[index:]...)...)
}

// IsFree 检查组是否为空
func (slf *Group[I]) IsFree() bool {
	return len(slf.items) == 0
}

// GetCount 获取组中剩余的内容数量
func (slf *Group[I]) GetCount() int {
	return len(slf.items)
}

// GetItem 获取组中的指定内容
func (slf *Group[I]) GetItem(index int) I {
	return slf.items[index]
}
