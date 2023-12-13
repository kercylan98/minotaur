package super

import (
	"github.com/kercylan98/minotaur/utils/generic"
	"math/bits"
	"strconv"
)

// BitMask 使用泛型类型 Bit 表示一个比特掩码，Bit 必须是整数类型。
type BitMask[Bit generic.Integer] uint64

// Mask 创建一个新的 BitMask，包含给定的比特位。
// 参数 bits 是一个可变参数列表，表示要设置的比特位。
//   - 使用按位或运算符 (|=) 将 bit 位置的比特设置为 1。
func Mask[Bit generic.Integer](bits ...Bit) BitMask[Bit] {
	var mask Bit
	for _, bit := range bits {
		mask |= 1 << bit
	}
	return BitMask[Bit](mask)
}

// Matches 检查当前 BitMask 是否与另一个 BitMask 完全匹配。
func (slf *BitMask[Bit]) Matches(bits BitMask[Bit]) bool {
	return *slf == bits
}

// Contains 检查当前 BitMask 是否包含另一个 BitMask。
//   - 使用按位与运算符 (&) 计算两个 BitMask 的公共部分，然后与 bits 进行比较。
func (slf *BitMask[Bit]) Contains(bits BitMask[Bit]) bool {
	return (*slf & bits) == bits
}

// Has 检查当前 BitMask 是否包含特定的比特位。
//   - 使用按位与运算符 (&) 检查 bit 位置的比特是否为 1。
func (slf *BitMask[Bit]) Has(bits ...Bit) bool {
	for _, bit := range bits {
		if *slf&(1<<bit) == 0 {
			return false
		}
	}
	return true
}

// Add 将给定的比特位添加到当前 BitMask 中。
//   - 使用按位或运算符 (|=) 将 bit 位置的比特设置为 1。
func (slf *BitMask[Bit]) Add(bit Bit) {
	*slf |= 1 << bit
}

// Del 从当前 BitMask 中删除给定的比特位。
//   - 使用按位与运算符 (&) 取反 (^) 然后按位与运算符 (&) 清除 bit 位置的比特。
func (slf *BitMask[Bit]) Del(bit Bit) {
	*slf &= ^(1 << bit)
}

// Toggle 反转指定比特位。
//   - 使用按位或运算符 (|=) 将 bit 位置的比特设置为 1，如果 bit 位置的比特已经为 1，则使用按位与运算符 (&^=) 将 bit 位置的比特设置为 0。
func (slf *BitMask[Bit]) Toggle(bit Bit) {
	*slf ^= 1 << bit
}

// Reset 将当前 BitMask 重置为 all 0。
func (slf *BitMask[Bit]) Reset() {
	*slf = 0
}

// Count 统计 BitMask 中被设置的比特位数量。
//   - 使用按位与运算符 (&) 将 BitMask 与 0xFFFFFFFF 进行按位与运算，然后使用 `bits.Len()` 函数统计 0 的数量。
func (slf *BitMask[Bit]) Count() int {
	return 64 - bits.Len64(uint64(*slf)&0xFFFFFFFF)
}

// Union 返回当前 BitMask 和另一个 BitMask 的并集。
//   - 使用按位或运算符 (&) 将两个 BitMask 进行合并。
func (slf *BitMask[Bit]) Union(other BitMask[Bit]) BitMask[Bit] {
	return *slf | other
}

// Intersection 返回当前 BitMask 和另一个 BitMask 的交集。
//   - 使用按位与运算符 (&) 将两个 BitMask 进行交集。
func (slf *BitMask[Bit]) Intersection(other BitMask[Bit]) BitMask[Bit] {
	return *slf & other
}

// Difference 返回当前 BitMask 和另一个 BitMask 的差集。
//   - 使用按位异或运算符 (^) 将两个 BitMask 进行异或。
func (slf *BitMask[Bit]) Difference(other BitMask[Bit]) BitMask[Bit] {
	return *slf ^ other
}

// String 返回当前 BitMask 的字符串表示形式。
//   - 使用 `strconv.FormatUint()` 函数将 BitMask 转换为二进制字符串。
func (slf *BitMask[Bit]) String() string {
	return strconv.FormatUint(uint64(*slf), 2)
}
