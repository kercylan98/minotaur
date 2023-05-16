package synchronization

// MapReadonly 并发安全的只读字典接口
type MapReadonly[Key comparable, Value any] interface {
	Get(key Key) Value
	Exist(key Key) bool
	GetExist(key Key) (Value, bool)
	Length() int
	Range(handle func(key Key, value Value))
	RangeSkip(handle func(key Key, value Value) bool)
	RangeBreakout(handle func(key Key, value Value) bool)
	RangeFree(handle func(key Key, value Value, skip func(), breakout func()))
	Keys() []Key
	Slice() []Value
	Map() map[Key]Value
	Size() int
	GetOne() (value Value)
}
