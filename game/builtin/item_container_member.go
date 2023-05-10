package builtin

import (
	"github.com/kercylan98/minotaur/game"
	"github.com/kercylan98/minotaur/utils/huge"
)

type ItemContainerMember[ItemID comparable] struct {
	game.Item[ItemID]
	sort     int
	count    *huge.Int
	bakCount *huge.Int
}

func (slf *ItemContainerMember[ItemID]) GetCount() *huge.Int {
	return slf.count.Copy()
}
