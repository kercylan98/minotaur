// Package game 定义了通用游戏相关的接口
//   - actor.go：游戏通用对象接口定义
//   - aoi2d.go：基于2D的感兴趣领域(Area Of Interest)接口定义
//   - aoi2d_entity.go：基于2D的感兴趣领域(Area Of Interest)对象接口定义
//   - attrs.go：游戏属性接口定义，属性通常为直接读取配置，是否合理暂不清晰，目前不推荐使用
//   - fsm.go：有限状态机接口定义
//   - fsm_state.go：有限状态机状态接口定义
//
// 其中 builtin 包内包含了内置的实现
package game
