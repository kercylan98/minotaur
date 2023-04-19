package runtimes

import (
	"log"
	"os"
	"path"
	"path/filepath"
	"runtime"
	"strings"
)

// GetWorkingDir 获取工作目录绝对路径
func GetWorkingDir() string {
	dir := GetExecutablePathByBuild()
	if strings.Contains(dir, GetTempDir()) {
		return GetExecutablePathByCaller()
	}
	return dir
}

// GetTempDir 获取系统临时目录
func GetTempDir() string {
	dir := os.Getenv("TEMP")
	if dir == "" {
		dir = os.Getenv("TMP")
	}
	res, _ := filepath.EvalSymlinks(dir)
	return res
}

// GetExecutablePathByBuild 获取当前执行文件绝对路径（go build）
func GetExecutablePathByBuild() string {
	exePath, err := os.Executable()
	if err != nil {
		log.Fatal(err)
	}
	res, _ := filepath.EvalSymlinks(filepath.Dir(exePath))
	return res
}

// GetExecutablePathByCaller 获取当前执行文件绝对路径（go run）
func GetExecutablePathByCaller() string {
	var abPath string
	_, filename, _, ok := runtime.Caller(3)
	if ok {
		abPath = path.Dir(filename)
	}
	return abPath
}
