package ecs

import "github.com/kercylan98/minotaur/toolkit/charproc"

type dynamicBitSetKey = string

func newDynamicBitSet() *dynamicBitSet {
	return &dynamicBitSet{bits: make([]uint64, 1)}
}

// dynamicBitSet 动态位集
type dynamicBitSet struct {
	bits []uint64
	key  string
}

// Key 返回 dynamicBitSet 表示的键
func (db *dynamicBitSet) Key() dynamicBitSetKey {
	if db.key != charproc.None {
		return db.key
	}
	var bytes = make([]byte, len(db.bits)*8)
	for i, v := range db.bits {
		bytes[i*8+0] = byte(v >> 0)
		bytes[i*8+1] = byte(v >> 8)
		bytes[i*8+2] = byte(v >> 16)
		bytes[i*8+3] = byte(v >> 24)
		bytes[i*8+4] = byte(v >> 32)
		bytes[i*8+5] = byte(v >> 40)
		bytes[i*8+6] = byte(v >> 48)
		bytes[i*8+7] = byte(v >> 56)
	}
	return dynamicBitSetKey(bytes)
}

// Copy 复制 dynamicBitSet
func (db *dynamicBitSet) Copy() *dynamicBitSet {
	bits := make([]uint64, len(db.bits))
	copy(bits, db.bits)
	return &dynamicBitSet{bits: bits}
}

// Set 设置给定位置的位为 1
func (db *dynamicBitSet) Set(pos ComponentId) {
	index, offset := pos/64, pos%64
	for len(db.bits) <= int(index) {
		db.bits = append(db.bits, 0)
	}
	db.bits[index] |= 1 << offset
	db.key = charproc.None
}

// Bits 返回位集
func (db *dynamicBitSet) Bits() []ComponentId {
	var bits = make([]ComponentId, 0, len(db.bits)*64)
	for i, v := range db.bits {
		for j := 0; j < 64; j++ {
			if v&(1<<j) != 0 {
				bits = append(bits, ComponentId(i*64+j))
			}
		}
	}
	return bits
}

// Clear 清除给定位置的位
func (db *dynamicBitSet) Clear(pos ComponentId) {
	index, offset := pos/64, pos%64
	if len(db.bits) > int(index) {
		db.bits[index] &^= 1 << offset
	}
	db.key = charproc.None
}

// IsSet 检查给定位置的位是否被设置
func (db *dynamicBitSet) IsSet(pos ComponentId) bool {
	index, offset := pos/64, pos%64
	if len(db.bits) <= int(index) {
		return false
	}
	return db.bits[index]&(1<<offset) != 0
}

// Equal 比较两个 dynamicBitSet 是否相等
func (db *dynamicBitSet) Equal(other dynamicBitSet) bool {
	// 如果 key 相同则直接返回 true
	if db.key != charproc.None && other.key != charproc.None && db.key == other.key {
		return true
	}

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

func parseDynamicBitSetKey(key dynamicBitSetKey) *dynamicBitSet {
	bits := make([]uint64, len(key)/8)
	for i := 0; i < len(bits); i++ {
		bits[i] = uint64(key[i*8+0]) |
			uint64(key[i*8+1])<<8 |
			uint64(key[i*8+2])<<16 |
			uint64(key[i*8+3])<<24 |
			uint64(key[i*8+4])<<32 |
			uint64(key[i*8+5])<<40 |
			uint64(key[i*8+6])<<48 |
			uint64(key[i*8+7])<<56
	}
	return &dynamicBitSet{bits: bits, key: key}
}
