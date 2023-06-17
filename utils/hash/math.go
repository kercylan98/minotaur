package hash

import "github.com/kercylan98/minotaur/utils/generic"

// Sum 计算一个 map 中的 value 总和
func Sum[K comparable, V generic.Number](m map[K]V) V {
	var sum V
	for _, v := range m {
		sum += v
	}
	return sum
}
