package fileproc

import (
	"os"
)

// CheckIsDir 检查指定路径是否是文件夹
func CheckIsDir(path string) bool {
	f, err := os.Stat(path)
	if err != nil {
		return false
	}
	return f.IsDir()
}
