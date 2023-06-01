package report

type GlobalBuriedOption[Data any] func(buried *GlobalBuried[Data])

// WithGlobalBuriedStorage 通过特定的存储模式创建全局埋点
//   - 默认情况下埋点数据存储在内存中
//   - 使用该方式可以将埋点存储存储在其他如数据库、消息队列中
func WithGlobalBuriedStorage[DataID comparable, Data any](getData func() Data, setData func(data Data)) GlobalBuriedOption[Data] {
	return func(buried *GlobalBuried[Data]) {
		buried.getData = getData
		buried.setData = setData
	}
}
