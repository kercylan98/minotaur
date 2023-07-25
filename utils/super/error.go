package super

import (
	"errors"
)

var errorMapper = make(map[error]int)
var errorMapperRef = make(map[error]error)

// RegError 通过错误码注册错误，返回错误的引用
func RegError(code int, message string) error {
	if code == 0 {
		panic("error code can not be 0")
	}
	err := errors.New(message)
	errorMapper[err] = code
	return err
}

// RegErrorRef 通过错误码注册错误，返回错误的引用
func RegErrorRef(code int, message string, ref error) error {
	if code == 0 {
		panic("error code can not be 0")
	}
	err := errors.New(message)
	errorMapper[err] = code
	errorMapperRef[ref] = err
	return ref
}

// GetErrorCode 通过错误引用获取错误码，如果错误不存在则返回 0
func GetErrorCode(err error) (int, error) {
	if ref, exist := errorMapperRef[err]; exist {
		err = ref
	}
	return errorMapper[err], err
}
