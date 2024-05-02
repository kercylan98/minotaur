package router

// MultistageOption 路由器选项
type MultistageOption[HandleFunc any] func(multistage *Multistage[HandleFunc])

// WithRouteTrim 路由修剪选项
//   - 将在路由注册前对路由进行对应处理
func WithRouteTrim[HandleFunc any](handle func(route any) any) MultistageOption[HandleFunc] {
	return func(multistage *Multistage[HandleFunc]) {
		multistage.trim = handle
	}
}
