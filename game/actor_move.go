package game

// ActorMove 是一个将Actor和Position结合的接口，表示可以移动的游戏对象。
// 通过实现此接口，游戏对象可以在游戏世界中进行移动和操作。
type ActorMove interface {
	Actor
	Position
	// MoveTo2D 将游戏对象移动到指定的位置。
	// 该方法将对象的位置设置为指定的坐标，即绝对移动。
	// 参数x、y分别表示对象在X、Y轴上的目标位置。
	MoveTo2D(x, y float64)
	// MoveBy2D 在当前位置基础上移动游戏对象。
	// 该方法将对象的位置增加指定的偏移量，即相对移动。
	// 参数dx、dy分别表示对象在X、Y轴上的偏移量。
	MoveBy2D(dx, dy float64)
	// MoveTo3D 将游戏对象移动到指定的位置。
	// 该方法将对象的位置设置为指定的坐标，即绝对移动。
	// 参数x、y和z分别表示对象在X、Y和Z轴上的目标位置。
	MoveTo3D(x, y, z float64)
	// MoveBy3D 在当前位置基础上移动游戏对象。
	// 该方法将对象的位置增加指定的偏移量，即相对移动。
	// 参数dx、dy和dz分别表示对象在X、Y和Z轴上的偏移量。
	MoveBy3D(dx, dy, dz float64)
	// MoveToX 将游戏对象在X轴上移动到指定位置。
	// 该方法将对象的X坐标设置为指定的位置，即绝对移动。
	// 参数x表示对象在X轴上的目标位置。
	MoveToX(x float64)
	// MoveByX 在当前位置基础上移动游戏对象在X轴上的位置。
	// 该方法将对象的X坐标增加指定的偏移量，即相对移动。
	// 参数dx表示对象在X轴上的偏移量。
	MoveByX(dx float64)
	// MoveToY 将游戏对象在Y轴上移动到指定位置。
	// 该方法将对象的Y坐标设置为指定的位置，即绝对移动。
	// 参数y表示对象在Y轴上的目标位置。
	MoveToY(y float64)
	// MoveByY 在当前位置基础上移动游戏对象在Y轴上的位置。
	// 该方法将对象的Y坐标增加指定的偏移量，即相对移动。
	// 参数dy表示对象在Y轴上的偏移量。
	MoveByY(dy float64)
	// MoveToZ 将游戏对象在Z轴上移动到指定位置。
	// 该方法将对象的Z坐标设置为指定的位置，即绝对移动。
	// 参数z表示对象在Z轴上的目标位置。
	MoveToZ(z float64)
	// MoveByZ 在当前位置基础上移动游戏对象在Z轴上的位置。
	// 该方法将对象的Z坐标增加指定的偏移量，即相对移动。
	// 参数dz表示对象在Z轴上的偏移量。
	MoveByZ(dz float64)
	// GetSpeed 获取对象的移动速度
	GetSpeed() float64
	// SetSpeed 设置对象的移动速度
	SetSpeed(speed float64)
}
