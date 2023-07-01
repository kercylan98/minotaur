package configexport_test

import (
	"fmt"
	"github.com/kercylan98/minotaur/config"
	"github.com/kercylan98/minotaur/planner/configexport"
	"github.com/kercylan98/minotaur/planner/configexport/example"
	"os"
	"path/filepath"
	"strings"
)

func ExampleNew() {
	var workdir = "./"
	files, err := os.ReadDir(workdir)
	if err != nil {
		panic(err)
	}
	var ces []*configexport.ConfigExport
	for _, file := range files {
		if file.IsDir() || !strings.HasSuffix(file.Name(), ".xlsx") || strings.HasPrefix(file.Name(), "~") {
			continue
		}

		ces = append(ces, configexport.New(filepath.Join(workdir, file.Name())))
	}

	c := configexport.Merge(ces...)
	outDir := filepath.Join(workdir, "example")
	c.ExportGo("", outDir)
	c.ExportServer("", outDir)
	c.ExportClient("", outDir)

	// 下方为配置加载代码
	// 使用生成的 LoadConfig 函数加载配置
	config.Init(outDir, example.LoadConfig, example.Refresh)

	fmt.Println("success")

	// Output:
	// success
}
