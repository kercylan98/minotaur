package reporter

type BuriedOption[Data any] func(buried *Buried[Data])

// WithErrorHandle 通过包含上报错误处理函数的方式创建数据埋点
func WithErrorHandle[Data any](handle func(buried *Buried[Data], err error)) BuriedOption[Data] {
	return func(buried *Buried[Data]) {
		buried.errorHandle = handle
	}
}
