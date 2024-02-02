package str

import (
	"net/url"
	"strings"
)

type String string

func New(s string) *String {
	return (*String)(&s)
}

// String 返回字符串
func (s *String) String() string {
	return string(*s)
}

// Append 追加字符串
func (s *String) Append(str *String) *String {
	*s = String(string(*s) + string(*str))
	return s
}

// Join 以指定字符连接字符串
func (s *String) Join(sep string) *String {
	*s = String(strings.Join([]string{string(*s), sep}, ""))
	return s
}

// QueryEscape 对字符串进行 URL 编码
func (s *String) QueryEscape() *String {
	*s = String(strings.ReplaceAll(url.QueryEscape(string(*s)), "+", "%20"))
	return s
}

// Replace 替换字符串
func (s *String) Replace(old, new string, n int) *String {
	*s = String(strings.Replace(string(*s), old, new, n))
	return s
}

// ReplaceAll 替换字符串
func (s *String) ReplaceAll(old, new string) *String {
	*s = String(strings.ReplaceAll(string(*s), old, new))
	return s
}

// Trim 去除字符串两端的指定字符
func (s *String) Trim(cs string) *String {
	*s = String(strings.Trim(string(*s), cs))
	return s
}

// TrimLeft 去除字符串左端的指定字符
func (s *String) TrimLeft(cs string) *String {
	*s = String(strings.TrimLeft(string(*s), cs))
	return s
}

// TrimRight 去除字符串右端的指定字符
func (s *String) TrimRight(cs string) *String {
	*s = String(strings.TrimRight(string(*s), cs))
	return s
}

// Default 返回字符串，如果字符串为空则返回默认值
func (s *String) Default(def string) *String {
	if string(*s) == "" {
		*s = String(def)
	}
	return s
}

// Format 格式化字符串
func (s *String) Format(formatter func(s *String) *String) *String {
	*s = *formatter(s)
	return s
}

// Index 返回字符串指定位置的字符
func (s *String) Index(i int) String {
	return String(string(*s)[i])
}

// Split 以指定字符分割字符串
func (s *String) Split(sep string) *Strings {
	slice := strings.Split(string(*s), sep)
	ss := make(Strings, len(slice))
	for i, v := range slice {
		ss[i] = New(v)
	}
	return &ss
}

// SplitN 以指定字符分割字符串，最多分割 n 次
func (s *String) SplitN(sep string, n int) *Strings {
	slice := strings.SplitN(string(*s), sep, n)
	ss := make(Strings, len(slice))
	for i, v := range slice {
		ss[i] = New(v)
	}
	return &ss
}

// ToUpper 将字符串转为大写
func (s *String) ToUpper() *String {
	*s = String(strings.ToUpper(string(*s)))
	return s
}

// ToLower 将字符串转为小写
func (s *String) ToLower() *String {
	*s = String(strings.ToLower(string(*s)))
	return s
}

// TrimSpace 去除字符串两端的空白字符
func (s *String) TrimSpace() *String {
	*s = String(strings.TrimSpace(string(*s)))
	return s
}

// TrimPrefix 去除字符串前缀
func (s *String) TrimPrefix(prefix string) *String {
	*s = String(strings.TrimPrefix(string(*s), prefix))
	return s
}

// TrimSuffix 去除字符串后缀
func (s *String) TrimSuffix(suffix string) *String {
	*s = String(strings.TrimSuffix(string(*s), suffix))
	return s
}

// TrimSpacePrefix 去除字符串两端的空白字符后再去除前缀
func (s *String) TrimSpacePrefix(prefix string) *String {
	*s = String(strings.TrimPrefix(strings.TrimSpace(string(*s)), prefix))
	return s
}

// TrimSpaceSuffix 去除字符串两端的空白字符后再去除后缀
func (s *String) TrimSpaceSuffix(suffix string) *String {
	*s = String(strings.TrimSuffix(strings.TrimSpace(string(*s)), suffix))
	return s
}
