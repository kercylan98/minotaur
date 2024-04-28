package server

// ConsoleParams 控制台参数
type ConsoleParams map[string][]string

// Get 获取参数值
func (slf ConsoleParams) Get(key string) string {
	if v, exist := slf[key]; exist && len(v) > 0 {
		return v[0]
	}
	return ""
}

// GetValues 获取参数值
func (slf ConsoleParams) GetValues(key string) []string {
	if v, exist := slf[key]; exist {
		return v
	}
	return []string{}
}

// GetValueNum 获取参数值数量
func (slf ConsoleParams) GetValueNum(key string) int {
	if v, exist := slf[key]; exist {
		return len(v)
	}
	return 0
}

// Has 是否存在参数
func (slf ConsoleParams) Has(key string) bool {
	_, exist := slf[key]
	return exist
}

// Add 添加参数
func (slf ConsoleParams) Add(key, value string) {
	if _, exist := slf[key]; !exist {
		slf[key] = []string{}
	}
	slf[key] = append(slf[key], value)
}

// Del 删除参数
func (slf ConsoleParams) Del(key string) {
	delete(slf, key)
}

// Clear 清空参数
func (slf ConsoleParams) Clear() {
	for k := range slf {
		delete(slf, k)
	}
}
