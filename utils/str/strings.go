package str

import "strings"

type Strings []*String

// Index 返回字符串指定位置的字符
func (s *Strings) Index(i int) *String {
	return (*s)[i]
}

// Append 追加字符串
func (s *Strings) Append(str *String) *Strings {
	*s = append(*s, str)
	return s
}

// Join 以指定字符连接字符串
func (s *Strings) Join(sep string) *String {
	var ss []string
	for _, v := range *s {
		ss = append(ss, v.String())
	}
	*s = append(*s, New(strings.Join(ss, sep)))
	return (*s)[len(*s)-1]
}
