package pce

// DataTmpl 数据导出模板
type DataTmpl interface {
	// Render 渲染模板
	Render(data map[any]any) (string, error)
}
