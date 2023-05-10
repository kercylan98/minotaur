package builtin

import (
	"github.com/kercylan98/minotaur/game"
	"github.com/kercylan98/minotaur/utils/huge"
)

type ItemContainerOption[ItemID comparable, Item game.Item[ItemID]] func(container *ItemContainer[ItemID, Item])

// WithItemContainerSizeLimit 通过特定的物品容器非堆叠数量上限创建物品容器
func WithItemContainerSizeLimit[ItemID comparable, Item game.Item[ItemID]](sizeLimit int) ItemContainerOption[ItemID, Item] {
	return func(container *ItemContainer[ItemID, Item]) {
		if sizeLimit <= 0 {
			return
		}
		container.sizeLimit = sizeLimit
	}
}

// WithItemContainerStackLimit 通过设置特定物品堆叠数量创建容器
func WithItemContainerStackLimit[ItemID comparable, Item game.Item[ItemID]](id ItemID, stackLimit *huge.Int) ItemContainerOption[ItemID, Item] {
	return func(container *ItemContainer[ItemID, Item]) {
		if stackLimit.LessThanOrEqualTo(huge.IntZero) {
			return
		}
		if container.stackLimit == nil {
			container.stackLimit = map[ItemID]*huge.Int{}
		}
		container.stackLimit[id] = stackLimit
	}
}
