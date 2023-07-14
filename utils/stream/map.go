package stream

import (
	"github.com/kercylan98/minotaur/utils/asynchronous"
	"github.com/kercylan98/minotaur/utils/hash"
	"github.com/kercylan98/minotaur/utils/synchronization"
	"reflect"
)

// WithMap 使用传入的 map 执行链式操作
//   - 该函数将会直接影响到传入的 map
func WithMap[K comparable, V any](m map[K]V) Map[K, V] {
	return m
}

// WithMapCopy 使用传入的 map 执行链式操作
//   - 该函数不会影响到传入的 map
func WithMapCopy[K comparable, V any](m map[K]V) Map[K, V] {
	return hash.Copy(m)
}

// WithHashMap 使用传入的 map 执行链式操作
func WithHashMap[K comparable, V any](m hash.Map[K, V]) Map[K, V] {
	return m.Map()
}

// Map 提供了 map 的链式操作
type Map[K comparable, V any] map[K]V

// Set 设置一个值
func (slf Map[K, V]) Set(key K, value V) Map[K, V] {
	slf[key] = value
	return slf
}

// Delete 删除一个值
func (slf Map[K, V]) Delete(key K) Map[K, V] {
	delete(slf, key)
	return slf
}

// Filter 过滤 handle 返回 false 的元素
func (slf Map[K, V]) Filter(handle func(key K, value V) bool) Map[K, V] {
	for k, v := range slf {
		if !handle(k, v) {
			delete(slf, k)
		}
	}
	return slf
}

// FilterKey 过滤特定的 key
func (slf Map[K, V]) FilterKey(keys ...K) Map[K, V] {
	for _, key := range keys {
		delete(slf, key)
	}
	return slf
}

// FilterValue 过滤特定的 value
func (slf Map[K, V]) FilterValue(values ...V) Map[K, V] {
	for k, v := range slf {
		for _, value := range values {
			if reflect.DeepEqual(v, value) {
				delete(slf, k)
			}
		}
	}
	return slf
}

// RandomKeep 随机保留 n 个元素
func (slf Map[K, V]) RandomKeep(n int) Map[K, V] {
	length := len(slf)
	if n >= length {
		return slf
	}
	for k := range slf {
		if n > 0 {
			n--
		} else {
			delete(slf, k)
		}
	}
	return slf
}

// RandomDelete 随机删除 n 个元素
func (slf Map[K, V]) RandomDelete(n int) Map[K, V] {
	var count int
	for k := range slf {
		if count < n {
			count++
			delete(slf, k)
		} else {
			return slf
		}
	}
	return slf
}

// RandomReplace 将 values 覆盖到当前的 map 中
//   - 如果 values 的长度大于当前 map 的长度，则只会覆盖当前 map 的长度
func (slf Map[K, V]) RandomReplace(values ...V) Map[K, V] {
	var record []K
	var valuesLen = len(values)
	for k := range slf {
		record = append(record, k)
		if len(record) >= valuesLen {
			break
		}
	}
	for i, k := range record {
		slf.Set(k, values[i])
	}
	return slf
}

// Distinct 去重，如果 handle 返回 true，则认为是重复的元素
func (slf Map[K, V]) Distinct(handle func(key K, value V) bool) Map[K, V] {
	for k, v := range slf {
		if handle(k, v) {
			delete(slf, k)
		}
	}
	return slf
}

// Range 遍历当前 Map, handle 返回 false 则停止遍历
func (slf Map[K, V]) Range(handle func(key K, value V) bool) Map[K, V] {
	for k, v := range slf {
		if !handle(k, v) {
			break
		}
	}
	return slf
}

// ValueOr 当 key 不存在时，设置一个默认值
func (slf Map[K, V]) ValueOr(key K, value V) Map[K, V] {
	if _, ok := slf[key]; !ok {
		slf[key] = value
	}
	return slf
}

// GetValueOr 当 key 不存在时，返回一个默认值
func (slf Map[K, V]) GetValueOr(key K, value V) V {
	if v, ok := slf[key]; ok {
		return v
	}
	return value
}

// Clear 清空当前 Map
func (slf Map[K, V]) Clear() Map[K, V] {
	for k := range slf {
		delete(slf, k)
	}
	return slf
}

// Merge 合并多个 Map
func (slf Map[K, V]) Merge(maps ...map[K]V) Map[K, V] {
	for _, m := range maps {
		for k, v := range m {
			slf[k] = v
		}
	}
	return slf
}

// ToSliceStream 将当前 Map stream 转换为 Slice stream
func (slf Map[K, V]) ToSliceStream() Slice[V] {
	return hash.ToSlice(slf)
}

// ToSliceStreamWithKey 将当前 Map stream 转换为 Slice stream，key 为 Slice 的元素
func (slf Map[K, V]) ToSliceStreamWithKey() Slice[K] {
	return hash.KeyToSlice(slf)
}

// ToSyncMap 将当前 Map 转换为 synchronization.Map
func (slf Map[K, V]) ToSyncMap() *synchronization.Map[K, V] {
	return synchronization.NewMap[K, V](synchronization.WithMapSource(slf))
}

// ToAsyncMap 将当前 Map 转换为 asynchronous.Map
func (slf Map[K, V]) ToAsyncMap() *asynchronous.Map[K, V] {
	return asynchronous.NewMap[K, V](asynchronous.WithMapSource(slf))
}

// ToMap 将当前 Map 转换为 map
func (slf Map[K, V]) ToMap() map[K]V {
	return slf
}
