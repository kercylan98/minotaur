package charproc

// FirstUpper 返回字符串的首字母大写形式
func FirstUpper(s string) string {
	if len(s) == 0 {
		return s
	}
	return string(s[0]-32) + s[1:]
}

// FirstLower 返回字符串的首字母小写形式
func FirstLower(s string) string {
	if len(s) == 0 {
		return s
	}
	return string(s[0]+32) + s[1:]
}

// LenWithChinese 返回含中文字符串的长度，一个中文字符长度为 1
func LenWithChinese(s string) int {
	return len([]rune(s))
}

// Snake 蛇形字符串
func Snake(str string) string {
	var snakeStr string
	vv := []rune(str)
	for i := 0; i < len(vv); i++ {
		if vv[i] >= 65 && vv[i] <= 90 {
			vv[i] += 32
			snakeStr += "_" + string(vv[i])
		} else {
			snakeStr += string(vv[i])
		}
	}
	return snakeStr
}

// BigCamel 大驼峰字符串
func BigCamel(str string) string {
	return FirstUpper(Camel(str))
}

// Camel 驼峰字符串
func Camel(str string) string {
	var camelStr string
	vv := []rune(str)
	for i := 0; i < len(vv); i++ {
		if vv[i] == '_' {
			i++
			if vv[i] >= 97 && vv[i] <= 122 {
				vv[i] -= 32
				camelStr += string(vv[i])
			} else {
				return str
			}
		} else {
			camelStr += string(vv[i])
		}
	}
	return camelStr
}
