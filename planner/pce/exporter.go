package pce

// NewExporter 创建导出器
func NewExporter() *Exporter {
	return &Exporter{}
}

// Exporter 导出器
type Exporter struct{}

// ExportStruct 导出结构
func (slf *Exporter) ExportStruct(tmpl Tmpl, tmplStruct ...*TmplStruct) ([]byte, error) {
	raw, err := tmpl.Render(tmplStruct)
	if err != nil {
		return nil, err
	}
	return []byte(raw), nil
}

// ExportData 导出数据
func (slf *Exporter) ExportData(tmpl DataTmpl, data map[any]any) ([]byte, error) {
	raw, err := tmpl.Render(data)
	if err != nil {
		return nil, err
	}

	return []byte(raw), nil
}
