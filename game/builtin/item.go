package builtin

func NewItem[ItemID comparable](id ItemID, options ...ItemOption[ItemID]) *Item[ItemID] {
	item := &Item[ItemID]{
		id: id,
	}
	for _, option := range options {
		option(item)
	}
	return item
}

type Item[ItemID comparable] struct {
	id   ItemID
	guid int64
}

func (slf *Item[ItemID]) GetID() ItemID {
	return slf.id
}

func (slf *Item[ItemID]) GetGUID() int64 {
	return slf.guid
}
