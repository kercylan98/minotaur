package log

import (
	"path/filepath"
	"strconv"
)

// CallerBasicFormat 返回调用者的基本格式
func CallerBasicFormat(file string, line int) (repFile, refLine string) {
	return filepath.Base(file), strconv.Itoa(line)
}
