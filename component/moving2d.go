package component

// Moving2D 2D移动功能接口定义
type Moving2D interface {
	// MoveTo 设置对象移动至特定位置
	MoveTo(entity Moving2DEntity, x float64, y float64)
	// StopMove 终止特定对象的移动
	StopMove(guid int64)

	// RegPosition2DChangeEvent 对象位置改变时将立即执行被注册的事件处理函数
	RegPosition2DChangeEvent(handle Position2DChangeEventHandle)
	OnPosition2DChangeEvent(entity Moving2DEntity, oldX, oldY float64)

	// RegPosition2DDestinationEvent 对象抵达终点时将立即执行被注册的事件处理函数
	RegPosition2DDestinationEvent(handle Position2DDestinationEventHandle)
	OnPosition2DDestinationEvent(entity Moving2DEntity)
}

type (
	Position2DChangeEventHandle      func(moving Moving2D, entity Moving2DEntity, oldX, oldY float64)
	Position2DDestinationEventHandle func(moving Moving2D, entity Moving2DEntity)
	Position2DStopMoveEventHandle    func(moving Moving2D, entity Moving2DEntity)
)
