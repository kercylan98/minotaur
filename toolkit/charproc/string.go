package charproc

import (
	"slices"
	"strings"
)

const (
	None      = ""  // 空字符串
	Dunno     = "?" // 未知
	CenterDot = "·" // 中点
	Dot       = "." // 点
	Slash     = "/" // 斜杠
)

const (
	NoneChar      = byte(0)   // 空字符
	DunnoChar     = byte('?') // 未知
	CenterDotChar = byte('·') // 中点
	DotChar       = byte('.') // 点
	SlashChar     = byte('/') // 斜杠
)

// IsUpper 判断字符是否为大写字母
func IsUpper(r rune) bool {
	return r >= 'A' && r <= 'Z'
}

// IsLower 判断字符是否为小写字母
func IsLower(r rune) bool {
	return r >= 'a' && r <= 'z'
}

// IsLetter 判断字符是否为字母
func IsLetter(r rune) bool {
	return IsUpper(r) || IsLower(r)
}

// IsDigit 判断字符是否为数字
func IsDigit(r rune) bool {
	return r >= '0' && r <= '9'
}

// IsSpace 判断字符是否为空白字符
func IsSpace(r rune) bool {
	return r == ' ' || r == '\t' || r == '\n' || r == '\r'
}

// SortJoin 将多个字符串排序后拼接
func SortJoin(delimiter string, s ...string) string {
	var strList = make([]string, 0, len(s))
	for _, str := range s {
		strList = append(strList, str)
	}
	slices.Sort(strList)
	return strings.Join(strList, delimiter)
}
