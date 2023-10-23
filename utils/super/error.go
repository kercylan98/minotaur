package super

import (
	"errors"
	"sync"
)

var errorManagerInstance *errorManager
var errorManagerMutex sync.Mutex

// RegError 通过错误码注册错误，返回错误的引用
func RegError(code int, message string) error {
	if code == 0 {
		return errors.New("error code can not be 0")
	}
	errorManagerMutex.Lock()
	defer errorManagerMutex.Unlock()
	if errorManagerInstance == nil {
		errorManagerInstance = new(errorManager).init()
	}
	err := errors.New(message)
	errorManagerInstance.errorMapper[err] = code
	return err
}

// RegErrorRef 通过错误码注册错误，返回错误的引用
//   - 引用将会被重定向到注册的错误信息
func RegErrorRef(code int, message string, ref error) error {
	if code == 0 {
		return errors.New("error code can not be 0")
	}
	errorManagerMutex.Lock()
	defer errorManagerMutex.Unlock()
	if errorManagerInstance == nil {
		errorManagerInstance = new(errorManager).init()
	}
	err := errors.New(message)
	errorManagerInstance.errorMapper[err] = code
	errorManagerInstance.errorMapperRef[ref] = err
	return ref
}

// GetError 通过错误引用获取错误码和真实错误信息，如果错误不存在则返回 0，如果错误引用不存在则返回原本的错误
func GetError(err error) (int, error) {
	if errorManagerInstance == nil {
		return 0, err
	}
	unw := errors.Unwrap(err)
	if unw == nil {
		unw = err
	}
	errorManagerMutex.Lock()
	defer errorManagerMutex.Unlock()
	if ref, exist := errorManagerInstance.errorMapperRef[unw]; exist {
		//err = fmt.Errorf("%w : %s", ref, err.Error())
		err = ref
	}
	unw = errors.Unwrap(err)
	if unw == nil {
		unw = err
	}
	code, exist := errorManagerInstance.errorMapper[unw]
	if !exist {
		return 0, errors.New("error not found")
	}
	return code, err
}

type errorManager struct {
	errorMapper    map[error]int
	errorMapperRef map[error]error
	mutex          sync.Mutex
}

func (slf *errorManager) init() *errorManager {
	slf.errorMapper = make(map[error]int)
	slf.errorMapperRef = make(map[error]error)
	return slf
}
