package matrix

type Match3Option[ItemType comparable, Item Match3Item[ItemType]] func(match3 *Match3[ItemType, Item])

func WithMatch3Generator[ItemType comparable, Item Match3Item[ItemType]](itemType ItemType, generator func() Item) Match3Option[ItemType, Item] {
	return func(match3 *Match3[ItemType, Item]) {
		match3.generators[itemType] = generator
	}
}
