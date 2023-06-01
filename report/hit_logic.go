package report

import "github.com/kercylan98/minotaur/utils/generic"

// HitLogic 埋点命中逻辑
//   - data: 当前数据
//   - input: 新输入的数据
//   - return: 设置当前数据
type HitLogic[Data any] func(data Data, input Data) Data

// WithHitLogicCustomize 通过自定义命中处理逻辑处理命中事件
//   - 通过对旧数据和新数据的处理，将返回的值作为新的数据值
func WithHitLogicCustomize[Data any](logic func(data Data, input Data) Data) HitLogic[Data] {
	return func(data Data, input Data) Data {
		return logic(data, input)
	}
}

// WithHitLogicCover 通过覆盖原数值的方式处理命中事件
func WithHitLogicCover[Data any]() HitLogic[Data] {
	return func(data Data, input Data) Data {
		return input
	}
}

// WithHitLogicOverlay 通过数值叠加的方式处理命中事件
//   - 例如原始值为100，新输入为200，最终结果为300
func WithHitLogicOverlay[Data generic.Number]() HitLogic[Data] {
	return func(data Data, input Data) Data {
		return data + input
	}
}
