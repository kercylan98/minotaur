package random

import (
	"fmt"
	"math/rand"
	"strconv"
)

// ChineseName 返回一个随机组成的中文姓名。
func ChineseName() string {
	var first string // 名
	// 随机产生2位或者3位的名
	for i := 0; i <= rand.Intn(1); i++ {
		first = fmt.Sprint(firstName[rand.Intn(firstNameLen-1)])
	}
	//返回姓名
	return fmt.Sprintf("%s%s", fmt.Sprint(lastName[rand.Intn(lastNameLen-1)]), first)
}

// EnglishName 返回一个随机组成的英文姓名。
func EnglishName() string {
	var englishName string
	for i := 0; i <= rand.Intn(1); i++ {
		englishName = fmt.Sprint(engName[rand.Intn(englishNameLen-1)])
	}
	first := engName[rand.Intn(englishNameLen-1)]
	last := englishName
	if first == last {
		return first
	} else {
		return fmt.Sprintf("%s %s", fmt.Sprint(engName[rand.Intn(englishNameLen-1)]), englishName)
	}
}

// Name 返回一个随机组成的中文或英文姓名
//   - 以1/2的概率决定生产的是中文还是英文姓名。
func Name() string {
	if Int64(0, 1000) > 500 {
		return EnglishName()
	} else {
		return ChineseName()
	}
}

// NumberString 返回一个介于min和max之间的string类型的随机数。
func NumberString(min int, max int) string {
	return strconv.Itoa(int(Int64(int64(min), int64(max))))
}

// NumberStringRepair 返回一个介于min和max之间的string类型的随机数
//   - 通过Int64生成一个随机数，当结果的字符串长度小于max的字符串长度的情况下，使用0在开头补齐。
func NumberStringRepair(min int, max int) string {
	result := strconv.Itoa(int(Int64(int64(min), int64(max))))
	for i := len(result); i < len(strconv.Itoa(max)); i++ {
		result = "0" + result
	}
	return result
}

// HostName 返回一个随机产生的hostname。
func HostName() string {
	str := "0123456789abcdefghijklmnopqrstuvwxyz"
	bytes := []byte(str)
	var hostname []byte
	r := rand.New(rand.NewSource(Int64(0, 999999)))
	for i := 0; i < 12; i++ {
		hostname = append(hostname, bytes[r.Intn(len(bytes))])
	}
	return string(hostname)
}
