package super

import (
	"errors"
	"sync"
)

var errorMapper = make(map[error]int)
var errorMapperRef = make(map[error]error)
var mutex sync.Mutex

// RegError 通过错误码注册错误，返回错误的引用
func RegError(code int, message string) error {
	if code == 0 {
		return errors.New("error code can not be 0")
	}
	mutex.Lock()
	defer mutex.Unlock()
	err := &ser{code: code, message: message}
	errorMapper[err] = code
	return err
}

// RegErrorRef 通过错误码注册错误，返回错误的引用
//   - 引用将会被重定向到注册的错误信息
func RegErrorRef(code int, message string, ref error) error {
	if code == 0 {
		return errors.New("error code can not be 0")
	}
	mutex.Lock()
	defer mutex.Unlock()
	err := &ser{code: code, message: message}
	errorMapper[err] = code
	errorMapperRef[ref] = err
	return ref
}

// GetError 通过错误引用获取错误码和真实错误信息，如果错误不存在则返回 0，如果错误引用不存在则返回原本的错误
func GetError(err error) (int, error) {
	unw := errors.Unwrap(err)
	if unw == nil {
		unw = err
	}
	mutex.Lock()
	defer mutex.Unlock()
	if ref, exist := errorMapperRef[unw]; exist {
		//err = fmt.Errorf("%w : %s", ref, err.Error())
		err = ref
	}
	unw = errors.Unwrap(err)
	if unw == nil {
		unw = err
	}
	code, exist := errorMapper[unw]
	if !exist {
		return 0, errors.New("error not found")
	}
	return code, err
}

type ser struct {
	code    int
	message string
}

func (slf *ser) Error() string {
	return slf.message
}
