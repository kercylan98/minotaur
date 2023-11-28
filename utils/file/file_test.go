package file_test

import (
	"fmt"
	"github.com/kercylan98/minotaur/utils/file"
	"strings"
	"testing"
	"time"
)

func TestFilePaths(t *testing.T) {
	var line int
	var fileCount int
	for _, path := range file.Paths(`D:\sources\minotaur`) {
		if !strings.HasSuffix(path, ".go") {
			continue
		}
		fmt.Println(path)
		line += file.LineCount(path)
		fileCount++
	}
	fmt.Println("total line:", line, "total file:", fileCount)
}

func TestNewIncrementReader(t *testing.T) {
	n, _ := file.ReadLineWithParallel(`./test/t.log`, 1*1024*1024*1024, func(s string) {
		t.Log(s)
	})

	time.Sleep(time.Second * 3)
	n, _ = file.ReadLineWithParallel(`./test/t.log`, 1*1024*1024*1024, func(s string) {
		t.Log(s)
	}, n)

}
