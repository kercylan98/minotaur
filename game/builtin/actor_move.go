package builtin

import (
	"minotaur/game"
)

type ActorMove struct {
	game.Actor
	game.Position
	speed float64 // 移动速度
}

func (slf *ActorMove) MoveTo2D(x, y float64) {
	//TODO implement me
	panic("implement me")
}

func (slf *ActorMove) MoveBy2D(dx, dy float64) {
	//TODO implement me
	panic("implement me")
}

func (slf *ActorMove) MoveTo3D(x, y, z float64) {
	//TODO implement me
	panic("implement me")
}

func (slf *ActorMove) MoveBy3D(dx, dy, dz float64) {
	//TODO implement me
	panic("implement me")
}

func (slf *ActorMove) MoveToX(x float64) {
	//TODO implement me
	panic("implement me")
}

func (slf *ActorMove) MoveByX(dx float64) {
	//TODO implement me
	panic("implement me")
}

func (slf *ActorMove) MoveToY(y float64) {
	//TODO implement me
	panic("implement me")
}

func (slf *ActorMove) MoveByY(dy float64) {
	//TODO implement me
	panic("implement me")
}

func (slf *ActorMove) MoveToZ(z float64) {
	//TODO implement me
	panic("implement me")
}

func (slf *ActorMove) MoveByZ(dz float64) {
	//TODO implement me
	panic("implement me")
}

func (slf *ActorMove) GetSpeed() float64 {
	//TODO implement me
	panic("implement me")
}

func (slf *ActorMove) SetSpeed(speed float64) {
	//TODO implement me
	panic("implement me")
}
