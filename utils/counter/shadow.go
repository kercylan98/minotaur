package counter

import (
	"github.com/kercylan98/minotaur/utils/generic"
	"github.com/kercylan98/minotaur/utils/hash"
)

func newShadow[K comparable, V generic.Number](counter *Counter[K, V]) *Shadow[K, V] {
	counter.lock.Lock()
	shadow := &Shadow[K, V]{
		Sub:                 counter.sub,
		DeduplicationRecord: hash.Copy(counter.dr),
		Counter:             counter.GetCounts(),
	}
	for k, c := range counter.GetSubCounters() {
		shadow.SubCounters[k] = newShadow(c)
	}
	counter.lock.Unlock()
	return shadow
}

// Shadow 计数器的影子计数器
type Shadow[K comparable, V generic.Number] struct {
	Sub                 bool                // 是否为子计数器
	DeduplicationRecord map[K]int64         // 最后一次写入时间
	Counter             map[K]V             // 计数器
	SubCounters         map[K]*Shadow[K, V] // 子计数器
}

// ToCounter 将影子计数器转换为计数器
func (slf *Shadow[K, V]) ToCounter() *Counter[K, V] {
	return slf.toCounter(nil)
}

// toCounter 将影子计数器转换为计数器
func (slf *Shadow[K, V]) toCounter(parent *Counter[K, V]) *Counter[K, V] {
	counter := &Counter[K, V]{
		sub: slf.Sub,
		c:   map[K]any{},
	}
	if slf.Sub {
		counter.dr = parent.dr
	} else {
		counter.dr = hash.Copy(slf.DeduplicationRecord)
	}
	for k, v := range slf.Counter {
		counter.c[k] = v
	}
	for k, s := range slf.SubCounters {
		counter.c[k] = s.toCounter(counter)
	}
	return counter
}
