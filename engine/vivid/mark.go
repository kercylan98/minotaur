package vivid

const (
	// MarkMessageImmutability 消息不可变性
	//
	// 在 actor 模型中，设计的核心原则是消息的不可变性，每个 actor 接收的消息应该是独立且不可修改的，
	// 当传递指针消息时可能导致竞态访问问题，例如将自身状态中的一个 map 发送给其他本地 Actor 时，当自身与目标 Actor 在同一时间段内操作该 map，将导致并发问题。
	MarkMessageImmutability = iota

	// MarkNonStrictConcurrencySafety 非严格并发安全
	//
	// 意味着该函数或者数据结构在正常情况下访问将会是并发安全的，当以其他如匿名函数的形式被调用时，可能会导致并发问题。
	//
	// 通常情况下，访问标准作用域内的数据结构都是并发安全的，但当它们被传递给其他函数时，可能会导致并发问题。
	MarkNonStrictConcurrencySafety

	// MarkConcurrentSafety 并发安全
	//
	// 意味着该函数或者数据结构的访问将会是并发安全的，并且不会出现并发问题。
	MarkConcurrentSafety
)
