package fileproc

import "os"

// CheckPathExist 检查指定路径是否存在
func CheckPathExist(path string) bool {
	if _, err := os.Stat(path); err == nil {
		return true
	}
	return false
}
