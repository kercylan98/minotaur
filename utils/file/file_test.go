package file_test

import (
	"fmt"
	"github.com/kercylan98/minotaur/utils/file"
	"strings"
	"testing"
)

func TestFilePaths(t *testing.T) {
	var line int
	var fileCount int
	for _, path := range file.FilePaths(`D:\sources\minotaur`) {
		if !strings.HasSuffix(path, ".go") {
			continue
		}
		fmt.Println(path)
		line += file.LineCount(path)
		fileCount++
	}
	fmt.Println("total line:", line, "total file:", fileCount)
}
