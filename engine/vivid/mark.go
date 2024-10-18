package vivid

const (
	// MarkMessageImmutability 消息不可变性
	//
	// 在 actor 模型中，设计的核心原则是消息的不可变性，每个 actor 接收的消息应该是独立且不可修改的，
	// 当传递指针消息时可能导致竞态访问问题，例如将自身状态中的一个 map 发送给其他本地 Actor 时，当自身与目标 Actor 在同一时间段内操作该 map，将导致并发问题。
	MarkMessageImmutability = iota
)
