package random

import (
	"sync"
)

var (
	chineseNameGenerator *NameGenerator
	chineseNameOnce      sync.Once

	englishNameGenerator *NameGenerator
	englishNameOnce      sync.Once
)

func ChineseName() string {
	chineseNameOnce.Do(func() {
		chineseNameGenerator = DefaultChineseNameGenerator()
	})
	return chineseNameGenerator.RandomName()
}

func EnglishName() string {
	englishNameOnce.Do(func() {
		englishNameGenerator = DefaultEnglishNameGenerator()
	})
	return englishNameGenerator.RandomName()
}

// HostName 返回一个随机产生的hostname。
func HostName() string {
	str := "0123456789abcdefghijklmnopqrstuvwxyz"
	bytes := []byte(str)
	var hostname []byte
	for i := 0; i < 12; i++ {
		hostname = append(hostname, bytes[IntN(len(bytes))])
	}
	return string(hostname)
}

// NumberVerificationCode 返回一个随机产生的数字验证码
func NumberVerificationCode(length int) string {
	str := "0123456789"
	bytes := []byte(str)
	var code []byte
	for i := 0; i < length; i++ {
		code = append(code, bytes[IntN(len(bytes))])
	}
	return string(code)
}

// VerificationCode 返回一个随机产生的验证码
func VerificationCode(length int) string {
	str := "0123456789abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ"
	bytes := []byte(str)
	var code []byte
	for i := 0; i < length; i++ {
		code = append(code, bytes[IntN(len(bytes))])
	}
	return string(code)
}
