package stream

import "strings"

// NewStrings 创建字符串切片
func NewStrings[S ~string](s ...S) *Strings[S] {
	return &Strings[S]{s}
}

// Strings 字符串切片
type Strings[S ~string] struct {
	s []S
}

// Elem 返回原始元素
func (s *Strings[S]) Elem() []S {
	return s.s
}

// Len 返回切片长度
func (s *Strings[S]) Len() int {
	return len(s.s)
}

// Append 添加字符串
func (s *Strings[S]) Append(ss ...S) *Strings[S] {
	s.s = append(s.s, ss...)
	return s
}

// Join 连接字符串
func (s *Strings[S]) Join(sep S) *String[S] {
	var cast = make([]string, len(s.s))
	for i, v := range s.s {
		cast[i] = string(v)
	}
	return NewString(S(strings.Join(cast, string(sep))))
}

// Choice 选择字符串
func (s *Strings[S]) Choice(i int) *String[S] {
	return NewString(s.s[i])
}

// Choices 选择多个字符串
func (s *Strings[S]) Choices(i ...int) *Strings[S] {
	var ss = make([]S, len(i))
	for j, v := range i {
		ss[j] = s.s[v]
	}
	return NewStrings(ss...)
}

// ChoiceInRange 选择范围内的字符串
func (s *Strings[S]) ChoiceInRange(start, end int) *Strings[S] {
	return NewStrings(s.s[start:end]...)
}

// Remove 移除字符串
func (s *Strings[S]) Remove(i int) *Strings[S] {
	s.s = append(s.s[:i], s.s[i+1:]...)
	return s
}

// Removes 移除多个字符串
func (s *Strings[S]) Removes(i ...int) *Strings[S] {
	var ss = make([]S, 0, len(s.s)-len(i))
	for j, v := range s.s {
		for _, i := range i {
			if j != i {
				ss = append(ss, v)
			}
		}
	}
	s.s = ss
	return s
}

// RemoveInRange 移除范围内的字符串
func (s *Strings[S]) RemoveInRange(start, end int) *Strings[S] {
	s.s = append(s.s[:start], s.s[end:]...)
	return s
}

// Clear 清空字符串
func (s *Strings[S]) Clear() *Strings[S] {
	s.s = []S{}
	return s
}

// First 第一个字符串
func (s *Strings[S]) First() *String[S] {
	return NewString(s.s[0])
}

// Last 最后一个字符串
func (s *Strings[S]) Last() *String[S] {
	return NewString(s.s[len(s.s)-1])
}
