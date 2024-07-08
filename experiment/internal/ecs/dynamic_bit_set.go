package ecs

import (
	"github.com/kercylan98/minotaur/toolkit/charproc"
	"github.com/kercylan98/minotaur/toolkit/convert"
	"strings"
)

type DynamicBitSetKey = string

func newDynamicBitSet() *DynamicBitSet {
	return &DynamicBitSet{bits: make([]uint64, 1)}
}

// DynamicBitSet 动态位集
type DynamicBitSet struct {
	bits []uint64
	key  string
}

// Key 返回 DynamicBitSet 表示的键
func (db *DynamicBitSet) Key() DynamicBitSetKey {
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
	return DynamicBitSetKey(bytes)
}

// Copy 复制 DynamicBitSet
func (db *DynamicBitSet) Copy() *DynamicBitSet {
	bits := make([]uint64, len(db.bits))
	copy(bits, db.bits)
	return &DynamicBitSet{bits: bits}
}

// Set 设置给定位置的位为 1
func (db *DynamicBitSet) Set(pos ComponentId) {
	index, offset := pos/64, pos%64
	for len(db.bits) <= int(index) {
		db.bits = append(db.bits, 0)
	}
	db.bits[index] |= 1 << offset
	db.key = charproc.None
}

// Bits 返回位集
func (db *DynamicBitSet) Bits() []ComponentId {
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
func (db *DynamicBitSet) Clear(pos ComponentId) {
	index, offset := pos/64, pos%64
	if len(db.bits) > int(index) {
		db.bits[index] &^= 1 << offset
	}
	db.key = charproc.None
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
func (db *DynamicBitSet) Equal(other *DynamicBitSet) bool {
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

// In 检查 DynamicBitSet 是否包含另一个 DynamicBitSet
func (db *DynamicBitSet) In(mask *DynamicBitSet) bool {
	// 如果 key 相同则直接返回 true
	if db.key != "" && mask.key != "" && db.key == mask.key {
		return true
	}

	// 逐位比较
	for i := range mask.bits {
		// 如果 db 的 bits 长度不够，直接返回 false
		if i >= len(db.bits) {
			return false
		}

		// 逐位检查
		if db.bits[i]&mask.bits[i] != mask.bits[i] {
			return false
		}
	}

	return true
}

func (db *DynamicBitSet) NotIn(mask *DynamicBitSet) bool {
	// 逐位比较
	for i := range mask.bits {
		if i >= len(db.bits) {
			continue
		}

		// 如果 mask 中的位在 db 中存在且设置，则返回 false
		if db.bits[i]&mask.bits[i] != 0 {
			return false
		}
	}
	return true
}

func (db *DynamicBitSet) String() string {
	var builder strings.Builder
	builder.WriteString("{")
	bits := db.Bits()
	for i, id := range bits {
		builder.WriteString(convert.Uint32ToString(id))
		if i < len(bits)-1 {
			builder.WriteString(", ")
		}
	}
	builder.WriteString("}")

	return builder.String()
}

func parseDynamicBitSetKey(key DynamicBitSetKey) *DynamicBitSet {
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
	return &DynamicBitSet{bits: bits, key: key}
}
