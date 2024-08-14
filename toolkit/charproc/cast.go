package charproc

import (
	"strings"
	"unicode"
)

const vowels = "aeiouAEIOU"

// FirstUpper 返回字符串的首字母大写形式
func FirstUpper(s string) string {
	if len(s) == 0 {
		return s
	}
	r := []rune(s)
	r[0] = unicode.ToUpper(r[0])
	return string(r)
}

// FirstLower 返回字符串的首字母小写形式
func FirstLower(s string) string {
	if len(s) == 0 {
		return s
	}
	r := []rune(s)
	r[0] = unicode.ToLower(r[0])
	return string(r)
}

// LenWithChinese 返回含中文字符串的长度，一个中文字符长度为 1
func LenWithChinese(s string) int {
	return len([]rune(s))
}

// Snake 蛇形字符串
func Snake(str string) string {
	var builder strings.Builder
	vv := []rune(str)
	for i := 0; i < len(vv); i++ {
		if vv[i] >= 65 && vv[i] <= 90 {
			vv[i] += 32
			if builder.Len() > 0 {
				builder.WriteString("_")
			}
		}
		builder.WriteRune(vv[i])
	}
	return builder.String()
}

// BigCamel 大驼峰字符串
func BigCamel(str string) string {
	return FirstUpper(Camel(str))
}

// Camel 驼峰字符串
func Camel(str string) string {
	var builder strings.Builder
	vv := []rune(str)
	for i := 0; i < len(vv); i++ {
		if vv[i] == '_' {
			i++
			if vv[i] >= 97 && vv[i] <= 122 {
				vv[i] -= 32
				builder.WriteRune(vv[i])
			} else {
				return str
			}
		} else {
			builder.WriteRune(vv[i])
		}
	}
	return builder.String()
}

// Plural 将最后一个单词转换为复数形式
func Plural(word string) string {
	if strings.HasSuffix(word, "y") && len(word) > 1 && !IsVowel(rune(word[len(word)-2])) {
		return word[:len(word)-1] + "ies"
	}
	if strings.HasSuffix(word, "s") || strings.HasSuffix(word, "x") || strings.HasSuffix(word, "z") || strings.HasSuffix(word, "ch") || strings.HasSuffix(word, "sh") {
		return word + "es"
	}
	return word + "s"
}

// IsVowel 判断是否为元音
func IsVowel(ch rune) bool {
	return strings.ContainsRune(vowels, ch)
}
