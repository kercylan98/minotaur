package genreadme

import (
	"io/fs"
	"path/filepath"
	"strings"
	"testing"
)

func TestBuilder_Generate(t *testing.T) {
	//b, err := New(`/Users/kercylan/Coding.localized/Go/minotaur/utils/buffer`, `/Users/kercylan/Coding.localized/Go/minotaur/utils/buffer/README.md`)
	//if err != nil {
	//	panic(err)
	//}
	//if err = b.Generate(); err != nil {
	//	panic(err)
	//}
	//return
	filepath.Walk("D:/sources/minotaur", func(path string, info fs.FileInfo, err error) error {
		if !info.IsDir() {
			return nil
		}
		if strings.Contains(strings.TrimPrefix(path, "D:/sources/minotaur"), ".") {
			return nil
		}
		b, err := New(
			path,
			filepath.Join(path, "README.md"),
		)
		if err != nil {
			return nil
		}
		if err = b.Generate(); err != nil {
			panic(err)
		}
		return nil
	})

}
