package stream

import (
	"sort"
)

// NewStrings 创建字符串切片
func NewStrings[S ~string](s ...S) Strings[S] {
	var slice = make(Strings[S], len(s))
	for i, v := range s {
		slice[i] = v
	}
	return slice
}

// Strings 字符串切片
type Strings[S ~string] []S

// Len 返回切片长度
func (s *Strings[S]) Len() int {
	return len(*s)
}

// Append 添加字符串
func (s *Strings[S]) Append(ss ...S) *Strings[S] {
	*s = append(*s, NewStrings(ss...)...)
	return s
}

// Clear 清空切片
func (s *Strings[S]) Clear() *Strings[S] {
	*s = make(Strings[S], 0)
	return s
}

// Copy 复制切片
func (s *Strings[S]) Copy() *Strings[S] {
	ss := make(Strings[S], len(*s))
	copy(ss, *s)
	return &ss
}

// Range 返回指定范围的切片
func (s *Strings[S]) Range(start, end int) *Strings[S] {
	*s = (*s)[start:end]
	return s
}

// First 返回第一个元素
func (s *Strings[S]) First() *String[S] {
	return NewString((*s)[0])
}

// Last 返回最后一个元素
func (s *Strings[S]) Last() *String[S] {
	return NewString((*s)[len(*s)-1])
}

// Index 返回指定的元素
func (s *Strings[S]) Index(i int) *String[S] {
	return NewString((*s)[i])
}

// Reverse 反转切片
func (s *Strings[S]) Reverse() *Strings[S] {
	for i, j := 0, len(*s)-1; i < j; i, j = i+1, j-1 {
		(*s)[i], (*s)[j] = (*s)[j], (*s)[i]
	}
	return s
}

// Desc 降序排序
func (s *Strings[S]) Desc() *Strings[S] {
	sort.Slice(*s, func(i, j int) bool {
		return (*s)[i] > (*s)[j]
	})
	return s
}

// Asc 升序排序
func (s *Strings[S]) Asc() *Strings[S] {
	sort.Slice(*s, func(i, j int) bool {
		return (*s)[i] < (*s)[j]
	})
	return s
}

// Sort 自定义排序
func (s *Strings[S]) Sort(f func(int, int) bool) *Strings[S] {
	sort.Slice(*s, func(i, j int) bool {
		return f(i, j)
	})
	return s
}

// Unique 去重
func (s *Strings[S]) Unique() *Strings[S] {
	m := map[S]struct{}{}
	for _, v := range *s {
		m[v] = struct{}{}
	}
	*s = make(Strings[S], 0, len(m))
	for k := range m {
		*s = append(*s, k)
	}
	return s
}

// Delete 删除指定位置的字符串
func (s *Strings[S]) Delete(i int) *Strings[S] {
	*s = append((*s)[:i], (*s)[i+1:]...)
	return s
}

// Each 遍历切片
func (s *Strings[S]) Each(f func(int, S) bool) *Strings[S] {
	for i, v := range *s {
		if !f(i, v) {
			break
		}
	}
	return s
}
