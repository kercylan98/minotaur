package main

import (
	"fmt"
	"github.com/kercylan98/minotaur/planner/configexport/internal"
	"github.com/kercylan98/minotaur/utils/file"
	"github.com/kercylan98/minotaur/utils/log"
	"github.com/kercylan98/minotaur/utils/str"
	"github.com/tealeg/xlsx"
	"go.uber.org/zap"
	"os"
	"os/exec"
	"path/filepath"
	"runtime/debug"
	"sync"
)

func New(xlsxPath string) *ConfigExport {
	ce := &ConfigExport{xlsxPath: xlsxPath}
	xlsxFile, err := xlsx.OpenFile(xlsxPath)
	if err != nil {
		panic(err)
	}
	for i := 0; i < len(xlsxFile.Sheets); i++ {
		ce.configs = append(ce.configs, internal.NewConfig(xlsxFile.Sheets[i]))
	}
	return ce
}

type ConfigExport struct {
	xlsxPath string
	configs  []*internal.Config
}

func (slf *ConfigExport) ExportJSON(outputDir string) {
	var errors []func()
	var wait sync.WaitGroup
	for _, config := range slf.configs {
		config := config
		go func() {
			wait.Add(1)
			defer func() {
				if err := recover(); err != nil {
					errors = append(errors, func() {
						log.Error("导出失败", zap.String("名称", slf.xlsxPath), zap.String("Sheet", config.GetName()), zap.Any("err", err))
						fmt.Println(debug.Stack())
					})
				}
			}()
			if err := file.WriterFile(filepath.Join(outputDir, fmt.Sprintf("%s.json", config.GetName())), config.GetJSON()); err != nil {
				panic(err)
			}
			wait.Done()
		}()
	}

	wait.Wait()

	for _, f := range errors {
		f()
	}
}

func (slf *ConfigExport) ExportGo(packageName string, outputDir string) {
	var vars string
	var varsMake string
	var types string
	var varsReplace string
	for _, config := range slf.configs {
		v := config.GetVariable()
		vars += fmt.Sprintf("var %s %s\nvar _%sReady %s\n", str.FirstUpper(config.GetName()), v, str.FirstUpper(config.GetName()), v)
		varsMake += fmt.Sprintf("_%sReady = make(%s)"+`
	if err := handle("%s.json", &_%sReady); err != nil {
		panic(err)
	}
`, str.FirstUpper(config.GetName()), v, str.FirstUpper(config.GetName()), str.FirstUpper(config.GetName()))
		types += fmt.Sprintf("%s\n", config.GetStruct())
		varsReplace += fmt.Sprintf("%s = _%sReady", str.FirstUpper(config.GetName()), str.FirstUpper(config.GetName()))
	}

	_ = os.MkdirAll(outputDir, 0666)
	if err := file.WriterFile(filepath.Join(outputDir, "config.struct.go"), []byte(fmt.Sprintf(internal.TemplateStructGo, packageName, types))); err != nil {
		panic(err)
	}
	if err := file.WriterFile(filepath.Join(outputDir, "config.go"), []byte(fmt.Sprintf(internal.TemplateGo, packageName, vars, varsMake, varsReplace))); err != nil {
		panic(err)
	}
	cmd := exec.Command("gofmt", "-w", filepath.Join(outputDir, "config.struct.go"))
	if err := cmd.Run(); err != nil {
		fmt.Println(err)
	}

	cmd = exec.Command("gofmt", "-w", filepath.Join(outputDir, "config.go"))
	if err := cmd.Run(); err != nil {
		fmt.Println(err)
	}

}
