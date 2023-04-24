package game

// ActorFollow 是一个将ActorMove和Position结合的接口，表示可以跟随目标移动的游戏对象。
// 通过实现此接口，游戏对象可以跟随指定的目标进行移动和操作。
type ActorFollow interface {
	ActorMove
	// FollowTarget2D 让游戏对象跟随目标在2D平面上移动。
	// 该方法将对象的位置设置为目标的位置加上指定的偏移量，即相对移动。
	// 参数target表示要跟随的目标位置，dx、dy表示对象在X、Y轴上的偏移量。
	FollowTarget2D(target Position, dx, dy float64)
	// FollowTarget3D 让游戏对象跟随目标在3D空间中移动。
	// 该方法将对象的位置设置为目标的位置加上指定的偏移量，即相对移动。
	// 参数target表示要跟随的目标位置，dx、dy、dz表示对象在X、Y、Z轴上的偏移量。
	FollowTarget3D(target Position, dx, dy, dz float64)
}
