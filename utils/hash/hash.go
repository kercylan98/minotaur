package hash

import "encoding/json"

// Exist 检查特定 key 是否存在
func Exist[K comparable, V any](m map[K]V, key K) bool {
	_, exist := m[key]
	return exist
}

// AllExist 检查多个 key 是否存在
func AllExist[K comparable, V any](m map[K]V, keys ...K) bool {
	for key := range m {
		if _, exist := m[key]; !exist {
			return false
		}
	}
	return true
}

// ToJson 将 map 转换为 json 字符串
func ToJson[K comparable, V any](m map[K]V) string {
	if data, err := json.Marshal(m); err == nil {
		return string(data)
	}
	return "{}"
}

// RandomGet 随机获取一个元素
func RandomGet[K comparable, V any](m map[K]V) (v V) {
	for _, v := range m {
		return v
	}
	return
}

// RandomGetKey 随机获取一个 key
func RandomGetKey[K comparable, V any](m map[K]V) (k K) {
	for k = range m {
		return k
	}
	return
}

// RandomGetN 随机获取 n 个元素
//   - 获取到的元素不会是重复的，当 map 的长度不足 n 时，返回的元素等同于 hash.ToSlice
func RandomGetN[K comparable, V any](m map[K]V, n int) (vs []V) {
	for _, v := range m {
		vs = append(vs, v)
		if len(vs) >= n {
			return
		}
	}
	return
}

// RandomGetKeyN 随机获取 n 个 key
//   - 获取到的元素不会是重复的，当 map 的长度不足 n 时，返回的元素等同于 hash.KeyToSlice
func RandomGetKeyN[K comparable, V any](m map[K]V, n int) (ks []K) {
	for k := range m {
		ks = append(ks, k)
		if len(ks) >= n {
			return
		}
	}
	return
}
