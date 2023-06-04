package matrix

type Match3Option[ItemType comparable, Item Match3Item[ItemType]] func(match3 *Match3[ItemType, Item])

func WithMatch3Generator[ItemType comparable, Item Match3Item[ItemType]](itemType ItemType, generator func() Item) Match3Option[ItemType, Item] {
	return func(match3 *Match3[ItemType, Item]) {
		match3.generators[itemType] = generator
	}
}

// WithMatch3Tactics 设置匹配策略
//   - 匹配策略用于匹配出对应成员
func WithMatch3Tactics[ItemType comparable, Item Match3Item[ItemType]](tactics func(matrix [][]Item) [][]Item) Match3Option[ItemType, Item] {
	return func(match3 *Match3[ItemType, Item]) {
		match3.matchStrategy[len(match3.matchStrategy)+1] = tactics
	}
}
