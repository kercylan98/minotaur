package str

import (
	"regexp"
	"strconv"
	"strings"
)

// HideSensitivity 返回防敏感化后的字符串
//   - 隐藏身份证、邮箱、手机号等敏感信息用*号替代
func HideSensitivity(str string) (result string) {
	if str == "" {
		return "***"
	}
	if strings.Contains(str, "@") {
		res := strings.Split(str, "@")
		if len(res[0]) < 3 {
			resString := "***"
			result = resString + "@" + res[1]
		} else {
			resRs := []rune(str)
			res2 := string(resRs[0:3])
			resString := res2 + "***"
			result = resString + "@" + res[1]
		}
		return result
	} else {
		reg := `^1[0-9]\d{9}$`
		rgx := regexp.MustCompile(reg)
		mobileMatch := rgx.MatchString(str)
		if mobileMatch {
			rs := []rune(str)
			result = string(rs[0:5]) + "****" + string(rs[7:11])

		} else {
			nameRune := []rune(str)
			lens := len(nameRune)

			if lens <= 1 {
				result = "***"
			} else if lens == 2 {
				result = string(nameRune[:1]) + "*"
			} else if lens == 3 {
				result = string(nameRune[:1]) + "*" + string(nameRune[2:3])
			} else if lens == 4 {
				result = string(nameRune[:1]) + "**" + string(nameRune[lens-1:lens])
			} else if lens > 4 {
				result = string(nameRune[:2]) + "***" + string(nameRune[lens-2:lens])
			}
		}
		return
	}
}

// ThousandsSeparator 返回将str进行千位分隔符处理后的字符串。
func ThousandsSeparator(str string) string {
	length := len(str)
	if length < 4 {
		return str
	}
	arr := strings.Split(str, ".") //用小数点符号分割字符串,为数组接收
	length1 := len(arr[0])
	if length1 < 4 {
		return str
	}
	count := (length1 - 1) / 3
	for i := 0; i < count; i++ {
		arr[0] = arr[0][:length1-(i+1)*3] + "," + arr[0][length1-(i+1)*3:]
	}
	return strings.Join(arr, ".") //将一系列字符串连接为一个字符串，之间用sep来分隔。
}

// KV 返回str经过转换后形成的key、value
//   - 这里tag表示使用什么字符串来区分key和value的分隔符。
//   - 默认情况即不传入tag的情况下分隔符为“=”。
func KV(str string, tag ...string) (string, string) {
	tagChar := "="
	if len(tag) > 0 {
		tagChar = tag[0]
	}
	kv := strings.SplitN(str, tagChar, 2)
	if len(kv) < 2 {
		return "", ""
	}
	return kv[0], kv[1]
}

// FormatSpeedyInt 返回numberStr经过格式化后去除空格和“,”分隔符的结果
//   - 当字符串为“123,456,789”的时候，返回结果为“123456789”。
//   - 当字符串为“123 456 789”的时候，返回结果为“123456789”。
//   - 当字符串为“1 23, 45 6, 789”的时候，返回结果为“123456789”。
func FormatSpeedyInt(numberStr string) (int, error) {
	return strconv.Atoi(strings.ReplaceAll(strings.ReplaceAll(numberStr, " ", ""), ",", ""))
}

// FormatSpeedyInt64 返回numberStr经过格式化后去除空格和“,”分隔符的结果
//   - 当字符串为“123,456,789”的时候，返回结果为“123456789”。
//   - 当字符串为“123 456 789”的时候，返回结果为“123456789”。
//   - 当字符串为“1 23, 45 6, 789”的时候，返回结果为“123456789”。
func FormatSpeedyInt64(numberStr string) (int64, error) {
	return strconv.ParseInt(strings.ReplaceAll(strings.ReplaceAll(numberStr, " ", ""), ",", ""), 10, 64)
}

// FormatSpeedyFloat32 返回numberStr经过格式化后去除空格和“,”分隔符的结果
//   - 当字符串为“123,456,789.123”的时候，返回结果为“123456789.123”。
//   - 当字符串为“123 456 789.123”的时候，返回结果为“123456789.123”。
//   - 当字符串为“1 23, 45 6, 789.123”的时候，返回结果为“123456789.123”。
func FormatSpeedyFloat32(numberStr string) (float64, error) {
	return strconv.ParseFloat(strings.ReplaceAll(strings.ReplaceAll(numberStr, " ", ""), ",", ""), 32)
}

// FormatSpeedyFloat64 返回numberStr经过格式化后去除空格和“,”分隔符的结果
//   - 当字符串为“123,456,789.123”的时候，返回结果为“123456789.123”。
//   - 当字符串为“123 456 789.123”的时候，返回结果为“123456789.123”。
//   - 当字符串为“1 23, 45 6, 789.123”的时候，返回结果为“123456789.123”。
func FormatSpeedyFloat64(numberStr string) (float64, error) {
	return strconv.ParseFloat(strings.ReplaceAll(strings.ReplaceAll(numberStr, " ", ""), ",", ""), 64)
}
