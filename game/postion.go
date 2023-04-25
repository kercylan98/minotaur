package game

// Position 是一个兼容2D和3D的位置接口，用于表示游戏中对象的空间位置。
// 该接口提供了获取和设置X、Y和Z轴坐标的方法，以便在2D和3D环境中灵活使用。
// 通过实现此接口，游戏对象可以方便地在不同的坐标系中进行转换和操作。
type Position interface {
	// GetX 获取X轴坐标
	GetX() float64
	// GetY 获取Y轴坐标
	GetY() float64
	// GetZ 获取Z轴坐标
	GetZ() float64
	// GetXY 获取X和Y轴坐标
	GetXY() (x, y float64)
	// GetXYZ 获取X、Y和Z轴坐标
	GetXYZ() (x, y, z float64)
	// SetX 设置X轴坐标
	SetX(x float64)
	// SetY 设置Y轴坐标
	SetY(y float64)
	// SetZ 设置Z轴坐标
	SetZ(z float64)
	// SetXY 设置X和Y轴坐标
	SetXY(x, y float64)
	// SetXYZ 设置X、Y和Z轴坐标
	SetXYZ(x, y, z float64)
	// Clone 克隆当前位置到新结构体
	Clone() Position
	// Compare 比较两个坐标是否相同
	Compare(position Position) bool
	// RegPositionChangeEvent 当位置发生改变时，将立即执行注册的事件处理函数
	RegPositionChangeEvent(handle PositionChangeEventHandle)
	OnPositionChangeEvent(old, new Position)
}

type (
	PositionChangeEventHandle func(old, new Position)
)
