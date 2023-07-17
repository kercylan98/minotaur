package configuration

// Loader 配置加载器
type Loader interface {
	// Load 加载配置
	//   - 加载后并不会刷新线上配置，需要执行 Refresh 函数对线上配置进行刷新
	Load()
	// Refresh 刷新线上配置
	Refresh()
}
