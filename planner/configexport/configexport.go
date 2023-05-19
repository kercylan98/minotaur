package configexport

import (
	"bytes"
	"fmt"
	"github.com/kercylan98/minotaur/planner/configexport/internal"
	"github.com/kercylan98/minotaur/utils/file"
	"github.com/kercylan98/minotaur/utils/str"
	"github.com/tealeg/xlsx"
	"path/filepath"
	"strings"
	"text/template"
)

func New(xlsxPath string) *ConfigExport {
	ce := &ConfigExport{xlsxPath: xlsxPath}
	xlsxFile, err := xlsx.OpenFile(xlsxPath)
	if err != nil {
		panic(err)
	}
	for i := 0; i < len(xlsxFile.Sheets); i++ {
		if config, err := internal.NewConfig(xlsxFile.Sheets[i]); err != nil {
			panic(err)
		} else {
			ce.configs = append(ce.configs, config)
		}
	}
	return ce
}

type ConfigExport struct {
	xlsxPath string
	configs  []*internal.Config
}

func (slf *ConfigExport) ExportClient(prefix, outputDir string) {
	for _, config := range slf.configs {
		config := config
		if len(prefix) > 0 {
			config.Prefix = fmt.Sprintf("%s.", prefix)
		}
		if err := file.WriterFile(filepath.Join(outputDir, fmt.Sprintf("%s%s.json", prefix, config.Name)), config.JsonClient()); err != nil {
			panic(err)
		}
	}
}

func (slf *ConfigExport) ExportServer(prefix, outputDir string) {
	for _, config := range slf.configs {
		config := config
		if len(prefix) > 0 {
			config.Prefix = fmt.Sprintf("%s.", prefix)
		}
		if err := file.WriterFile(filepath.Join(outputDir, fmt.Sprintf("%s%s.json", prefix, config.Name)), config.JsonServer()); err != nil {
			panic(err)
		}
	}
}

func (slf *ConfigExport) ExportGo(prefix, outputDir string) {
	if len(prefix) > 0 {
		for _, config := range slf.configs {
			config.Prefix = fmt.Sprintf("%s.", prefix)
		}
	}
	slf.exportGoConfig(outputDir)
	slf.exportGoDefine(outputDir)
}

func (slf *ConfigExport) exportGoConfig(outputDir string) {
	var v struct {
		Package string
		Configs []*internal.Config
	}
	v.Package = filepath.Base(outputDir)

	for _, config := range slf.configs {
		v.Configs = append(v.Configs, config)
	}

	tmpl, err := template.New("struct").Parse(internal.GenerateGoConfigTemplate)
	if err != nil {
		panic(err)
	}

	var buf bytes.Buffer
	if err = tmpl.Execute(&buf, &v); err != nil {
		panic(err)
	}
	var result string
	_ = str.RangeLine(buf.String(), func(index int, line string) error {
		if len(strings.TrimSpace(line)) == 0 {
			return nil
		}
		result += fmt.Sprintf("%s\n", strings.ReplaceAll(line, "\t\t", "\t"))
		if len(strings.TrimSpace(line)) == 1 {
			result += "\n"
		}
		return nil
	})

	if err := file.WriterFile(filepath.Join(outputDir, "config.go"), []byte(result)); err != nil {
		panic(err)
	}
}

func (slf *ConfigExport) exportGoDefine(outputDir string) {
	var v struct {
		Package string
		Configs []*internal.Config
	}
	v.Package = filepath.Base(outputDir)

	for _, config := range slf.configs {
		v.Configs = append(v.Configs, config)
	}

	tmpl, err := template.New("struct").Parse(internal.GenerateGoDefineTemplate)
	if err != nil {
		panic(err)
	}

	var buf bytes.Buffer
	if err = tmpl.Execute(&buf, &v); err != nil {
		panic(err)
	}
	var result string
	_ = str.RangeLine(buf.String(), func(index int, line string) error {
		if len(strings.TrimSpace(line)) == 0 {
			return nil
		}
		result += fmt.Sprintf("%s\n", strings.ReplaceAll(line, "\t\t", "\t"))
		if len(strings.TrimSpace(line)) == 1 {
			result += "\n"
		}
		return nil
	})

	if err := file.WriterFile(filepath.Join(outputDir, "config.define.go"), []byte(result)); err != nil {
		panic(err)
	}
}
