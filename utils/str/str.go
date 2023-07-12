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
	vv := []rune(str) // 后文有介绍
	for i := 0; i < len(vv); i++ {
		if i == 0 {
			if vv[i] >= 97 && vv[i] <= 122 { // 后文有介绍
				vv[i] -= 32 // string的码表相差32位
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
