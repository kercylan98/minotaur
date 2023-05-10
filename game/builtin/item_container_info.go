package builtin

import (
	"github.com/kercylan98/minotaur/game"
	"github.com/kercylan98/minotaur/utils/huge"
)

type itemContainerInfo[ItemID comparable, Item game.Item[ItemID]] struct {
	item     Item
	count    *huge.Int
	bakCount *huge.Int
}
