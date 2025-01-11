package marks

const (
	// ConcurrencyUnsafe 该标记表示所标记内容不是并发安全的
	//   - 在并发环境下，该内容可能会出现数据竞争
	ConcurrencyUnsafe = iota

	// ConcurrencySafe 该标记表示所标记内容是并发安全的
	//   - 在并发环境下，该内容不会出现数据竞争
	ConcurrencySafe

	// ExcludeDuplicateElements 该标记表示所标记内容不会出现重复元素
	//   - 该标记通常用于表示集合或者列表
	//   - 通常来说，发生重复的元素将会被忽略
	//
	// 更严格的发生重复应产生 panic 的标记可参考: marks.PanicOnDuplicateElements
	ExcludeDuplicateElements

	// PanicOnDuplicateElements 该标记表示所标记内容在发生重复元素时应发生 panic
	//   - 该标记通常用于表示集合或者列表
	PanicOnDuplicateElements
)
