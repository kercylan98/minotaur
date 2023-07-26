package hash

// NewSet 创建一个 Set 集合
func NewSet[K comparable](ks ...K) Set[K] {
	s := make(Set[K])
	s.AddAll(ks...)
	return s
}

// Set 集合
type Set[K comparable] map[K]struct{}

// Exist 检查特定 key 是否存在
func (s Set[K]) Exist(key K) bool {
	_, exist := s[key]
	return exist
}

// AllExist 检查多个 key 是否存在
func (s Set[K]) AllExist(keys ...K) bool {
	for _, key := range keys {
		if _, exist := s[key]; !exist {
			return false
		}
	}
	return true
}

// Add 添加元素
func (s Set[K]) Add(key K) {
	s[key] = struct{}{}
}

// AddAll 添加多个元素
func (s Set[K]) AddAll(keys ...K) {
	for _, key := range keys {
		s[key] = struct{}{}
	}
}

// Remove 移除元素
func (s Set[K]) Remove(key K) {
	delete(s, key)
}

// RemoveAll 移除多个元素
func (s Set[K]) RemoveAll(keys ...K) {
	for _, key := range keys {
		delete(s, key)
	}
}

// Clear 清空集合
func (s Set[K]) Clear() {
	for key := range s {
		delete(s, key)
	}
}

// Size 集合长度
func (s Set[K]) Size() int {
	return len(s)
}

// ToSlice 转换为切片
func (s Set[K]) ToSlice() []K {
	keys := make([]K, 0, len(s))
	for key := range s {
		keys = append(keys, key)
	}
	return keys
}

// ToMap 转换为 map
func (s Set[K]) ToMap() map[K]struct{} {
	return s
}

// ToJson 转换为 json 字符串
func (s Set[K]) ToJson() string {
	return ToJson(s)
}

// RandomGet 随机获取一个元素
func (s Set[K]) RandomGet() (k K) {
	for k = range s {
		return k
	}
	return
}
