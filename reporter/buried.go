package reporter

type BuriedHitHandle[Data any] func(actions *HitOperationSet[Data])
type BuriedReportingHandle[Data any] func(name string, data Data) error

// NewBuried 新建数据埋点
//   - 通常数据埋点应该在全局创建唯一埋点实例进行使用
//   - 需要指定数据埋点名称，由 reporter 包中的全局上报器统一调度上报行为
//   - hitHandle：当埋点命中时将执行该函数，该函数适用于数据处理及持久化等行为，不应该在该函数中执行上报操作
//   - reportingHandle：当数据上报时将执行该函数
//
// 数据埋点必须注册后才会生效
func NewBuried[Data any](name string, hitHandle BuriedHitHandle[Data], reportingHandle BuriedReportingHandle[Data], options ...BuriedOption[Data]) *Buried[Data] {
	if hitHandle == nil {
		panic("data buried point without hit processing function")
	}
	if reportingHandle == nil {
		panic("data buried point without reporting processing function")
	}
	buried := &Buried[Data]{
		name:            name,
		operationSet:    &HitOperationSet[Data]{},
		hitHandle:       hitHandle,
		reportingHandle: reportingHandle,
	}
	for _, option := range options {
		option(buried)
	}
	return buried
}

// Buried 埋点
type Buried[Data any] struct {
	name            string
	operationSet    *HitOperationSet[Data]
	hitHandle       BuriedHitHandle[Data]
	reportingHandle BuriedReportingHandle[Data]
	errorHandle     func(buried *Buried[Data], err error)
}

// GetName 获取埋点数据名称
func (slf *Buried[Data]) GetName() string {
	return slf.name
}

// Hit 命中数据埋点
func (slf *Buried[Data]) Hit(data Data) {
	slf.operationSet.input = data
	slf.hitHandle(slf.operationSet)
}

// Reporting 上报数据埋点
func (slf *Buried[Data]) Reporting() error {
	return slf.reportingHandle(slf.name, slf.operationSet.GetData())
}

// Register 注册到上报器
//   - 注册数据埋点必须要指定最少一个策略
//   - 相同策略可注册多个（ ReportStrategyInstantly 除外）
func (slf *Buried[Data]) Register(strategies ...ReportStrategy[Data]) *Buried[Data] {
	if len(strategies) == 0 {
		panic("invalid data buried without any strategy")
	}
	if registerBuried.Exist(slf.name) {
		panic("repeated data buried point registration")
	}
	registerBuried.Set(slf.name, true)
	for _, strategy := range strategies {
		strategy(slf)
	}
	return slf
}
