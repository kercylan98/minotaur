package configxport

import (
	"fmt"
	"github.com/kercylan98/minotaur/planner/configexport/internal"
	"github.com/kercylan98/minotaur/utils/file"
	"github.com/kercylan98/minotaur/utils/str"
	"github.com/tealeg/xlsx"
	"os"
	"os/exec"
	"path/filepath"
	"strings"
)

func New(xlsxPath string) *ConfigExport {
	ce := &ConfigExport{xlsxPath: xlsxPath}
	xlsxFile, err := xlsx.OpenFile(xlsxPath)
	if err != nil {
		panic(err)
	}
	for i := 0; i < len(xlsxFile.Sheets); i++ {
		if config := internal.NewConfig(xlsxFile.Sheets[i]); config != nil {
			ce.configs = append(ce.configs, config)
		}
	}
	return ce
}

type ConfigExport struct {
	xlsxPath string
	configs  []*internal.Config
}

func (slf *ConfigExport) ExportJSON(outputDir string) {
	for _, config := range slf.configs {
		config := config
		if err := file.WriterFile(filepath.Join(outputDir, fmt.Sprintf("%s.json", config.GetName())), config.GetJSON()); err != nil {
			panic(err)
		}
		if err := file.WriterFile(filepath.Join(outputDir, fmt.Sprintf("%s.client.json", config.GetName())), config.GetJSONC()); err != nil {
			panic(err)
		}
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
		if config.GetIndexCount() == 0 {
			varsMake += fmt.Sprintf("_%sReady = new(%s)"+`
	if err := handle("%s.json", &_%sReady); err != nil {
		panic(err)
	}
`, str.FirstUpper(config.GetName()), strings.TrimPrefix(v, "*"), str.FirstUpper(config.GetName()), str.FirstUpper(config.GetName()))
		} else {
			varsMake += fmt.Sprintf("_%sReady = make(%s)"+`
	if err := handle("%s.json", &_%sReady); err != nil {
		panic(err)
	}
`, str.FirstUpper(config.GetName()), v, str.FirstUpper(config.GetName()), str.FirstUpper(config.GetName()))
		}
		types += fmt.Sprintf("%s\n", config.GetStruct())
		varsReplace += fmt.Sprintf("%s = _%sReady\n", str.FirstUpper(config.GetName()), str.FirstUpper(config.GetName()))
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
