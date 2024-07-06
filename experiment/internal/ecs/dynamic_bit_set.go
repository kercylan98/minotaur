package ecs

// DynamicBitSet 动态位集
type DynamicBitSet struct {
	bits []uint64
}

// Set 设置给定位置的位为 1
func (db *DynamicBitSet) Set(pos ComponentId) {
	index, offset := pos/64, pos%64
	for len(db.bits) <= int(index) {
		db.bits = append(db.bits, 0)
	}
	db.bits[index] |= 1 << offset
}

// Clear 清除给定位置的位
func (db *DynamicBitSet) Clear(pos ComponentId) {
	index, offset := pos/64, pos%64
	if len(db.bits) > int(index) {
		db.bits[index] &^= 1 << offset
	}
}

// IsSet 检查给定位置的位是否被设置
func (db *DynamicBitSet) IsSet(pos ComponentId) bool {
	index, offset := pos/64, pos%64
	if len(db.bits) <= int(index) {
		return false
	}
	return db.bits[index]&(1<<offset) != 0
}

// Equal 比较两个 DynamicBitSet 是否相等
func (db *DynamicBitSet) Equal(other DynamicBitSet) bool {
	// 先比较长度
	if len(db.bits) != len(other.bits) {
		return false
	}

	// 逐位比较
	for i := range db.bits {
		if db.bits[i] != other.bits[i] {
			return false
		}
	}
	return true
}
