package builtin

import "github.com/kercylan98/minotaur/game"

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
