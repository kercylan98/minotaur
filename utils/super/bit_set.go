package super

import (
	"fmt"
	"github.com/kercylan98/minotaur/utils/generic"
	"math/bits"
)

// NewBitSet 通过指定的 Bit 位创建一个 BitSet
//   - 当指定的 Bit 位存在负数时，将会 panic
func NewBitSet[Bit generic.Integer](bits ...Bit) *BitSet[Bit] {
	set := &BitSet[Bit]{set: make([]uint64, 0, 1)}
	for _, bit := range bits {
		if bit < 0 {
			panic(fmt.Errorf("bit %v is negative", bit))
		}
		set.Set(bit)
	}
	return set
}

// BitSet 是一个可以动态增长的比特位集合
//   - 默认情况下将使用 64 位无符号整数来表示比特位，当需要表示的比特位超过 64 位时，将自动增长
type BitSet[Bit generic.Integer] struct {
	set []uint64 // 比特位集合
}

// Set 将指定的位 bit 设置为 1
func (slf *BitSet[Bit]) Set(bit Bit) *BitSet[Bit] {
	word := bit >> 6
	for word >= Bit(len(slf.set)) {
		slf.set = append(slf.set, 0)
	}
	slf.set[word] |= 1 << (bit & 0x3f)
	return slf
}

// Del 将指定的位 bit 设置为 0
func (slf *BitSet[Bit]) Del(bit Bit) *BitSet[Bit] {
	word := bit >> 6
	if word < Bit(len(slf.set)) {
		slf.set[word] &^= 1 << (bit & 0x3f)
	}
	return slf
}

// Shrink 将 BitSet 中的比特位集合缩小到最小
//   - 正常情况下当 BitSet 中的比特位超出 64 位时，将自动增长，当 BitSet 中的比特位数量减少时，可以使用该方法将 BitSet 中的比特位集合缩小到最小
func (slf *BitSet[Bit]) Shrink() *BitSet[Bit] {
	switch len(slf.set) {
	case 0:
		return slf
	case 1:
		if slf.set[0] == 0 {
			slf.set = nil
			return slf
		}
	}
	index := len(slf.set) - 1
	if slf.set[index] != 0 {
		return slf
	}

	for i := index - 1; i >= 0; i-- {
		if slf.set[i] != 0 {
			slf.set = slf.set[:i+1]
			return slf
		}
	}
	return slf
}

// Cap 返回当前 BitSet 中可以表示的最大比特位数量
func (slf *BitSet[Bit]) Cap() int {
	return len(slf.set) * 64
}

// Has 检查指定的位 bit 是否被设置为 1
func (slf *BitSet[Bit]) Has(bit Bit) bool {
	word := bit >> 6
	return word < Bit(len(slf.set)) && slf.set[word]&(1<<(bit&0x3f)) != 0
}

// Clear 清空所有的比特位
func (slf *BitSet[Bit]) Clear() *BitSet[Bit] {
	slf.set = nil
	return slf
}

// Len 返回当前 BitSet 中被设置的比特位数量
func (slf *BitSet[Bit]) Len() int {
	var count int
	for _, word := range slf.set {
		count += bits.OnesCount64(word)
	}
	return count
}

// Bits 返回当前 BitSet 中被设置的比特位
func (slf *BitSet[Bit]) Bits() []Bit {
	bits := make([]Bit, 0, slf.Len())
	for i, word := range slf.set {
		for j := 0; j < 64; j++ {
			if word&(1<<j) != 0 {
				bits = append(bits, Bit(i*64+j))
			}
		}
	}
	return bits
}

// Reverse 反转当前 BitSet 中的所有比特位
func (slf *BitSet[Bit]) Reverse() *BitSet[Bit] {
	for i, word := range slf.set {
		slf.set[i] = bits.Reverse64(word)
	}
	return slf
}

// Not 返回当前 BitSet 中所有比特位的反转
func (slf *BitSet[Bit]) Not() *BitSet[Bit] {
	for i, word := range slf.set {
		slf.set[i] = ^word
	}
	return slf
}

// And 将当前 BitSet 与另一个 BitSet 进行按位与运算
func (slf *BitSet[Bit]) And(other *BitSet[Bit]) *BitSet[Bit] {
	for i, word := range other.set {
		if i < len(slf.set) {
			slf.set[i] &= word
		} else {
			slf.set = append(slf.set, word)
		}
	}
	return slf
}

// Or 将当前 BitSet 与另一个 BitSet 进行按位或运算
func (slf *BitSet[Bit]) Or(other *BitSet[Bit]) *BitSet[Bit] {
	for i, word := range other.set {
		if i < len(slf.set) {
			slf.set[i] |= word
		} else {
			slf.set = append(slf.set, word)
		}
	}
	return slf
}

// Xor 将当前 BitSet 与另一个 BitSet 进行按位异或运算
func (slf *BitSet[Bit]) Xor(other *BitSet[Bit]) *BitSet[Bit] {
	for i, word := range other.set {
		if i < len(slf.set) {
			slf.set[i] ^= word
		} else {
			slf.set = append(slf.set, word)
		}
	}
	return slf
}

// Sub 将当前 BitSet 与另一个 BitSet 进行按位减运算
func (slf *BitSet[Bit]) Sub(other *BitSet[Bit]) *BitSet[Bit] {
	for i, word := range other.set {
		if i < len(slf.set) {
			slf.set[i] &^= word
		}
	}
	return slf
}

// IsZero 检查当前 BitSet 是否为空
func (slf *BitSet[Bit]) IsZero() bool {
	for _, word := range slf.set {
		if word != 0 {
			return false
		}
	}
	return true
}

// Clone 返回当前 BitSet 的副本
func (slf *BitSet[Bit]) Clone() *BitSet[Bit] {
	other := &BitSet[Bit]{set: make([]uint64, len(slf.set))}
	copy(other.set, slf.set)
	return other
}

// Equal 检查当前 BitSet 是否与另一个 BitSet 相等
func (slf *BitSet[Bit]) Equal(other BitSet[Bit]) bool {
	if len(slf.set) != len(other.set) {
		return false
	}
	for i, word := range slf.set {
		if word != other.set[i] {
			return false
		}
	}
	return true
}

// Contains 检查当前 BitSet 是否包含另一个 BitSet
func (slf *BitSet[Bit]) Contains(other BitSet[Bit]) bool {
	for i, word := range other.set {
		if i >= len(slf.set) || slf.set[i]&word != word {
			return false
		}
	}
	return true
}

// ContainsAny 检查当前 BitSet 是否包含另一个 BitSet 中的任意比特位
func (slf *BitSet[Bit]) ContainsAny(other BitSet[Bit]) bool {
	for i, word := range other.set {
		if i < len(slf.set) && slf.set[i]&word != 0 {
			return true
		}
	}
	return false
}

// ContainsAll 检查当前 BitSet 是否包含另一个 BitSet 中的所有比特位
func (slf *BitSet[Bit]) ContainsAll(other BitSet[Bit]) bool {
	for i, word := range other.set {
		if i >= len(slf.set) || slf.set[i]&word != word {
			return false
		}
	}
	return true
}

// Intersect 检查当前 BitSet 是否与另一个 BitSet 有交集
func (slf *BitSet[Bit]) Intersect(other BitSet[Bit]) bool {
	for i, word := range other.set {
		if i < len(slf.set) && slf.set[i]&word != 0 {
			return true
		}
	}
	return false
}

// Union 检查当前 BitSet 是否与另一个 BitSet 有并集
func (slf *BitSet[Bit]) Union(other *BitSet[Bit]) bool {
	for i, word := range other.set {
		if i < len(slf.set) && slf.set[i]&word != 0 {
			return true
		}
	}
	return false
}

// Difference 检查当前 BitSet 是否与另一个 BitSet 有差集
func (slf *BitSet[Bit]) Difference(other *BitSet[Bit]) bool {
	for i, word := range other.set {
		if i < len(slf.set) && slf.set[i]&word != 0 {
			return true
		}
	}
	return false
}

// SymmetricDifference 检查当前 BitSet 是否与另一个 BitSet 有对称差集
func (slf *BitSet[Bit]) SymmetricDifference(other *BitSet[Bit]) bool {
	for i, word := range other.set {
		if i < len(slf.set) && slf.set[i]&word != 0 {
			return true
		}
	}
	return false
}

// Subset 检查当前 BitSet 是否为另一个 BitSet 的子集
func (slf *BitSet[Bit]) Subset(other *BitSet[Bit]) bool {
	for i, word := range other.set {
		if i >= len(slf.set) || slf.set[i]&word != word {
			return false
		}
	}
	return true
}

// Superset 检查当前 BitSet 是否为另一个 BitSet 的超集
func (slf *BitSet[Bit]) Superset(other *BitSet[Bit]) bool {
	for i, word := range slf.set {
		if i >= len(other.set) || other.set[i]&word != word {
			return false
		}
	}
	return true
}

// Complement 检查当前 BitSet 是否为另一个 BitSet 的补集
func (slf *BitSet[Bit]) Complement(other *BitSet[Bit]) bool {
	for i, word := range slf.set {
		if i >= len(other.set) || other.set[i]&word != word {
			return false
		}
	}
	return true
}

// Max 返回当前 BitSet 中最大的比特位
func (slf *BitSet[Bit]) Max() Bit {
	for i := len(slf.set) - 1; i >= 0; i-- {
		if slf.set[i] != 0 {
			return Bit(i*64 + bits.Len64(slf.set[i]) - 1)
		}
	}
	return 0
}

// Min 返回当前 BitSet 中最小的比特位
func (slf *BitSet[Bit]) Min() Bit {
	for i, word := range slf.set {
		if word != 0 {
			return Bit(i*64 + bits.TrailingZeros64(word))
		}
	}
	return 0
}

// String 返回当前 BitSet 的字符串表示
func (slf *BitSet[Bit]) String() string {
	return fmt.Sprintf("[%v] %v", slf.Len(), slf.Bits())
}

// MarshalJSON 实现 json.Marshaler 接口
func (slf *BitSet[Bit]) MarshalJSON() ([]byte, error) {
	return MarshalJSONE(slf.set)
}

// UnmarshalJSON 实现 json.Unmarshaler 接口
func (slf *BitSet[Bit]) UnmarshalJSON(data []byte) error {
	return UnmarshalJSON(data, &slf.set)
}
