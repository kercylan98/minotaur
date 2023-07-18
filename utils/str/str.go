package str

const (
	None      = ""  // 空字符串
	Dunno     = "?" // 未知
	CenterDot = "·" // 中点
	Dot       = "." // 点
	Slash     = "/" // 斜杠
)

var (
	NoneBytes      = []byte("")  // 空字符串
	DunnoBytes     = []byte("?") // 未知
	CenterDotBytes = []byte("·") // 中点
	DotBytes       = []byte(".") // 点
	SlashBytes     = []byte("/") // 斜杠
)

// FirstUpper 首字母大写
func FirstUpper(str string) string {
	var upperStr string
	vv := []rune(str)
	for i := 0; i < len(vv); i++ {
		if i == 0 {
			if vv[i] >= 97 && vv[i] <= 122 {
				vv[i] -= 32
				upperStr += string(vv[i])
			} else {
				return str
			}
		} else {
			upperStr += string(vv[i])
		}
	}
	return upperStr
}

// FirstLower 首字母小写
func FirstLower(str string) string {
	var lowerStr string
	vv := []rune(str)
	for i := 0; i < len(vv); i++ {
		if i == 0 {
			if vv[i] >= 65 && vv[i] <= 90 {
				vv[i] += 32
				lowerStr += string(vv[i])
			} else {
				return str
			}
		} else {
			lowerStr += string(vv[i])
		}
	}
	return lowerStr
}

// FirstUpperBytes 首字母大写
func FirstUpperBytes(str []byte) []byte {
	var upperStr []byte
	vv := []rune(string(str))
	for i := 0; i < len(vv); i++ {
		if i == 0 {
			if vv[i] >= 97 && vv[i] <= 122 {
				vv[i] -= 32
				upperStr = append(upperStr, string(vv[i])...)
			} else {
				return str
			}
		} else {
			upperStr = append(upperStr, string(vv[i])...)
		}
	}
	return upperStr
}

// FirstLowerBytes 首字母小写
func FirstLowerBytes(str []byte) []byte {
	var lowerStr []byte
	vv := []rune(string(str))
	for i := 0; i < len(vv); i++ {
		if i == 0 {
			if vv[i] >= 65 && vv[i] <= 90 {
				vv[i] += 32
				lowerStr = append(lowerStr, string(vv[i])...)
			} else {
				return str
			}
		} else {
			lowerStr = append(lowerStr, string(vv[i])...)
		}
	}
	return lowerStr
}

// IsEmpty 判断字符串是否为空
func IsEmpty(str string) bool {
	return len(str) == 0
}

// IsEmptyBytes 判断字符串是否为空
func IsEmptyBytes(str []byte) bool {
	return len(str) == 0
}

// IsNotEmpty 判断字符串是否不为空
func IsNotEmpty(str string) bool {
	return !IsEmpty(str)
}

// IsNotEmptyBytes 判断字符串是否不为空
func IsNotEmptyBytes(str []byte) bool {
	return !IsEmptyBytes(str)
}

// SnakeString 蛇形字符串
func SnakeString(str string) string {
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

// SnakeStringBytes 蛇形字符串
func SnakeStringBytes(str []byte) []byte {
	var snakeStr []byte
	vv := []rune(string(str))
	for i := 0; i < len(vv); i++ {
		if vv[i] >= 65 && vv[i] <= 90 {
			vv[i] += 32
			snakeStr = append(snakeStr, '_')
			snakeStr = append(snakeStr, string(vv[i])...)
		} else {
			snakeStr = append(snakeStr, string(vv[i])...)
		}
	}
	return snakeStr
}

// CamelString 驼峰字符串
func CamelString(str string) string {
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

// CamelStringBytes 驼峰字符串
func CamelStringBytes(str []byte) []byte {
	var camelStr []byte
	vv := []rune(string(str))
	for i := 0; i < len(vv); i++ {
		if vv[i] == '_' {
			i++
			if vv[i] >= 97 && vv[i] <= 122 {
				vv[i] -= 32
				camelStr = append(camelStr, string(vv[i])...)
			} else {
				return str
			}
		} else {
			camelStr = append(camelStr, string(vv[i])...)
		}
	}
	return camelStr
}
