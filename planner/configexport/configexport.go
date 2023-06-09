package configexport

import (
	"bytes"
	"fmt"
	"github.com/kercylan98/minotaur/planner/configexport/internal"
	"github.com/kercylan98/minotaur/utils/file"
	"github.com/kercylan98/minotaur/utils/log"
	"github.com/kercylan98/minotaur/utils/str"
	"github.com/tealeg/xlsx"
	"go.uber.org/zap"
	"os/exec"
	"path/filepath"
	"strings"
	"text/template"
)

// New 创建一个导表配置
func New(xlsxPath string) *ConfigExport {
	ce := &ConfigExport{xlsxPath: xlsxPath, exist: make(map[string]bool)}
	xlsxFile, err := xlsx.OpenFile(xlsxPath)
	if err != nil {
		panic(err)
	}
	for i := 0; i < len(xlsxFile.Sheets); i++ {
		sheet := xlsxFile.Sheets[i]
		if config, err := internal.NewConfig(sheet, ce.exist); err != nil {
			switch err {
			case internal.ErrReadConfigFailedSame:
				log.Warn("ConfigExport",
					zap.String("File", xlsxPath),
					zap.String("Sheet", sheet.Name),
					zap.String("Info", "A configuration with the same name exists, skipped"),
				)
			case internal.ErrReadConfigFailedIgnore:
				log.Info("ConfigExport",
					zap.String("File", xlsxPath),
					zap.String("Sheet", sheet.Name),
					zap.String("Info", "Excluded non-configuration table files"),
				)
			default:
				log.ErrorHideStack("ConfigExport",
					zap.String("File", xlsxPath),
					zap.String("Sheet", sheet.Name),
					zap.String("Info", "Excluded non-configuration table files"),
				)
			}
		} else {
			if config == nil {
				continue
			}
			ce.configs = append(ce.configs, config)
			ce.exist[config.Name] = true

			log.Info("ConfigExport",
				zap.String("File", xlsxPath),
				zap.String("Sheet", sheet.Name),
				zap.String("Info", "Export successfully"),
			)
		}
	}
	return ce
}

type ConfigExport struct {
	xlsxPath string
	configs  []*internal.Config
	exist    map[string]bool
}

// Merge 合并多个导表配置
func Merge(exports ...*ConfigExport) *ConfigExport {
	if len(exports) == 0 {
		return nil
	}
	if len(exports) == 1 {
		return exports[0]
	}
	var export = exports[0]
	for i := 1; i < len(exports); i++ {
		ce := exports[i]
		for _, config := range ce.configs {
			if _, ok := export.exist[config.Name]; ok {
				log.Warn("ConfigExport",
					zap.String("File", ce.xlsxPath),
					zap.String("Sheet", config.Name),
					zap.String("Info", "A configuration with the same name exists, skipped"),
				)
				continue
			}
			export.configs = append(export.configs, config)
			export.exist[config.Name] = true
		}
	}
	return export
}

func (slf *ConfigExport) ExportClient(prefix, outputDir string) {
	for _, config := range slf.configs {
		config := config
		if len(prefix) > 0 {
			config.Prefix = fmt.Sprintf("%s.", prefix)
		}
		if err := file.WriterFile(filepath.Join(outputDir, fmt.Sprintf("%s%s.json", config.Prefix, config.Name)), config.JsonClient()); err != nil {
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
		if err := file.WriterFile(filepath.Join(outputDir, fmt.Sprintf("%s%s.json", config.Prefix, config.Name)), config.JsonServer()); err != nil {
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

	filePath := filepath.Join(outputDir, "config.go")
	if err := file.WriterFile(filePath, []byte(result)); err != nil {
		panic(err)
	}

	cmd := exec.Command("gofmt", "-w", filePath)
	_ = cmd.Run()
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

	filePath := filepath.Join(outputDir, "config.define.go")
	if err := file.WriterFile(filePath, []byte(result)); err != nil {
		panic(err)
	}

	cmd := exec.Command("gofmt", "-w", filePath)
	_ = cmd.Run()
}
