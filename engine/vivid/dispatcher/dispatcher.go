package dispatcher

// Dispatcher 是用于分发邮箱消息的接口，不同的实现可以用于不同的场景
//   - 通常分发器应该作为单例提供，当无法满足时应做出说明
type Dispatcher interface {
	Dispatch(f func())
}
