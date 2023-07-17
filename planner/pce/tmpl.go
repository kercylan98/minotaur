package pce

// Tmpl 配置结构模板接口
type Tmpl interface {
	// Render 渲染模板
	Render(templates []*TmplStruct) (string, error)
}
