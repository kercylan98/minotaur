package hash

// Map 提供了map集合接口
type Map[Key comparable, Value any] interface {
	Set(key Key, value Value)
	Get(key Key) Value
	// AtomGetSet 原子方式获取一个值并在之后进行赋值
	AtomGetSet(key Key, handle func(value Value, exist bool) (newValue Value, isSet bool))
	// Atom 原子操作
	Atom(handle func(m Map[Key, Value]))
	Exist(key Key) bool
	GetExist(key Key) (Value, bool)
	Delete(key Key)
	DeleteGet(key Key) Value
	DeleteGetExist(key Key) (Value, bool)
	DeleteExist(key Key) bool
	Clear()
	ClearHandle(handle func(key Key, value Value))
	Range(handle func(key Key, value Value))
	RangeSkip(handle func(key Key, value Value) bool)
	RangeBreakout(handle func(key Key, value Value) bool)
	RangeFree(handle func(key Key, value Value, skip func(), breakout func()))
	Keys() []Key
	Slice() []Value
	Map() map[Key]Value
	Size() int
	// GetOne 获取一个
	GetOne() (value Value)
}
