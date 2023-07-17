package pce

import (
	"github.com/kercylan98/minotaur/utils/file"
	"os/exec"
	"path/filepath"
)

// NewExporter 创建导出器
func NewExporter(filePath string) *Exporter {
	return &Exporter{
		filePath: filePath,
	}
}

// Exporter 导出器
type Exporter struct {
	filePath string
}

// ExportStruct 导出结构
func (slf *Exporter) ExportStruct(tmpl Tmpl, tmplStruct ...*TmplStruct) error {
	filePath, err := filepath.Abs(slf.filePath)
	if err != nil {
		return err
	}
	raw, err := tmpl.Render(tmplStruct)
	if err != nil {
		return err
	}
	if err = file.WriterFile(filePath, []byte(raw)); err != nil {
		return err
	}

	cmd := exec.Command("gofmt", "-w", filePath)
	return cmd.Run()
}

// ExportData 导出数据
func (slf *Exporter) ExportData(tmpl DataTmpl, data map[any]any) error {
	filePath, err := filepath.Abs(slf.filePath)
	if err != nil {
		return err
	}
	raw, err := tmpl.Render(data)
	if err != nil {
		return err
	}
	if err = file.WriterFile(filePath, []byte(raw)); err != nil {
		return err
	}
	return nil
}
