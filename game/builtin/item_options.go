package builtin

type ItemOption[ItemID comparable] func(item *Item[ItemID])

// WithItemGuid 通过特定的guid创建物品
func WithItemGuid[ItemID comparable](guid int64) ItemOption[ItemID] {
	return func(item *Item[ItemID]) {
		item.guid = guid
	}
}
