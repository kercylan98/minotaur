package game

// AOI2D 感兴趣的领域(Area Of Interest)接口定义
type AOI2D interface {
	// AddEntity 添加对象
	AddEntity(entity AOIEntity2D)
	// DeleteEntity 移除对象
	DeleteEntity(entity AOIEntity2D)
	// Refresh 刷新对象焦点
	Refresh(entity AOIEntity2D)
	// GetFocus 获取对象焦点列表
	GetFocus(guid int64) map[int64]AOIEntity2D
	// SetSize 设置总区域大小
	//  - 将会导致区域的重新划分
	SetSize(width, height int)
	// SetAreaSize 设置区域大小
	//  - 将会导致区域的重新划分
	SetAreaSize(width, height int)

	// RegEntityJoinVisionEvent 在新对象进入视野时将会立刻执行被注册的事件处理函数
	RegEntityJoinVisionEvent(handle EntityJoinVisionEventHandle)
	OnEntityJoinVisionEvent(entity, target AOIEntity2D)
	// RegEntityLeaveVisionEvent 在对象离开视野时将会立刻执行被注册的事件处理函数
	RegEntityLeaveVisionEvent(handle EntityLeaveVisionEventHandle)
	OnEntityLeaveVisionEvent(entity, target AOIEntity2D)
}

type (
	EntityJoinVisionEventHandle  func(entity, target AOIEntity2D)
	EntityLeaveVisionEventHandle func(entity, target AOIEntity2D)
)
