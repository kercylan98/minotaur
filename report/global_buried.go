package report

// NewGlobalBuried 创建一个全局埋点
func NewGlobalBuried[Data any](hitLogic HitLogic[Data], options ...GlobalBuriedOption[Data]) *GlobalBuried[Data] {
	buried := &GlobalBuried[Data]{
		hitLogic: hitLogic,
	}
	buried.setData = func(data Data) {
		buried.data = data
	}
	buried.getData = func() Data {
		return buried.data
	}
	for _, option := range options {
		option(buried)
	}
	return buried
}

// GlobalBuried 全局埋点
//   - 全局埋点适用于活跃用户数等统计
type GlobalBuried[Data any] struct {
	data     Data
	hitLogic HitLogic[Data]
	getData  func() Data
	setData  func(data Data)
}

// Hit 命中数据埋点
func (slf *GlobalBuried[Data]) Hit(data Data) {
	slf.setData(slf.hitLogic(slf.getData(), data))
}

// GetData 获取数据
func (slf *GlobalBuried[Data]) GetData() Data {
	return slf.getData()
}
