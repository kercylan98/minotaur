package builtin

import (
	"github.com/kercylan98/minotaur/game"
	"github.com/kercylan98/minotaur/utils/huge"
)

func NewItemContainerMember[ItemID comparable, I game.Item[ItemID]](guid int64, item I) *ItemContainerMember[ItemID, I] {
	return &ItemContainerMember[ItemID, I]{
		item: item,
		guid: guid,
	}
}

type ItemContainerMember[ItemID comparable, I game.Item[ItemID]] struct {
	item     I
	guid     int64
	sort     int
	count    *huge.Int
	bakCount *huge.Int
}

func (slf *ItemContainerMember[ItemID, I]) GetID() ItemID {
	return slf.item.GetID()
}

func (slf *ItemContainerMember[ItemID, I]) GetGUID() int64 {
	return slf.guid
}

func (slf *ItemContainerMember[ItemID, I]) GetCount() *huge.Int {
	return slf.count.Copy()
}

func (slf *ItemContainerMember[ItemID, I]) GetItem() I {
	return slf.item
}
