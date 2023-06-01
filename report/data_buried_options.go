package report

type DataBuriedOption[DataID comparable, Data any] func(buried *DataBuried[DataID, Data])

// WithDataBuriedStorage 通过特定的存储模式创建数据埋点
//   - 默认情况下埋点数据存储在内存中
//   - 使用该方式可以将埋点存储存储在其他如数据库、消息队列中
func WithDataBuriedStorage[DataID comparable, Data any](getData func(id DataID) Data, setData func(id DataID, data Data)) DataBuriedOption[DataID, Data] {
	return func(buried *DataBuried[DataID, Data]) {
		buried.getData = getData
		buried.setData = setData
	}
}
