package hash

// RandomDrop 随机删除哈希表中的指定数量的元素
//   - 该函数会修改原哈希表，如果不想修改原哈希表，请使用 RandomDropCopy
func RandomDrop[K comparable, V any](n int, hash map[K]V) map[K]V {
	if n <= 0 || len(hash) <= n {
		return hash
	}
	for k := range hash {
		if n <= 0 {
			break
		}
		delete(hash, k)
		n--
	}
	return hash
}

// RandomDropCopy 随机删除哈希表中的指定数量的元素
//   - 该函数不会修改原哈希表，如果想修改原哈希表，请使用 RandomDrop
func RandomDropCopy[K comparable, V any](n int, m map[K]V) map[K]V {
	if n <= 0 || len(m) <= n {
		return map[K]V{}
	}
	var nm = make(map[K]V, len(m))
	for k, v := range m {
		if n <= 0 {
			break
		}
		nm[k] = v
		n--
	}
	return nm
}

// DropBy 从哈希表中删除指定的元素
func DropBy[K comparable, V any](m map[K]V, fn func(key K, value V) bool) map[K]V {
	for k, v := range m {
		if fn(k, v) {
			delete(m, k)
		}
	}
	return m
}

// DropByCopy 与 DropBy 功能相同，但是该函数不会修改原哈希表
func DropByCopy[K comparable, V any](m map[K]V, fn func(key K, value V) bool) map[K]V {
	var nm = make(map[K]V, len(m))
	for k, v := range m {
		if fn(k, v) {
			continue
		}
		nm[k] = v
	}
	return nm
}
