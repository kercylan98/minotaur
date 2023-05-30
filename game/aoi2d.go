package game

// AOI2D 基于2D定义的AOI功能接口
//   - AOI（Area Of Interest）翻译过来叫感兴趣的区域，是大型多人在线的游戏服务器中一个非常重要的基础模块，用于游戏对象在场景中的视野管理
//   - 透过 AOI 系统可以在其他玩家或怪物进入视野时得到感知，从而进行相应处理
//   - 内置实现：builtin.AOI2D
//   - 构建函数：builtin.NewAOI2D
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
