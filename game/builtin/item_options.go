package builtin

import "github.com/kercylan98/minotaur/utils/huge"

type ItemOption[ID comparable] func(item *Item[ID])

// WithItemStackLimit 通过特定堆叠上限创建物品
//   - 默认无限制
func WithItemStackLimit[ID comparable](stackLimit *huge.Int) ItemOption[ID] {
	return func(item *Item[ID]) {
		if stackLimit == nil || stackLimit.LessThanOrEqualTo(huge.IntZero) {
			return
		}
		item.stackLimit = stackLimit
	}
}
