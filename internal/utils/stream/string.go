package stream

import (
	"strconv"
	"strings"
)

// NewString 创建字符串流
func NewString[S ~string](s S) *String[S] {
	return &String[S]{s}
}

// String 字符串流
type String[S ~string] struct {
	str S
}

// Elem 返回原始元素
func (s *String[S]) Elem() S {
	return s.str
}

// String 返回字符串
func (s *String[S]) String() string {
	return string(s.str)
}

// Index 返回字符串指定位置的字符，当索引超出范围时将会触发 panic
func (s *String[S]) Index(i int) *String[S] {
	return NewString(S(s.str[i]))
}

// Range 返回字符串指定范围的字符
func (s *String[S]) Range(start, end int) *String[S] {
	return NewString(s.str[start:end])
}

// TrimSpace 返回去除字符串首尾空白字符的字符串
func (s *String[S]) TrimSpace() *String[S] {
	return NewString(S(strings.TrimSpace(string(s.str))))
}

// Trim 返回去除字符串首尾指定字符的字符串
func (s *String[S]) Trim(cs string) *String[S] {
	return NewString(S(strings.Trim(string(s.str), cs)))
}

// TrimPrefix 返回去除字符串前缀的字符串
func (s *String[S]) TrimPrefix(prefix string) *String[S] {
	return NewString(S(strings.TrimPrefix(string(s.str), prefix)))
}

// TrimSuffix 返回去除字符串后缀的字符串
func (s *String[S]) TrimSuffix(suffix string) *String[S] {
	return NewString(S(strings.TrimSuffix(string(s.str), suffix)))
}

// ToUpper 返回字符串的大写形式
func (s *String[S]) ToUpper() *String[S] {
	return NewString(S(strings.ToUpper(string(s.str))))
}

// ToLower 返回字符串的小写形式
func (s *String[S]) ToLower() *String[S] {
	return NewString(S(strings.ToLower(string(s.str))))
}

// Equal 返回字符串是否相等
func (s *String[S]) Equal(ss S) bool {
	return s.str == ss
}

// HasPrefix 返回字符串是否包含指定前缀
func (s *String[S]) HasPrefix(prefix S) bool {
	return strings.HasPrefix(string(s.str), string(prefix))
}

// HasSuffix 返回字符串是否包含指定后缀
func (s *String[S]) HasSuffix(suffix S) bool {
	return strings.HasSuffix(string(s.str), string(suffix))
}

// Len 返回字符串长度
func (s *String[S]) Len() int {
	return len(s.str)
}

// Contains 返回字符串是否包含指定子串
func (s *String[S]) Contains(sub S) bool {
	return strings.Contains(string(s.str), string(sub))
}

// Count 返回字符串包含指定子串的次数
func (s *String[S]) Count(sub S) int {
	return strings.Count(string(s.str), string(sub))
}

// Repeat 返回重复 count 次的字符串
func (s *String[S]) Repeat(count int) *String[S] {
	return NewString(S(strings.Repeat(string(s.str), count)))
}

// Replace 返回替换指定子串后的字符串
func (s *String[S]) Replace(old, new S, n int) *String[S] {
	return NewString(S(strings.Replace(string(s.str), string(old), string(new), n)))
}

// ReplaceAll 返回替换所有指定子串后的字符串
func (s *String[S]) ReplaceAll(old, new S) *String[S] {
	return NewString(S(strings.ReplaceAll(string(s.str), string(old), string(new))))
}

// Append 返回追加指定字符串后的字符串
func (s *String[S]) Append(ss S) *String[S] {
	return NewString(S(string(s.str) + string(ss)))
}

// Prepend 返回追加指定字符串后的字符串，追加的字符串在前
func (s *String[S]) Prepend(ss S) *String[S] {
	return NewString(S(string(ss) + string(s.str)))
}

// Clear 返回清空字符串后的字符串
func (s *String[S]) Clear() *String[S] {
	return NewString(S(""))
}

// Reverse 返回反转字符串后的字符串
func (s *String[S]) Reverse() *String[S] {
	var str = []rune(string(s.str))
	for i, j := 0, len(str)-1; i < j; i, j = i+1, j-1 {
		str[i], str[j] = str[j], str[i]
	}
	return NewString(S(string(str)))
}

// Queto 返回带引号的字符串
func (s *String[S]) Queto() *String[S] {
	return NewString(S(strconv.Quote(string(s.str))))
}

// QuetoToASCII 返回带引号的字符串
func (s *String[S]) QuetoToASCII() *String[S] {
	return NewString(S(strconv.QuoteToASCII(string(s.str))))
}

// FirstUpper 返回首字母大写的字符串
func (s *String[S]) FirstUpper() *String[S] {
	var upperStr string
	vv := []rune(string(s.str))
	for i := 0; i < len(vv); i++ {
		if i == 0 {
			if vv[i] >= 97 && vv[i] <= 122 {
				vv[i] -= 32
				upperStr += string(vv[i])
			} else {
				return s
			}
		} else {
			upperStr += string(vv[i])
		}
	}
	return NewString(S(upperStr))
}

// FirstLower 返回首字母小写的字符串
func (s *String[S]) FirstLower() *String[S] {
	var lowerStr string
	vv := []rune(string(s.str))
	for i := 0; i < len(vv); i++ {
		if i == 0 {
			if vv[i] >= 65 && vv[i] <= 90 {
				vv[i] += 32
				lowerStr += string(vv[i])
			} else {
				return s
			}
		} else {
			lowerStr += string(vv[i])
		}
	}
	return NewString(S(lowerStr))
}

// SnakeCase 返回蛇形命名的字符串
func (s *String[S]) SnakeCase() *String[S] {
	var str = string(s.str)
	var result string
	for i, v := range str {
		if v >= 65 && v <= 90 {
			if i != 0 {
				result += "_"
			}
			result += string(v + 32)
		} else {
			result += string(v)
		}
	}
	return NewString(S(result))
}

// CamelCase 返回驼峰命名的字符串
func (s *String[S]) CamelCase() *String[S] {
	var str = string(s.str)
	var result string
	var upper = false
	for _, v := range str {
		if v == 95 {
			upper = true
		} else {
			if upper {
				result += string(v - 32)
				upper = false
			} else {
				result += string(v)
			}
		}
	}
	return NewString(S(result))
}

// KebabCase 返回短横线命名的字符串
func (s *String[S]) KebabCase() *String[S] {
	var str = string(s.str)
	var result string
	for i, v := range str {
		if v >= 65 && v <= 90 {
			if i != 0 {
				result += "-"
			}
			result += string(v + 32)
		} else {
			result += string(v)
		}
	}
	return NewString(S(result))
}

// TitleCase 返回标题命名的字符串
func (s *String[S]) TitleCase() *String[S] {
	var str = string(s.str)
	var result string
	var upper = true
	for _, v := range str {
		if v == 95 || v == 45 || v == 32 {
			upper = true
		} else {
			if upper {
				result += string(v - 32)
				upper = false
			} else {
				result += string(v)
			}
		}
	}
	return NewString(S(result))
}

// Bytes 返回字符串的字节数组
func (s *String[S]) Bytes() []byte {
	return []byte(s.str)
}

// Runes 返回字符串的字符数组
func (s *String[S]) Runes() []rune {
	return []rune(s.str)
}

// Default 当字符串为空时设置默认值
func (s *String[S]) Default(def S) *String[S] {
	if s.str == "" {
		return NewString(def)
	}
	return s
}

// Handle 处理字符串
func (s *String[S]) Handle(f func(S)) *String[S] {
	f(s.str)
	return s
}

// Update 更新字符串
func (s *String[S]) Update(f func(S) S) *String[S] {
	s.str = f(s.str)
	return s
}

// Split 返回字符串切片
func (s *String[S]) Split(sep string) *Strings[S] {
	slice := strings.Split(string(s.str), sep)
	rep := make([]S, len(slice))
	for i, v := range slice {
		rep[i] = S(v)
	}
	return NewStrings(rep...)
}

// SplitN 返回字符串切片
func (s *String[S]) SplitN(sep string, n int) *Strings[S] {
	slice := strings.SplitN(string(s.str), sep, n)
	rep := make([]S, len(slice))
	for i, v := range slice {
		rep[i] = S(v)
	}
	return NewStrings(rep...)
}

// Batched 将字符串按照指定长度分组，最后一组可能小于指定长度
func (s *String[S]) Batched(size int) *Strings[S] {
	var str = string(s.str)
	var result = make([]S, 0, len(str)/size+1)
	for len(str) >= size {
		result = append(result, S(str[:size]))
		str = str[size:]
	}
	if len(str) > 0 {
		result = append(result, S(str))
	}
	return NewStrings(result...)
}
