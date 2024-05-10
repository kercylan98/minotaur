package mask

import (
	"github.com/kercylan98/minotaur/toolkit"
	"github.com/kercylan98/minotaur/toolkit/ident"
	"strconv"
	"strings"
	"unsafe"
)

// DynamicMask 可变且不限制大小的掩码，它的底层实现是一个 uint64 的切片
//   - 虽然无界掩码是无限，但是内存和位大小是受限的
//   - 越大的掩码会导致越多的内存占用
type DynamicMask struct {
	data      unsafe.Pointer
	length    uint64
	bitsCache []uint64
}

// scaling 伸缩掩码的大小
func (m *DynamicMask) scaling(newSize uint64) {
	if newSize <= m.length {
		return
	}

	newSlice := make([]uint64, newSize)
	if m.data != nil {
		copy(newSlice, (*[1 << 30]uint64)(m.data)[:m.length])
	}

	m.data = unsafe.Pointer(&newSlice[0])
	m.length = newSize
}

// Set 设置指定位置的位
func (m *DynamicMask) Set(bit uint64) {
	index := bit >> 6
	offset := bit & 63

	if index >= m.length {
		m.scaling((index + 1) << 1)
	}

	*(*uint64)(unsafe.Pointer(uintptr(m.data) + uintptr(index)*8)) |= 1 << uint(offset)
	m.bitsCache = nil
}

// Del 删除指定位置的位
func (m *DynamicMask) Del(bit uint64) {
	index := bit >> 6
	offset := bit & 63

	if index >= m.length {
		return
	}

	*(*uint64)(unsafe.Pointer(uintptr(m.data) + uintptr(index)*8)) &= ^(1 << uint(offset))
	m.bitsCache = nil
}

// Has 检查指定位是否存在
func (m *DynamicMask) Has(bit uint64) bool {
	index := bit >> 6
	offset := bit & 63

	if index >= m.length {
		return false
	}

	return *(*uint64)(unsafe.Pointer(uintptr(m.data) + uintptr(index)*8))&(1<<uint(offset)) != 0
}

// HasAll 检查所有位是否均存在
func (m *DynamicMask) HasAll(bits []uint64) bool {
	if len(bits) == 0 {
		return m.length == 0
	}
	for _, bit := range bits {
		if !m.Has(bit) {
			return false
		}
	}
	return true
}

// HasAny 检查是否有任意一个位存在
func (m *DynamicMask) HasAny(bits []uint64) bool {
	for _, bit := range bits {
		if m.Has(bit) {
			return true
		}
	}
	return false
}

// EqualBits 检查是否匹配所有位
func (m *DynamicMask) EqualBits(bits []uint64) bool {
	if m.length == 0 && len(bits) == 0 {
		return true
	}
	m.Bits()
	if len(m.bitsCache) != len(bits) {
		return false
	}
	for i, bit := range bits {
		if m.bitsCache[i] != bit {
			return false
		}
	}
	return true
}

// Equal 检查是否匹配另一个掩码
func (m *DynamicMask) Equal(other DynamicMask) bool {
	if m.length != other.length {
		return false
	}

	mSlice := (*[1 << 30]uint64)(m.data)[:m.length]
	oSlice := (*[1 << 30]uint64)(other.data)[:other.length]
	if m.length != other.length {
		return false
	}

	for i := uint64(0); i < m.length; i++ {
		if mSlice[i] != oSlice[i] {
			return false
		}
	}

	return true
}

// Bits 返回所有位的值
func (m *DynamicMask) Bits() []uint64 {
	if m.bitsCache != nil {
		return m.bitsCache
	}
	var bits = make([]uint64, 0, m.length<<6)
	for i := uint64(0); i < m.length; i++ {
		mask := *(*uint64)(unsafe.Pointer(uintptr(m.data) + uintptr(i)<<3))
		for j := uint64(0); j < 64; j++ {
			if mask&(1<<uint(j)) != 0 {
				bits = append(bits, i<<6+j)
			}
		}
	}
	m.bitsCache = bits
	return bits
}

// Clear 清空所有位
func (m *DynamicMask) Clear() {
	m.data = nil
	m.length = 0
	m.bitsCache = nil
}

// Clone 克隆一个掩码
func (m *DynamicMask) Clone() DynamicMask {
	var clone DynamicMask
	clone.scaling(m.length)
	copy((*[1 << 30]uint64)(clone.data)[:m.length], (*[1 << 30]uint64)(m.data)[:m.length])
	return clone
}

// String 返回掩码的适合阅读的字符串表示，例如 [3] 1, 2, 3
func (m *DynamicMask) String() string {
	bits := m.Bits()
	var builder strings.Builder
	builder.WriteByte('[')
	builder.WriteString(strconv.Itoa(len(bits)))
	builder.WriteByte(']')
	builder.WriteByte(' ')
	for i, bit := range bits {
		builder.WriteString(strconv.Itoa(int(bit)))
		if i == len(bits)-1 {
			break
		}
		builder.WriteString(", ")
	}
	return builder.String()
}

// StringIdentical 返回掩码的唯一字符串表示，该字符串不具备可读性，但是是可以用于唯一标识的高性能字符串
func (m *DynamicMask) StringIdentical() string {
	slice := (*[1 << 30]uint64)(m.data)[:m.length]
	return ident.GenerateOrderedUniqueIdentStringWithUInt64(slice...)
}

func (m *DynamicMask) MarshalJSON() ([]byte, error) {
	slice := (*[1 << 30]uint64)(m.data)[:m.length]
	return toolkit.MarshalJSONE(slice)
}

func (m *DynamicMask) UnmarshalJSON(data []byte) error {
	var slice []uint64
	err := toolkit.UnmarshalJSONE(data, &slice)
	if err != nil {
		return err
	}
	m.scaling(uint64(len(slice)))
	copy((*[1 << 30]uint64)(m.data)[:m.length], slice)
	return nil
}
