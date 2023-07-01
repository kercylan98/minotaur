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
