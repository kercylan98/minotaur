package super

import "os/exec"

// GoFormat go 代码格式化
func GoFormat(filePath string) {
	cmd := exec.Command("gofmt", "-w", filePath)
	_ = cmd.Run()
}
